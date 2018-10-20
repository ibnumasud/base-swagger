package upload_s3

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	middleware "github.com/go-openapi/runtime/middleware"
	minio "github.com/minio/minio-go"
)

//Upload buat ke s3
func Upload(params UploadS3Params) middleware.Responder {

	var Buf bytes.Buffer
	// in your case file would be fileupload
	file, header, err := params.HTTPRequest.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])

	// Copy the file data to my buffer
	// io.Copy(&Buf, file)
	// do something with the contents...
	// I normally have a struct defined and unmarshal into a struct, but this will
	// work as an example
	// contents := Buf.String()
	// fmt.Println(contents)
	// I reset the buffer in case I want to use it again
	// reduces memory allocations in more intense projects
	// Buf.Reset()
	// handle err

	endpoint := os.Getenv("ENDPOINT")
	accessKeyID := os.Getenv("ACCESSKEYID")
	secretAccessKey := os.Getenv("SECRETACCESSKEY")
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("lolos create minio client")

	n, err := minioClient.PutObject("ibnumasud", header.Filename, file, header.Size, minio.PutObjectOptions{ContentType: "application/pdf"})
	if err != nil {
		log.Println("push ke minio")
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", header.Filename, n)

	return NewUploadS3OK()
}
