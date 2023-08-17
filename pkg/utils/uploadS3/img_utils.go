package uploadS3

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/promatch/structs"
)

func ImageUpload(fileName string, data string) structs.ImageUploadResponse {

	decoded, _ := base64.StdEncoding.DecodeString(data)

	tmpFileName := fmt.Sprintf(`/tmp/%s`, fileName)

	fileErr := ioutil.WriteFile(tmpFileName, []byte(decoded), 0644)
	if fileErr != nil {
		log.Fatalf("Failed to save file : %s %v\n", fileName, fileErr)
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
