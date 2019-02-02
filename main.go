package main

// How to create upload file with GoLang
// !> https://stackoverflow.com/questions/20205796/golang-post-data-using-the-content-type-multipart-form-data
// https://kenyaappexperts.com/blog/how-to-upload-files-in-go-step-by-step/
// !> https://stackoverflow.com/questions/18639929/accessing-uploaded-files-in-golang
// !> https://stackoverflow.com/questions/45541656/golang-send-file-via-post-request

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/upload", UploadFileHandler).Methods("POST")
	log.Println("Service Started")
	log.Println("Listening Port: 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("IndexHandler called")
	response.Write([]byte("Hello, world"))
}

func UploadFileHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("UploadFileHandler called")

	srcFile, srcFileHeader, srcFileError := request.FormFile("file")
	if srcFileError != nil {
		errorDescription := "Was not able to access the uploaded file: " + srcFileError.Error()
		log.Printf(errorDescription)
		response.Write([]byte(errorDescription))
		return
	}
	defer srcFile.Close()

	destFile, destFileError := os.Create("./upload/" + srcFileHeader.Filename)
	if destFileError != nil {
		errorDescription := "Couldn't create file: " + destFileError.Error()
		log.Printf(errorDescription)
		response.Write([]byte(errorDescription))
		return
	}
	defer destFile.Close()

	io.Copy(destFile, srcFile)
	log.Println("Upload file success!!")
	response.Write([]byte("Upload file success!!"))
}
