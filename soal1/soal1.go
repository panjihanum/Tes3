// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "sort"

func main() {
	var consonant = []string{"m", "s"}
	fmt.Println("Huruf Mati  : ", consonant)
	var word string = "osama"
	fmt.Println("\nKata : ", word)

	var result = sortingArrayByVowel(word, consonant)
	fmt.Println("\nSorting Berdasarkan abjad dan diawali dengan huruf Hidup")
	fmt.Println("Hasil : ",result)
}

func sortingArrayByVowel(word string, consonant []string) []string {
	var arrayvowel []string
	var arrayconsonant []string

	x := 0
	for x < len(word) {
		yo := string([]rune(word)[x])
		if indexOf(yo, consonant) != -1 {
			arrayconsonant = append(arrayconsonant, yo)
		} else {
			arrayvowel = append(arrayvowel, yo)
		}
		x += 1
	}
	
	sort.Strings(arrayvowel)
	sort.Strings(arrayconsonant)
	
	var result = concat(arrayvowel, arrayconsonant);
	return result
}

func concat(array1 []string, array2 []string) []string {
	
	result := []string{}
	result = array1
	for key, _ := range array2{
		result = append(result, array2[key])
	}
	
	return result
}


func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
