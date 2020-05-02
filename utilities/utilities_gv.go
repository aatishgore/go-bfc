package utilities

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// GvNetwork give global variable for network
type GvNetwork struct {
	Port string
	Host string
}

// GvDatabase is globalvaribale for database connection
type GvDatabase struct {
	User     string
	Password string
	Database string
	Host     string
}

// GvNetworkVariable give gloabl variable for network
var GvNetworkVariable *GvNetwork

// GvDatabaseVariable give global variale got database
var GvDatabaseVariable *GvDatabase

// Environment tell project currecnt environment
var Environment string

// GetGlobalVariable function get global variable
func GetGlobalVariable() {

	// load global environment variable
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// assigning value for environment
	Environment = os.Getenv("ENVIRONMENT")

	// assigning value for network global varibale
	GvNetworkVariable = &GvNetwork{
		Port: os.Getenv("PORT"),
	}

	// assigning value for database variable
	GvDatabaseVariable = &GvDatabase{
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Database: os.Getenv("DATABASE"),
	}

}