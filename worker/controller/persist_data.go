package controller

import (
	"database/sql"
	"fmt"
	"rabbit-worker/model"
)

// PersistData abre uma conexao com banco e persiste dados enviados para funcao, retorna um bool de acorco com sucesso
func PersistData(dataBase *sql.DB, data model.DataRabbit) bool {
	stmt, err := dataBase.Prepare("INSERT INTO rabbitmq(rabbit_uuid, rabbit_name) values(?, ?)")
	if err != nil {
		fmt.Println("Erro ao preparar query", err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Rabbit_uuid, data.Rabbit_name)
	if err != nil {
		fmt.Println("Erro ao executar query", err)
		return false
	}
	
	return true
}
