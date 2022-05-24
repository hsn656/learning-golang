package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["cs"] = "csharp"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println(languages)
	fmt.Println(languages["js"])
	delete(languages, "rb")
	fmt.Println(languages)

	for k, v := range languages {
		println(k, ":", v)
	}
}
