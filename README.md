# Store-Api Using Golang
RestfullApi using golang httpRouter and Postgres as Database, This Project Using Layer Architecture:
1. Controller
2. Service
3. Repository

## Entity Relational Diagram
```mermaid
erDiagram
    users{
        string id
        string username
        string password
        string role
        timestamp created_at
        bool is_deleted
    }
    products ||--|{ categories: have_1
    products{
        string id
        string product_name
        string price
        string category_id
        timestamp created_at
        bool is_deleted
    }
    shopping_carts ||--|{ shopping_cart_items: contains
    users||--o{ shopping_carts: have
    shopping_carts {
        string id
        string user_id
        string status
        timestamp created_at
        bool is_deleted
    }
    shopping_cart_items ||--|{ products: have_1
    shopping_cart_items{
        string id
        string product_id
        int quantity
    }
    categories{
        string id
        string category_name
        timestamp created_at
        bool is_deleted
    }
```
## This Project Has 4 Domain
1. user
2. product
3. shopping cart
4. product category

## How to Run This Project
1. Clone the project using `git clone [URL]` 
2. Create Postgres Database `store-api`
3. Install Dependencies `go mod tidy`
4. Run the Migration using `go run migration/migration.go`
5. Run app using `go run main.go`

# REST API

## Register User
### Request

`POST localhost:3000/api/users`

    curl --location 'localhost:3000/api/users' \
    --header 'x-api-key: RAHASIA' \
    --header 'Content-Type: application/json' \
    --data '{
        "username": "admin",
        "password": "password",
        "role": "ADMIN" 
    }'

## Login
### Request

`POST localhost:3000/api/users/login`
    curl --location 'localhost:3000/api/users/login' \
    --header 'x-api-key: d155d392-ff7f-4569-9465-1387afca7684' \
    --header 'Content-Type: application/json' \
    --data '{
        "username": "admin",
        "password": "password"
    }'

