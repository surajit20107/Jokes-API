# Jokes-API 

A simple and lightweight RESTful API built with Go, designed to provide random jokes for your applications and projects.

## Features

- Fetch random jokes via an HTTP endpoint.
- Built using Go for high performance and easy deployment.
- Simple interface, designed for quick integration into any app or website.
- Ideal for learning about Go web APIs or for use in chatbots, fun apps, and more.

## Getting Started

### Prerequisites

- Go (1.22+ recommended)

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
GET http://localhost:8080/jokes
```

## API Endpoints

- `GET /jokes` - Returns 5 jokes in JSON format.
- `POST /jokes/create` - Creates a new joke.
- `PATCH /jokes/:id` - Updates a perticular joke.
- `DELETE /jokes/:id` - Removes a joke.

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

> Made with ❤️ & lots of ☕ by Surajit in GoLang 🦦
