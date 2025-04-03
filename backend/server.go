package main 

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)


func registerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}

		var user Register
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "JSON invalide : "+err.Error(), http.StatusBadRequest)
			return
		}

		// Hashage du mot de passe APRÈS décodage
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Erreur hashage mot de passe", http.StatusInternalServerError)
			return
		}
		user.Password = string(hashedPassword)

		err = register(Collection, user)
		if err != nil {
			http.Error(w, "Erreur inscription : "+err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}


func loginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}

		var login Login
		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			http.Error(w, "JSON invalide : "+err.Error(), http.StatusBadRequest)
			return
		}

		user, err := getUserForLog(login)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur d'authentification : %v", err), http.StatusUnauthorized)
			return
		}

		user.Password = ""
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func enableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // autoriser toutes les origines, ou mets "http://localhost:5173"
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}



// --- SERVER ---

func listenServer() {
	http.HandleFunc("/register", registerHandler())
	http.HandleFunc("/login", loginHandler())

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("🚀 Serveur en écoute sur http://localhost%s\n", addr)

	// Utilisation du middleware CORS ici
	if err := http.ListenAndServe(addr, enableCors(http.DefaultServeMux)); err != nil {
		log.Fatalf("❌ Erreur serveur HTTP : %v", err)
	}
}
