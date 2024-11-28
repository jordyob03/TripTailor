# TripTailor

## Overview

**TripTailor** is a platform designed to revolutionize travel planning by offering a seamless and personalized experience for users to create, share, and explore travel itineraries. Combining personalization, community engagement, and intuitive features, TripTailor empowers users to craft itineraries tailored to their unique interests, demographics, and budgets.

## Features

- **Personalized Itineraries**: Create travel plans with detailed information such as destinations, events, timings, and budgets.
- **Community Engagement**: Like, comment, and save itineraries into user-defined boards.
- **Advanced Search**: Dynamic search functionality with tag-based filtering and ranking for relevant results.
- **Scalable Architecture**: Microservice-based design for modularity, fault tolerance, and scalability.
- **Profile Customization**: Tailor profiles with travel interests, languages spoken, and tags for tailored recommendations.

## Technology Stack

### Frontend
- **React.js**: Component-driven framework for building dynamic user interfaces.
- **JavaScript (ES6+)**: Used for frontend logic and interaction handling.
- **CSS**: Modular and responsive styling for consistent design.

### Backend
- **Golang**: Efficient, high-performance language for backend logic.
- **Gin Framework**: RESTful API creation with lightweight and fast routing.
- **Docker**: Containerized services for consistent environments and easy deployment.

### Database
- **PostgreSQL**: Relational database supporting JSONB for structured and semi-structured data.
- **PgAdmin**: GUI for managing and monitoring the database.

## Services Architecture

### Microservices
- **Main Service**: Central service for database management and API exposure.
- **Authentication Service**: Secure user signup, login, and JWT-based authorization.
- **Profile Service**: Handles user profile creation and updates.
- **Itinerary Service**: Manages creation, storage, and retrieval of travel itineraries.
- **Feed Service**: Displays curated itineraries based on user-selected tags.
- **Search Service**: Enables advanced itinerary search with relevance-based ranking.
- **Save Service**: Allows users to save and manage boards with curated itineraries and posts.

## Core Principles

- **Single Responsibility Principle**: Each microservice is dedicated to a specific feature.
- **Loose Coupling**: Independent services minimize dependencies and enhance scalability.
- **Testability**: Modular design ensures robust unit and integration testing capabilities.

---

TripTailor combines innovation and simplicity to create an enjoyable and efficient travel planning experience for users, making their journey from planning to execution as seamless as possible.
