# Dog App 🐶

A simple web application that allows you to fetch and display random dog facts and pictures. The backend is written in Go using the Gin framework, while the frontend is a minimalist page based on pure HTML and JavaScript. The application fetches data asynchronously from two external public APIs.

## Technologies Used

* **Backend:** Go (Golang) 1.25, [Gin Web Framework](https://github.com/gin-gonic/gin)
* **Frontend:** HTML, JavaScript (Fetch API)
* **External APIs:**
  * Dog pictures: [Dog API](https://dog.ceo/dog-api/)
  * Dog facts: [Dog Facts API](https://dogapi.dog/)
* **Deployment:** Docker (Multi-stage build)


## How It Works

The backend has two main endpoints:
1. `GET /` - Returns the main page (`main.html`).
2. `GET /api/dog` - API endpoint that concurrently (using goroutines and channels) queries external services for a dog picture and a dog fact, and then returns them in JSON format.

## How to Run the Project

You can run the application in two ways: using Docker or directly locally using the Go environment.

### Option 1: Running with Docker

Make sure you have [Docker](https://www.docker.com/) installed.

1. Build the image based on the Dockerfile in the main project directory:
    docker build -t dog-app .

2. Run the container, mapping port 8080:
    docker run -p 8080:8080 dog-app

3. Open your browser and go to: `http://localhost:8080`

### Option 2: Running locally (without Docker)

Requires the Go environment to be installed.

1. Navigate to the backend directory:
    cd backend

2. Download dependencies:
    go mod download

3. Run the application:
    go run main.go

4. Open your browser and go to: `http://localhost:8080`

*(Note: Because the Gin server loads views from `../frontend/*`, running the program must be done from within the `backend` folder so that relative paths work correctly).*