package main

import (
	"fmt"
	"strings"
)

type asjad struct {
	Name string
}

func main() {
	//t := "2021-06-21T22:06:56+00:00"
	//t2, err := time.Parse(time.RFC3339, t)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(t2.Year())


	data := "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2MjQ0ODgzNTksIklzc3VlZEF0IjoxNjI0MzE1NTU5LCJ1c2VySUQiOjV9.cZvQxAK3cJhZ4H6ZI-kSOqIiw2wgkqQE4AZ0LFk4ZQA"
	token := strings.Split(data, "bearer")
	fmt.Println(strings.TrimSpace(token[1]))
}
