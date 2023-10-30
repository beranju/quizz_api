# Api Spec

Base Url => `https://quiz-app.osc-fr1.scalingo.io`

## Authentication

All endpoint must use this authentication except login and register

Request :
- Header :
    - Authorization : "Bearer YOUR_TOKEN"

##  Register

Request :
- Method : POST
- Endpoint : `/users/register`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "name" : "string",
    "email" : "string, unique",
    "password" : "string",
    "image_url" : "string"
}
```

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": {
        "ID": "int, unique",
        "CreatedAt": "time.Time",
        "UpdatedAt": "time.Time",
        "DeletedAt": "time.Time",
        "name": "string",
        "email": "string, unique",
        "password": "string, hash",
        "image_url": "string"
    }
}
```

## Login

Request :
- Method : POST
- Endpoint : `/users/login`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "email" : "string, unique",
    "password" : "string",
}
```

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": {
        "ID": "int, unique",
        "CreatedAt": "time.Time",
        "UpdatedAt": "time.Time",
        "name": "string",
        "email": "string, unique",
        "image_url": "string",
        "token": "string unique"
    }
}
```

## Update User

Request :
- Method : PUT
- Endpoint : `/users/update/{id_user}`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "name" : "string",
    "email" : "string, unique",
    "password" : "string",
    "image_url" : "string"
}
```

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": "null"
}
```

## Add Quiz

Request :
- Method : POST
- Endpoint : `/quizzes/`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "quiz_name" : "string",
    "description": "string",
    "duration" : "int"
}
```

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": {
        "ID": "int, unique",
        "CreatedAt": "time.Time",
        "UpdatedAt": "time.Time",
        "DeletedAt": "time.Time",
        "quiz_name" : "string",
        "description": "string",
        "duration" : "int"
    }
}
```

## Quiz Detail

Request :
- Method : GET
- Endpoint : `/quizzes/{quiz_id}`
- Header :
    - Content-Type: application/json
    - Accept: application/json

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": {
        "ID": "int, unique",
        "CreatedAt": "time.Time",
        "UpdatedAt": "time.Time",
        "DeletedAt": "time.Time",
        "quiz_name" : "string",
        "description": "string",
        "duration" : "int",
        "Question" : "[]"
    }
}
```

## All Quiz

Request :
- Method : GET
- Endpoint : `/quizzes`
- Header :
    - Content-Type: application/json
    - Accept: application/json

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": [
        {
            "ID": "int, unique",
            "CreatedAt": "time.Time",
            "UpdatedAt": "time.Time",
            "quiz_name" : "string",
            "description": "string",
            "duration" : "int",
        },
        {
            "ID": "int, unique",
            "CreatedAt": "time.Time",
            "UpdatedAt": "time.Time",
            "quiz_name" : "string",
            "description": "string",
            "duration" : "int",
        }
    ]
}
```

## Update Quiz

Request :
- Method : PUT
- Endpoint : `/quizzes/{quiz_id}`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "quiz_name" : "string",
    "description": "string",
    "duration" : "int"
}
```

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": {
        "ID": "int, unique",
        "CreatedAt": "time.Time",
        "UpdatedAt": "time.Time",
        "quiz_name" : "string",
        "description": "string",
        "duration" : "int",
    }
}
```

## Delete Quiz

