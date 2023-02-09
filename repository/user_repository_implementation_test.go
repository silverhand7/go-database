package repository

import (
	"context"
	"fmt"
	go_database "go-database"
	"go-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestUserRepository(t *testing.T) {
	userRepository := NewUserRepository(go_database.GetConnection())
	ctx := context.Background()

	user := entity.User{
		Username: "repository_test",
		Password: "repository_password_test",
	}

	result, err := userRepository.Insert(ctx, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	userRepository := NewUserRepository(go_database.GetConnection())
	ctx := context.Background()

	user, err := userRepository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func TestFindAll(t *testing.T) {
	userRepository := NewUserRepository(go_database.GetConnection())
	ctx := context.Background()

	users, err := userRepository.All(ctx)
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(user.Id, user.Username, user.Password)
	}
}
