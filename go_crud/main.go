package main

import (
	"context"
	"fmt"
	"github.com/PyMarcus/go_crud/configs"
	"github.com/PyMarcus/go_crud/repository"
	"log"
)

func main() {
	log.Println("Try to connect with database")
	container := repository.New(repository.Option{WriterSqlx: configs.GetWriterSqlx(), ReadSqlx: configs.GetReaderSqlx()})

	users, err := container.User.GetAll(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	for _, user := range users {
		fmt.Printf("User -> ID: %d Nickname: %s Name: %s Email: %s Create: %s",
			user.ID, user.NickName, user.Email, user.CreateAt)
	}
}
