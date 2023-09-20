# APIS

### Estrutura de diretórios.

Guia para organização de pastas: https://github.com/golang-standards/project-layout

`api`: Swagger doc, documentações no geral.
`cmd`: Código da aplicação em si, dentro dele é comum encontrar uma pasta com o nome do projeto.
`configs`: Arquivos de configuração, conexão com banco de dados...
`internal`: Código da aplicação não reutilizável.
`pkg`: Código da aplicação que é reutilizável.
`test`: Arquivos que auxiliam nos testes, não são necessáriamente arquivos .go.

### Criando arquivo de configuração.

Na pasta `/configs` foi criado um arquivo `config.go`, em que nele foi criada uma struct que comportará todos os dados
de banco, JWT e webservice.

### Finalizando configuração.

Agora, utilizando o pacote `github.com/spf13/viper`, criamos uma função e dizemos para o viper qual é o nosso arquivo
de configuração, aonde que ele vai estar localizado, qual o tipo do arquivo e falamos para ele automaticamente setar o
environment.

```GO
func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}
```

Em seguida, dizemos para o viper através de notations qual é o campo dentro do arquivo `.env` que será
representado na nossa struct:

```GO
type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}
```

### Outras possibilidades de configuração.

Neste caso vimos que é possível abordar o arquivo de configuração de outra maneira, transformando e suas propriedades em
atributos privados, criando getters para as propriedades porém sem a possibilidade de alterá-las.

### Criando a entidade user.

Neste caso é bem simples, basicamente criamos uma struct com alguns campos como ID, usuário, senha e e-mail e criamos
uma função para instanciar os campos de usuário, senha e e-mail.

Para o ID, foi criado dentro da pasta `pkg` um tipo para comportar o UUID do pacote do google, assim na nossa struct
definimos que o ID será do tipo uuid, e nele também criamos as funções para criar um novo UUID e parsea-lo.

Já para a senha foi utilizado sistema de hash do pacote `golang.org/x/crypto/bcrypt` para persistir o dado em banco.
Junto a geração do hash, também foi criada a função para validar se o hash e a senha batem.

```GO
type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}
```

### Testando a entidade user.

Para testar a entidade user foi criado então o arquivo de teste, e nele foi novamente utilizado o pacote `assert` do
`testify` para comprar os valores.

Na primeira função foi testado a geração de um novo usuário:

```GO
func TestNewUser(t *testing.T) {
	user, err := NewUser("Dev", "dev@dev.com", "asdf000")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Dev", user.Name)
	assert.Equal(t, "dev@dev.com", user.Email)
}
```

Na segunda função foi testado a validação da senha:

```GO
func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Dev", "dev@dev.com", "asdf000")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("asdf000"))
	assert.False(t, user.ValidatePassword("asdf00"))
	assert.NotEqual(t, "asdf000", user.Password)
}
```

Para a entidade Produto seguiu-se praticamente o mesmo raciocínio, com a exceção de que no produto criamos uma função de
validate para validar os campos que estão sendo inseridos.

### Criando UserDB

Na criação do UserDB criamos dentro da pasta `internal`, a pasta `infra` seguida da pasta `database`, nela criamos
a seguinte interface dentro do pacte `database`:

```GO
type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
```

Em seguida criamos o arquivo `user_db.go` no qual criamos a struct User e implementamos as funções da interface
`UserInterface`. Além disso criamos uma função chamada `NewUser` que recebe um DB como parâmetro e retorna uma entidade
de usuário.

```GO
type User struct {
	DB *gorm.DB
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := u.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}
```

### Criando a entidade Product

Para criar a entidade product, seguiu-se o mesmo padrão do user, com a diferença que o product tem as funções de
validação, que executa uma validação em cima de todos os campos que foram passados ao criar um NewProduct. Além disso
o `ProductDB` conta com todas as funções do user mais as funções de `Update`, `Delete` e `FindAll`, desta forma:

```GO
func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	var products []entity.Product
	var err error
	if page > 0 && limit > 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}
```

E para os testes a estrutura também ficou bem parecida com a de Users, com a diferença dos testes das funções novas:

```GO
func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Iphone 15", 7300.00)
	productDB := NewProduct(db)

	productDB.Create(product)

	product.Name = "Iphone 14 pro"
	product.Price = 6999
	productDB.Update(product)

	productFound, err := productDB.FindByID(product.ID.String())

	assert.Nil(t, err)
	assert.Equal(t, productFound.ID, product.ID)
	assert.Equal(t, productFound.Name, product.Name)
	assert.Equal(t, productFound.Price, product.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Iphone 15", 7300.00)
	productDB := NewProduct(db)
	productDB.Create(product)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, productFound)
}
```

### Criando handler para Produto

Dentro do `main.go` do `cmd/server` foi inicializado o servidor com o pacote `http` como visto em outras aulas na porta
8000. Depois foi criado uma struct responsável por ser o handler de produtos, desta forma:

```GO
type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}
```

E então criamos a função de criação de produto a este handler:

```GO
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := entity.NewProduct(productDto.Name, productDto.Price)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
```

