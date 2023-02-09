package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-database/entity"
	"strconv"
)

type userRepositoryImplementation struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImplementation{DB: db}
}

func (repository *userRepositoryImplementation) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	script := "INSERT INTO users(username, password) VALUES(?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, user.Username, user.Password)
	defer repository.DB.Close()

	if err != nil {
		return user, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}
	user.Id = int32(id)
	return user, nil
}

func (repository *userRepositoryImplementation) FindById(ctx context.Context, id int32) (entity.User, error) {
	script := "SELECT id, username, password FROM users WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	user := entity.User{}
	if err != nil {
		return user, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password)
		return user, nil
	} else {
		return user, errors.New("id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repository *userRepositoryImplementation) All(ctx context.Context) ([]entity.User, error) {
	script := "SELECT id, username, password FROM users"
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		user := entity.User{}

		rows.Scan(&user.Id, &user.Username, &user.Password)

		users = append(users, user)
	}
	return users, nil
}
