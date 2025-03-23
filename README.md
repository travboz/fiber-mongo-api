# Go Fiber and MongoDB API
<img src="https://raw.githubusercontent.com/egonelbre/gophers/63b1f5a9f334f9e23735c6e09ac003479ffe5df5/vector/friends/stovepipe-hat-front.svg" alt="Stove-Pipe-Hat-Wearing Gopher" width="300">


A user management API built with Go, Fiber, and MongoDB. This project was primarily a refactor effort to improve the codebase by decoupling components, leveraging dependency injection, and implementing the repository pattern for better maintainability and scalability. It enables basic CRUD operations for handling user data through an HTTP server and utilizes Docker Compose for service orchestration.
Commit history includes multiple changes to the code. 


## Features
TBD

## Getting Started

### Prerequisites
- Docker
- Docker Compose
- Go (1.22+ recommended)

## Installation

1. Clone this repository:
   ```sh
   git clone https://github.com/travboz/fiber-mongo-api.git
   cd fiber-mongo-api
   ```
2. Run docker container:
    ```sh
    make compose-up
   ```
3. Run server:
    ```sh
    make run
    ```
4. Navigate to `http://localhost:6000` and call an endpoint

### `.env` file
This server uses a `.env` file for basic configuration.
Here is an example of the `.env`:
   ```sh
   TBD=true
   ```
   
## API endpoints

| Method   | Endpoint        | Description                 |
|----------|-----------------|-----------------------------|
| `POST`   | `/user`         | Create a new user           |
| `GET`    | `/user/:userId` | Get user by ID              |
| `PUT`    | `/user/:userId` | Update a user by ID         |
| `DELETE` | `/user/:userId` | Delete a user by ID         |
| `GET`    | `/users`        | Get all users               |
| `GET`    | `/health`       | Health check / service status|


## Example usage

TBD

## Contributing
Feel free to fork and submit PRs!

## License:
`MIT`


This should work for GitHub! Let me know if I can make any tweaks. 


## Image
Image by [Egon Elbre](https://github.com/egonelbre), used under CC0-1.0 license.