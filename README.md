# Microservice Test

## Run server

#### Auth server
go to folder auth, please enter command below
```
go run main.go
```

#### Logistic server
go to folder logistic, please enter command below
```
go run main.go
```

## Run with Docker Compose
```
docker compose up
```

## End point
#### Auth

| Service                   | Method | URL                                    |
| ------------------------- | ------ | -------------------------------------- |
| Login                     | POST   | http://localhost:3000/auth/            |
| Register                  | POST   | http://localhost:3000/auth/register    |
| Check Token               | GET    | http://localhost:3000/auth/check-token |

#### Logistic

| Service                   | Method | URL                                    |
| ------------------------- | ------ | -------------------------------------- |
| Get All                   | GET    | http://localhost:3001/logistic/all    |
| Get Data                  | GET    | http://localhost:3001/logistic/         |
