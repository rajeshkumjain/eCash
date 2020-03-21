package main

import (
	"fmt"
	"strings"
	// "github.com/antzucaro/matchr"
)

func main() {
	var Text string
	var Res string
	Text = "Food, Groceries and Dining"

	TextUpper := strings.ToUpper(Text)
	TextUpper = strings.ReplaceAll(TextUpper, " AND ", "")
	TextUpper = strings.ReplaceAll(TextUpper, ",", "")
	TextUpper = strings.ReplaceAll(TextUpper, "/", "")
	fmt.Println("Input: ", Text, " Output: ", TextUpper)

	for _, ch := range TextUpper {
		if string(ch) == " " || string(ch) == "A" || string(ch) == "E" || string(ch) == "I" || string(ch) == "O" || string(ch) == "U" {
			//	fmt.Println("Nothing")
		} else {
			fmt.Println(string(ch))
			Res = Res + string(ch)
		}
	}

	fmt.Println(Res)
	// fmt.Println(matchr.Soundex(Res))

}
