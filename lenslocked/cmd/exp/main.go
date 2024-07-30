package main

import (
	"html/template"
	"os"
)

type User struct {
	Name    string
	Widgets []Widget
	Keys    map[string]int
	BoolVal bool
}

type Widget struct {
	Name  string
	Price int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Mark W",
		Widgets: []Widget{
			Widget{"Red Widget", 12},
			Widget{"Green Widget", 11},
			Widget{"Blue Widget", 10},
		},
		Keys: map[string]int{
			"rsc": 3711,
			"r":   2138,
			"gri": 1908,
			"adg": 912,
		},
		BoolVal: false,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
