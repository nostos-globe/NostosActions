# Nostos Actions Service

The **Nostos Actions Service** manages social interactions across the platform, including likes, favorites, and user follows. It ensures a seamless social experience through robust action tracking and integration with authentication and profile services.

---

## 🚀 Features

* Like functionality for trips and media
* User favorites for media content
* Prevents duplicate likes per user per item
* Like/favorite counters with efficient database updates
* Event publishing for notifications (`event.like.new`)
* Follow/unfollow other users
* Integrates with Profile and Auth services for identity and access

---

## 📌 Endpoints

### 🔹 Likes

* **Like a Trip**
  `POST /api/likes/trip/:id`
  Likes a specific trip.

* **Unlike a Trip**
  `DELETE /api/likes/trip/:id`
  Removes like from a trip.

* **Get Likes for a Trip**
  `GET /api/likes/trip/:id`
  Retrieves the list of users who liked the trip.

* **My Likes**
  `GET /api/likes/myLikes`
  Returns all content liked by the current user.

* **Likes by User ID**
  `GET /api/likes/userID/:id`
  Retrieves likes made by a specific user.

### 🔹 Favorites

* **Get Favorite Status (Media)**
  `GET /api/favourites/media/:id`
  Checks if the current user has favorited a media item.

* **Add to Favorites**
  `POST /api/favourites/media/:id`
  Marks a media item as a favorite.

* **Remove from Favorites**
  `DELETE /api/favourites/media/:id`
  Unmarks a media item as a favorite.

### 🔹 User Actions

* **Create Action (e.g. Follow)**
  `POST /api/actions/create`
  Records a user action such as following another user.

---

## ⚙️ Installation and Configuration

### Prerequisites

* Go installed
* PostgreSQL
* Docker and Docker Compose (for local development)
* Auth service with JWT support

### Installation

```bash
git clone https://github.com/nostos-globe/NostosActions.git
cd NostosActions
go mod download
```

### Configuration

Ensure the following environment variables or Vault secrets are set:

* `DATABASE_URL`
* `JWT_SECRET`
* `NATS_URL` (for event messaging)

Vault can be accessed using a token, AppRole, or Kubernetes auth.

---

## ▶️ Running the Application

```bash
go run cmd/main.go
```

---

## 🧱 Technologies Used

* **Language**: Go
* **Framework**: Gin
* **Database**: PostgreSQL (GORM)
* **Authentication**: JWT via Auth Service
* **Messaging**: NATS
* **Orchestration**: Docker

---

## 🏗️ Project Structure

```
NostosActions/
├── cmd/                  # Application entry point
│   └── main.go
├── internal/
│   ├── api/              # HTTP route handlers
│   ├── db/               # Database access logic
│   ├── models/           # Data models
│   └── service/          # Business logic
├── pkg/
│   ├── config/           # Config and secret management
│   └── messaging/        # NATS event publishing
├── Dockerfile
├── go.mod
└── README.md
```
