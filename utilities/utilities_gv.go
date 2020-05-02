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
	Type     string
	Port     string
}

// GvNetworkVariable give gloabl variable for network
var GvNetworkVariable *GvNetwork

// GvDatabaseVariable give global variale got database
var GvDatabaseVariable *GvDatabase

// Environment tell project currecnt environment
var Environment string

// InstanceAPIURL API
var InstanceAPIURL string

// InstanceAPIUserName API
var InstanceAPIUserName string

// InstanceAPIPassword API
var InstanceAPIPassword string

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
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		Type:     os.Getenv("DB_TYPE"),
		Port:     os.Getenv("DB_PORT"),
	}

	InstanceAPIURL = os.Getenv("INSTANCE_API_URL")
	InstanceAPIUserName = os.Getenv("INSTANCE_API_USERNAME")
	InstanceAPIPassword = os.Getenv("INSTANCE_API_PASSWORD")

}
