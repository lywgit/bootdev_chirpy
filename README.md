# Learn HTTP Servers in Go

This repository contains the code developed as I worked through the "[Learn HTTP Servers in Go](https://www.boot.dev/courses/learn-http-servers-golang)" course on [Boot.dev](https://www.boot.dev). The lessons guide you through the building of simple a HTTP web server *Chirpy* from scratch, which is a simplified Twitter-like service.


*Note: The Usage and API docs were gereated with the help of Gemini*

## Usage

Follow the guide if you would like to run the Chirpy server locally.

### Prerequisites

Before you begin, ensure you have the following installed:

*   **Go:** The programming language used for this project.
*   **PostgreSQL:** The database used to store Chirpy's data.
*   **`goose`:** A database migration tool for managing schema changes. You can install it with:
    ```bash
    go install github.com/pressly/goose/v3/cmd/goose@latest
    ```
*   **An HTTP Client:** A tool like [cURL](https://curl.se/) or [Postman](https://www.postman.com/) for interacting with the API.

### Installation & Setup

1.  **Clone the Repository:**
    ```bash
    git clone <repository-url>
    cd bootdev_chirpy
    ```

2.  **Configure Environment Variables:**
    Create a `.env` file in the root of the project by copying the example file:
    ```bash
    cp env.example .env
    ```
    Now, open the `.env` file and update the following variables with your local database and JWT information:
    *   `DB_URL`: Your PostgreSQL connection string.
    *   `JWT_SECRET`: A secret key for signing JWTs.
    *   `POLKA_KEY`: Your API key for the Polka service.
    *   `PLATFORM`: Set to `dev` for development mode or `production` for production.

3.  **Apply Database Migrations:**
    Navigate to the `sql/schema` directory and run `goose` to set up the database tables:
    ```bash
    cd sql/schema
    goose postgres "postgres://user:password@localhost:5432/chirpy" up
    ```
    *(Replace the connection string with your own.)*

### Running the Server

1.  **Start the Server:**
    From the project's root directory, run the following command:
    ```bash
    go run ./main.go
    ```

2.  **Access the Application:**
    The server will be running at `http://localhost:8080`. You can access the main application at `http://localhost:8080/app` and the API endpoints under `http://localhost:8080/api/`.


## APIs

The following endpoints are available: 

### Users

*   **Create a User**
    *   **Method:** `POST`
    *   **Endpoint:** `/api/users`
    *   **Request Body:**
        ```json
        {
            "email": "example@example.com",
            "password": "yourpassword"
        }
        ```
    *   **Success Response:** `201 Created` with the new user object (without the password).

*   **Update a User**
    *   **Method:** `PUT`
    *   **Endpoint:** `/api/users`
    *   **Authentication:** Requires a valid JWT Bearer token.
    *   **Request Body:**
        ```json
        {
            "email": "new-email@example.com",
            "password": "new-password"
        }
        ```
    *   **Success Response:** `200 OK` with the updated user object.

### Authentication

*   **Login**
    *   **Method:** `POST`
    *   **Endpoint:** `/api/login`
    *   **Request Body:**
        ```json
        {
            "email": "user@example.com",
            "password": "password",
            "expires_in_seconds": 3600
        }
        ```
    *   **Success Response:** `200 OK` with a JWT access token and a refresh token.

*   **Refresh Token**
    *   **Method:** `POST`
    *   **Endpoint:** `/api/refresh`
    *   **Authentication:** Requires a valid JWT Bearer token (using the refresh token).
    *   **Success Response:** `200 OK` with a new JWT access token.

*   **Revoke Token**
    *   **Method:** `POST`
    *   **Endpoint:** `/api/revoke`
    *   **Authentication:** Requires a valid JWT Bearer token (using the refresh token).
    *   **Success Response:** `204 No Content`

### Chirps

*   **Create a Chirp**
    *   **Method:** `POST`
    *   **Endpoint:** `/api/chirps`
    *   **Authentication:** Requires a valid JWT Bearer token.
    *   **Request Body:**
        ```json
        {
            "body": "This is a new chirp!"
        }
        ```
    *   **Success Response:** `201 Created` with the newly created chirp object.

*   **Get All Chirps**
    *   **Method:** `GET`
    *   **Endpoint:** `/api/chirps`
    *   **Success Response:** `200 OK` with an array of all chirp objects.

*   **Get a Chirp by ID**
    *   **Method:** `GET`
    *   **Endpoint:** `/api/chirps/{chirpID}`
    *   **Success Response:** `200 OK` with the chirp object.

*   **Delete a Chirp**
    *   **Method:** `DELETE`
    *   **Endpoint:** `/api/chirps/{chirpID}`
    *   **Authentication:** Requires a valid JWT Bearer token. The user must be the author of the chirp.
    *   **Success Response:** `204 No Content`

### Polka Webhooks

*   **Polka Webhook**
    *   **Method:** `POST`
    *   **Endpoint:** `/api/polka/webhooks`
    *   **Authentication:** Requires a valid API key from Polka.
    *   **Request Body:**
        ```json
        {
            "event": "user.upgraded",
            "data": {
                "user_id": 123
            }
        }
        ```
    *   **Success Response:** `204 No Content`

### Admin

*   **Get Metrics**
    *   **Method:** `GET`
    *   **Endpoint:** `/admin/metrics`
    *   **Success Response:** `200 OK` with an HTML page displaying the number of fileserver hits.

*   **Reset Server**
    *   **Method:** `POST`
    *   **Endpoint:** `/admin/reset`
    *   **Success Response:** `200 OK` with a message indicating the server has been reset. This endpoint is only available in `dev` mode.


## My Takeways

Some of the lessons are pretty challenging for me as many concepts are new.
Key concepts learned:

- Routing and handler
- Http requests and information (header and json body)
- Http respond status code
- Authentication process: password, JWT and refresh tokens
- UUID in database
- goose: a database migration tool
- sqlc: a tool that converts SQL queries into Go functions

### A diagram of the layered structure showing the request flow

```mermaid
    graph TD
    subgraph "Client"
        A[Browser / HTTP Client]
    end
    subgraph "Presentation Layer (Go net/http)"
        B["/"<br>HTTP Router in main.go] --> C{Request Mux};
        C --> D["/api/..."<br>API Handlers];
        C --> E["/app"<br>File Server];
    end


    subgraph "Business & Service Layer"
        F["internal/auth"<br>JWT & Password Logic];
        G["profane_words.go"<br>Input Validation];
    end

    subgraph "Data Access Layer (DAL)"
        H["internal/database<br>Repository Pattern (sqlc)"]
    end

    subgraph "Database"
        I[(PostgreSQL)]
    end


    %% Request Flow
    A --> B;
    D --> F;
    D --> G;
    D --> H;
    H --> I;
```


## Some Qeustions

- No Goroutine is explicitly used. Perhaps this is automatically handled under http.Server?


