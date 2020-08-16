package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rakyll/statik/fs"
	"github.com/ryomak/qrme/web/src"
	_ "github.com/ryomak/qrme/web/statik"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	sfs, err := statikFS.Open("/profile.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	defer sfs.Close()
	profileTmplByte, err := ioutil.ReadAll(sfs)
	if err != nil {
		log.Fatal(err)
	}
	sfs, err = statikFS.Open("/notfound.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	defer sfs.Close()
	notfoundTmplByte, err := ioutil.ReadAll(sfs)
	if err != nil {
		log.Fatal(err)
	}
	profileTmpl := template.Must(template.New("profile.html").Parse(string(profileTmplByte)))
	notfoundTmpl := template.Must(template.New("notfound.html").Parse(string(notfoundTmplByte)))

	router := chi.NewRouter()
	router.Route("/profile", func(r chi.Router) {
		r.Get("/{uid}", func(w http.ResponseWriter, r *http.Request) {
			uid := chi.URLParam(r, "uid")
			img := src.NewProfileImage(uid, "")
			bytes, err := img.Download(r.Context())
			if err != nil {
				log.Println(err)
				notfoundTmpl.Execute(w, map[string]string{
					"ID": uid,
				})
				return
			}
			profileTmpl.Execute(w, map[string]string{
				"Image": base64.StdEncoding.EncodeToString(bytes),
			})
			w.WriteHeader(http.StatusOK)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var params struct {
				UID   string `json:"uid"`
				Image string `json:"image"`
			}
			if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
				ErrorJSON(w, "bad image", http.StatusBadRequest)
				return
			}
			if params.UID == "" {
				ErrorJSON(w, "id cannot be empty", http.StatusBadRequest)
				return
			}
			img := src.NewProfileImage(params.UID, params.Image)
			if err := img.Upload(r.Context()); err != nil {
				ErrorJSON(w, "cannot upload image", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		})
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			notfoundTmpl.Execute(w, map[string]string{
				"ID": "QRME",
			})
			w.WriteHeader(http.StatusOK)
			return
		})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func ErrorJSON(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]string{
		"error": err,
	})
}
