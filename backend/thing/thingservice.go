package thing

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Post handle the post request
func Post(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var t Thing
	for {
		if err := dec.Decode(&t); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	Create(t)
}

// GetAll things and send it
func GetAll(w http.ResponseWriter, r *http.Request) {
	things := ReadAll()

	res, err := json.Marshal(things)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}