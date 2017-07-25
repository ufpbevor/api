package core

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"
)

//Handler ...
type Handler struct {
	DB *mgo.Session
}

func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		log.Println("Erro: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("respond:", err)
	}
}

//HandleZen ...
func (h *Handler) HandleZen(w http.ResponseWriter, r *http.Request) {
	data := SuccessMessage{
		Message: "Keep it logically awesome.",
	}
	log.Println("Ping /zen")
	respond(w, r, http.StatusOK, data)
	return
}

//HandleVersion ...
func (h *Handler) HandleVersion(w http.ResponseWriter, r *http.Request) {
	data := VersionMessage{
		AppID:          os.Getenv("HEROKU_APP_ID"),
		AppName:        os.Getenv("HEROKU_APP_NAME"),
		ServerID:       os.Getenv("HEROKU_DYNO_ID"),
		CreatedAt:      os.Getenv("HEROKU_RELEASE_CREATED_AT"),
		ReleaseVersion: os.Getenv("HEROKU_RELEASE_VERSION"),
		Commit:         os.Getenv("HEROKU_SLUG_COMMIT"),
		Description:    os.Getenv("HEROKU_SLUG_DESCRIPTION"),
	}
	log.Println("GET Version ", data)
	respond(w, r, http.StatusOK, data)
	return
}
