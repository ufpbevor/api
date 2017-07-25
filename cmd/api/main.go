package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ufpblor/api/core"
)

var (
	apiPort  string
	mongoURL string
)

func init() {
	flag.StringVar(&apiPort, "api-port", os.Getenv("PORT"), "api port")
	flag.StringVar(&mongoURL, "mongo-url", os.Getenv("MONGO_URL"), "mongo url")
}

func main() {
	flag.Parse()

	fmt.Printf("===========================================================\n")
	fmt.Printf("========= UFPB LOR API Version %s | Port %s \n", os.Getenv("HEROKU_SLUG_DESCRIPTION"), apiPort)
	fmt.Printf("========= Mongo URL=%s\n", mongoURL)
	fmt.Printf("===========================================================\n")

	fmt.Printf("INFO: Conecting Mongo... ")
	db := core.GetMongoConnection()
	db.Connect(core.GetDialInfo(mongoURL, false))
	defer db.Close()
	_, dbSession := core.GetDatabase()

	h := &core.Handler{DB: dbSession}
	router := core.Router(h)

	//Verifica se foi definida a porta ou utiliza a porta padrão 8080
	port := map[bool]string{true: apiPort, false: "8080"}[apiPort != ""]

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Println("Erro ListenAndServe err=", err)
		panic(err)
	}

}
