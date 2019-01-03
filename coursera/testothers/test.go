package main

import "fmt"

func main() {
	blogArticleViews := map[string]int{
		"unix":         0,
		"python":       1,
		"go":           2,
		"javascript":   3,
		"testing":      4,
		"philosophy":   5,
		"startups":     6,
		"productivity": 7,
		"hn":           8,
		"reddit":       9,
		"C++":          10,
	}
	for key, views := range blogArticleViews {
		fmt.Println("There are", views, "views for", key)
	}
}
