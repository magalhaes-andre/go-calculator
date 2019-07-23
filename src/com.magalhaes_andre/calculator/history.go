package calculator

var History = []Operation{}

func AddOperation(operation Operation){
	History = append(History, operation)
}