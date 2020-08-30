package main

import "github.com/vashish1/Treasuro/database"

func main() {
	_,_,db:=database.Createdb()
	var ques database.Question
	ques.Id=1
	ques.Answer="answer"
	database.InsertQuestion(db,ques)
}