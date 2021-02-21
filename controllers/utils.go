package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func MultipleFileUpload(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
	fhs := req.MultipartForm.File["files"]
	var urls []string
	for _, fh := range fhs {
		f, err := fh.Open()
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(Response{
				Data:      nil,
				IsSuccess: true,
				Errors:    []string{"Файл поврежден"},
			})
			return
		}
		// f is one of the files

		data, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile(fmt.Sprintf("./public/uploads/%s", fh.Filename), data, 0644)
		if err != nil {
			fmt.Println(err)
		}
		urls = append(urls, req.URL.Host+"/public/uploads/"+fh.Filename)
	}
	json.NewEncoder(w).Encode(Response{
		Data:      urls,
		IsSuccess: true,
		Errors:    nil,
	})
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `file`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, h, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		json.NewEncoder(w).Encode(Response{
			Data:      nil,
			IsSuccess: true,
			Errors:    []string{"Файл поврежден"},
		})
		return
	}
	defer file.Close()
	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	var data []byte
	_, _ = file.Read(data)
	err = ioutil.WriteFile(fmt.Sprintf("/public/uploads/%s", h.Filename), data, 0644)
	if err != nil {
		fmt.Println(err)
	}
	url := r.URL.Host + r.URL.RequestURI() + h.Filename
	json.NewEncoder(w).Encode(Response{
		Data:      url,
		IsSuccess: true,
		Errors:    nil,
	})
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {

}

func GetUploadedImages(w http.ResponseWriter, r *http.Request) {
	fileName := mux.Vars(r)["name"]
	img, err := os.Open(fmt.Sprintf("./public/uploads/%s", fileName))
	if err != nil {
		log.Fatal(err)
	}
	var bytes []byte
	_, err = img.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	//data, err := ioutil.ReadFile(fmt.Sprintf("public/uploads/%s", "2.jpg"))
	//log.Println(data)
	//defer img.Close()
	w.Header().Set("Content-Type", "text")
	io.Copy(w, img)
}

type Response struct {
	Data      interface{} `json:"data"`
	IsSuccess bool        `json:"is_success"`
	Errors    interface{} `json:"errors"`
}
