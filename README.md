
# 🧩 Golang SaaS Starter — Auth API + Frontend

Projet minimaliste d’authentification SaaS avec :

- ✅ Un backend en **Go**
- ✅ Une base de données **MongoDB**
- ✅ Un frontend **React (Vite)**

---

## 🗂️ Structure du projet

```
Golang-SaaS-Template/
├── backend/              # Code source Go
│   ├── main.go
│   ├── server.go         # Routes HTTP
│   ├── db.go             # Connexion MongoDB
│   └── go.mod, go.sum
├── frontend/             # Frontend React (Vite)
│   ├── src/
│   ├── index.html
│   └── ...
├── docker-compose.yml
├── backend.Dockerfile
├── frontend.Dockerfile
└── README.md
```

---

## ⚙️ Démarrage rapide

### 1. 🐳 Lancer le projet

Assure-toi d’avoir **Docker** et **Docker Compose** installés.

```bash
docker-compose up --build
```

- Backend disponible sur : [http://localhost:8080](http://localhost:8080)
- Frontend disponible sur : [http://localhost](http://localhost)

---

## 🔒 API Backend

Développée en **Go**, avec MongoDB pour la persistence.

### POST `/register`

**Objectif** : Créer un utilisateur.

#### ✅ Body JSON :
```json
{
  "pseudo": "john_doe",
  "email": "john@example.com",
  "password": "motdepasse",
  "admin": false
}
```

#### ✅ Réponse :
- `201 Created` si succès
- `400 Bad Request` si email/pseudo déjà utilisé

---

### POST `/login`

**Objectif** : Authentifier un utilisateur.

#### ✅ Body JSON :
```json
{
  "login": "john_doe",  // pseudo ou email
  "password": "motdepasse"
}
```

#### ✅ Réponse :
```json
{
  "id": "xxxx",
  "pseudo": "john_doe",
  "email": "john@example.com",
  "admin": false
}
```

- `401 Unauthorized` si les identifiants sont incorrects

---

## 🌐 Frontend React

Développé avec [Vite](https://vitejs.dev), `react-router-dom`, `lucide-react`.

### Pages incluses :

| URL               | Description           |
|-------------------|-----------------------|
| `/`               | Page d’accueil        |
| `/login`          | Connexion utilisateur |
| `/register`       | Inscription           |
| `/dashboard`      | Dashboard mock        |

---

## 🔧 Configuration

### Variables d’environnement (`.env`)

Crée un fichier `.env` à la racine du projet avec :

```env
MONGO_INITDB_ROOT_USERNAME=mongoadmin
MONGO_INITDB_ROOT_PASSWORD=secret
MONGO_INITDB_DATABASE=saas
DATABASE_URL=mongodb://mongoadmin:secret@db:27017/
```

> Ces variables sont utilisées automatiquement par Docker & Go

---

## 🚀 Technologies utilisées

- [Go 1.22](https://go.dev/)
- [MongoDB](https://www.mongodb.com/)
- [React 18 + Vite](https://vitejs.dev/)
- [Docker](https://www.docker.com/)
- [Lucide Icons](https://lucide.dev/)

---

## ✅ À venir (roadmap possible)

- [ ] Authentification avec JWT
- [ ] Connexion sécurisée via HTTPS
- [ ] Système de rôles / permissions
- [ ] Déploiement automatique (Render, Railway...)
