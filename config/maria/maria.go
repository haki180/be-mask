package maria

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect( /**conf *psh.RuntimeConfig*/ ) (*gorm.DB, error) {
	var host, user, password, dbname, port string
	host = "localhost"
	user = "root"
	password = "password"
	dbname = "main"
	port = "33062"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	// Accessing the database relationship Credentials struct
	// credentials, err := conf.Credentials("database")
	// if err != nil {
	// 	panic(err)
	// }

	// // Using the sqldsn formatted credentials package
	// formatted, err := sqldsn.FormattedCredentials(credentials)
	// if err != nil {
	// 	panic(err)
	// }

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
