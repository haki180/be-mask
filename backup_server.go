package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	// "github.com/gin-gonic/gin/binding"
// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-sql-driver/mysql"
// 	_ "github.com/go-sql-driver/mysql"
// 	psh "github.com/platformsh/config-reader-go/v2"
// 	sqldsn "github.com/platformsh/config-reader-go/v2/sqldsn"
// 	"github.com/platformsh/template-golang/handler/datahdl"
// 	"github.com/platformsh/template-golang/handler/mockhdl"
// 	"github.com/platformsh/template-golang/handler/userhdl"
// 	"github.com/platformsh/template-golang/repository/datarps"
// 	"github.com/platformsh/template-golang/repository/userrps"
// 	"github.com/platformsh/template-golang/service/datasvc"
// 	"github.com/platformsh/template-golang/service/usersvc"
// )

// type Test struct {
// 	Username string `json:"username"`
// }

// func main() {

// 	// The Config Reader library provides Platform.sh environment information mapped to Go structs.
// 	config, err := psh.NewRuntimeConfig()
// 	if err != nil {
// 		panic("Not in a Platform.sh Environment.")
// 	}

// 	// // Set up an extremely simple web server response.
// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

// 	// 	// Name the template on front page and print hello message
// 	// 	fmt.Println("WELCOM")
// 	// 	fmt.Fprintf(w, "Hello, world! - A simple Go template for Platform.sh\n\n")

// 	// 	// Run some background SQL, just to prove we can.
// 	// 	trySql(config, w)

// 	// })

// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

// 	// 	if r.Method != http.MethodPost {
// 	// 		w.Write([]byte("method salah"))
// 	// 		return
// 	// 	}

// 	// 	request := new(Test)
// 	// 	if err := binding.Bind(r, request.FieldMap()); err.Handle(w) {
// 	// 		return
// 	// 	}

// 	// 	w.WriteHeader(http.StatusCreated)
// 	// 	w.Header().Set("Content-Type", "application/json")
// 	// 	resp := make(map[string]interface{})
// 	// 	resp["message"] = "Status Created"
// 	// 	resp["data"] = request
// 	// 	jsonResp, err := json.Marshal(resp)
// 	// 	if err != nil {
// 	// 		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
// 	// 	}
// 	// 	w.Write(jsonResp)
// 	// 	return

// 	// })

// 	// // The port to listen on is defined by Platform.sh.
// 	// log.Fatal(http.ListenAndServe(":"+config.Port(), nil))

// 	port := config.Port()

// 	if port == "" {
// 		log.Fatal("$PORT must be set")
// 	}

// 	postgres, err := mysql.Connect()
// 	if err != nil {
// 		log.Panic("can't connect to postgres")
// 	}

// 	router := gin.New()
// 	router.Use(gin.Logger(), cors.New(cors.Config{
// 		AllowOrigins: []string{"*"},
// 		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
// 		AllowHeaders: []string{"Authorization", "Access-Control-Allow-Origin", "Content-Type", "User-Agent"},
// 	}))

// 	// initialize repository
// 	userRepo := userrps.NewUserRepository(postgres)
// 	dataRepo := datarps.NewDataRepository(postgres)

// 	// initialize service
// 	userService := usersvc.NewUserService(userRepo)
// 	dataService := datasvc.NewDataService(dataRepo)

// 	// initialize handler
// 	userHandler := userhdl.NewUserHandler(userService)
// 	dataHandler := datahdl.NewDataHandler(dataService, userService)

// 	router.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "welcome",
// 		})
// 	})
// 	router.GET("/health", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "healthy",
// 		})
// 	})

// 	// mock api
// 	mockApi := router.Group("/mock")
// 	mockApi.POST("/login", mockhdl.Login)
// 	mockApi.POST("/logout", mockhdl.Logout)
// 	mockApi.POST("/datas/chart", mockhdl.ChartData)
// 	mockApi.POST("/datas/paginate", mockhdl.PaginateData)

// 	// user api
// 	userApi := router.Group("/users")
// 	userApi.POST("/", userHandler.Create)

