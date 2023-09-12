### Testify

Neste caso utiliza-se uma lib (pacote) para auxiliar nos testes. O pacote que foi utilizado é o
`github.com/stretchr/testify/assert`, que faz com que os testes fiquem semelhantes a outras bibliotecas de outras
linguagens com o sistema de assertions e mocking. Exemplo:"

```GO
package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "Amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
}
```

### Trabalhando com mocks

No exemplo criamos a função `CalculateTaxAndSave` que chama a função `CalculateTax` e a partir
de um objeto `repository`, chama o `SaveTax`.

```GO
type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)

	return repository.SaveTax(tax)
}
```

Neste caso a ideia é somente testar a função `CalculateTaxAndSave`, então não nos interessa o que a função `SaveTax` vai
retornar, então podemos "mocka-la".
Veja que no exemplo acima, criamos uma interface que possui a função `SaveTax`, sendo assim, precisamos criar uma
estrutura que possua esta função implementada. No caso dos nossos testes, vamos criar esta estrutura utilizando o pacote
**github.com/stretchr/testify/mock**:

```GO
type TaxRepositoryMock struct {
	mock.Mock
}

func (mock TaxRepositoryMock) SaveTax(tax float64) error {
	args := mock.Called(tax)
	return args.Error(0)
}
```

No mock acima, informamos ao objeto mock, através do `Called` que ele irá receber o parâmetro tax que é um float64 e
irá retornar Error, pois é isso que a interface espera.
Em seguida desenvolvemos nosso teste:

```GO
func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{} //Instância do mock
	repository.On("SaveTax", 10.0).Return(nil) // Se tax for 10, SaveTax retornará nil.
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax")) // Se tax for 0, SaveTax retornará errors.

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2) // Verifica se o SaveTax realmente foi chamado 2 vezes.
}
```