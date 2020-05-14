package main

import (
	"github.com/pmccau/TriviaGo/server/datamanagement"
	"github.com/pmccau/TriviaGo/server/router"
	"net/http"

	//"github.com/pmccau/TriviaGo/server/middleware"
	//"github.com/pmccau/TriviaGo/server/router"
	"log"
	//"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

func readCredFromFile(pathToFile string) string {
	data, err := ioutil.ReadFile(pathToFile)
	datamanagement.Check(err)
	return strings.TrimSpace(string(data))
}

func main() {
	//cred := readCredFromFile("assets/cred.txt")
	//uri := "mongodb+srv://admin:" + cred + "@cluster0-ciqct.mongodb.net/test?retryWrites=true&w=majority"
	//client, ctx := middleware.ConnectToMongo(uri)
	//middleware.PingClient(ctx, client)
	//triviaDb := middleware.GetDb("trivia", client)
	//questionsCollection := middleware.GetCollection("questions", triviaDb)
	//questions := middleware.GetAllDocuments(ctx, questionsCollection)
	//fmt.Println(questions)
	//

	r := router.Router()
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
