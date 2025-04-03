# Étape 1 : Compilation de l'application Go
FROM golang:1.24rc1-alpine AS builder


WORKDIR /app

# Copier les fichiers de dépendances Go en premier
COPY backend/go.mod backend/go.sum ./backend/

# Télécharger les dépendances
RUN cd backend && go mod tidy

# Copier le reste du code backend
COPY backend ./backend

# Compiler l'application
RUN cd backend && go build -o /app/backend .

# Étape 2 : Image finale légère
FROM alpine:latest

WORKDIR /app

# Copier le binaire compilé
COPY --from=builder /app/backend .

EXPOSE 8080

ENTRYPOINT ["./backend"]
