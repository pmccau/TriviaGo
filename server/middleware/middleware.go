package middleware

import (
	"github.com/pmccau/TriviaGo/server/data"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Response struct {
	ID			primitive.ObjectID
	Task 		string
	Status		bool
}

func Test(w http.ResponseWriter, r *http.Request) {
	message := "this is a test from golang"
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var res Response
	_ = json.NewDecoder(r.Body).Decode(&res)
	fmt.Println(res)
	json.NewEncoder(w).Encode(message)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	// Request the questions from the API
	var client = &http.Client{Timeout: 10*time.Second}
	apiRes, err := client.Get("https://opentdb.com/api_category.php")
	if err != nil {
		log.Fatal(err)
	}
	defer apiRes.Body.Close()

	// Parse API response to get the questions (in the 'results'). It's a bit roundabout to
	// marshal then unmarshal, but couldn't get it to work using other methods
	var temp interface{}
	json.NewDecoder(apiRes.Body).Decode(&temp)
	jsonStr, err := json.Marshal(temp)
	var parsedResponse map[string]interface{}
	err = json.Unmarshal(jsonStr, &parsedResponse)
	results := parsedResponse["trivia_categories"]
	categories := parseCategories(results)

	// Send the response back to the calling server
	w.Header().Set("Context-Type", "application/results-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(categories)
}

// GetQuestions will route a response of trivia questions from the source DB to the requester
// in the form of a Question array
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	// Request the questions from the API
	var client = &http.Client{Timeout: 10*time.Second}
	apiRes, err := client.Get("https://opentdb.com/api.php?amount=10&category=9&difficulty=easy")
	if err != nil {
		log.Fatal(err)
	}
	defer apiRes.Body.Close()

	// Parse API response to get the questions (in the 'results'). It's a bit roundabout to
	// marshal then unmarshal, but couldn't get it to work using other methods
	var temp interface{}
	json.NewDecoder(apiRes.Body).Decode(&temp)
	jsonStr, err := json.Marshal(temp)
	var parsedResponse map[string]interface{}
	err = json.Unmarshal(jsonStr, &parsedResponse)
	results := parsedResponse["results"]
	questions := parseQuestions(results)

	// Send the response back to the calling server
	w.Header().Set("Context-Type", "application/results-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(questions)
}

func parseCategories(results interface{}) []*data.Category {
	var categories []*data.Category

	// Handle the []interface{} returned by response results
	switch result := results.(type) {
	case []interface{}:
		for _, val := range result {
			// Iterate over each map returned
			switch val := val.(type) {
			case map[string]interface {}:
				name := interfaceToString(val["name"])
				id := interfaceToInt(val["id"])
				c := data.NewCategory(id, name)
				categories = append(categories, c)
			}
		}
	}
	return categories
}


// parseQuestions is a helper that returns an array of pointers to questions
// and is meant to be used only in the GetQuestions function
func parseQuestions(results interface{}) []*data.Question {
	var questions []*data.Question

	// Handle the []interface{} returned by response results
	switch result := results.(type) {
	case []interface{}:
		for _, val := range result {

			// Iterate over each map returned
			switch val := val.(type) {
			case map[string]interface {}:
				category := interfaceToString(val["category"])
				answer := interfaceToString(val["correct_answer"])
				difficulty := interfaceToString(val["difficulty"])
				text := interfaceToString(val["question"])
				q := data.NewQuestion(text, answer, category, difficulty)
				questions = append(questions, q)
			default:
				fmt.Println("ERROR: Found type", reflect.TypeOf(result), "but expected map[string]interface {}")
			}
		}
	default:
		// Shouldn't happen
		fmt.Println("ERROR: Found type", reflect.TypeOf(result), "but expected []interface{}")
	}

	return questions
}

// interfaceToString is a quick helper to convert from an ambiguous string to a real one
func interfaceToString(toConvert interface{}) string {
	switch a := toConvert.(type) {
	case string:
		return a
	}
	return ""
}

func interfaceToInt(toConvert interface{}) int {
	switch a := toConvert.(type) {
	case int:
		return a
	case float64:
		return int(a)
	}
	return -1
}