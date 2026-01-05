package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	employeeList := []Employee{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 40},
		{Name: "Chris", Age: 50},
	}

	jsonList, _ := json.MarshalIndent(employeeList, "", "	")
	fmt.Println(string(jsonList))
}