// 	// auth api
// 	authApi := router.Group("/auth")
// 	authApi.POST("/login", userHandler.Login)
// 	authApi.POST("/logout", userHandler.Logout)

// 	// data api
// 	dataApi := router.Group("/datas")
// 	dataApi.POST("/", dataHandler.Create)
// 	dataApi.DELETE("/:uuid", dataHandler.Delete)
// 	dataApi.GET("/chart", dataHandler.ChartData)
// 	dataApi.GET("/paginate", dataHandler.PaginateData)

// 	router.Run(":" + port)

// }

// // trySql simply connects to a MySQL server defined by Platform.sh and
// // writes and reads from it.  This is not particularly useful code,
// // but demonstrates how you can leverage the Platform.sh Config Reader library.
// func trySql(conf *psh.RuntimeConfig, w http.ResponseWriter) {

// 	// Accessing the database relationship Credentials struct
// 	credentials, err := conf.Credentials("database")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Using the sqldsn formatted credentials package
// 	formatted, err := sqldsn.FormattedCredentials(credentials)
// 	if err != nil {
// 		panic(err)
// 	}

// 	db, err := sql.Open("mysql", formatted)
// 	checkErr(err)

// 	// Force MySQL into modern mode.
// 	db.Exec("SET NAMES=utf8")
// 	db.Exec("SET sql_mode = 'ANSI,STRICT_TRANS_TABLES,STRICT_ALL_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,ONLY_FULL_GROUP_BY'")

// 	_, err = db.Exec("DROP TABLE IF EXISTS userinfo")
// 	checkErr(err)

// 	_, err = db.Exec(`CREATE TABLE userinfo (
// 			uid INT(10) NOT NULL AUTO_INCREMENT,
// 			username VARCHAR(64) NULL DEFAULT NULL,
// 			departname VARCHAR(128) NULL DEFAULT NULL,
// 			created DATE NULL DEFAULT NULL,
// 			PRIMARY KEY (uid)
// 			) DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;`)
// 	checkErr(err)

// 	// insert
// 	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
// 	checkErr(err)

// 	res, err := stmt.Exec("platform", "Deploy Friday", "2019-06-17")
// 	checkErr(err)

// 	id, err := res.LastInsertId()
// 	checkErr(err)

// 	fmt.Println(id)
// 	// update
// 	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
// 	checkErr(err)

// 	res, err = stmt.Exec("goPlatformsh", id)
// 	checkErr(err)

// 	affect, err := res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println(affect)

// 	// query
// 	rows, err := db.Query("SELECT * FROM userinfo")
// 	checkErr(err)

// 	for rows.Next() {
// 		var uid int
// 		var username string
// 		var department string
// 		var created string
// 		err = rows.Scan(&uid, &username, &department, &created)
// 		checkErr(err)
// 		// Section for MySQL tests
// 		fmt.Fprintf(w, "\nMySQL Tests:\n\n")
// 		fmt.Fprintf(w, "* Connect and add row:\n")
// 		fmt.Fprintf(w, "   - Row ID (1): ")
// 		fmt.Fprintln(w, uid)
// 		fmt.Fprintf(w, "   - Username (goPlatformsh): ")
// 		fmt.Fprintln(w, username)
// 		fmt.Fprintf(w, "   - Department (Deploy Friday): ")
// 		fmt.Fprintln(w, department)
// 		fmt.Fprintf(w, "   - Created (2019-06-17): ")
// 		fmt.Fprintln(w, created)
// 	}

// 	// delete
// 	stmt, err = db.Prepare("delete from userinfo where uid=?")
// 	checkErr(err)

// 	res, err = stmt.Exec(id)
// 	checkErr(err)

// 	affect, err = res.RowsAffected()
// 	checkErr(err)

// 	fmt.Fprintf(w, "\n* Delete row:\n")
// 	fmt.Fprintf(w, "   - Status (1): ")
// 	fmt.Fprintln(w, affect)

// 	db.Close()
// }

// // checkErr is a simple wrapper for panicking on error.
// // It likely should not be used in a real application.
// func checkErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
