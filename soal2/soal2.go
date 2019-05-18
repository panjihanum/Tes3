package main

import (
	"fmt"
	"net/http"
	"sort"
	"html/template"
	"strings"
)


func index(i http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(i, "Apa kabar!")
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        var tmpl = template.Must(template.New("result").ParseFiles("view.html"))

        if err := r.ParseForm(); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

		var consonant = []string{"b", "m", "s", "z"}
		var word string = r.FormValue("word")
		var result = sortingArrayByVowel(word, consonant)

        var data = map[string]string{"word": word, "consonant": strings.Join(consonant, ", "), "result": strings.Join(result, ", ")}

        if err := tmpl.Execute(w, data); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func main() {

	var consonant = []string{"b", "m", "s", "z"}
	var word string = "osama"
	var result = sortingArrayByVowel(word, consonant)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// fmt.Fprintln(i, "Huruf Mati  : ", consonant)
		// fmt.Fprintln(i, "\nKata : ", word)
		// fmt.Fprintln(i, "\nSorting Berdasarkan abjad dan diawali dengan huruf Hidup")
		// fmt.Fprintln(i, "Hasil : ",result)

		var data = map[string]string{
			"word" : word,
            "consonant": strings.Join(consonant, ", "),
            "result":strings.Join(result, ", "),
		}
		
		var t, err = template.ParseFiles("template.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }

        t.Execute(w, data)
	})
	http.HandleFunc("/index", index)
	http.HandleFunc("/process", routeSubmitPost)

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)

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
