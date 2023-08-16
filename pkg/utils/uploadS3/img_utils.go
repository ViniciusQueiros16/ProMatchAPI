package uploadS3

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"github.com/promatch/structs"
)

func imageUpload(bodyRequest *structs.ImageRequestBody, data []byte) structs.ImageUploadResponse {
	tmpFileName := fmt.Sprintf(`/tmp/%s`, bodyRequest.FileName)

	fileErr := ioutil.WriteFile(tmpFileName, []byte(data), 0644)
	if fileErr != nil {
		log.Fatalf("Failed to save file : %s %v\n", bodyRequest.FileName, fileErr)
	}
	res := UploadImage(tmpFileName)
	os.Remove(tmpFileName)

	return res
}

// IMAGE UPLOAD
func UploadImage(fileName string) structs.ImageUploadResponse {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Failed to open file : %s %v\n", fileName, err)
	}

	// Upload to S3
	name := filepath.Base(fileName)
	return uploadToS3Bucket(file, name)
}