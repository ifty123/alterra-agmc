package factory

import (
	config "day2-task1/database"
	"day2-task1/internal/repository"
)

type Factory struct {
	UserRepository repository.UserDatabase
}

type FakeFactory struct {
	UserRepository repository.UserDatabase
}

func NewFactory() *Factory {
	db := config.InitDB()
	return &Factory{
		repository.NewUsersRepository(db),
	}
}
