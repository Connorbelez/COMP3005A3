package main

import (
	"fmt"

	"github.com/Connorbelez/COMP3005A3/kisley"
)

func main() {
	db, _ := kisley.NewDBConn()
	query1, err := db.GetAllStudents()
	if err != nil {
		fmt.Println("ERROR : ", err)
	}

	for i, s := range query1 {
		fmt.Println(i, " : ", s)
	}

	db.UpdateStudentEmail(1, "updatedEmail@gmail.com")
	if err != nil {
		fmt.Println("ERROR : ", err)
	}

	query1, _ = db.GetAllStudents()
	for i, s := range query1 {
		fmt.Println(i, " : ", s)
	}
}
