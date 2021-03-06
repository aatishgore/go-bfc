/**
 * this is used to for common db operation
 */

package utilities

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // not export
)

// DbConnect is help to connect a database and retun response as db object
func DbConnect(db *gorm.DB) *gorm.DB {
	// connecting to database
	dbString := GvDatabaseVariable.User + ":" + GvDatabaseVariable.Password + "@tcp(" + GvDatabaseVariable.Host + ":" + GvDatabaseVariable.Port + ")/" + GvDatabaseVariable.Database
	db, err := gorm.Open(GvDatabaseVariable.Type, dbString)

	// checking a error in connection
	if err != nil {
		log.Fatalln("Unable to connect a db please check credentials")
		return nil
	}

	return db
}
