package matching_api

import "github.com/swaggo/swag"

var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8082",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Bi-Taksi API Matching Api",
	Description:      "Bi-Taksi servisleri için API dokümantasyonu",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

const docTemplate = `{
  "swagger": "2.0",
  "info": {
    "description": "{{.Description}}",
    "title": "{{.Title}}",
    "termsOfService": "",
    "contact": {},
    "license": {},
    "version": "{{.Version}}"
  },
  "host": "{{.Host}}",
  "basePath": "{{.BasePath}}",
  "schemes": {{marshal .Schemes}},
  "paths": {
    "/login": {
      "post": {
        "tags": ["User"],
        "summary": "User login",
        "description": "Allows an existing user to login",
        "operationId": "UserLogin",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "parameters": [
          {
            "in": "body",
            "name": "user",
            "description": "User login details",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserLoginSchema"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Login successful",
            "schema": {
              "$ref": "#/definitions/UserResponse"
            }
          },
          "400": {
            "description": "Bad request - validation error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "tags": ["User"],
        "summary": "User registration",
        "description": "Registers a new user in the system",
        "operationId": "UserRegister",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "parameters": [
          {
            "in": "body",
            "name": "user",
            "description": "User registration details",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserRegisterSchema"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Registration successful",
            "schema": {
              "$ref": "#/definitions/UserResponse"
            }
          },
          "400": {
            "description": "Bad request - validation error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/match-driver": {
      "post": {
        "tags": ["Matching"],
        "summary": "Find nearest driver",
        "description": "Finds the nearest driver to a specified latitude and longitude within a given radius",
        "operationId": "MatchDriver",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "parameters": [
          {
            "in": "body",
            "name": "matchRequest",
            "description": "Matching parameters",
            "required": true,
            "schema": {
              "$ref": "#/definitions/MatchRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Matched driver location",
            "schema": {
              "$ref": "#/definitions/Driver"
            }
          },
          "400": {
            "description": "Bad request - validation error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "UserLoginSchema": {
      "type": "object",
      "properties": {
        "username": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "UserRegisterSchema": {
      "type": "object",
      "properties": {
        "username": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "email": {
          "type": "string",
          "format": "email"
        }
      }
    },
    "MatchRequest": {
      "type": "object",
      "properties": {
        "latitude": {
          "type": "number",
          "format": "double"
        },
        "longitude": {
          "type": "number",
          "format": "double"
        },
        "radius": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "Driver": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "location": {
          "type": "object",
          "properties": {
            "latitude": {
              "type": "number",
              "format": "double"
            },
            "longitude": {
              "type": "number",
              "format": "double"
            }
          }
        }
      }
    },
    "UserResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        },
        "username": {
          "type": "string"
        },
        "email": {
          "type": "string"
        }
      }
    },
    "ApiResponse": {
      "type": "object",
      "properties": {
        "status": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`
