# Actions Service

## Description

The Nostos Actions Service manages user interactions such as "likes" and "favorites" for albums and media content. It optimizes calculation and storage through Redis caching and generates events for notifications, enhancing the social experience of the platform.

---

## Features

- Registration of likes on albums, trips, and media content  
- User favorites functionality for media content  
- Prevention of duplicate likes by the same user on a single publication  
- Fast calculation and updating of like counters without database reloading  
- Redis caching for popular content tracking  
- Event generation for notifications (`event.like.new`)  
- User following/unfollowing functionality  
- Cross-service integration with Profile and Auth services  

---

## Technologies Used

- **Language**: Go  
- **Framework**: Gin  
- **Database**: PostgreSQL with GORM  
- **Cache**: Redis  
- **Authentication**: JWT via Auth Service  
- **Messaging**: NATS  
- **Orchestration**: Docker  

---

## Architecture

The service follows a clean architecture pattern with the following components:

- **API Controllers**: Handle HTTP requests and responses  
- **Services**: Implement business logic  
- **Repositories**: Handle database operations  
- **Models**: Define data structures  
- **Clients**: Communicate with other microservices  

---

## Database Schema

The service uses the following schema in PostgreSQL:

- `activity.likes`: Stores like information  
- `activity.actions`: Stores user actions (likes, favorites, follows)  

---

## Action Features

### Like Management

Users can like and unlike various content types (trips, media) with proper tracking and prevention of duplicates.

### Favorites System

Media content can be marked as favorite by users, allowing them to build personal collections of preferred content.

### User Following

The service manages user follow relationships, enabling social networking features within the platform.

### Cross-Service Integration

The Actions Service integrates with:

- **Profile Service**: To fetch user profiles for action attribution  
- **Auth Service**: For authentication and authorization  

---

## Security

- **Authentication**: Implemented using JWT tokens from the Auth Service  
- **Duplicate Prevention**: Users can't like the same content multiple times  
- **Performance**: Redis used for optimizing high-traffic operations  

---

