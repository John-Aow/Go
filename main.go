package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"
	"time"

	"example.com/greetings/todo"
	jwt "github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

var index int
var tasks map[int]*Task = make(map[int]*Task)

type Task struct {
	Title string
	Done  bool
}

type NewTaskTodo struct {
	Task string `json:"task"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/auth", func(rw http.ResponseWriter, r *http.Request) {
		mySigningKey := []byte("password")
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
			Issuer:    "test",
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(map[string]string{
			"token": ss,
		})
	})
	api := r.NewRoute().Subrouter()
	api.Use(authMiddleware)
	api.HandleFunc("/todos", todo.GetTask).Methods(http.MethodGet)

	api.HandleFunc("/todos", todo.AddTask).Methods(http.MethodPut)

	api.HandleFunc("/todos/{index}", todo.SetTask).Methods(http.MethodPut)

	http.ListenAndServe(":8080", r)
}

func List() map[int]*Task {
	return tasks
}

func New(task string) {
	defer func() {
		index++
	}()

	tasks[index] = &Task{
		Title: task,
		Done:  false,
	}
}

func Power(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.ReplaceAll(tokenString, "bearer", "")
		mySigningKey := []byte("password")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpect %v", token.Header["alg"])
			}
			return mySigningKey, nil
		})
		if err != nil {

		}

	})
}
