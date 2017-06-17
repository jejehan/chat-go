package main

import "github.com/gorilla/websocket"

//client akan menampilkan single chat dari user
type client struct {
	socket *websocket.Conn //websocket menggunakan gorrilla/websocket
	send   chan []byte     // send dalam bentuk chanell untuk mengirimkan data
	room   *room           // room adalah room tempat dimana para user berchat
}

// read
// Method ini membaca pesan yang dikirimkan ke server.
// dan mengirimkan ke room melalui channel forward
// kalo socket mati maka akan keluar dari loop dan socker akan meniutup
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

// write
// write akan secara terus menerus mengecek apakah ada data yang akan dikirimkan
// dari channel send lalu diteruskan dengan ditulis semuaya melalui socker.WriteMessage
// jika menulis socker gagal maka akan keluar dari loop dan socket close
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}

// defer
// defer adalah gini, ketika gw mau exit dari fungsi yang digunakan
// tetapi gw gak tau kapan tepatnya mau exit maka gw menggunakan defer
// disini kita gak butuh untuk menuliskan close dimana mana karena dengan
// defer udah cukup untuk semua.
// ada yang komplen menggunakan defer mengurangi preforma, menjadi lebih lambat
// perlu diingat saat ini yang butukan belum untuk optimasi, menulis kode yang
// bersih dan jelas lebih diperlukan. optimasi bisa kapan saja dilakukan jika
// aplikasi yang dibuat sudah bisa diterima oleh pegguna
