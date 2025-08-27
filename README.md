# Jokes-API 

A simple and lightweight RESTful API built with Go, designed to provide random jokes for your applications and projects.

## Features

- Fetch random jokes via an HTTP endpoint.
- Built using Go for high performance and easy deployment.
- Simple interface, designed for quick integration into any app or website.
- Ideal for learning about Go web APIs or for use in chatbots, fun apps, and more.

## Getting Started

### Prerequisites

- Go (1.16+ recommended)

### Installation

Clone this repository:

```bash
git clone https://github.com/surajit20107/Jokes-API.git
cd Jokes-API
```

### Running the API

```bash
go run main.go
```

The server will start (by default on port 8080). You can then access the jokes API at:

```
GET http://localhost:8080/joke
```

## API Endpoints

- `GET /jokes` - Returns all jokes in JSON format.
- `GET /jokes/:id` - Returns a single joke wkth same id.
- `GET /jokes/random` - Returns a random joke in JSON format.
- `POST /jokes/create` - Creates a new joke in jokes slice.
- `PATCH /jokes/:id` - Updates a perticular joke with the same id.
- `DELETE /jokes/:id` - Removes a joke with the same id of parameters.

## Example Response

```json
{
  "joke": "Why don't scientists trust atoms? Because they make up everything!"
}
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is open source and available under the MIT License.

---

> Made with Go 🦦
