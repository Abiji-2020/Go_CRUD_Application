

# Go CRUD APPLICATION (GO + REDIS)

This project implements a CRUD (Create, Read, Update, Delete) microservice using Golang, Chi Router, and Redis as the database

---

## ✨Features

- CRUD operations for managing resources.
- Uses Redis as the database for storage.
- Utilizes the Chi router for handling HTTP requests.

---
## API Endpoints

- `GET /orders`: Get all Orders.
- `GET /orders/{id}`: Get a Orders by ID.
- `POST /orders`: Create a new Order.
- `PUT /orders/{id}`: Update a Order by ID.
- `DELETE /orders/{id}`: Delete a Order by ID.

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
5. **Run the application in  a new terminal**
   ```bash
   go run main.go
   ```
   The application will run on `PORT : 3000`
---

## Example Request 

### Create an order

```bash 
curl --header "Content-Type: application/json" \
     --request POST \
     --data @example.json \
     localhost:3000/orders
```

### List Orders
```bash
curl localhost:3000/orders
```
