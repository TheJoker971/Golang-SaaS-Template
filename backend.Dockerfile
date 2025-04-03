# Étape 1 : Compilation de l'application Go
FROM golang:1.22-alpine AS builder

# Définir le répertoire de travail
WORKDIR /app

# Copier le code source du backend uniquement
COPY backend ./backend

# Installer les dépendances
RUN go mod download

# Compiler l'exécutable Go à partir du dossier backend
RUN go build -o backend ./backend

# Étape 2 : Container minimal pour exécution
FROM alpine:latest

WORKDIR /app

# Copier uniquement l'exécutable du builder
COPY --from=builder /app/backend .

# Exposer le port utilisé par ton backend
EXPOSE 8080

# Exécuter l'application Go
ENTRYPOINT ["./backend"]
