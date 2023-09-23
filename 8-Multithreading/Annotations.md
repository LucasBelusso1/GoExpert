### Concorrência x Paralelismo

**Concorrência** é quando temos apenas um núcleo de processamento (CPU/Core) executando tarefas diferentes porém de
forma que não é necessário uma tarefa terminar para que a outra seja executada.
Exemplo: Temos a tarefa 1 e tarefa 2 rodando, a tarefa 1 pode rodar por 100ms e passar para a tarefa 2 rodar por mais
100ms e depois voltar para a tarefa 1 e assim sucessivamente, porém nunca rodando ambas ao mesmo tempo.

**Paralelismo** é quando as tarefas são executadas simultâneamente, necessitando assim de múltiplos núcleos de
processamento (CPU/Core).
Exemplo: Neste cenário, pegando o exemplo acima, a tarefa 1 e tarefa 2 seriam executadas simultaneamente.

No caso do GO, utilizam-se as **duas estratégias** de execução ao mesmo tempo, sendo assim, o GO consegue separar
tarefas diferentes em núcleos diferentes e mesmo assim executar o código de forma concorrente em cada um dos núcleos.

**Problemas**: Por conta do paralelismo e da concorrência, é possível que threads diferentes executando linhas de código
diferentes estejam querendo utilizadar o mesmo endereço de memória, ocasionando assim a chamada **Race condition**.

### Multithreading

Trabalhar com múltiplas threads pode parecer muito vantajoso, entretanto há alguns poréns a se considerar para se
trabalhar utilizando paralelismo:

- **Race conditions:** Como já comentado anteriormente, é necessário tomar muito cuidado ao desenvolver aplicações
com paralelismo para não ter inconsistência com os dados.

- **Memória:** Para cada Thread criada, é necessário que a aplicação chame o sistema operacional para solicitar uma
nova thread, e isso custa memória. Além disso, cada thread criada inicia com 1MB alocado de memória que posteriormente
pode escalar, o que também é muito custoso. E há também a troca de contexto de uma aplicação concorrente, que também
gera um custo de memória significativo a se considerar.

### Schedulers

Há dois tipos de schedulers dentro de um sistema operacional, são eles:

`Preemptivos`: Determina um tempo para que uma tarefa seja executada e depois passa a executar outra tarefa e assim por
diante.

`Cooperativos`: Determina a execução das tarefas em ordem de finalização, ou seja, espera uma tarefa executar para
executar a próxima.

### Como funciona no GO?

A linguagem GO possui seu **próprio scheduler** que lida com a criação e gerenciamento de threads de outra maneira, sem
precisar se comunicar com o sistema operacional e sem precisar alocar 1MB de memória para aquela thread.
Ao solicitar uma thread nova em GO, a **RUNTIME** do GO cria uma `Green Thread` ou `Threads in Userland`, que
praticamente não possui custo operacional e aloca apenas 2Kb de memória, possibilitando assim a criação de milhares de
threads para rodar processos em paralelo.
Além disso, o scheduler do GO executa as tarefas de modo `Cooperativo`, entretanto existem situações em que a própria
linguagem reconhece que aquela thread travará o sistema e de forma `Preemptiva` interrompe a execução.