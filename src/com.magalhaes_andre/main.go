package main

import (
	"net/http"
	"./calculator"
	"strconv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)

func welcomeHandler(response http.ResponseWriter, request *http.Request){
	fmt.Fprintf(response, "Welcome to your web calculator!\n\nFor usage you have two main paths:\n- /calc/{operationGoesHere}/{firstNumberHere}/{secondNumberHere}\n- /calc/history\n\nOperations Accepted:\n- sum\n- sub\n- divide\n- multiply")
}

func calcHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	operation := vars["operation"]
	first := vars["firstNumber"]
	second := vars["secondNumber"]
	
	firstNumber, firstErr := strconv.ParseFloat(first, 64)
	if firstErr != nil{
		http.Error(response, firstErr.Error(), http.StatusBadRequest)
	}
	
	secondNumber, secondErr := strconv.ParseFloat(second, 64)
	if secondErr != nil{
		http.Error(response, secondErr.Error(), http.StatusBadRequest)
	}

	if firstErr == nil && secondErr == nil{
		calcResult := calculator.Calculate(calculator.NewOperation(firstNumber, secondNumber, operation))
		calc, calcErr := json.Marshal(calcResult)
		if calcErr != nil {
			http.Error(response, calcErr.Error(), http.StatusInternalServerError)
		}
		response.Header().Set("Content-Type", "application/json")
		response.Write(calc)
	}
}

func historyHandler(response http.ResponseWriter, request *http.Request){
	history, err := json.Marshal(calculator.History)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(history)
}


var serverPort string = "8080"
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", welcomeHandler)
	router.HandleFunc("/calc/{operation}/{firstNumber}/{secondNumber}", calcHandler)
	router.HandleFunc("/calc/history", historyHandler)
	fmt.Printf("Your calculator is up on localhost:"+serverPort)
	http.ListenAndServe(":"+serverPort, router)
	
}
