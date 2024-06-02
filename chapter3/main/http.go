package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	Db sqlx.DB
}

type Value struct {
	V string
}

func (h *Handler) Get(w http.ResponseWriter, req *http.Request) {
	fmt.Println("VALUE HERE")
	var v Value
	v.V = "Taichi"
	// err := h.Db.Get(&v, "SELECT * FROM person WHERE id=$1", req.URL.Query().Get("id"))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	json.NewEncoder(w).Encode(v)
}

func Register(h *Handler) {
	mux := http.NewServeMux()
	//HandlerFuncとして登録するものはfunc(w http.ResponseWriter, req *http.Request)のsignatureを持つもの
	mux.HandleFunc("/value", h.Get)
	//Handler = nilにしたらDefaultのMutexが使われるけど今回は自分のMutexを使う
	//Handlerとして登録するものはServeHTTP(http.ResponseWriter, *http.Request)を実装すればいい
	//基本的にDefaultMutexだろうがそうでなかろうが、内部でやってることはURLパターンに応じたFunctionを持ってくる処理をしてるだけで実装も同じ
	server := http.Server{Addr: ":8080", Handler: mux}
	server.ListenAndServe()
}

// 上二つをまとめてClosureで実装すると以下のようになる
func RegisterWithClosure(h *Handler) {
	mux := http.NewServeMux()
	mux.HandleFunc("/value", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("VALUE HERE")
		var v Value
		v.V = "Taichi"
		// err := h.Db.Get(&v, "SELECT * FROM person WHERE id=$1", req.URL.Query().Get("id"))
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		json.NewEncoder(w).Encode(v)
	})
	server := http.Server{
        Addr:    ":8081",
		Handler: mux,
	}
	server.ListenAndServe()
}
