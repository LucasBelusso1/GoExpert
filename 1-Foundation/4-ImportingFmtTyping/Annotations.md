### Importando pacote e tipagem

Para importar um pacote no Go, no início do arquivo é necessário utilizar a palavra reservada `import` juntamente como o
`package`, desta forma:

```GO
import "fmt"

import ( // Múltiplas importações.
	"fmt"
)
```

OBS.: Da mesma forma que as variáveis, não é possível importar um pacote e não utilizá-lo.