Request :
- Method : DELETE
- Endpoint : `/quizzes/{quiz_id}`
- Header :
    - Content-Type: application/json
    - Accept: application/json

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": "null"
}
```

## Add Question

Request :
- Method : POST
- Endpoint : `/quizzes/{quiz_id}/questions`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
  "quiz_id":"int, unique",
  "text":"string",
  "score":"int",
  "answer":"int, index of the answer",
  "options_1":"string",
  "options_2":"string",
  "options_3":"string",
  "options_4":"string",
  "image_url":"string"
}
```

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": {
        "ID": "int, unique",
        "CreatedAt": "time.Time",
        "UpdatedAt": "time.Time",
        "quiz_id":"int, unique",
        "text":"string",
        "score":"int",
        "answer":"int, index of the answer",
        "options_1":"string",
        "options_2":"string",
        "options_3":"string",
        "options_4":"string",
        "image_url":"string"
    }
}
```

## All Question

Request :
- Method : GET
- Endpoint : `/quizzes/{quiz_id}/questions`
- Header :
    - Content-Type: application/json
    - Accept: application/json

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": [
        {
            "ID": "int, unique",
            "CreatedAt": "time.Time",
            "UpdatedAt": "time.Time",
            "quiz_id":"int, unique",
            "text":"string",
            "score":"int",
            "answer":"int, index of the answer",
            "options_1":"string",
            "options_2":"string",
            "options_3":"string",
            "options_4":"string",
            "image_url":"string"
        },
        {
            "ID": "int, unique",
            "CreatedAt": "time.Time",
            "UpdatedAt": "time.Time",
            "quiz_id":"int, unique",
            "text":"string",
            "score":"int",
            "answer":"int, index of the answer",
            "options_1":"string",
            "options_2":"string",
            "options_3":"string",
            "options_4":"string",
            "image_url":"string"
        }
    ]
}
```

## Detail Question

Request :
- Method : GET
- Endpoint : `/quizzes/{quiz_id}/questions/{question_id}`
- Header :
    - Content-Type: application/json
    - Accept: application/json

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": {
        "ID": "int, unique",
        "CreatedAt": "time.Time",
        "UpdatedAt": "time.Time",
        "quiz_id":"int, unique",
        "text":"string",
        "score":"int",
        "answer":"int, index of the answer",
        "options_1":"string",
        "options_2":"string",
        "options_3":"string",
        "options_4":"string",
        "image_url":"string"
    }
}
```

## Update Question

Request :
- Method : PUT
- Endpoint : `/quizzes/{quiz_id}/questions/{question_id}`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
  "quiz_id":"int, unique",
  "text":"string",
  "score":"int",
  "answer":"int, index of the answer",
  "options_1":"string",
  "options_2":"string",
  "options_3":"string",
  "options_4":"string",
  "image_url":"string"
}
```

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": {
        "ID": "int, unique",
        "CreatedAt": "time.Time",
        "UpdatedAt": "time.Time",
        "DeleteAt": "time.Time",
        "quiz_id":"int, unique",
        "text":"string",
        "score":"int",
        "answer":"int, index of the answer",
        "options_1":"string",
        "options_2":"string",
        "options_3":"string",
        "options_4":"string",
        "image_url":"string"
    }
}
```

## Delete Question
Request :
- Method : DELETE
- Endpoint : `/quizzes/{quiz_id}/questions/{question_id}`
- Header :
    - Content-Type: application/json
    - Accept: application/json

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": "null"
}
```

## Add Result

Request :
- Method : POST
- Endpoint : `/quizzes/{quiz_id}/result`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
    "user_id":"int, unique",
    "quiz_id":"int, unique",
    "user_score":"int"
}
```

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": {
        "ID": "int, unique",
        "CreatedAt": "time.Time",
        "UpdatedAt": "time.Time",
        "user_id":"int, unique",
        "quiz_id":"int, unique",
        "user_score":"int"
    }
}
```

## Get Result

Request :
- Method : GET
- Endpoint : `/quizzes/{quiz_id}/result`
- Header :
    - Content-Type: application/json
    - Accept: application/json

Response :

```json
{
    "status": "string",
    "message": "string",
    "data": {
        "ID": "int, unique",
        "CreatedAt": "time.Time",
        "UpdatedAt": "time.Time",
        "user_id":"int, unique",
        "quiz_id":"int, unique",
        "user_score":"int"
    }
}
```

# Tech Stack
- Language : Go
- ORM : GORM
- Framework : Echo
- Design : MVC
- Auth : JWT


# Host
- database : db4free.net
- code : scalingo.com