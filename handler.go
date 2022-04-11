package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	mux "github.com/julienschmidt/httprouter"
)

func Get(w http.ResponseWriter, r *http.Request, ps mux.Params) {

	id, err := strconv.Atoi(ps.ByName("key"))

	HandleError(err)

	key := RepositoryGet(id)

	if err := json.NewEncoder(w).Encode(key); err != nil {
		panic(err)
	}
}

func Put(w http.ResponseWriter, r *http.Request, _ mux.Params) {

	var Mymap MyMap

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &Mymap); err != nil {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	RepositoryCreate(Mymap)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
