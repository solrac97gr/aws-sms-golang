package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	. "github.com/solrac97gr/sendsms/api/database"
	. "github.com/solrac97gr/sendsms/api/messagesender"
	. "github.com/solrac97gr/sendsms/api/publicgenerator"
)

var err error

func RUN() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	CreateConection()
	public := GetPublic()
	AWS_ACCESS_KEY_ID := os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY := os.Getenv("AWS_SECRET_ACCESS_KEY")
	SendSMS_ABTEST(public, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY)
}
