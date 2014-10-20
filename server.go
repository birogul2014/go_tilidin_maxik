package main
import "net/http"


func main() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/tor",tormahiri)
    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}

func tormahiri(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("sadkfljsafd"))
	w.Write([]byte("hi tormahiri"))
}