


# Vision Backend API
This is the Vision Backend API.
  

## Informations

### Version

v0.1.0

### License

[Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### Contact

  

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## Access control

### Security Schemes

#### BasicAuth



> **Type**: basic

## All endpoints

###  user

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/users | [create user](#create-user) | Create user |
| DELETE | /api/v1/users/{id} | [delete user](#delete-user) | Delete user |
| GET | /api/v1/users/{id} | [get user](#get-user) | Get user |
| GET | /api/v1/users | [list user](#list-user) | List users |
| PUT | /api/v1/users/{id} | [update user](#update-user) | Update user |
  


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

#### Security Requirements
  * BasicAuth

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

### <span id="delete-user"></span> Delete user (*deleteUser*)

```
DELETE /api/v1/users/{id}
```

Delete specified user by ID

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-user-200) | OK | Success |  | [schema](#delete-user-200-schema) |
| [400](#delete-user-400) | Bad Request | Bad Request |  | [schema](#delete-user-400-schema) |
| [401](#delete-user-401) | Unauthorized | Unauthorized |  | [schema](#delete-user-401-schema) |
| [404](#delete-user-404) | Not Found | Not Found |  | [schema](#delete-user-404-schema) |
| [429](#delete-user-429) | Too Many Requests | Too Many Requests |  | [schema](#delete-user-429-schema) |
| [500](#delete-user-500) | Internal Server Error | Internal Server Error |  | [schema](#delete-user-500-schema) |

#### Responses


##### <span id="delete-user-200"></span> 200 - Success
Status: OK

###### <span id="delete-user-200-schema"></span> Schema
   
  



##### <span id="delete-user-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="delete-user-400-schema"></span> Schema
   
  

any

##### <span id="delete-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-user-401-schema"></span> Schema
   
  

any

##### <span id="delete-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="delete-user-404-schema"></span> Schema
   
  

any

##### <span id="delete-user-429"></span> 429 - Too Many Requests
Status: Too Many Requests

###### <span id="delete-user-429-schema"></span> Schema
   
  

any

##### <span id="delete-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="delete-user-500-schema"></span> Schema
   
  

any

### <span id="get-user"></span> Get user (*getUser*)

```
GET /api/v1/users/{id}
```

Get user information by user ID

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-user-200) | OK | Success |  | [schema](#get-user-200-schema) |
| [400](#get-user-400) | Bad Request | Bad Request |  | [schema](#get-user-400-schema) |
| [401](#get-user-401) | Unauthorized | Unauthorized |  | [schema](#get-user-401-schema) |
| [404](#get-user-404) | Not Found | Not Found |  | [schema](#get-user-404-schema) |
| [429](#get-user-429) | Too Many Requests | Too Many Requests |  | [schema](#get-user-429-schema) |
| [500](#get-user-500) | Internal Server Error | Internal Server Error |  | [schema](#get-user-500-schema) |

#### Responses


##### <span id="get-user-200"></span> 200 - Success
Status: OK

###### <span id="get-user-200-schema"></span> Schema
   
  

[EntityUser](#entity-user)

##### <span id="get-user-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="get-user-400-schema"></span> Schema
   
  

any

##### <span id="get-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-user-401-schema"></span> Schema
   
  

any

##### <span id="get-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-user-404-schema"></span> Schema
   
  

any

##### <span id="get-user-429"></span> 429 - Too Many Requests
Status: Too Many Requests

###### <span id="get-user-429-schema"></span> Schema
   
  

any

##### <span id="get-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="get-user-500-schema"></span> Schema
   
  

any

### <span id="list-user"></span> List users (*listUser*)

```
GET /api/v1/users
```

List all users

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-user-200) | OK | Success |  | [schema](#list-user-200-schema) |
| [400](#list-user-400) | Bad Request | Bad Request |  | [schema](#list-user-400-schema) |
| [401](#list-user-401) | Unauthorized | Unauthorized |  | [schema](#list-user-401-schema) |
| [404](#list-user-404) | Not Found | Not Found |  | [schema](#list-user-404-schema) |
| [429](#list-user-429) | Too Many Requests | Too Many Requests |  | [schema](#list-user-429-schema) |
| [500](#list-user-500) | Internal Server Error | Internal Server Error |  | [schema](#list-user-500-schema) |

#### Responses


##### <span id="list-user-200"></span> 200 - Success
Status: OK

###### <span id="list-user-200-schema"></span> Schema
   
  

[EntityUser](#entity-user)

##### <span id="list-user-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="list-user-400-schema"></span> Schema
   
  

any

##### <span id="list-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="list-user-401-schema"></span> Schema
   
  

any

##### <span id="list-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="list-user-404-schema"></span> Schema
   
  

any

##### <span id="list-user-429"></span> 429 - Too Many Requests
Status: Too Many Requests

###### <span id="list-user-429-schema"></span> Schema
   
  

any

##### <span id="list-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-user-500-schema"></span> Schema
   
  

any

### <span id="update-user"></span> Update user (*updateUser*)

```
PUT /api/v1/users/{id}
```

Update the specified user

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | User ID |
| user | `body` | [RequestUpdateUserRequest](#request-update-user-request) | `models.RequestUpdateUserRequest` | | ✓ | | Updated user |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-user-200) | OK | Success |  | [schema](#update-user-200-schema) |
| [400](#update-user-400) | Bad Request | Bad Request |  | [schema](#update-user-400-schema) |
| [401](#update-user-401) | Unauthorized | Unauthorized |  | [schema](#update-user-401-schema) |
| [404](#update-user-404) | Not Found | Not Found |  | [schema](#update-user-404-schema) |
| [429](#update-user-429) | Too Many Requests | Too Many Requests |  | [schema](#update-user-429-schema) |
| [500](#update-user-500) | Internal Server Error | Internal Server Error |  | [schema](#update-user-500-schema) |

#### Responses


##### <span id="update-user-200"></span> 200 - Success
Status: OK

###### <span id="update-user-200-schema"></span> Schema
   
  

[EntityUser](#entity-user)

##### <span id="update-user-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="update-user-400-schema"></span> Schema
   
  

any

##### <span id="update-user-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-user-401-schema"></span> Schema
   
  

any

##### <span id="update-user-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-user-404-schema"></span> Schema
   
  

any

##### <span id="update-user-429"></span> 429 - Too Many Requests
Status: Too Many Requests

###### <span id="update-user-429-schema"></span> Schema
   
  

any

##### <span id="update-user-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="update-user-500-schema"></span> Schema
   
  

any

## Models

### <span id="entity-user"></span> entity.User


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creation_timestamp | string| `string` |  | |  |  |
| email | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| password | string| `string` |  | |  |  |
| update_timestamp | string| `string` |  | |  |  |
| username | string| `string` |  | |  |  |



### <span id="request-create-user-request"></span> request.CreateUserRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| name | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |
| username | string| `string` | ✓ | |  |  |



### <span id="request-update-user-request"></span> request.UpdateUserRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| name | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |


