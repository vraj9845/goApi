package main

import (
	"database/sql"
	"fmt"
	productHttp "goApi/http"
	productService "goApi/services"
	productStore "goApi/stores"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// conf := MySQLConfig{Host: "localhost", User: "root", Password: "password", Port: "2051", Db: "goApi"}
	// db, err := ConnectToMySQL(conf)
	// if err != nil {
	// 	log.Println("could not connect to sql, err:", err)
	// 	return
	// }

	db, err := sql.Open("mysql", "root:password@(localhost:2051)/goApi") //sql.Open("mysql", "root:password1@tcp(0.0.0.0:2051)/goApi")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB Connected!")
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	productStore := productStore.New(db)
	productService := productService.New(productStore)
	productHttp := productHttp.New(&productService)

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/product", productHttp.Create).Methods("POST")
	// router.HandleFunc("/product", productHttp.Read).Methods("GET")
	// // router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	// router.HandleFunc("/events/{id}", productHttp.Update).Methods("PUT")
	// router.HandleFunc("/events/{id}", productHttp.Delete).Methods("DELETE")

	http.HandleFunc("/product", productHttp.Handler)
	fmt.Println(http.ListenAndServe(":21000", nil))
}

type MySQLConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Db       string
}

// ConnectToMySQL takes mysql config, forms the connection string and connects to mysql.
func ConnectToMySQL(conf MySQLConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.User, conf.Password, conf.Host, conf.Port, conf.Db)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
