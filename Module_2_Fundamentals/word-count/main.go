package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := `"That's the password: 'PASSWORD 123'!", \tcried the Special Agent.\nSo I fled.`
	re := regexp.MustCompile(`\\[nt]`)
	text = re.ReplaceAllString(text, "")
	fmt.Println(text)
	re1 := regexp.MustCompile(`[^\w']`)
	text = re1.ReplaceAllString(text, " ")
	fmt.Println(text)
}
