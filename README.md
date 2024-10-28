# Driver Location API

A **Driver Location API** that uses location data stored in a MongoDB collection,

# Matching API Service

A **Matching API Service** that finds the nearest driver with the rider location using the Driver Location API.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Importing Initial Data](#importing-initial-data)
- [Postman Collection](#postman-collection)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/volkankocaali/bi-taksi-case.git
    cd bi-taksi-case
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up the environment variables:
   Copy the `.env.example` file to `.env` and adjust the values as needed.
    ```sh
    cp deployments/.env.example deployments/.env
    ```

## Configuration

The application uses environment variables for configuration. Below is a list of the variables you need to set in the `.env` file:


- `MONGO_INITDB_ROOT_USERNAME`: Mongo DB username.
- `MONGO_INITDB_ROOT_PASSWORD`: Mongo DB password.
- `MONGO_PORT`: Mongo DB host.
- `MONGO_LOGLEVEL`: Mongo DB log level.
- `DRIVER_LOCATION_PORT`: Driver location API port.
- `MATCHING_API_PORT`: Matching API port.
- `DRIVER_LOCATION_API_URL`: Driver location API URL.

## Usage

1. Run the application via docker:
    ```sh
    docker-compose -f deployments/docker-compose.yml -p bi-taksi-case up --build
    ```

2. The server will start on the port specified in the `deployments/.env` file.

## API Endpoints

Swagger documentation is available at the following URL:
````
http://localhost:8083/

http://localhost:8084/
````

## Importing Initial Data

Run the command to add the driver location add, You need to remove services from docker before running the command.

```bash
go run cmd/import-data/main.go
```
# Matching Api

----
### User Registration 

- **Base URL**: `http://localhost:8081/api`
- **Version**: `v1`
- **URL**: `http://localhost:8081/api/v1/`

----
- **URL**: `/register`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
        "username": "volkan.kocaali",
        "password": "bitaksi",
        "password_confirmation": "bitaksi"
    }
    ```
- **Response**:
    ```json
    {
      "status": 200,
      "message": "User logged in successfully",
      "token": "jwt-token-here",
      "user": {
          "id": "60f4b1b3b3b3b3b3b3b3b3b3",
          "username": "volkan.kocaali"
      }
    }
    ```


### User Login

- **URL**: `/login`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
        "username": "volkan.kocaali",
        "password": "bitaksi"
    }
    ```
- **Response**:
    ```json
    {
      "status": 200,
      "message": "User logged in successfully",
      "token": "jwt-token-here",
      "user": {
          "id": "60f4b1b3b3b3b3b3b3b3b3b3",
          "username": "volkan.kocaali"
      }
    }
    ```

### Match Driver

- **URL**: `/match-driver`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
        "lat": 41.0082,
        "lon": 28.9784,
        "radius": 5
    }
    ```
- **Response**:
    ```json
    {
        "id": "driver-8be0e1d5-c44f-455b-9bee-a909a076d50e",
        "latitude": 28.97953,
        "longitude": 41.015137
    }
    ```


# Driver Location Api

----
### User Registration

- **Base URL**: `http://localhost:8081/api`
- **Version**: `v1`
- **URL**: `http://localhost:8081/api/v1/`

----
- **URL**: `/register`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
        "username": "john.doe",
        "password": "bitaksi",
        "password_confirmation": "bitaksi"
    }
    ```
- **Response**:
    ```json
    {
      "status": 200,
      "message": "User logged in successfully",
      "token": "jwt-token-here",
      "user": {
          "id": "60f4b1b3b3b3b3b3b3b3b3b3",
          "username": "john.doe"
      }
    }
    ```


### User Login

- **URL**: `/login`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
        "username": "john.doe",
        "password": "bitaksi"
    }
    ```
- **Response**:
    ```json
    {
      "status": 200,
      "message": "User logged in successfully",
      "token": "jwt-token-here",
      "user": {
          "id": "60f4b1b3b3b3b3b3b3b3b3b3",
          "username": "volkan.kocaali"
      }
    }
    ```

### Driver Locations Create Update

- **URL**: `/driver-locations`
- **Method**: `POST` or `PUT`
- **Request Body**:
    ```json
    [
    {
        "driver_id": "driver-8be0e1d5-c44f-455b-9bee-a909a076d50e",
        "location": {
            "type": "Point",
            "coordinates": [28.979530, 41.015137]
        },
        "timestamp": "2024-10-24T10:00:00Z"
    },
    {
        "driver_id": "driver-8be0e1d5-c44f-455b-9bee-a909a076d50e",
        "location": {
            "type": "Point",
            "coordinates": [32.859741, 39.933365]
        },
        "timestamp": "2024-10-24T10:05:00Z"
    }
]
    ```
- **Response**:
    ```json
    {
        "status": 200,
        "message": "locations processed successfully"
    }
    ```


### Driver Locations Get

- **URL**: `/driver-locations/{driver_id}`
- **Method**: `GET`
  - **Response**:
      ```json
      {
              "driver_id": "driver-8be0e1d5-c44f-455b-9bee-a909a076d50e",
              "location": {
                  "type": "Point",
                  "coordinates": [
                      28.97953,
                      41.015137
                  ]
              },
              "timestamp": "2024-10-28T03:15:50.754Z"
        }
    ```
    
### Find Driver Within Radius

- **URL**: `/find-driver-within-radius`
- **Method**: `POST`
  - **Request Body**:
      ```json
       {
             "lat" : 28.97953,
             "lon" : 41.015137,
             "radius": 5,
             "page" : 1,
             "page_size" : 1000
        }
    ```
  - **Response**:
      ```json
      [
          {
            "id": "driver123",
            "latitude": 28.97953,
            "longitude": 41.015137
          }
      ]
      ```

### Default Credentials

#### Matching Api Service
```yaml
username: volkan.kocaali
password: bitaksi
``` 

#### Driver Location Api Service
```yaml
username: john.doe
password: bitaksi
``` 

## Postman Collection

You can download the Postman collection for this project [here](./bi-taksi-case.postman_collection.json).


## License

This project is licensed under the MIT License.
