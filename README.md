# Accounta Backend in golang and gin, gorm

This is a simple backend for accounta webapp. It is written in golang and uses gin and gorm.

## Getting Started

1. Clone the repository
2. Create a `.env` file in the root directory and add the following environment variables:

```
PORT=8080
MODE=debug
   
# JWT Secret
JWT_SECRET=secret

# MySQL Config
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=root
MYSQL_DB=test
```


3. Run the following command to start the server:

```
go run main.go
```


4. Postman collection is available [here](https://documenter.getpostman.com/view/34577808/2sAXjQ1qDF)

