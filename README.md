

# Go CRUD APPLICATION (GO + REDIS)

This project implements a CRUD (Create, Read, Update, Delete) microservice using Golang, Chi Router, and Redis as the database

---

## ✨Features

- CRUD operations for managing resources.
- Uses Redis as the database for storage.
- Utilizes the Chi router for handling HTTP requests.

---
## API Endpoints

- `GET /orders`: List all Orders
- `GET /orders/{id}`: Get an Order by ID.
- `POST /orders`: Create a new Order.
- `PUT /orders/{id}`: Update an Order by ID.
- `DELETE /orders/{id}`: Delete an Order by ID.

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

# Example Request 


## POSTMAN type Requests

### Get All Orders

```bash
GET http://localhost:3000/orders
```

The request will return with the reponse of: 
```json
{
  "items": [
    {
      "order_id": 1.2470706902630818e+19,
      "customer_id": "550e8400-e29b-41d4-a716-446655440000",
      "line_items": [
        {
          "item_id": "123e4567-e89b-12d3-a456-426614174000",
          "quantity": 2,
          "price": 1000
        },
        {
          "item_id": "123e4567-e89b-12d3-a456-426614174001",
          "quantity": 1,
          "price": 500
        }
      ],
      "created_at": "2024-12-11T09:19:09.15246014Z",
      "shipped_at": null,
      "completed_at": null
    }
  ]
}
```

### Create an order

```bash
POST http://localhost:3000/orders
```
Payload for the request:

```json
{
  {
    "order_id": 12345,
    "customer_id": "550e8400-e29b-41d4-a716-446655440000",
    "line_items": [
      {
        "item_id": "123e4567-e89b-12d3-a456-426614174000",
        "quantity": 2,
        "price": 1000
      },
      {
        "item_id": "123e4567-e89b-12d3-a456-426614174001",
        "quantity": 1,
        "price": 500
      }
    ]
  }
  
```
The reponse will be same with the status:`OK`

### To update the status of the order

```bash
PUT http://localhost:3000/orders/{id}
```
#### Payload to be sent:
```json
{
    "status":"shipped"
}
```

#### Response of the update:
```json
{
  "order_id": 1234,
  "customer_id": "550e8400-e29b-41d4-a716-446655440000",
  "line_items": [
    {
      "item_id": "123e4567-e89b-12d3-a456-426614174000",
      "quantity": 2,
      "price": 1000
    },
    {
      "item_id": "123e4567-e89b-12d3-a456-426614174001",
      "quantity": 1,
      "price": 500
    }
  ],
  "created_at": "2024-12-11T09:40:33.534526754Z",
  "shipped_at": "2024-12-11T10:01:24.974032869Z",
  "completed_at": null
}
```

---

## cURL Request

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

---



