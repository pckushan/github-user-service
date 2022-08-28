package repository

import (
	"database/sql"
	"fmt"
	"github-user-service/internal/domain/adaptors/logger"
	"github-user-service/internal/domain/models"
	"github.com/google/uuid"
	"log"
)

type Repository struct {
	db  *sql.DB
	log logger.Logger
}

func NewRepository(logger logger.Logger) Repository {
	psqlConn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode= disable",
		Config.Host, Config.Port, Config.User, Config.Password, Config.UserName)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		log.Fatalf("error opening db connection")
	}
	defer db.Close()

	return Repository{
		db:  db,
		log: logger,
	}
}

func (r Repository) Insert(record interface{}) (id uuid.UUID, err error) {
	user := record.(models.User)
	userid := uuid.New()
	rows := r.db.QueryRow(fmt.Sprintf(`INSERT INTO users(login, name, unique_id,followers,repos)
	VALUES(%s, %s, %s,%d,%d) RETURNING unique_id`, user.Login, user.Name, userid, user.Followers, user.PublicRepos)).Scan(&userid)

	if rows == nil {
		r.log.Error(fmt.Sprintf("insert record error for user_id, %s", userid))
		return userid, fmt.Errorf("error inserting record for user_id: %s", userid)
	}

	return userid, nil
}

func (r Repository) Update(record interface{}) error {
	user := record.(models.User)
	var userid uuid.UUID
	err := r.db.QueryRow(fmt.Sprintf(`UPDATE users SET 
                 login = %s, 
                 name = %s, 
                 followers = %d,
                 repos = %d) RETURNING unique_id`, user.Login, user.Name, user.Followers, user.PublicRepos)).Scan(&userid)

	if err != nil {
		r.log.Error(fmt.Sprintf("check record error with, %s", err))
		return err
	}
	return nil
}

func (r Repository) Check(id uuid.UUID) bool {
	v, err := r.db.Query(`SELECT id FROM users WHERE unique_id = $1`, id)
	if err != nil {
		r.log.Error(fmt.Sprintf("check record error with, %s", err))
		return false
	}

	if v == nil {
		return false
	}

	return true

}
