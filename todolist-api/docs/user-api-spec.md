# User API

## Registration

### User registration to create a new user: POST ``/api/register``

- Request sample
    ```json
    ### HEADER
    Content-Type: application/json
  
    ### BODY
    {
      "name": "John Doe",
      "email": "john@doe.com",
      "password": "johndoe"
    }
    ```

- Response sample
    ```json
    ### HEADER
    Content-Type: application/json
  
    ### BODY
    {
      "status": 201,
      "data": {
          "id": "5096125b-aba3-4c62-bd02-b15c9be8c5f2",
          "name": "John Doe",
          "email": "john@doe.com",
          "password": "$2a$12$UcCZUbSGVaJVXQSTOeocgOGdtMI8G/2CGqCi53.doiWcCgErYtPm.",
          "created_at": 1727589409
      }
    }
    ```

- Response sample (client error)
    ```json
    ### HEADER
    Content-Type: application/json
  
    ### BODY
    {
      "status": 400,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Validation errors"
      "errors": [
          {
            "field": "name",
            "rejected_value": string,
            "message": // "REQUIRED", "TO_LONG"
          },
          {
            "field": "email",
            "rejected_value": string,
            "message": // "REQUIRED", "EMAIL_FORMAT", "TO_LONG", "UNIQUE"
          },
          {
            "field": "password",
            "rejected_value": string,
            "message": // "REQUIRED", "TO_SHORT"
          },
      ]
    } 
    ```

- Response sample (client error)
    ```json
    ### HEADER
    Content-Type: application/json
  
    ### BODY
    {
      "status": 409,
      "timestamp": "22-07-2022 06:49:25",
      "message": "email already registered"
    } 
    ```

- Response sample (server error)
    ```json
    ### HEADER
    Content-Type: application/json
  
    ### BODY
    {
      "status": 500,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Internal server error" 
    }
    ```

## Login
### Login to authenticate the user and generate a token: POST ``/api/login``
-  Request Sample
    ```json
   ### HEADER
   Content-Type: application/json
   
   ### BODY
   {
        "email": "john@doe.com",
        "password": "johndoe"
    }
   ```
   
-  Response sample
    ```json
    { 
      "token": "ccc6413d-4cdd-46a5-8264-9a38fdc2639a"
   }
   ```

-  Response sample (client error)
    ```json
    ### HEADER
    Content-Type: application/json

    ### BODY
    {
      "status": 400,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Validation errors",
      "errors": [
        {
          "field": "email",
          "rejected_value": string,
          "message": // "REQUIRED", "EMAIL_FORMAT", "TO_LONG", "UNIQUE"
        },
        {
          "field": "password",
          "rejected_value": string,
          "message": // "REQUIRED", "TO_SHORT"
        },
      ]
    } 
    ```

- Response sample (server error)
    ```json
    ### HEADER
    Content-Type: application/json
  
    ### BODY
    {
      "status": 500,
      "timestamp": "22-07-2022 06:49:25",
      "message": "Internal server error" 
    }
    ```