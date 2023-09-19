# Car-sale-api
this is a web-API for a car salling system (Bama.ir clone) with golang

## Project Summery
I used Docker foy services like Redis, Postgres, Elastic, Kibana and more

all endpoins are documented with swagger

it is visible at localhost:4000/swagger/index.html


## How to Use

1. Make sure you have the docker installed
   
        docker --version
2. CD to project directory and run the docker

       docker compose -f "docker\docker-compose.yaml" up -d --buil
3. then run the man.go file

        cd src/cmd/ && go run main.go

> A Database migration will execute to add some primary data like admin User and some temp data will be add to postgres as soon as you run the project



## 
### Config

All project settings/configs are written in **.yml** files so nothong is hard coded

first i fetch the the project ENV called "APP_ENV" project environment includes "development", "docker", "production" 

i used Viper to load all the configs and covert them into proper Structs

every change to project config is as easy as editing a config.yml file
## 

### Users
Users can login/logout/Sinup in two ways

First Users can Use **PhoneNumber and OTP**
- users send a request to /users/otp
- an otp Code will be sent to user and the otp is stored in **Redis** with 120Second TTL
- users will send otp and phone number to the proper endpoint
- if user has an account with that PhoneNumber we login the user if not We Create an account fo user

Second users can Create account or login with username and password
- To create account users should send Username, Password and information like Firstname, lastname, email, ...
- Login with username is available too
- users send username/password to login and recive JWT
## 

### JWT
JWT tokens will be generated every time a user login

users can use refresh tokens to get new access tokens and refresh token

every refresh token can be used once 

if users logout i store the token in Redis blacklist with TTL of expire time of the token
## 

### Middlewares
- CORS middleware
- Limiter middleware is added so users can't send more than 2 requests per second (i store users ip)
- OTP Limiter so users request 1 otp every 120 seconds with one IP
- Logger request endpoint requst bodySize and ... are logged so you can see and access them in kibana and elastic
- Auth this middleware checks for jwt Token in header and prevent unauthorized users to use api
## 

### Validation
- password validation is used
- phone nu,ber validation will chaeck that only Iraninan ohoneNumbers can use the service

## 

### Logging
i implemented a **Logger interface** so changing the logger wont be so hard

**ZapLogger** and **ZeroLogger** is used in project

default logger is ZapLogger but you can change the logger in cfg Files 

logs will be saved under src/logs/logs.go

then Using **ELK** stack they are availble in elastic so you can filter them and sort logs
## 

### Handlers and views
i used a GinEngin for routing the api

business logic is implemented in services files

i created a Base **Generic** service for CRUD and filtering

all end points are documented

Users can get a car-model that comtains information like:
- built year
- price history for every model
- Color
- properties
- Price
- Comapny
- Images
- User comments
users can add comments to cars

every car has a galary of images 
## 

### Response 

the api response is unified so every request will return the same JSON format

3 types of Respose is possible
- GenerateBaseResponse
- GenerateBaseResponseWithError
- GenerateBaseResponseWithValidationError
## 

##### filtering

users can **paginate/filter/Sort** the responses

```{
  "filter": {
    "additionalProp1": {
      "filterType": "string",
      "from": "string",
      "to": "string",
      "type": "string"
    }
  },
  "page_number": 0,
  "page_size": 0,
  "sort": [
    {
      "colId": "string",
      "sort": "string"
    }
  ]
}
```
## 

### part of the endpoints in swagger 
**total 70 different endpoints**
![image](https://github.com/Arshia-Izadyar/Car-sale-api/assets/110552657/62a7d150-ef6f-4362-95c4-bbad5b6ec157)














