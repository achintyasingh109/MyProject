package main

import (
	"encoding/json"
	"fmt"
)

//const url = "https://gobyexample.com/"

func main() {

	// fmt.Println("Go web request")
	// response, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("The response is of type %T", response)
	// defer response.Body.Close() //  caller's responsibility to close conection
	// //extremely important

	// databyte, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(databyte))

	type User struct {
		Name  string `json:"name"`
		Email string
	}

	user := User{"Alice", "alice@example.com"}

	// Marshal converts the struct to JSON
	jsonData, _ := json.Marshal(user)

	var user1 User
	json.Unmarshal([]byte(jsonData), &user1)

	// JSON is stored in memory as a []byte
	fmt.Println(string(jsonData)) // Output: {"name":"Alice","email":"alice@example.com"}
	fmt.Println(user1.Email, user.Name)

}
