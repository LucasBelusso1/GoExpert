### Unit of Work

O padrão de projeto **"Unit of Work"** vai de encontro ao design pattern de `repository`, onde para cada entidade
há uma interface que define as operações que o repositório irá executar, separando a camanda de domínio da aplicação
da camada de dados.

Utilizar o padrão **Unit of Work** em cima do `repository` faz com que possamos executar queries de múltiplos
`repositories` dentro de uma transação.

### Criando interface UOW

Dentro de `/pkg/uow` criamos o arquivo `uow.go` com o seguinte código:

```GO
type RepositoryFactory func(tx *sql.Tx) interface{} // Criamos o tipo do factory que retornará repositórios

type UowInterface interface {
	Register(name string, fc RepositoryFactory) // Registra um repositório
	Unregister(name string) // Remove um repositório
	GetRepository(ctx context.Context, name string) (interface{}, error) // Obtém um repositório
	Do(ctx context.Context, fn func(uow *UowInterface) error) error // Executa os repositórios
	CommitOrRollback() error // Commita em caso de sucesso ou executa o rollback em caso de erro
	Rollback() error // Executa o roolback separadamente
}
```

### Registrando repositórios

Para registrar e remover do registro os repositórios, escrevemos o seguinte código:

```GO
type Uow struct { // Criação da struct que receberá:
	Db           *sql.DB // Conexão com banco
	Tx           *sql.Tx // Transação
	Repositories map[string]RepositoryFactory // Map contendo nome do repositório e um RepositoryFactory de saida
}

func NewUow(ctx context.Context, db *sql.DB) *Uow { // Construtor
	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, fc RepositoryFactory) { // Registra dentro do map
	u.Repositories[name] = fc
}

func (u *Uow) Unregister(name string) { // Remove do map
	delete(u.Repositories, name)
}
```

### Implementando métodos principais

Agora iremos implementar os métodos `Rollback`, `CommitOrRollback` e o `Do`:

```GO
func (u *Uow) Rollback() error {
	if u.Tx == nil { // Verifica se existe uma transação
		return errors.New("no transaction to rollback")
	}
	err := u.Tx.Rollback() // Executa o rollback

	if err != nil { // Retorna o erro caso a rollback falhar
		return err
	}

	u.Tx = nil // Limpa a transação em caso de sucesso
	return nil
}

func (u *Uow) CommitOrRollback() error {
	err := u.Tx.Commit() // Commita a transação

	if err != nil {
		errRb := u.Tx.Rollback() // Caso o commit der erro, executa um rollback

		if errRb != nil { // Caso ambos commit e rollback derem erro, retorna o erro.
			return errors.New(fmt.Sprintf("Commit error: %s. Rollback error: %s", err.Error(), errRb.Error()))
		}

		return err
	}

	u.Tx = nil // Em caso de sucesso, limpa a transação
	return nil
}

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {
	if u.Tx != nil { // Verifica se há uma transação em andamento
		return fmt.Errorf("Transaction already started")
	}

	tx, err := u.Db.BeginTx(ctx, nil) // Começa a transação

	if err != nil {
		return err
	}
	u.Tx = tx // Propriedade da Uow recebe a transação para controle

	err = fn(u) // Executa as operações de acordo com os repositórios registrados

	if err != nil {
		errRb := u.Rollback() // Caso houver algum erro executa o rollback

		if errRb != nil { // Caso houver erro no rollback retorna ambos os erros
			return errors.New(fmt.Sprintf("Original error: %s. Rollback error: %s", err.Error(), errRb.Error()))
		}
	}

	return u.CommitOrRollback() // Executa o commit ou rollback caso o commit falhar
}
```

### Implementando GetRepository

Para o GetRepository faremos o seguinte:

```GO
func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.Tx == nil { // Se a transação não foi inicializada, inicia uma nova transação
		tx, err := u.Db.BeginTx(ctx, nil)

		if err != nil {
			return nil, err
		}

		u.Tx = tx
	}

	repo := u.Repositories[name](u.Tx) // Obtém o repositório passando o contexto
	return repo, nil
}
```

### Criando novo caso de uso

Agora adaptamos o caso de uso `add_course` para o seguinte:

```GO
type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUseCaseUow(uow uow.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error { // Chama a função que executará os repositórios dentro de uma transação
		category := entity.Category{ // Cria a entidade de categoria
			Name: input.CategoryName,
		}

		repoCategory := a.getCategoryRepository(ctx) // Obtém o repositório de categorias
		err := repoCategory.Insert(ctx, category) // Realiza o insert

		if err != nil {
			return err
		}

		course := entity.Course{ // Cria a entidade de curso
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}

		repoCourse := a.getCourseRepository(ctx) // Obtém o repositório de cursos
		err = repoCourse.Insert(ctx, course) // Realiza o insert

		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository") // Obtém repositório de categorias

	if err != nil {
		panic(err)
	}

	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository") // Obtém repositório de cursos

	if err != nil {
		panic(err)
	}

	return repo.(repository.CourseRepositoryInterface)
}
```

Agora no arquivo de teste inserimos o seguinte:

```GO
func TestAddCourseUow(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses") // Conecta com o banco
	assert.NoError(t, err)

	// Dropa as tabelas caso existam
	dbt.Exec("DROP TABLE if exists `courses`;")
	dbt.Exec("DROP TABLE if exists `categories`;")

	// Cria novas tables
	dbt.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dbt.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	ctx := context.Background()

	uow := uow.NewUow(ctx, dbt) // Inicializa o Unit of Work (UOW)
	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} { // Registra o repositório de categoria
		repo := repository.NewCategoryRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} { // Registra o repositório de cursos
		repo := repository.NewCourseRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	input := InputUseCase{ // Executa o usecase criado anteriormente
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryID: 1,
	}

	useCase := NewAddCourseUseCaseUow(uow)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
```

Neste cenário, caso alterarmos o valor de `CourseCategoryID` para 2, um erro será retornado e será feito o rollback da
alteração.