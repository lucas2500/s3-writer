package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Unable to load .env file!!")
	}
}

func main() {

	UploadObjetToS3()

	DeleteObjectFromS3()
}

func CreateS3Client() *s3.Client {

	sdkConfig, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatal("Unable to load default configs!!")
	}

	s3Client := s3.NewFromConfig(sdkConfig)

	return s3Client
}

func UploadObjetToS3() {

	S3Client := CreateS3Client()

	// Create file in memory
	file := strings.NewReader("Uploading file!!")

	_, err := S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String("file.txt"),
		Body:   file,
	})

	if err != nil {
		log.Fatal("Unable to upload object!!")
	}

	fmt.Println("Object uploaded successfully!!")
}

func DeleteObjectFromS3() {

	var ObjectIds []types.ObjectIdentifier

	S3Client := CreateS3Client()

	// A loop can be made here to append multiples objects IDs
	ObjectIds = append(ObjectIds, types.ObjectIdentifier{
		Key: aws.String("file.txt"),
	})

	_, err := S3Client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Delete: &types.Delete{Objects: ObjectIds},
	})

	if err != nil {
		log.Fatal("Unable to delete object!!")
	}

	fmt.Println("Object deleted successfully!!")
}
