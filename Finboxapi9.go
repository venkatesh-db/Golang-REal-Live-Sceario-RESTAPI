package main

import (
	"log"
	"mime"
	"net/http"
)

func enforceJSONHandler( next http.Handler ) http.Handler {

         log.Print("balaram avatar")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

               log.Print("kantara")

		contentType := r.Header.Get("Content-Type")

		if contentType != "" {

			mt, _, err := mime.ParseMediaType(contentType)

			if err != nil {

				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return

			}

			if mt != "application/json" {

				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return

			}
		}

		       next.ServeHTTP(w, r)

		          log.Print("kantara part1 ends ")
	})
}

func final(w http.ResponseWriter, r *http.Request) {

       log.Print("adheera-sanjay dutt")
        
	w.Write([]byte("OK"))

}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", enforceJSONHandler(finalHandler))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}	