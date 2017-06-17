package main

type room struct {
	forward chan []byte // channel yang menahan data yang nantinya akan dikirimkan ke client
}
