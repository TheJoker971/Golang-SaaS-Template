package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Définition des structures avec des champs exportés et des tags BSON
type User struct {
	ID       string  `bson:"_id,omitempty"`
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

var Collection,Client = connectDb()

// ConnectDb établit une connexion à la base de données MongoDB et retourne une référence à la collection "users"
func connectDb() (*mongo.Collection, *mongo.Client) {
	uri, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		uri = "mongodb://mongoadmin:secret@localhost:27017/"
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("saas").Collection("users")

	// ⚠️ Créer les index uniques
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "pseudo", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err = collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Fatal("Erreur création index uniques : ", err)
	}

	return collection, client
}


// Register insère un nouvel utilisateur dans la collection
func register(collection *mongo.Collection, user Register) error {
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
		log.Printf("Erreur lors de l'insertion de l'utilisateur : %v", err)
		return err
	}

	fmt.Println("Utilisateur inséré dans la base de données")
	return nil
}


func getUserForLog(login Login) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			{"pseudo": login.Login},
			{"email": login.Login},
		},
	}

	var user User
	err := Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, fmt.Errorf("utilisateur introuvable")
		}
		return user, err
	}

	// Comparaison sécurisée du mot de passe
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return user, fmt.Errorf("mot de passe incorrect")
	}

	fmt.Printf("Bienvenue à %s !\n", user.Pseudo)
	return user, nil
}


func getAllUsers(collection *mongo.Collection) <-chan User {
	out := make(chan User)

	go func() {
		defer close(out)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		filter := bson.M{"admin": false}
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			log.Panic("❌ Erreur Find:", err)
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var user User
			if err := cursor.Decode(&user); err != nil {
				log.Panic("⚠️ Erreur Decode:", err)
				continue
			}
			out <- user
		}

		if err := cursor.Err(); err != nil {
			log.Println("❌ Erreur Cursor:", err)
		}
	}()

	return out
}