


# Vision Backend API
This is the Vision Backend API.
  

## Informations

### Version

v0.1.0

### Contact

  

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  user

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/users | [create user](#create-user) | Create user |
  


## Paths

### <span id="create-user"></span> Create user (*createUser*)

```
POST /api/v1/users
```

Create a new user

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user | `body` | [RequestCreateUserRequest](#request-create-user-request) | `models.RequestCreateUserRequest` | | ✓ | | Created user |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-user-200) | OK | Success |  | [schema](#create-user-200-schema) |
| [400](#create-user-400) | Bad Request | Bad Request |  | [schema](#create-user-400-schema) |
| [401](#create-user-401) | Unauthorized | Unauthorized |  | [schema](#create-user-401-schema) |
| [404](#create-user-404) | Not Found | Not Found |  | [schema](#create-user-404-schema) |
| [429](#create-user-429) | Too Many Requests | Too Many Requests |  | [schema](#create-user-429-schema) |
| [500](#create-user-500) | Internal Server Error | Internal Server Error |  | [schema](#create-user-500-schema) |

#### Responses


##### <span id="create-user-200"></span> 200 - Success
Status: OK

###### <span id="create-user-200-schema"></span> Schema
   
  

[EntityUser](#entity-user)

##### <span id="create-user-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="create-user-400-schema"></span> Schema
   
  

any

##### <span id="create-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="create-user-401-schema"></span> Schema
   
  

any

##### <span id="create-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="create-user-404-schema"></span> Schema
   
  

any

##### <span id="create-user-429"></span> 429 - Too Many Requests
Status: Too Many Requests

###### <span id="create-user-429-schema"></span> Schema
   
  

any

##### <span id="create-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="create-user-500-schema"></span> Schema
   
  

any

## Models

### <span id="entity-user"></span> entity.User


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creation_timestamp | string| `string` |  | |  |  |
| email | string| `string` |  | |  |  |
| id | integer| `int64` |  | |  |  |
| name | string| `string` |  | |  |  |
| password | string| `string` |  | |  |  |
| update_timestamp | string| `string` |  | |  |  |
| username | string| `string` |  | |  |  |



### <span id="request-create-user-request"></span> request.CreateUserRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| password | string| `string` |  | |  |  |
| username | string| `string` |  | |  |  |


