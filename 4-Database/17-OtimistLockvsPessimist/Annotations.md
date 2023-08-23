### Lock otimista vs Pessimista

No lock otimista, a ideia é que o registro fique disponível para alteração a qualquer momento em qualquer processo que
esteja sendo executado, e neste caso o controle de alterações é versionado, fazendo assim com que no final da operação
seja verificado se a versão do dado permanece a mesma, se sim o dado é atualizado, senão é feito um rollback e a
operação executa novamente em cima da nova versão.

No lock pessimista, é feito uma trava no registro caso alguém esteja ocupando-o, desta forma é possível garantir que o
registro será atualizado, entretanto como o registro fica travado, outros processos não conseguem executar, podendo
gerar lentidão.

Nesta aula foi dado um exemplo de lock pessimista, desta maneira:

```GO
tx := db.Begin()
var c Category
err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
if err != nil {
	panic(err)
}
c.Name = "Eletronicos"
tx.Debug().Save(&c)
tx.Commit()
```

Neste exemplo acima, é criada uma transaction que adiciona uma cláusula de lock no registro, fazendo assim com que
ele só possa ser alterado após executar o UPDATE.