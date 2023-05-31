package main

import "fmt"

func main() {

	type Book struct {
		Title  string
		Author string
	}
	person := map[string]string{
		"name":    "John",
		"address": "Jakarta",
		"age":     "20",
	}
	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["age"])

	fmt.Println(len(person))

	fmt.Println("==========")
	// book := make(map[string]string)

	books1 := []Book{{Title: "Belajar Go-Lang", Author: "John"}, {Title: "Belajar Python", Author: "Doe"}}

	ourBooks := []Book{}

	for _, val := range books1 {
		ourBooks = append(ourBooks, val)
	}

	fmt.Println(ourBooks)
	library := map[int]Book{}

	fmt.Println(library)
	for key, val := range ourBooks {
		library[key] = val
	}

	// access map value
	fmt.Println(library)
	fmt.Println(library[1])

	// access struct value
	fmt.Println(library[1].Author)
	// book["title"] = myBook.Title
	// book["author"] = "John"
	// fmt.Println(book)

	// fmt.Println(book)
}
