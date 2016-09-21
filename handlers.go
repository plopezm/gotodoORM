package main

import (
	"encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"io"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
)

func getDB() (*gorm.DB, error){
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db, err;
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	var todos Todos;

	db, err:= getDB();
	defer db.Close();

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError);
		fmt.Fprintf(w, "error: %s",err)
		return;
	}

	db.Find(&todos);

	w.Header().Set("Content-Type", "application/json; charset=UTF-8");
	w.WriteHeader(http.StatusOK);

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	var todo Todo;

	vars := mux.Vars(r);
	todoId := vars["todoId"];

	db, err := getDB();

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError);
		fmt.Fprintf(w, "error: %s",err)
		return;
	}

	db.Find(&todo, todoId);

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}

}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo;
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return;
	}

	db, err:= getDB();
	defer db.Close();

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError);
		fmt.Fprintf(w, "error: %s",err)
		return;
	}

	db.Create(&todo);

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func TodoComplete(w http.ResponseWriter, r *http.Request){
	var todo Todo;
	vars := mux.Vars(r);
	todoId := vars["todoId"];

	i, err := strconv.Atoi(todoId);
	if err != nil {
		w.WriteHeader(http.StatusBadRequest);
		fmt.Fprintf(w, "<todoId> must be integer")
		return;
	}
	db, err:= getDB();
	defer db.Close();

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError);
		fmt.Fprintf(w, "error: %s",err)
		return;
	}

	db.Find(&todo, i);
	todo.completeTask();
	db.Update(&todo);

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func TodoRemove(w http.ResponseWriter, r *http.Request) {
	var todo Todo;
	vars := mux.Vars(r);
	todoId := vars["todoId"];

	i, err := strconv.Atoi(todoId);
	if err != nil {
		w.WriteHeader(http.StatusBadRequest);
		fmt.Fprintf(w, "<todoId> must be integer")
		return;
	}
	db, err:= getDB();
	defer db.Close();

	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError);
		fmt.Fprintf(w, "error: %s",err)
		return;
	}

	db.Delete(&todo, i);
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}



