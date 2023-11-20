### Clean Architecture

- `Orientado a casos de uso`: Desenvolvemos com base na intenção da ação a ser executada, exemplo: Comprar um produto,
registrar uma categoria...
- `Baixo acoplamento entre camadas`: A ideia é dividir o sistema em camadas e fazer a comunicação entre as camadas
utilizando contratos.

[Livro recomendado (Em português)](https://www.amazon.com.br/Arquitetura-Limpa-Artes%C3%A3o-Estrutura-Software/dp/8550804606/ref=asc_df_8550804606/?tag=googleshopp00-20&linkCode=df0&hvadid=379787347388&hvpos=&hvnetw=g&hvrand=4221360949575654971&hvpone=&hvptwo=&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=1031481&hvtargid=pla-809227152896&psc=1)
[Versão em inglês](https://www.amazon.com.br/Clean-Architecture-Craftsmans-Software-Structure/dp/0134494164/ref=asc_df_0134494164/?tag=googleshopp00-20&linkCode=df0&hvadid=379726160779&hvpos=&hvnetw=g&hvrand=4221360949575654971&hvpone=&hvptwo=&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=1031481&hvtargid=pla-423658477418&psc=1)


### Pontos importantes sobre arquitetura

- Formato que o software terá.
- Divisão de componentes.
- Comunicação entre componentes.
- Uma boa arquitetura facilita o desenvolvimento, deploy, operação e manutenção.
- Deixar opções em aberto, ter flexibilidade dentro da arquitetura do software.

### Keep options open

Dar suporte ao ciclo de vida do sistema. O objetivo final é minimizar o custo de vida útil do sistema e maximizar a
produtividade do programador.

- Regras de negócio trazem o real valor para o software.
- Detalhes ajudam a suportar as regras, não devemos nos atentar tanto aos detalhes, devemos dar mais valor as regras
de negócio do software.
- Detalhes não devem impactar as regras de negócio. Quando um detalhe influencia na regra de negócio, isso significa que
não foi delimitada uma "camada".
- Framework, API's, banco de dados... não devem impactar nas regras, isso deve estar bem isolado dentro de suas devidas
camadas para que uma camada não interfira diretamente na outra.

### Use cases

- Intenção.
- Clareza de cada comportamento de software.
- Detalhes não devem impactar as regras de negócio.

**Use cases x Single Responsability Principle (SRP)**

- Temos a tendência de "reaproveitar" use cases por serem muito parecidos. Ex: Alterar ou inserir um registro no banco.
Ambos os processos são muito parecidos entre si, principalmente em cenários em que há a possibilidade de alterar a
entidade como um todo. Em ambos os casos, faremos uma conexão com o banco e validaremos todos os registros, então muito
do processo poderia ser reaproveitado (DRY - Don't Repeat Yourself). Entretanto, quando estamos falando em "alterar"
algo ou "inserir" algo, estamos falando de intenções diferentes, sendo assim não podemos reaproveitar o comportamento de
um no comportamento do outro, mesmo sendo idêntico.

- Como saber se estamos ferindo o princípio da responsábilidade única?
R: Quando alteramos o mesmo código por razões diferentes, sendo assim, aquele código não possui uma única
responsabilidade.

**Duplicação real vs acidental**

Nada é "escrito em pedra", mas geralmente a duplicação real é perceptível, quando o mesmo trecho de código aparece em
muitos lugares do sistema executando a exata mesma tarefa, enquanto que a duplicação acidental é uma duplicação
necessária por se tratar de fluxos diferentes, casos de uso diferentes.

### Fluxo de use cases

Use cases contam uma história, a partir de uma intenção, há uma série de passos a serem executados, validações a serem
feitas, informações a serem persistidas ou buscadas...Ou seja, há uma série de regras de negócios que compõem um use
case.

### Limites arquiteturais

Tudo que não impacta diretamente as regras de negócio, deve estar em um limite arquitetural diferente. Exemplo: Não será
o frontend ou o banco de dados que mudarão as regras de negócio da aplicação.

### Entendendo DTO's (Data Transfer Object)

- Ajuda a trafegar os dados entre os limites arquiteturais.
- Não possui comportamento, é um objeto "anêmico", ou seja, não possui regras, apenas dados.
- Contém dados (input ou output).
- Não faz nada, apenas guarda dados de uma forma específica e estruturada da maneira que quisermos.
- Cada intenção, geralmente terá um DTO diferente.

### Presenters

- Objetos de transformação
- Adequa o DTO de output no formato correto para entregar o resultado.
- Lembrando: Um sistema pode ter diversos formatos de entrega. Ex: XML, JSON, Protobuf, GraphQL, CLI, etc.

### Entities x DDD

Entidades na Clean Architecture representa uma camada enquanto que no Domain Driven Design, uma entidade representa algo
único na aplicação. Não há uma definição explicita de como criar uma entidade no Clean Architecture, não existem regras
de como criar uma entidade.