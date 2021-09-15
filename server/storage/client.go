package storage

import (
	"context"
	"io"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
	"github.com/youngeek-0410/mottake/server/config"
	"google.golang.org/api/option"
)

var Client *storage.Client

func Init() {
	var err error
	opt := option.WithCredentialsFile(config.Config.FirebaseSecret)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println(err)
	}
	Client, err = app.Storage(context.Background())
	if err != nil {
		log.Println(err)
	}
}

func Save(filePath string, file io.Reader) error {
	bucket, err := Client.Bucket(config.Config.FirebaseBucket)
	if err != nil {
		return err
	}
	writer := bucket.Object(filePath).NewWriter(context.Background())
	defer writer.Close()
	_, err = io.Copy(writer, file)
	return err
}
