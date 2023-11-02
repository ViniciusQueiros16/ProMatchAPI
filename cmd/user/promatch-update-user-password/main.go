package main

import (
    "context"
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/promatch/pkg/database"
    "github.com/promatch/pkg/utils/response"
    "github.com/promatch/structs"
    "golang.org/x/crypto/bcrypt"
)

func UpdatePassword(db *sql.DB, userID int64, newPassword string) error {
    // Criptografa a nova senha
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("UpdatePassword: %w", err)
    }

    // Atualiza a senha no banco de dados
    stmt, err := db.Prepare("UPDATE users SET password = ? WHERE id = ?")
    if err != nil {
        return fmt.Errorf("UpdatePassword: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(string(hashedPassword), userID)
    if err != nil {
        return fmt.Errorf("UpdatePassword: %w", err)
    }

    return nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    db, err := database.InitDB()
    if err != nil {
        return response.ApiResponse(http.StatusInternalServerError, structs.ErrorBody{
            ErrorMsg: aws.String(err.Error()),
        })
    }
    defer database.CloseDB()

    var requestBody struct {
        UserID      int64  `json:"user_id"`
        NewPassword string `json:"new_password"`
    }

    err = json.Unmarshal([]byte(request.Body), &requestBody)
    if err != nil {
        return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
            ErrorMsg: aws.String(err.Error()),
        })
    }

    err = UpdatePassword(db, requestBody.UserID, requestBody.NewPassword)
    if err != nil {
        return response.ApiResponse(http.StatusBadRequest, structs.ErrorBody{
            ErrorMsg: aws.String(err.Error()),
        })
    }

    return response.ApiResponse(http.StatusOK, "Password updated successfully")
}

func main() {
    lambda.Start(handler)
}
