

# Go CRUD APPLICATION (GO + REDIS)

This project implements a CRUD (Create, Read, Update, Delete) microservice using Golang, Chi Router, and Redis as the database

---

## ✨Features

- CRUD operations for managing resources.
- Uses Redis as the database for storage.
- Utilizes the Chi router for handling HTTP requests.

---
## API Endpoints

- `GET /resources`: Get all resources.
- `GET /resources/{id}`: Get a resource by ID.
- `POST /resources`: Create a new resource.
- `PUT /resources/{id}`: Update a resource by ID.
- `DELETE /resources/{id}`: Delete a resource by ID.

---

## 🚀 Getting Started  

### Open Using Daytona  

1. **Install Daytona**: Follow the [Daytona installation guide](https://www.daytona.io/docs/installation/installation/).

2. **Create the Workspace**:  
   ```bash  
   daytona create https://github.com/Abiji-2020/Go_CRUD_Application.git
   ```  
3. **Run the Redis server**:
4. ```bash
   redis-server
   ```  
5. **Run the application**
   ```bash
   go run main.go
   ```
---
