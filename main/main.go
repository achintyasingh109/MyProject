package main

import (
	// add "MYGOLANG/Addition"
	// sub "MYGOLANG/Addition"
	// "errors"
	"fmt"

	// "math"
	// "sort"
	// "strings"
	"io"
	"os"
)

func main() {

	file, err := os.Create("./abc.txt")
	content := "Hi this is my file"

	if err != nil {
		panic(err)
	}
	len, err := io.WriteString(file, content)
	if err != nil {

		panic(err)
	}
	fmt.Println("the length is:", len)
	defer file.Close()

	readFile(file.Name())

	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Result:", result)

	// var x string = "hello"
	// //var s string
	// s := "hi"
	// y := 10
	// fmt.Printf("%v , %v , y is of type %T and value is, %v\n", s, x, y, y)
	// var arr = [3]int{1, 2, 3}

	// fmt.Println(arr, len(arr))

	// arr[1] = 100
	// //arr=append(arr,19)   wont work
	// fmt.Println(arr)
	// var arr2 = []int{1, 2, 3}
	// arr2 = append(arr2, 12)

	// var str string = "Hello Friends of Go"
	// fmt.Println(strings.Contains(str, "Go"))
	// fmt.Println(strings.ReplaceAll(str, "Go", "GoLang"))
	// fmt.Println(str[:5])
	// fmt.Println(strings.Split(str, " "))
	// arr3 := []int{50, 32, 45, 34, 65}
	// sort.Ints(arr3)
	// fmt.Println(arr3)

	// fmt.Println(arr2)

	// // i := 0
	// // for i <= 5 {
	// // 	fmt.Println(i)
	// // 	i++
	// // }

	//names := []string{"one", "two", "three", "four", "five"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("the value at %v is %v \n", index, value)
	// }

	// age := 46 // Boolean checking is done like this
	// fmt.Println(age <= 50)
	// fmt.Println(age <= 55)
	// fmt.Println(age != 50)

	// if age < 30 { // Conditionals are done like this
	// 	fmt.Println("Age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println(("Age is greater than 30 and less than 45"))
	// } else {
	// 	fmt.Println("Age is greater than 45")
	// }

	//sayGreeting("Hello")
	// sayBye("bye")
	// add.Add(2, 3)
	// sub.Sub(4, 5)
	// sayArray(names)
	// cycleNames(names, sayGreeting)
	// fmt.Println(circleArea(10.55))

	// fmt.Println(getInitials("John Doe"))
	// fm1, fm2 := getInitials(("sam curran"))
	// fmt.Println(fm1, fm2)

	// name := "John"
	// m := &name
	// updateName(m)
	// fmt.Println(name)

	// menu := map[string]float64{
	// 	//key : value
	// 	"xyz": 12.33,
	// 	"abc": 45,
	// }
	// fmt.Println(menu)
	// fmt.Println(menu["abc"])

	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	// }

	// menu["xyz"] = 21

	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	//}

	// defer fmt.Println("World")
	// fmt.Println("Hello")
	// defer fmt.Println("defer 2")
	openFile(file.Name())

}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic((err))
	}

	fmt.Println(string(databyte))
	// file, err := os.Open(filename)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// defer file.Close()
	// buf := make([]byte, 512)
	// for {
	// 	n, err := file.Read(buf)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		fmt.Println("Error reading file:", err)
	// 		return
	// 	}
	// 	fmt.Print(string(buf[:n]))
	// }
}

func openFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	fmt.Println("File opened successfully", file)

}

// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("division by zero")
// 	}
// 	return a / b, nil
// }

// func updateName(n *string) {
// 	*n = "rahul"
// }

//	func sayArray(names []string) {
//		fmt.Printf("%v \n", names)
//	}
// func sayGreeting(n string) {
// 	fmt.Printf(" %v \n", n)
// }

// func sayBye(n string) {
// 	fmt.Printf("%v \n", n)
// }

// func cycleNames(names []string, f func(string)) {
// 	for _, v := range names {
// 		f(v)

// 	}
// }
// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }

// func getInitials(n string) (string, string) {

// 	s := strings.ToUpper(n)
// 	str := strings.Split(s, " ")
// 	var initials []string
// 	for _, v := range str {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}
// 	return initials[0], " "

// }
