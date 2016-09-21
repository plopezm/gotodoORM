package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var routes = Routes{
    Route{
        "TodoIndex",
        "GET",
        "/api/v1/todos",
        TodoIndex,
    },
    Route{
        "TodoShow",
        "GET",
        "/api/v1/todos/{todoId}",
        TodoShow,
    },
    Route{
	"TodoCreate",
	"PUT",
	"/api/v1/todos",
	TodoCreate,
    },
    Route{
	"TodoComplete",
	"POST",
	"/api/v1/todos/complete/{todoId}",
	TodoComplete,
    },
    Route{
	"TodoRemove",
	"DELETE",
	"/api/v1/todos/{todoId}",
	TodoRemove,
    },
}

func initDatabase(){
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Todo{});

	// Create
	if err := db.AutoMigrate(Todo{}).Error; err != nil {
		txt := "AutoMigrate Job table failed"
		panic( fmt.Sprintf( "%s: %s", txt, err ) )
	}
}

func main() {
	initDatabase();

	port := ":8080";
	fmt.Println("====================================");
	fmt.Println("Starting server at port "+port);
	fmt.Println("====================================");

	router := NewRouter(routes);

	//Adding path as web-page server
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("webapp")));

	//Open new mongodb session
	//mdbOpenSession("localhost");

	fmt.Println("====================================");
	log.Fatal(http.ListenAndServe(port, router));
}
