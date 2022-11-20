package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Fujito"}
	post2 := Post{Id: 2, Content: "Hoge", Author: "Hoge"}

	store(post1)
	store(post2)

	for _, post := range PostByAuthor["Fujito"] {
		fmt.Println(post)
	}
	for _, post := range PostByAuthor["Hoge"] {
		fmt.Println(post)
	}
}
