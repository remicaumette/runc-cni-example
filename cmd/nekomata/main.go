package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	log.Printf("hello world\n")
}
