package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/speps/go-hashids"
	"log"
)

//in memory object to store the urls
var ShortenedUrls = make(map[string]MyUrl)

type MyUrl struct {
	ID       string `json:"id,omitempty"`
	LongUrl  string `json:"long_url,omitempty"`
	ShortUrl string `json:"short_url,omitempty"`
}

//Retrieve lists the urls stored in memory
func Retrieve(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ShortenedUrls)
}

//Create creates a shorten url for the requested long url
func Create(w http.ResponseWriter, req *http.Request) {
	var (
		url          MyUrl
		shortenedUrl MyUrl
	)

	err := json.NewDecoder(req.Body).Decode(&url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error while decoding the data", err)
		json.NewEncoder(w).Encode("error while decoding the data")
		return
	}

	//if the requested url already contains short url then return
	_, ok := ShortenedUrls[url.LongUrl]
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("this url is already exist")
		return
	}

	//generate the shorten url
	uniqueID, err := getUniqueShortID()
	if err != nil {
		log.Println("error while generating the unique id", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("error while shortening the long url")
		return
	}

	shortenedUrl.ID = uniqueID
	shortenedUrl.ShortUrl = "http://localhost:8090/" + shortenedUrl.ID
	ShortenedUrls[url.LongUrl] = shortenedUrl

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shortenedUrl)
}

//Redirect will redirect you the actual page when you access the shortened url
func Redirect(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for k, v := range ShortenedUrls {
		if v.ID == params["endpoint"] {
			http.Redirect(w, req, k, http.StatusMovedPermanently)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("requested url does not exist")
	return

}

//Delete deletes the the requested url
func Delete(w http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()

	_, ok := ShortenedUrls[query["endpoint"][0]]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("requested url does not exist")
		return
	}

	delete(ShortenedUrls, query["endpoint"][0])
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("url deleted successfully")
	return

}

//Update updates the requested long url with the new shortened url
func Update(w http.ResponseWriter, req *http.Request) {
	var url MyUrl
	_ = json.NewDecoder(req.Body).Decode(&url)

	//check if the requested url exist or not
	_, ok := ShortenedUrls[url.LongUrl]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("requested url does not exist")
		return
	}

	//if the requested url is exist then generate a new short url and update
	uniqueID, err := getUniqueShortID()
	if err != nil {
		log.Println("error while generating the unique id", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("error while shortening the long url")
		return
	}

	url.ID = uniqueID
	url.ShortUrl = "http://localhost:8090/" + url.ID
	ShortenedUrls[url.LongUrl] = url
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(url)

}

//getUniqueShortID returns the unique string for shortening the url
func getUniqueShortID() (string, error) {
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	return h.Encode([]int{int(now.Unix())})
}
