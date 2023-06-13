package main

import (
	"auth"
	database "db"
	"fmt"
	"log"
	"login"
	"structs"
	"time"
	"users"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() {

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

	//---------------------------------------------------------------

	createToken, err := auth.CreateAuthToken(db, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Users found: %v\n", createToken)

	//---------------------------------------------------------------
	err = auth.DeleteAuthToken(db, "gW10uSzXyFTKwkjMZ4ecQgl9WrjKOWulJ-5FNJzJtBHRWw17H1VExpHZF7wGjyosTVO03HW3goc84wpxoZ-ZqA==")
	if err != nil {
		fmt.Printf("Erro ao excluir o token: %v\n", err)
	} else {
		fmt.Println("Token excluído com sucesso!")
	}

	//---------------------------------------------------------------
	verifyToken, err := auth.VerifyAuthToken(db, "gW10uSzXyFTKwkjMZ4ecQgl9WrjKOWulJ-5FNJzJtBHRWw17H1VExpHZF7wGjyosTVO03HW3goc84wpxoZ-ZqA==")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token found: %v\n", verifyToken)

	//---------------------------------------------------------------
	err = users.DeleteUser(db, 11)
	if err != nil {
		fmt.Printf("Erro ao excluir o usuário: %v\n", err)

	} else {
		fmt.Println("Usuário excluído com sucesso!")

	}
	//---------------------------------------------------------------

	updatedUser := structs.Users{
		ID:        1,
		Name:      "Novo Nome",
		Email:     "novoemail@example.com",
		Password:  "novasenha",
		CreatedAt: time.Now(),
	}

	err = users.UpdateUser(db, updatedUser)
	if err == nil {
		fmt.Printf("Usuário atualizado com sucesso!")

	} else {
		fmt.Printf("Erro ao atualizar o usuário: %v\n", err)

	}

}

func main() {
	lambda.Start(handler)
}
