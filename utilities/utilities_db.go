/**
 * this is used to for common db operation
 */

package utilities

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // not export
)

// DbConnect is ...
// this is help to connect a database and retun response as db object
func DbConnect(db *gorm.DB) *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/bfc")

	// checking a error in connection
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return db
}
