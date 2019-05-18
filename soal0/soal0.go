// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "strconv"

func main() {
	var consonant = []string{"m","z","i","e"}
	fmt.Println("Huruf Mati  : ", consonant)
	var word string = "omama"
	fmt.Println("\nKata : ", word)
	var arrayvowel[] string;
	var arrayconsonant[] string;

	x := 0
	for x < len(word) {
		yo := string([] rune(word)[x])
		if indexOf(yo, consonant) != -1 {
			arrayconsonant = append(arrayconsonant, yo)
		} else {
			arrayvowel = append(arrayvowel, yo)

		}
		x += 1
	}
	arrayvowel = removeDuplicatesUnordered(arrayvowel)
	fmt.Println("\nHuruf Hidup : " , strconv.Itoa(len(arrayvowel)) , arrayvowel , "\nHuruf Mati  : " , strconv.Itoa(len(arrayconsonant)) , arrayconsonant)


}


func removeDuplicatesUnordered(elements[] string)[] string {
	encountered := map[string] bool {}

	for v := range elements {
		encountered[elements[v]] = true
	}

	result := [] string {}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

func indexOf(element string, data[] string)(int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
