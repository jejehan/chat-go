//golang memiliki keunikan dalam menjalankan programnya
//jika kita menginginkan menjalankan program secara langsung dengan
//mengetikan run main.go secara langsung dari command line kita
//bisa gunakan "package main"
//dan jika menginginkan package ini menjadi reusable maka harus diganti
//menjadi "package chat"
package main

import (
	"log"
	"net/http"
)

func main() {

	var port = ":9090"

	//http HandleFunc
	//jika user mengakses path patten / fungsi dibawah akan teresekusi
	//sehingga nantinya akan tampil data yang ditulis seperti dibawah ini
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title> Go Chat </title>
				</head>
				<body>
					Sedang dalam pengembangan...
				</body>
			</html>
		`))
	})

	//start webserver
	log.Println("ListenAndServe: localhost:", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println("Failed ListenAndServe:", err)
	}

}
