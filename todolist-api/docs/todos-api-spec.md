# TODO API

## Create: POST ``/api/todos``
-  Request sample
    ```json
   ### HEADER
   Content-Type: application/json
   Authorization: Bearer (access_token)
   
   ### BODY
    {
        "title": "Buy groceries",
        "description": "Buy milk, eggs, and bread"
    }
    ```
-  Response sample
    ```json
    {
      "status": 201,
      "data": {
           "id": "f8c3de3d-1fea-4d7c-a8b0-29f63c4c3454",
           "id_user":"5096125b-aba3-4c62-bd02-b15c9be8c5f2", 
           "title": "Buy groceries",
           "description": "Buy milk, eggs, and bread",
           "status": "in-progress",
           "created_at": 1727589409,
           "updated_at": 1727589409
      }  
    }
   ```
   
-  Response sample (client error)
    ```json
    {
      "status": 401,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Unauthorized"
    }
    ```

-  Response sample (client error)
    ```json
    {
      "status": 400,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Validation errors",
      "errors": [
        {
          "field": "Title",
          "rejected_value": string,
          "message": // "REQUIRED", "TO_LONG"
        },
        {
          "field": "Description",
          "rejected_value": string,
          "message": // "REQUIRED", "TO_LONG"
        },
      ]
    }
    ```

-  Response sample (server error)
    ```json
    {
      "status": 500,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Internal server error"
    }
    ```

---
## Update: PUT ``/api/todos/:id``
-  Request sample
    ```json
   ### Query parameter
   mark=done (optional)
   
   ### HEADER
   Content-Type: application/json
   Authorization: Bearer (access_token)
   
   ### BODY
    {
        "title": "Buy groceries",
        "description": "Buy milk, eggs, and bread, and cheese"
    }
    ```
-  Response sample
    ```json
    {
      "status": 200,
      "data": {
           "id": "f8c3de3d-1fea-4d7c-a8b0-29f63c4c3454",
           "id_user":"5096125b-aba3-4c62-bd02-b15c9be8c5f2", 
           "title": "Buy groceries",
           "description": "Buy milk, eggs, and bread, and scheese",
           "status": "in-progress",
           "created_at": 1727589409,
           "updated_at": 1727589409
      }  
    }
   ```

-  Response sample (client error)
    ```json
    {
      "status": 400,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Validation errors",
      "errors": [
        {
          "field": "Title",
          "rejected_value": string,
          "message": // "REQUIRED", "TO_LONG"
        },
        {
          "field": "Description",
          "rejected_value": string,
          "message": // "REQUIRED", "TO_LONG"
        },
      ]
    }
   
    {
      "status": 400,
      "timestamp": "22-07-2022 06:49:25",
      "message": "query parameter must be value 'done'"
    }
    ```

-  Response sample (client error)
    ```json
    {
      "status": 401,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Unauthorized"
    }
    ```

-  Response sample (client error)
    ```json
    {
      "status": 403,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Forbidden"
    }
    ```

-  Response sample (server error)
    ```json
    {
      "status": 500,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Internal server error"
    }
    ```


---
## Delete: DELETE ``/api/todos/:id``
-  Request sample
   ```json
   ### HEADER
   Authorization: Bearer (access_token)
   ```

-  Response sample: ``204``

-  Response sample (client error)
   ```json
   {
      "status": 401,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Unauthorized"
   }
   ```

-  Response sample (client error)
   ```json
   {
      "status": 404,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Not found"
   }
   ```

-  Response sample (server error)
    ```json
    {
      "status": 500,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Internal server error"
    }
    ```

---
## List: GET ``/api/todos``
- Request sample  
   ```json
  ### Query Parameter filtering
  status=string (in-progress or done) default in-progress
  
  ### Query Parameter sorting
  sort=string (created_at) default created_at
  order=string (asc, desc) default asc
  
  ### Query Parameter pagination
  page=integer, default 1
  limit=integer, default 20
  
   ### HEADER
   Content-Type: application/json
   Authorization: Baerer access_token
   ```

-  Response sample
   ```json
   {
      "status": 200,
      "data": [
          {
            "id": "f8c3de3d-1fea-4d7c-a8b0-29f63c4c3454",
            "id_user":"5096125b-aba3-4c62-bd02-b15c9be8c5f2", 
            "title": "Buy groceries",
            "description": "Buy milk, eggs, and bread",
            "status": "in-progress",
            "created_at": 1727589409,
            "updated_at": 1727589409
          },
          {...},
          {...},
          {...},
          {...},
          {...}
      ],     
      "paging": {
          "total_data": 100,
          "total_page": 5,
          "page": 1,
          "limit": 20
      }
   }
   ```

-  Response sample (client error)
   ```json
   {
      "status": 400,
      "errors": {
         "order": ["MUST_BE_ALPHABET", "MUST_BE_ASC", "MUST_BE_DESC"], 
         "status": ["MUST_BE_ALPHABET", "MUST_BE_DONE", "MUST_BE_IN-PROGRESS"],
         "page": ["MUST_BE_NUMERIC", "MUST_BE_UNSIGNED"],
         "limit": ["MUST_BE_NUMERIC", "MUST_BE_UNSIGNED"],
      }
   }
   ```

-  Response sample (server error)
    ```json
    {
      "status": 500,
      "errors": "Internal server error"
    }
    ```