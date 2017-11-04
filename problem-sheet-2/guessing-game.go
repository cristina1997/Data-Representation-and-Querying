package main

import (
	//"fmt"
	//"os"
	"net/http"
	"text/template"
	"math/rand"
	"time"
)

type tempData struct {
	GuessNo int
	MaxNumber int
	//UserGuess int
}

//Code adapted from: https://gobyexample.com/reading-files
func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

func templateHandler(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("go-static/guess.html")	
	const maxNum = 20

	//Random number generated according to the current time
	rand.Seed(time.Now().UnixNano())
	unknown := rand.Intn(maxNum)+1

	tmpl.Execute(w, tempData{GuessNo: unknown, MaxNumber: maxNum})

	checkError(err)

}

func indexHTML(w http.ResponseWriter, r * http.Request){
	http.ServeFile(w, r, "go-static/index.html")
}

func main() {
	//Code adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	//Code adapted from: 
						//https://github.com/Masterminds/go-fileserver/blob/master/fs.go	
						//https://github.com/Masterminds/go-in-practice/blob/master/chapter7/servefile.go
						//https://stackoverflow.com/questions/15407719/in-gos-http-package-how-do-i-get-the-query-string-on-a-post-request
						
	http.HandleFunc("/", indexHTML)
	http.HandleFunc("/guess/", templateHandler)

	http.ListenAndServe(":9999", nil)
}
