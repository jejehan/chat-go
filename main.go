//golang memiliki keunikan dalam menjalankan programnya
//jika kita menginginkan menjalankan program secara langsung dengan
//mengetikan run main.go secara langsung dari command line kita
//bisa gunakan "package main"
//dan jika menginginkan package ini menjadi reusable maka harus diganti
//menjadi "package chat"
package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// Memisahkan View dari Logic menggunakan Template
// Seperti website pada umumnya, biasanya kita akan memisahkan
// Bisnis logic dengan view agar mudah dimaintain dan pekerjaan
// menjadi lebih terfokuskan.
// Kali ini kita akan coba membuat template dimana nantinya,
// View akan menampilkan data nama yang berasal dari layer logic
// dan menempatkanya pada template sebagai layer persentasi / view
//
// disini kita akan menggunakan beberapa package lain
// golang memiliki standart library html/template dimana package ini
// dapat membaca tag-tag html yang dapat dibaca oleh browser.
// package ini juga turunan dari text/template dimana text/template
// hanya dapat digunakan untuk merender text saja bukan tag html
//
// disini kita juga menggunakan package sync.Once, package ini akan
// memastikan bahwa dia akan tampil dalam satu action
//
// tambahan package selanjunta adalah paht/filepath
// tugas package ini dapat mencari file yang dinginkan untuk diparse
// oleh http/template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServerHTTP ini akan menangangi http request
// ini adalah method turunan dari struct templateHandler
// ServeHTTP method yang harus ada karena nantinya akan digunakan
// pada http.Handle, http.Handle sendiri memiliki parameter
// pattern dengan tipe String dan handle dengan tipe Handler
// tipe Handler adalah sebuah interface yang membutukan method ServeHTTP
//
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {

	var port = ":9090"

	// http Handle
	// mengganti HandleFunc menjadi Handle karena kebutuhan parameter, hanya menggunakan struct
	// struct yang harus memiliki method ServeHTTP
	http.Handle("/", &templateHandler{filename: "chat.html"})

	//start webserver
	log.Println("ListenAndServe: localhost:", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println("Failed ListenAndServe:", err)
	}

}
