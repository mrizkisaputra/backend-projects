# API Documentation #

api version: v0.0.1  
server: ``http://localhost:8080``

___

### Create a new blog post: POST ``/api/posts``

- **Request samples**
  ```json5
  ### request header
  Content-Type: application/json
  
  
  ### request body
  {
    "title": "My First Blog Post",
    "content": "This is the content of my first blog post.",
    "category": "Technology",
    "tags": ["Tech", "Programming"]
  }
   ```
- **Response samples**
  ```json5
  ### response header
  Content-Type: application/json

  ### response body
  {
    "status": 201,
    "data": {
      "id": 1727071310,
      "title": "My First Blog Post",
      "content": "This is the content of my first blog post.",
      "category": "Technology",
      "tags": ["Tech", "Programming"],
      "createdAt": "2021-09-01T12:00:00Z",
      "updatedAt": "2021-09-01T12:00:00Z"
    }
  }
  ```
- **Response samples (client error)**
  ```json5
  {
    "status": 400,
    "errors": {
      "title": ["REQUIRED", "TO_LONG"],
      "content": ["REQUIRED"],
      "category": ["TO_LONG", "MUST_BE_ALPHABET"]  
    },
  }
  ```

- **Response samples (server error)**
  ```json5
  {
    "status": 500,
    "errors": "Internal server error",
  }
  ```

### Get all blog post: GET ``/api/posts``
- **Request samples**
  ```json5
  ### request header
  Content-Type: application/json
  
  ### request body
   ```
- **Response samples**
  ```json5
  ### response header
  Content-Type: application/json

  ### response body
  {
    "status": 200,
    "data": [
      {
        "id": 1727071310,
        "title": "My First Blog Post",
        "content": "This is the content of my first blog post.",
        "category": "Technology",
        "tags": ["Tech", "Programming"],
        "createdAt": "2021-09-01T12:00:00Z",
        "updatedAt": "2021-09-01T12:00:00Z"
      },
      {
        "id": 1825061220,
        "title": "My Second Blog Post",
        "content": "This is the content of my second blog post.",
        "category": "Technology",
        "tags": ["Tech", "Programming"],
        "createdAt": "2021-09-01T12:30:00Z",
        "updatedAt": "2021-09-01T12:30:00Z"
      }
    ]
  }
  ```
- **Response samples (client error)**
  ```json5
  {
    "status": 404,
    "errors": "Not found",
  }
  ```

- **Response samples (server error)**
  ```json5
  {
    "status": 500,
    "errors": "Internal server error",
  }
  ```

___

### Get a single blog post: GET ``/api/posts/{id}``
- **Request samples**  
  ``/api/posts/1727071310``
  ```json5
  ### request header
  Content-Type: application/json
  
  ### request body
   ```
- **Response samples**
  ```json5
  ### response header
  Content-Type: application/json

  ### response body
  {
    "status": 200,
    "data": {
        "id": 1727071310,
        "title": "My First Blog Post",
        "content": "This is the content of my first blog post.",
        "category": "Technology",
        "tags": ["Tech", "Programming"],
        "createdAt": "2021-09-01T12:00:00Z",
        "updatedAt": "2021-09-01T12:00:00Z"
      },
  }
  ```
- **Response samples (client error)**
  ```json5
  {
    "status": 404,
    "errors": "Not found",
  }
  ```

- **Response samples (server error)**
  ```json5
  {
    "status": 500,
    "errors": "Internal server error",
  }
  ```
  
___

### Update an existing blog post: PUT ``/api/posts/{id}``
- **Request samples**  
  ```json5
  ### request header
  Content-Type: application/json
  
  ### request body
  {
    "title": "My Updated Blog Post",
    "content": "This is the updated content of my first blog post.",
    "category": "Technology",
    "tags": ["Tech", "Programming"]
  }
   ```
- **Response samples**
  ```json5
  ### response header
  Content-Type: application/json

  ### response body
  {
    "status": 200,
    "data": {
        "id": 1727071310,
        "title": "My Updated Blog Post",
        "content": "This is the updated content of my first blog post.",
        "category": "Technology",
        "tags": ["Tech", "Programming"],
        "createdAt": "2021-09-01T12:00:00Z",
        "updatedAt": "2021-09-01T12:30:00Z"
      },
  }
  ```
- **Response samples (client error)**
  ```json5
  ### blog post was not found
  {
    "status": 404,
    "errors": "Not found",
  }
  
  ### validation error
  {
    "status": 400,
    "errors": {
      "title": ["REQUIRED", "TO_LONG"],
      "content": ["REQUIRED"],
      "category": ["TO_LONG", "MUST_BE_ALPHABET"]  
    }
  }
  ```

- **Response samples (server error)**
  ```json5
  {
    "status": 500,
    "errors": "Internal server error",
  }
  ```
___

### Delete an existing blog post: DELETE ``/api/posts/{id}`` 
If the blog post was successfully deleted return ``204 No Content``

- **Response samples (client error)**
  ```json5
  {
    "status": 404,
    "errors": "Not found",
  }
  ```

- **Response samples (server error)**
  ```json5
  {
    "status": 500,
    "errors": "Internal server error",
  }
  ```
___

### Filter blog posts by a search term: GET ``/api/posts?term={term}&page={page}&limit{limit}``
- **Request samples**  
  ```json5
  ### Query Param
  term={title or category}
  page={integer} default 0
  limit={integer} default 10
  
  ```
- **Response samples**  
  ```json5
    {
    "status": 200,
    "data": [
      {...},
      {...},
      {...},
      {...}
    ],
    "paging": {
      "totalData": 4, //Ini menunjukkan total jumlah postingan yang tersedia di database tanpa memperhatikan pagination.
      "page": 0,
      "limit": 10,
      "totalPage": 1 // Ini menunjukkan total jumlah halaman yang ada berdasarkan limit dan totalData
    }
  }
  ```

- **Response samples (client error)**
  ```json5
  {
    "status": 404,
    "errors": "Not found",
  }
  ```

- **Response samples (server error)**
  ```json5
  {
    "status": 500,
    "errors": "Internal server error",
  }
  ```