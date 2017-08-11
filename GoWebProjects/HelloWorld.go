package main

import (
"fmt"
"net/http"
)

func main() {

	http.HandleFunc("/bny", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ricardo Quaresma", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}


/*
Öncelikle request handler yazıyoruz  : func (w http.ResponseWriter, r *http.Request)

Bu fonksiyon 2 parametre almaktadır...

1- http.ResponseWriter : response yazdığımız yer.
2- http.request : http request all info.

HTTP Server 'a request handler yazmak;

http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you've requested: %s", r.URL.Path)
})


The request handler tek başına dışarından bir http bağlantısını kabul edemez.
HTTP Server request handler üzerinden bağlantıları dinler.
Aşağıdaki kod Go dilinin default HTTP Server'ı başlatır ve 80 portundan dinler.

http.ListenAndServe(":80", nil)

*/