package main

import (
	database "db"
	"fmt"
	"log"
	"login"
	"structs"
	"time"
	"users"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	//----------------------------------

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := login.Hello("Vinny")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	//----------------Buscar user pelo nome------------------

	userByName, err := users.UserByUsers(db, "Tony")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Users found: %v\n", userByName)

	//----------------Adiciona um user novo ------------------

	albID, err := users.AddUser(db, structs.Users{
		Name:      "Sergio",
		Email:     "sergio@example.com",
		Password:  "senha6",
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added user: %v\n", albID)
}
