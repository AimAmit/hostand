package controller

import (
	"github.com/aimamit/hostand/main-server/api/cache"
	"github.com/aimamit/hostand/main-server/api/remote"
	"log"
	"net/http"
)

type response struct {
	Domain string `json:"domain"`
	Version string `json:"version"`
}


func CreateApp(w http.ResponseWriter, r *http.Request) {
	if e:= r.ParseMultipartForm(32 << 20); e != nil {
		log.Printf("Error reading file: %v", e)
		http.Error(w, "can't read file", http.StatusBadRequest)
		return
	}

	file , _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error reading file: %v", err)
		http.Error(w, "can't read file", http.StatusBadRequest)
		return
	}

	res := response{}
	res.Domain = r.FormValue("domain")
	res.Version = r.FormValue("version")

	log.Println(res.Version, res.Domain)
	isExist, err := cache.Cache.SIsMember(res.Domain, res.Version).Result()
	if err != nil {
		log.Println("Connection issue with cache server: ", err)
		http.Error(w, "Connection issue with cache server", http.StatusInternalServerError)
		return
	}
	if isExist {
		http.Error(w, "App name already exists", http.StatusBadRequest)
		return
	}
	cache.Cache.SetXX(res.Domain, res.Version, 1000)
	cache.Cache.Set(res.Domain, res.Version, 1000)

	err = remote.InterServerCall.FileUpload(file, res.Domain, res.Version)
	err = remote.InterServerCall.ContainerCreate(res.Domain, res.Version)
	if err != nil {
		http.Error(w, "Error creating remote image", http.StatusBadRequest)
		return
	}
}