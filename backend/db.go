package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Définition des structures avec des champs exportés et des tags BSON
type User struct {
	ID       int64  `bson:"_id,omitempty"`
	Pseudo   string `bson:"pseudo"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Admin    bool   `bson:"admin"`
}

type Register struct {
	Pseudo   string
	Email    string
	Password string
	Admin    bool
}

type Login struct {
	Login    string
	Password string
}

// ConnectDb établit une connexion à la base de données MongoDB et retourne une référence à la collection "users"
func ConnectDb() (*mongo.Collection, *mongo.Client) {
	uri, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		log.Fatal("Database URL not found")
	}

	// Crée une nouvelle instance de client
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Établit la connexion avec un contexte de timeout
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Vérifie la connexion
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("saas").Collection("users")
	return collection, client
}

// Register insère un nouvel utilisateur dans la collection
func Register(collection *mongo.Collection, user Register) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newUser := User{
		Pseudo:   user.Pseudo,
		Email:    user.Email,
		Password: user.Password,
		Admin:    user.Admin,
	}

	_, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User inserted in database")
}

// GetUserForLog recherche un utilisateur par pseudo ou email et mot de passe
func GetUserForLog(collection *mongo.Collection, login Login) User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"$and": []bson.M{
			{"$or": []bson.M{
				{"pseudo": login.Login},
				{"email": login.Login},
			}},
			{"password": login.Password},
		},
	}

	var user User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Fatal("User not found")
		}
		log.Fatal(err)
	}

	return user
}

func GetAllUsers(collection *mongo.Collection) <-chan []User {
	out := make(chan []User)

	go func() {
		defer close(out)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		filter := bson.M{"admin": false}
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			log.Println("❌ Erreur Find:", err)
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var user User
			if err := cursor.Decode(&user); err != nil {
				log.Println("⚠️ Erreur Decode:", err)
				continue
			}
			out <- user // envoie le user directement dans le channel
		}

		if err := cursor.Err(); err != nil {
			log.Println("❌ Erreur Cursor:", err)
		}
	}()

	return out
}