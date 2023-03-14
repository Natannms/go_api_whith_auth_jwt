# Go API with JWT Authentication

This is a sample RESTful API built with Go that uses JWT authentication.

## Requirements

-   Go 1.16 or higher
-   PostgreSQL 9.5 or higher

## Installation

1.  Clone the repository:
``git clone https://github.com/Natannms/go_api_whith_auth_jwt.git``
1.  Navigate to the project directory:

    bashCopy code

    `cd go_api_whith_auth_jwt`

2.  Create a `.env` file in the root directory and add the following environment variables:

    bashCopy code

    `DATABASE_URL=postgres://user:password@localhost:5432/db_name?sslmode=disable
    JWT_SECRET=my_secret_key`

3.  Install the dependencies:

    bashCopy code

    `go mod download`

4.  Build and run the application:

    bashCopy code

    `go run main.go`

5.  The API is now running on `http://localhost:8000`
## Usage

### Sign Up

To create a new user account, send a `POST` request to `/api/v1/signup` with the following JSON payload:

jsonCopy code

`{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "password": "secret"
}`

### Login

To log in, send a `POST` request to `localhost:port/login` with the following JSON payload:

jsonCopy code

`{
  "email": "john.doe@example.com",
  "password": "secret"
}`

The server will respond with a JWT token that can be used to authenticate future requests.

### Protected Endpoint

To access a protected endpoint, include the JWT token in the `Authorization` header of the request:

makefileCopy code

`Authorization: Bearer <JWT token>`

Send a `GET` request to `/api/v1/profile` to retrieve the user profile information.

### Logout

To log out, send a `POST` request to `localhost:port/logout`. The server will invalidate the JWT token.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
