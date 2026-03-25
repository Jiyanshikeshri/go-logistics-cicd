# Go Logistics CI/CD Pipeline Project

##  Project Overview

This project demonstrates a **professional CI/CD pipeline** for a Go (Golang) based Logistics application using:

- GitHub Actions (CI/CD)
- Docker (Containerization)
- Docker Hub (Image Registry)

The pipeline automates:
- Code testing
- Application build
- Docker image creation
- Image push to Docker Hub
- Deployment simulation

---

## Tech Stack

- Go (Golang)
- Git & GitHub
- GitHub Actions
- Docker
- Docker Hub

---

## Project Structure
go-logistics-app/
│
├── cmd/
│ └── main.go
│
├── internal/
│ └── handlers/
│
├── tests/
│
├── Dockerfile
├── go.mod
├── .gitignore
└── README.md


---

## Features

- REST API for logistics tracking
- Unit testing for handlers
- Dockerized Go application
- Automated CI/CD pipeline
- Production-ready workflow

---

## CI/CD Pipeline Flow


Code Push → Test → Build → Docker Build → Docker Push → Deploy


### Pipeline Steps

1. **Test Stage**
   - Runs `go test ./...`

2. **Build Stage**
   - Builds Go application

3. **Docker Stage**
   - Builds Docker image
   - Pushes image to Docker Hub

4. **Deploy Stage**
   - Simulated deployment

---

## Docker Usage

### Build Image

```bash
docker build -t logistics-app .
▶ Run Container
docker run -p 9090:8080 logistics-app
 API Endpoints
➤ Get All Shipments
GET /shipments
➤ Create Shipment
POST /create
Sample Request:
{
  "id": "1",
  "status": "In Transit",
  "location": "Delhi"
}

 GitHub Secrets Used
DOCKER_USERNAME
DOCKER_PASSWORD
 Docker Hub Repository
jiyanshikeshrii/go-logistics-app