Para a criação do produto, utilizou-se o padrão de projeto de DTO, que nada mais é do que uma estrutura criada apenas
para representar os dados que deverão ser trafegados.
Para criar o DTO, criamos uma pasta chamada `/dto` dentro de `/internal` e nela criamos o arquivo `dto.go` contendo o
seguinte:

```GO
type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
```

Desta forma, ao fazermos o decode do JSON recebido, garantimos que os dados devem ser enviados desta forma... qualquer
dado diferente resultará em um 400 bad request.

Agora para chamar o handler fazemos desta forma:

```GO
productDB := database.NewProduct(db)
productHandler := NewProductHandler(productDB)

http.HandleFunc("/products", productHandler.CreateProduct)
```

Após realizada a prova de conceito, criamos dentro de `internal/infra` a pasta `/webserver/handlers` que conterá o
arquivo `product_handlers.go`, no qual passaremos a lógica do handler pra dentro.

### Roteadores

Neste projeto utilizaremos o roteador `github.com/go-chi/chi/v5`, que facilitará no desenvolvimento de rotas permitindo
a utilização de middlewares prontos, personalização de rotas e uma série de outras features.
Para inicializar com o `chi`, escrevemos o seguinte código:

```GO
r := chi.NewRouter()
r.Use(middleware.Logger)
r.Post("/products", productHandler.CreateProduct)
http.ListenAndServe(":8000", r)
```

Primeiramente iniciamos um NewRouter do `chi`, opcionalmente dizemos para o chi utilizar o middleware de logs, depois
definimos a rota de `/products` da mesma forma que foi feito utilizando o pacote `http` e iniciamos o servidor
normalmente com o pacote `http` passando o router como handler

### CRUD products

Da mesma forma que fizemos para o `create`, agora vamos criar o as funções de `getById`, `getAll`, `update` e `delete`,
desta forma:

```GO
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := entityPkg.ParseID(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := entityPkg.ParseID(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
```

### Token JWT

Neste caso, criamos a lógica para obtenção de um token JWT dentro das rotas de usuário, que em suma ficaram muito
parecidas com as rotas de produto.
Então para gerarmos um token, primeiro implementamos 2 outros campos em nosso `UserHandler`, para que ele receba o token
gerado dentro do `main.go` com o secret que definimos no `.env` e também o expiresIn também vindo do `.env`:

```GO
type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExipresIn int
}
```

Em seguida implementamos a função:

```GO
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var jwtDTO dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&jwtDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(jwtDTO.Email)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !user.ValidatePassword(jwtDTO.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExipresIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
```

Na função acima definimos mais um DTO para comportar os campos de "login", buscamos o usuário pelo e-mail fornecido no
corpo da requisição. Com o usuário retornado, validamos a senha que foi retornada e a que foi informada no corpo e só
então geramos o token JTW e retornamos.

### Protegendo as rotas de produto.

Aqui apenas informamos ao `chi` para utilizar o middleware de autenticação. Da seguinte forma:

```GO
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.Delete)
	})
```

Acima, definimos um grupo de rotas e nela informamos ao `chi` para que utilize o nosso token gerado com o secret do
`.env` como verificador e informamos a ele para que autentique a requisição.

### Criando middlewares

Para implementar um middleware, utilizamos o conceito de `Chain of Responsibility`, que consiste em uma fila de handlers
a serem executados um após o outro. Veja o exemplo de um middleware de log:

```GO
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
```

### Gerando a primeira documentação

Primeiramente é necessário instalar op binário encontrado em `https://github.com/swaggo/swag` através do comando:

```SHELL
go install github.com/swaggo/swag/cmd/swag@latest
```

Após instalado, inserimos algumas annotations em formato de comentário em nosso código que é o que irá gerar a doc do
swagger e rodamos o comando `swag init -g cmd/server/main.go`. Dependendo de onde estão as annotations, é possível rodar
apenas `swag init main.go` por exemplo.

Para acessar a documentação, precisamos primeiramente importar o pacote `httpSwagger "github.com/swaggo/http-swagger"` e
também precisamos importar a nossa pasta que contém a doc gerada
(ex: `_ "github.com/LucasBelusso1/GoExpert/7-APIS/docs"`). Feito isso criamos uma rota e com o `http-swagger` apontamos
para o nosso `doc.json`:

```GO
r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
```

Agora ao acessar `http://localhost:8000/docs/index.html` teremos acesso a nossa doc.

### Documentando user.

Para documentar a criação de users, utiliamos as seguintes annotations:

```GO
// Create		user doc
// @Summary		Create user
// @Description	Create user
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		request	body dto.CreateUserInput true "user request"
// @Success 	201
// @Failure 	500	{object}	Error
// @Router 		/users [post]
```

Para documentar a geração do JWT foram utilizadas as seguintes annotations:

```GO
// GetJWT doc
// @Summary		Get a user JWT
// @Description	Get a user JWT
// @tags		users
// @Accept		json
// @Produce		json
// @Param		request body dto.GetJWTInput true "user credentials"
// @Success		200 {object} dto.GetJWTOutput
// @Failure		404 {object} Error
// @Failure		500 {object} Error
// @Router		/users/generate_token [post]
```
Note que no caso acima, precisamos criar um DTO para o output do token JWT, seguindo a mesma lógica do input.

### Documentando products.

