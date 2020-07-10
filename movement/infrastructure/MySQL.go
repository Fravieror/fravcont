package infrastructure

import (
	"api-test/movement/domain"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
	Db *sql.DB
}

func NewMySQLRepository() *MySQLRepository {
	repo := &MySQLRepository{}
	db, err := sql.Open("mysql", os.Getenv("USERMYSQL")+":"+os.Getenv("PASSMYSQL")+"@tcp("+os.Getenv("IPMYSQL")+")/"+os.Getenv("DBMYSQL"))
	if err != nil {
		return nil
	}
	repo.Db = db
	return repo
}

func (m *MySQLRepository) Save(movement domain.Movement) error {
	repo := NewMySQLRepository()
	query := fmt.Sprintf("INSERT INTO movement VALUES ( %q, %q, %q, %v, %q)", movement.Id, movement.Client, movement.Description, movement.Value, movement.Date)
	insert, err := repo.Db.Query(query)
	defer insert.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
func (m *MySQLRepository) GetList(client string) ([]domain.Movement, error) {
	repo := NewMySQLRepository()

	query := fmt.Sprintf("SELECT * FROM movement where client=%q", client)

	rows, err := repo.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movements []domain.Movement
	for rows.Next() {
		var movement domain.Movement
		err := rows.Scan(&movement.Id, &movement.Client, &movement.Description, &movement.Value, &movement.Date)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		movements = append(movements, movement)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return movements, nil
}
func (m *MySQLRepository) Get(id string) (*domain.Movement, error) {
	// db, err := sql.Open("mysql", "root:Fravier10*@tcp(127.0.0.1:3306)/Movements")
	// if err != nil {
	// 	return nil, err
	// }
	repo := NewMySQLRepository()
	query := fmt.Sprintf("SELECT * FROM movement WHERE Id = %q", id)

	rows, err := repo.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movement domain.Movement
	if !rows.Next() {
		return nil, nil
	}
	err = rows.Scan(&movement.Id, &movement.Client, &movement.Description, &movement.Value, &movement.Date)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &movement, nil
}
