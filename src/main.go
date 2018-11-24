package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

// Configure the upgrader
var upgrader = websocket.Upgrader{}

type requestType string

type request struct {
	//Type RequestType `json:"type"`
	Query string `json:"query"`
}

type response struct {
	Recipes []recipe `json:"recipes"`
}

type recipe struct {
	Name string `json:"name"`
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	publicDir := strings.Join([]string{dir, "public"}, "/")
	log.Printf(publicDir)

	// Create a simple file server
	fs := http.FileServer(http.Dir(publicDir))
	http.Handle("/", fs)

	// Confgure websocket route
	http.HandleFunc("/ws", HandleConnections)

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	//Make sure we close the connection when the function returns
	defer ws.Close()

	for {
		var msg request
		//Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: error when reading JSON; %v", err)
			break
		}
		recipes := GetRecipes(msg.Query)
		ws.WriteJSON(response{Recipes: recipes})
	}
}

func GetRecipes(query string) []recipe {
	return []recipe{recipe{"test"}}
}
