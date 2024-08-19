#A CRUD API with Golang

Here's a `README.md` file for your project:

```markdown
# CRUD Movie API

This is a simple RESTful API built with Go that manages a list of movies. The API allows users to create, retrieve, update, and delete movies. The project uses the `gorilla/mux` package for routing.

## Features

- **Get All Movies:** Retrieve a list of all movies.
- **Get a Single Movie:** Retrieve details of a specific movie by its ID.
- **Create a Movie:** Add a new movie to the list.
- **Update a Movie:** Modify the details of an existing movie by its ID.
- **Delete a Movie:** Remove a movie from the list by its ID.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/harshal-rembhotkar/CRUD-API.git
   cd CRUD-API
   ```

2. **Install dependencies:**

   Make sure you have Go installed. Then run:

   ```bash
   go mod tidy
   ```

3. **Run the server:**

   ```bash
   go run main.go
   ```

4. **Access the API:**

   The server will start at `http://localhost:8000`. You can use tools like `curl` or Postman to interact with the API.

## API Endpoints

### 1. Get All Movies

- **Endpoint:** `/movies`
- **Method:** `GET`
- **Response:** JSON array of movies.

### 2. Get a Single Movie

- **Endpoint:** `/movies/{id}`
- **Method:** `GET`
- **Path Variable:** `id` - The ID of the movie.
- **Response:** JSON object of the movie.

### 3. Create a Movie

- **Endpoint:** `/movies`
- **Method:** `POST`
- **Request Body:** JSON object of the movie (excluding `id`).
- **Response:** JSON object of the created movie with an ID.

### 4. Update a Movie

- **Endpoint:** `/movies/{id}`
- **Method:** `PUT`
- **Path Variable:** `id` - The ID of the movie.
- **Request Body:** JSON object of the updated movie.
- **Response:** JSON object of the updated movie.

### 5. Delete a Movie

- **Endpoint:** `/movies/{id}`
- **Method:** `DELETE`
- **Path Variable:** `id` - The ID of the movie.
- **Response:** No content (success).

## Example Movie Object

```json
{
  "id": "1",
  "isbn": "438",
  "director": {
    "firstname": "D",
    "lastname": "Rembhotkar"
  }
}
```

## Dependencies

- [Gorilla Mux](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for Go.


```

