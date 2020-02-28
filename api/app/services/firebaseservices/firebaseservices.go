package firebaseservices

import (
	"log"
	"os"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func init() {
	firebaseJSON := os.Getenv("FIREBASE_JSON")
	opt := option.WithCredentialsFile(firebaseJSON)

	var err error
	FirebaseApp, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("error initializing firebases app :", err.Error())
	}

}
