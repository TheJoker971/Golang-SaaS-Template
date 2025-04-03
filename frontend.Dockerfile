# Étape 1 : Build
FROM node:22.0.0-alpine AS builder

WORKDIR /app

COPY frontend/my-front/package*.json ./
RUN npm install

COPY frontend/my-front ./
RUN npm run build

# Étape 2 : Serveur statique (ex: avec nginx ou un conteneur léger)
FROM nginx:alpine AS runner

# Copier les fichiers construits dans le dossier par défaut de nginx
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]
