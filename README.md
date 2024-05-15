# Store-Api Using Golang
Simple RestfullApi using golang httpRouter and Postgres as Database, This Project Using Layer Architecture:
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
    shopping_carts ||--|{ products: have_1
    users||--o{ shopping_carts: have
    shopping_carts {
        string id
        string user_id
        string status
        string product_id
        int quantity
        timestamp created_at
        bool is_deleted
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

## What Implemented in This Repo?
1. restfull api using httprouter
2. integration with postgres using [GORM](https://gorm.io/)

## How to Run This Project
1. Clone the project using `git clone [URL]` 
2. Create Postgres Database `store-api`
3. Install Dependencies `go mod tidy`
4. Run the Migration using `go run migration/migration.go`
5. Run app using `go run main.go`

## How to Run Using Docker Compose
1. Clone the project using `git clone [URL]` 
2. Create env file for docker compose `.sample.env` already exists for the sample
3. Run Command `docker compose up`
4. Create Migration `edit migration db URI and forwarded port` in **migration/migration.go**
5. Run the Migration using `go run migration/migration.go`
6. App run properly

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

## Create Product
Only User With Role `ADMIN` can access this API
### Request

`POST localhost:3000/api/products`

    curl --location 'localhost:3000/api/products' \
    --header 'X-API-Key: RAHASIA' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU1MjQ5NDMsImlkIjoiMTM2YzRmNDEtZmM5ZS00NWIyLWJhMGQtZTk0NjU0OTUwNjIwIiwicm9sZSI6IkFETUlOIiwidXNlcm5hbWUiOiJhZG1pbiJ9.z1yWvXrRPhSXbi5zwW4gqhn3bd7vai9a3vRJNYYyM0Q' \
    --data '{
        "name": "contoh",
        "categoryId": "c2befba5-91ee-49e3-9b4b-bf159741b404",
        "price": 55000
    }'

## List Product
### Request
`query: category_id (optional)` and `query: q (optional as search by name)`

`GET localhost:3000/api/products?category_id=c2befba5-91ee-49e3-9b4b-bf159741b404`

    curl --location 'localhost:3000/api/products?category_id=c2befba5-91ee-49e3-9b4b-bf159741b404' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU1NzI3MjEsImlkIjoiMTM2YzRmNDEtZmM5ZS00NWIyLWJhMGQtZTk0NjU0OTUwNjIwIiwicm9sZSI6IkFETUlOIiwidXNlcm5hbWUiOiJhZG1pbiJ9.uKj3146OBSAt1cBYAM5_Fg2OIZBOuK5f3Rnet8AyjGM'

## Find Product By Id
### Request
`params: productId (required)`

`GET localhost:3000/api/products/:productId`

    curl --location 'localhost:3000/api/products/b4c1d2fc-c5d1-4736-b162-b4715583adc2' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU1ODI5NDksImlkIjoiMTM2YzRmNDEtZmM5ZS00NWIyLWJhMGQtZTk0NjU0OTUwNjIwIiwicm9sZSI6IkFETUlOIiwidXNlcm5hbWUiOiJhZG1pbiJ9.ccoJSoG2jwtCXGuiVQQjdoNZ4O7kjPN23cvKPlp4dwQ'

## Create Product Categories
Only User With Role `ADMIN` can access this API
### Request

`POST localhost:3000/api/categories` 

    curl --location 'localhost:3000/api/categories' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU1MjM1NjgsImlkIjoiMTM2YzRmNDEtZmM5ZS00NWIyLWJhMGQtZTk0NjU0OTUwNjIwIiwicm9sZSI6IkFETUlOIiwidXNlcm5hbWUiOiJhZG1pbiJ9.-Qlopme_1jAX7mOAjsdaHwgnTO347mz8h_Z-Mo2r5WA' \
    --data '{
        "categoryName": "pakaian"
    }'

## List Product Categories
### Request

`GET localhost:3000/api/categories`

    curl --location 'localhost:3000/api/categories' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU1MjM1NjgsImlkIjoiMTM2YzRmNDEtZmM5ZS00NWIyLWJhMGQtZTk0NjU0OTUwNjIwIiwicm9sZSI6IkFETUlOIiwidXNlcm5hbWUiOiJhZG1pbiJ9.-Qlopme_1jAX7mOAjsdaHwgnTO347mz8h_Z-Mo2r5WA'

## Create Cart
### Request

`POST localhost:3000/api/shopping-cart`

    curl --location 'localhost:3000/api/shopping-cart' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU1OTA5MjgsImlkIjoiMTM2YzRmNDEtZmM5ZS00NWIyLWJhMGQtZTk0NjU0OTUwNjIwIiwicm9sZSI6IkFETUlOIiwidXNlcm5hbWUiOiJhZG1pbiJ9.EHmnqyqcnxfoCHil-fXo6XEYzYgmKnRwPtUGzQcx-kA' \
    --data '{
        "items": [
            {
                "productId": "b4c1d2fc-c5d1-4736-b162-b4715583adc2",
                "qty": 4
            }
        ]
    }'

# List Cart
### Request
`query: status (one of PENDING OR SUCCESS) optional`

`GET localhost:3000/api/shopping-cart`

    curl --location 'localhost:3000/api/shopping-cart?status=PENDING' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU1OTA5MjgsImlkIjoiMTM2YzRmNDEtZmM5ZS00NWIyLWJhMGQtZTk0NjU0OTUwNjIwIiwicm9sZSI6IkFETUlOIiwidXNlcm5hbWUiOiJhZG1pbiJ9.EHmnqyqcnxfoCHil-fXo6XEYzYgmKnRwPtUGzQcx-kA'

# Checkout Cart
### Request 

`PATCH localhost:3000/api/shopping-cart/checkout`

    curl --location --request PATCH 'localhost:3000/api/shopping-cart/checkout' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU2NTU3MzIsImlkIjoiMTM2YzRmNDEtZmM5ZS00NWIyLWJhMGQtZTk0NjU0OTUwNjIwIiwicm9sZSI6IkFETUlOIiwidXNlcm5hbWUiOiJhZG1pbiJ9.BHWG4isucwQKJh4IPaVD_7jjf0hbMSptfNlJqAzaNzc' \
    --data '{
        "shoppingCartIds": [
            "59cd59a7-1353-4316-bda5-0fa43d26a9fa"
        ]
    }'
### Checkout Workflow
```mermaid
sequenceDiagram
    Client->>Cart: Update Status Cart
    Cart->>DB:Start DB Transaction
    Cart->>DB: find cart
    alt cart not found
        Cart->>DB: Rollback transaction
    else cart found
        Cart->>User: Checking User Balance
    alt balance not enough
    Cart->>DB: Rollback data
    else balance enough
    Cart->>DB: update status cart
    Cart->>User: adjust user balance
    Cart->>DB: commit if succes and rollback if error
    end
    end
    DB-->>Cart:response commit success or rollback
    Cart-->>Client:response data
```
