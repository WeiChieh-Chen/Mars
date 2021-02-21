package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Yimismi/sql2go"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/cors"
)

type (
	Input struct {
		SQL string
	}

	Output struct {
		Code   int
		Result string
	}
)

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080","https://sql2go.ricestdiotech.com"},
		AllowedMethods: []string{"POST"},
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/transfer", sqlToStruct)

	handler := c.Handler(mux)

	log.Fatal(http.ListenAndServe("127.0.0.1:36988", handler))
}

func sqlToStruct(w http.ResponseWriter, r *http.Request) {
	defer recovery(w)
	defer r.Body.Close()

	sql := getContent(r.Body)

	args := sql2go.NewConvertArgs().SetGenJson(true).SetGenXorm(true)

	code, err := sql2go.FromSql(sql, args)
	if err != nil {
		panic(err)
	}

	o := Output{Code: 999, Result: string(code)}

	w.WriteHeader(200)
	w.Write(jsonMarshal(o))
}

func getContent(input io.Reader) string {
	b, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}

	var i Input

	jsonUnmarshal(b, &i)

	return i.SQL
}

func jsonMarshal(input interface{}) []byte {
	if b, err := jsoniter.Marshal(input); err != nil {
		panic(err)
	} else {
		return b
	}
}

func jsonUnmarshal(input []byte, bind interface{}) {
	if err :=  jsoniter.Unmarshal(input, bind); err != nil {
		panic(err)
	}
}

func recovery(w http.ResponseWriter) {
	if err := recover(); err != nil {
		w.WriteHeader(400)
		str := err.(error).Error()
		o := Output{Code: 444, Result: str}
		w.Write(jsonMarshal(o))
	}
}
