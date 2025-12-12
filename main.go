package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const bucketName = "webs3-go"

var s3Client *s3.Client

func main() {
	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("Unable to load AWS config:", err)
	}

	s3Client = s3.NewFromConfig(cfg)

	// routes
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/files", listFilesHandler)

	fmt.Println("Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// UPLOAD TO S3
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(header.Filename),
		Body:   file,
	})

	if err != nil {
		http.Error(w, "Error uploading to S3: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// DOWNLOAD FROM S3
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "Missing filename", http.StatusBadRequest)
		return
	}

	resp, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	})

	if err != nil {
		http.Error(w, "File not found in S3", http.StatusNotFound)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")

	io.Copy(w, resp.Body)
}

// ---------------------- LIST FILES FROM S3 ----------------------
func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		http.Error(w, "Unable to list files in S3", http.StatusInternalServerError)
		return
	}

	for _, obj := range resp.Contents {
		name := *obj.Key
		fmt.Fprintf(w, `<li><a href="/download?filename=%s">%s</a></li>`, name, name)
	}
}
