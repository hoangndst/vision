


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

###  blog

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/blogs | [create blog](#create-blog) | Create blog |
| DELETE | /api/v1/blogs/{id} | [delete blog](#delete-blog) | Delete blog |
| GET | /api/v1/blogs/{id} | [get blog](#get-blog) | Get blog |
| GET | /api/v1/blogs | [list blog](#list-blog) | List blogs |
| POST | /api/v1/blogs/sync | [sync blogs](#sync-blogs) | Sync blogs |
| PUT | /api/v1/blogs/{id} | [update blog](#update-blog) | Update blog |
  


###  user

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /api/v1/users | [create user](#create-user) | Create user |
| DELETE | /api/v1/users/{id} | [delete user](#delete-user) | Delete user |
| GET | /api/v1/users/{id} | [get user](#get-user) | Get user |
| GET | /api/v1/users | [list user](#list-user) | List users |
| PUT | /api/v1/users/{id} | [update user](#update-user) | Update user |
  


## Paths

### <span id="create-blog"></span> Create blog (*createBlog*)

```
POST /api/v1/blogs
```

Create a new blog

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| blog | `body` | [RequestCreateBlogRequest](#request-create-blog-request) | `models.RequestCreateBlogRequest` | | ✓ | | Created blog |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-blog-200) | OK | Success |  | [schema](#create-blog-200-schema) |
| [400](#create-blog-400) | Bad Request | Bad Request |  | [schema](#create-blog-400-schema) |
| [401](#create-blog-401) | Unauthorized | Unauthorized |  | [schema](#create-blog-401-schema) |
| [404](#create-blog-404) | Not Found | Not Found |  | [schema](#create-blog-404-schema) |
| [429](#create-blog-429) | Too Many Requests | Too Many Requests |  | [schema](#create-blog-429-schema) |
| [500](#create-blog-500) | Internal Server Error | Internal Server Error |  | [schema](#create-blog-500-schema) |

#### Responses


##### <span id="create-blog-200"></span> 200 - Success
Status: OK

###### <span id="create-blog-200-schema"></span> Schema
   
  

[CreateBlogOKBody](#create-blog-o-k-body)

##### <span id="create-blog-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="create-blog-400-schema"></span> Schema
   
  

any

##### <span id="create-blog-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="create-blog-401-schema"></span> Schema
   
  

any

##### <span id="create-blog-404"></span> 404 - Not Found
Status: Not Found

###### <span id="create-blog-404-schema"></span> Schema
   
  

any

##### <span id="create-blog-429"></span> 429 - Too Many Requests
Status: Too Many Requests

###### <span id="create-blog-429-schema"></span> Schema
   
  

any

##### <span id="create-blog-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="create-blog-500-schema"></span> Schema
   
  

any

###### Inlined models

**<span id="create-blog-o-k-body"></span> CreateBlogOKBody**


  


* composed type [HandlerResponse](#handler-response)
* inlined member (*createBlogOKBodyAO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | [EntityBlog](#entity-blog)| `models.EntityBlog` |  | |  |  |



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
   
  

[CreateUserOKBody](#create-user-o-k-body)

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

###### Inlined models

**<span id="create-user-o-k-body"></span> CreateUserOKBody**


  


* composed type [HandlerResponse](#handler-response)
* inlined member (*createUserOKBodyAO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | [EntityUser](#entity-user)| `models.EntityUser` |  | |  |  |



### <span id="delete-blog"></span> Delete blog (*deleteBlog*)

```
DELETE /api/v1/blogs/{id}
```

Delete specified blog by ID

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Blog ID |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-blog-200) | OK | Success |  | [schema](#delete-blog-200-schema) |
| [400](#delete-blog-400) | Bad Request | Bad Request |  | [schema](#delete-blog-400-schema) |
| [401](#delete-blog-401) | Unauthorized | Unauthorized |  | [schema](#delete-blog-401-schema) |
| [404](#delete-blog-404) | Not Found | Not Found |  | [schema](#delete-blog-404-schema) |
| [429](#delete-blog-429) | Too Many Requests | Too Many Requests |  | [schema](#delete-blog-429-schema) |
| [500](#delete-blog-500) | Internal Server Error | Internal Server Error |  | [schema](#delete-blog-500-schema) |

#### Responses


##### <span id="delete-blog-200"></span> 200 - Success
Status: OK

###### <span id="delete-blog-200-schema"></span> Schema
   
  

[DeleteBlogOKBody](#delete-blog-o-k-body)

##### <span id="delete-blog-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="delete-blog-400-schema"></span> Schema
   
  

any

##### <span id="delete-blog-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-blog-401-schema"></span> Schema
   
  

any

##### <span id="delete-blog-404"></span> 404 - Not Found
Status: Not Found

###### <span id="delete-blog-404-schema"></span> Schema
   
  

any

##### <span id="delete-blog-429"></span> 429 - Too Many Requests
Status: Too Many Requests

###### <span id="delete-blog-429-schema"></span> Schema
   
  

any

##### <span id="delete-blog-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="delete-blog-500-schema"></span> Schema
   
  

any

###### Inlined models

**<span id="delete-blog-o-k-body"></span> DeleteBlogOKBody**


  


* composed type [HandlerResponse](#handler-response)
* inlined member (*deleteBlogOKBodyAO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | string| `string` |  | |  |  |



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
   
  

[DeleteUserOKBody](#delete-user-o-k-body)

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

###### Inlined models

**<span id="delete-user-o-k-body"></span> DeleteUserOKBody**


  


* composed type [HandlerResponse](#handler-response)
* inlined member (*deleteUserOKBodyAO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | string| `string` |  | |  |  |



### <span id="get-blog"></span> Get blog (*getBlog*)

```
GET /api/v1/blogs/{id}
```

Get blog information by blog ID

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Blog ID |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-blog-200) | OK | Success |  | [schema](#get-blog-200-schema) |
| [400](#get-blog-400) | Bad Request | Bad Request |  | [schema](#get-blog-400-schema) |
| [401](#get-blog-401) | Unauthorized | Unauthorized |  | [schema](#get-blog-401-schema) |
| [404](#get-blog-404) | Not Found | Not Found |  | [schema](#get-blog-404-schema) |
| [429](#get-blog-429) | Too Many Requests | Too Many Requests |  | [schema](#get-blog-429-schema) |
| [500](#get-blog-500) | Internal Server Error | Internal Server Error |  | [schema](#get-blog-500-schema) |

#### Responses


##### <span id="get-blog-200"></span> 200 - Success
Status: OK

###### <span id="get-blog-200-schema"></span> Schema
   
  

[GetBlogOKBody](#get-blog-o-k-body)

##### <span id="get-blog-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="get-blog-400-schema"></span> Schema
   
  

any

##### <span id="get-blog-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-blog-401-schema"></span> Schema
   
  

any

##### <span id="get-blog-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-blog-404-schema"></span> Schema
   
  

any

##### <span id="get-blog-429"></span> 429 - Too Many Requests
Status: Too Many Requests

###### <span id="get-blog-429-schema"></span> Schema
   
  

any

##### <span id="get-blog-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="get-blog-500-schema"></span> Schema
   
  

any

###### Inlined models

**<span id="get-blog-o-k-body"></span> GetBlogOKBody**


  


* composed type [HandlerResponse](#handler-response)
* inlined member (*getBlogOKBodyAO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | [EntityBlog](#entity-blog)| `models.EntityBlog` |  | |  |  |



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
   
  

[GetUserOKBody](#get-user-o-k-body)

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

###### Inlined models

**<span id="get-user-o-k-body"></span> GetUserOKBody**


  


* composed type [HandlerResponse](#handler-response)
* inlined member (*getUserOKBodyAO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | [EntityUser](#entity-user)| `models.EntityUser` |  | |  |  |



### <span id="list-blog"></span> List blogs (*listBlog*)

```
GET /api/v1/blogs
```

List all blogs

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-blog-200) | OK | Success |  | [schema](#list-blog-200-schema) |
| [400](#list-blog-400) | Bad Request | Bad Request |  | [schema](#list-blog-400-schema) |
| [401](#list-blog-401) | Unauthorized | Unauthorized |  | [schema](#list-blog-401-schema) |
| [404](#list-blog-404) | Not Found | Not Found |  | [schema](#list-blog-404-schema) |
| [429](#list-blog-429) | Too Many Requests | Too Many Requests |  | [schema](#list-blog-429-schema) |
| [500](#list-blog-500) | Internal Server Error | Internal Server Error |  | [schema](#list-blog-500-schema) |

#### Responses


##### <span id="list-blog-200"></span> 200 - Success
Status: OK

###### <span id="list-blog-200-schema"></span> Schema
   
  

[ListBlogOKBody](#list-blog-o-k-body)

##### <span id="list-blog-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="list-blog-400-schema"></span> Schema
   
  

any

##### <span id="list-blog-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="list-blog-401-schema"></span> Schema
   
  

any

##### <span id="list-blog-404"></span> 404 - Not Found
Status: Not Found

###### <span id="list-blog-404-schema"></span> Schema
   
  

any

##### <span id="list-blog-429"></span> 429 - Too Many Requests
Status: Too Many Requests

###### <span id="list-blog-429-schema"></span> Schema
   
  

any

##### <span id="list-blog-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="list-blog-500-schema"></span> Schema
   
  

any

###### Inlined models

**<span id="list-blog-o-k-body"></span> ListBlogOKBody**


  


* composed type [HandlerResponse](#handler-response)
* inlined member (*listBlogOKBodyAO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | [][EntityBlog](#entity-blog)| `[]*models.EntityBlog` |  | |  |  |



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
   
  

[ListUserOKBody](#list-user-o-k-body)

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

###### Inlined models

**<span id="list-user-o-k-body"></span> ListUserOKBody**


  


* composed type [HandlerResponse](#handler-response)
* inlined member (*listUserOKBodyAO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | [][EntityUser](#entity-user)| `[]*models.EntityUser` |  | |  |  |



### <span id="sync-blogs"></span> Sync blogs (*syncBlogs*)

```
POST /api/v1/blogs/sync
```

Sync blogs information from GitHub repository

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#sync-blogs-200) | OK | Success |  | [schema](#sync-blogs-200-schema) |
| [400](#sync-blogs-400) | Bad Request | Bad Request |  | [schema](#sync-blogs-400-schema) |
| [401](#sync-blogs-401) | Unauthorized | Unauthorized |  | [schema](#sync-blogs-401-schema) |
| [404](#sync-blogs-404) | Not Found | Not Found |  | [schema](#sync-blogs-404-schema) |
| [429](#sync-blogs-429) | Too Many Requests | Too Many Requests |  | [schema](#sync-blogs-429-schema) |
| [500](#sync-blogs-500) | Internal Server Error | Internal Server Error |  | [schema](#sync-blogs-500-schema) |

#### Responses


##### <span id="sync-blogs-200"></span> 200 - Success
Status: OK

###### <span id="sync-blogs-200-schema"></span> Schema
   
  

[SyncBlogsOKBody](#sync-blogs-o-k-body)

##### <span id="sync-blogs-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="sync-blogs-400-schema"></span> Schema
   
  

any

##### <span id="sync-blogs-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="sync-blogs-401-schema"></span> Schema
   
  

any

##### <span id="sync-blogs-404"></span> 404 - Not Found
Status: Not Found

###### <span id="sync-blogs-404-schema"></span> Schema
   
  

any

##### <span id="sync-blogs-429"></span> 429 - Too Many Requests
Status: Too Many Requests

###### <span id="sync-blogs-429-schema"></span> Schema
   
  

any

##### <span id="sync-blogs-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="sync-blogs-500-schema"></span> Schema
   
  

any

###### Inlined models

**<span id="sync-blogs-o-k-body"></span> SyncBlogsOKBody**


  


* composed type [HandlerResponse](#handler-response)
* inlined member (*syncBlogsOKBodyAO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | string| `string` |  | |  |  |



### <span id="update-blog"></span> Update blog (*updateBlog*)

```
PUT /api/v1/blogs/{id}
```

Update the specified blog

#### Consumes
  * application/json

#### Produces
  * application/json

#### Security Requirements
  * BasicAuth

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | Blog ID |
| blog | `body` | [RequestUpdateBlogRequest](#request-update-blog-request) | `models.RequestUpdateBlogRequest` | | ✓ | | Updated blog |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-blog-200) | OK | Success |  | [schema](#update-blog-200-schema) |
| [400](#update-blog-400) | Bad Request | Bad Request |  | [schema](#update-blog-400-schema) |
| [401](#update-blog-401) | Unauthorized | Unauthorized |  | [schema](#update-blog-401-schema) |
| [404](#update-blog-404) | Not Found | Not Found |  | [schema](#update-blog-404-schema) |
| [429](#update-blog-429) | Too Many Requests | Too Many Requests |  | [schema](#update-blog-429-schema) |
| [500](#update-blog-500) | Internal Server Error | Internal Server Error |  | [schema](#update-blog-500-schema) |

#### Responses


##### <span id="update-blog-200"></span> 200 - Success
Status: OK

###### <span id="update-blog-200-schema"></span> Schema
   
  

[UpdateBlogOKBody](#update-blog-o-k-body)

##### <span id="update-blog-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="update-blog-400-schema"></span> Schema
   
  

any

##### <span id="update-blog-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-blog-401-schema"></span> Schema
   
  

any

##### <span id="update-blog-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-blog-404-schema"></span> Schema
   
  

any

##### <span id="update-blog-429"></span> 429 - Too Many Requests
Status: Too Many Requests

###### <span id="update-blog-429-schema"></span> Schema
   
  

any

##### <span id="update-blog-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="update-blog-500-schema"></span> Schema
   
  

any

###### Inlined models

**<span id="update-blog-o-k-body"></span> UpdateBlogOKBody**


  


* composed type [HandlerResponse](#handler-response)
* inlined member (*updateBlogOKBodyAO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | [EntityBlog](#entity-blog)| `models.EntityBlog` |  | |  |  |



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
   
  

[UpdateUserOKBody](#update-user-o-k-body)

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

###### Inlined models

**<span id="update-user-o-k-body"></span> UpdateUserOKBody**


  


* composed type [HandlerResponse](#handler-response)
* inlined member (*updateUserOKBodyAO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| data | [EntityUser](#entity-user)| `models.EntityUser` |  | |  |  |



## Models

### <span id="entity-blog"></span> entity.Blog


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creation_timestamp | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| path | string| `string` |  | |  |  |
| raw_data | string| `string` |  | |  |  |
| update_timestamp | string| `string` |  | |  |  |



### <span id="entity-user"></span> entity.User


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creation_timestamp | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| email | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| password | string| `string` |  | |  |  |
| update_timestamp | string| `string` |  | |  |  |
| username | string| `string` |  | |  |  |



### <span id="handler-duration"></span> handler.Duration


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| handler.Duration | integer| int64 | |  |  |



### <span id="handler-response"></span> handler.Response


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| costTime | [HandlerResponse](#handler-response)| `HandlerResponse` |  | | Time taken for the request. |  |
| data | [interface{}](#interface)| `interface{}` |  | | Data payload. |  |
| endTime | string| `string` |  | | Request end time. |  |
| message | string| `string` |  | | Descriptive message. |  |
| startTime | string| `string` |  | | Request start time. |  |
| success | boolean| `bool` |  | | Indicates success status. |  |
| traceID | string| `string` |  | | Trace identifier. |  |



### <span id="request-create-blog-request"></span> request.CreateBlogRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| description | string| `string` |  | | Description is a human-readable description of the blog. |  |
| path | string| `string` | ✓ | | Path is the path of the blog. |  |
| raw_data | string| `string` | ✓ | | RawData is the raw data of the blog. |  |



### <span id="request-create-user-request"></span> request.CreateUserRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| name | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |
| username | string| `string` | ✓ | |  |  |



### <span id="request-update-blog-request"></span> request.UpdateBlogRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| description | string| `string` |  | | Description is a human-readable description of the blog. |  |
| path | string| `string` | ✓ | | Path is the path of the blog. |  |
| raw_data | string| `string` | ✓ | | RawData is the raw data of the blog. |  |



### <span id="request-update-user-request"></span> request.UpdateUserRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | string| `string` | ✓ | |  |  |
| name | string| `string` | ✓ | |  |  |
| password | string| `string` | ✓ | |  |  |


