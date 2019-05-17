// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var dead = []string{"m", "z", "i", "e"}
	fmt.Println("Huruf Mati  : " ,dead)
	var word string = "omama"
	fmt.Println("\nKata : ", word)

	var arraylife []string;
	var arraydead []string;
	
	x:=0
	for x < len(word) {
		yo := string([]rune(word)[x])
		if indexOf(yo, dead) != -1 {
			arraydead = append(arraydead, yo)
		}else{
			arraylife= append(arraylife, yo)

		}
		x+=1
	}
	arraylife = removeDuplicatesUnordered(arraylife)
	fmt.Println("\nHuruf Hidup : ", len(arraylife), arraylife,  "\nHuruf Mati  : ", len(arraydead), arraydead)
	
}


func removeDuplicatesUnordered(elements []string) []string {
    encountered := map[string]bool{}

    for v:= range elements {
        encountered[elements[v]] = true
    }

    result := []string{}
    for key, _ := range encountered {
        result = append(result, key)
    }
    return result
}

func indexOf(element string, data[]string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}