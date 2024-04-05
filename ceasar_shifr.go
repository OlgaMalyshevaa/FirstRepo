package main

import "fmt"
import "unicode"

func main() {
	seq := "ABCDEFGHIJKLMNOPQRSTUVWXYZ&*()abcdefghijklmnopqrstuvwxyz"
	num := 1
	res := Cipher(seq, num)
	fmt.Println(res)
}

func Cipher(s string, shift int) string {
	var result []rune
	for _, char := range s {
		if unicode.IsLetter(char) {
			var newChar = char + rune(shift)
			if newChar > 'z'  {
				newChar -= 26
			}
			if newChar > 'Z' && newChar < 'a' {
				newChar -= 26
			}
			
			result = append(result, newChar)	
		} else {
			result = append(result, char)
		}
		
	}
	return string(result)
}