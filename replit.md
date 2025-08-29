# Overview

This is a RESTful Jokes API built with Go that provides endpoints for managing and retrieving jokes. The API is designed to be lightweight and simple, offering basic CRUD operations for joke management. It's structured as a learning project that demonstrates Go web API development patterns and can be easily integrated into applications, chatbots, or other services that need joke content.

# User Preferences

Preferred communication style: Simple, everyday language.

# System Architecture

## Backend Architecture

The application follows a simple monolithic architecture built with Go's standard HTTP library. The main entry point is `main.go` which sets up the HTTP server running on port 8080. The architecture is designed for simplicity and ease of deployment, making it suitable for learning purposes and quick integrations.

## API Design

The REST API follows standard HTTP conventions with the following endpoint structure:

- **GET /jokes** - Retrieves 5 random jokes in JSON format
- **POST /jokes/create** - Creates a new joke entry
- **PATCH /jokes/:id** - Updates a specific joke by ID
- **DELETE /jokes/:id** - Removes a joke by ID

The API uses JSON for all request and response payloads, making it easy to integrate with various client applications.

## Data Storage

Currently, the application appears to use an in-memory data storage approach, which is suitable for development and testing but would need to be enhanced with a persistent database for production use. The simple architecture allows for easy extension to include database integration when needed.

## Development Environment

The project is configured for Go 1.22+ and includes development tooling support through the Go toolchain. The structure supports hot-reloading during development, making it developer-friendly for rapid iteration.

# External Dependencies

## Core Dependencies

- **Go Standard Library** - HTTP server functionality and basic JSON handling
- **Go Toolchain** - Development and build tools (Go 1.21+ versions detected in telemetry)

## Development Tools

- **Go Language Server (gopls)** - IDE support for development, as evidenced by the telemetry data
- **Air** - Live reload functionality for development (referenced in gopath dependencies)

## Potential Future Dependencies

The gopath structure suggests the project may be prepared to integrate additional libraries commonly used in Go web development:

- **Gin Framework** - Web framework (dependencies present in gopath)
- **GORM** - Object-relational mapping for database operations
- **Protocol Buffers** - For advanced serialization if needed
- **Sonic** - High-performance JSON library for better performance

The modular structure allows for easy integration of these dependencies as the application grows in complexity.