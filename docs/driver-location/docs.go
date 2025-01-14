package driver_location

import (
	"github.com/swaggo/swag"
)

// SwaggerInfo holds the swagger metadata
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8081",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Bi-Taksi API Driver Location Api",
	Description:      "Bi-Taksi servisleri için API dokümantasyonu",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

const docTemplate = `{
  "swagger": "2.0",
  "info": {
    "description": "Bi-Taksi servisleri için API dokümantasyonu",
    "title": "Bi-Taksi API Driver Location Api",
    "version": "1.0.0"
  },
  "host": "localhost:8081",
  "basePath": "/api/v1",
  "schemes": ["http"],
  "paths": {
    "/login": {
      "post": {
        "summary": "User Login",
        "description": "User login endpoint",
        "parameters": [{
          "in": "body",
          "name": "user",
          "schema": { "$ref": "#/definitions/UserLoginSchema" }
        }],
        "responses": {
          "200": { "description": "Successful login", "schema": { "$ref": "#/definitions/UserLoginResponse" } },
          "400": { "description": "Failed to decode request body" }
        }
      }
    },
    "/register": {
      "post": {
        "summary": "User Register",
        "description": "User registration endpoint",
        "parameters": [{
          "in": "body",
          "name": "user",
          "schema": { "$ref": "#/definitions/UserRegisterSchema" }
        }],
        "responses": {
          "200": { "description": "User registered and logged in successfully" },
          "400": { "description": "Failed to decode request body" }
        }
      }
    },
    "/driver-locations/{driver_id}": {
      "get": {
        "summary": "Get Latest Driver Location",
        "description": "Retrieves the latest known location of the specified driver",
        "parameters": [{
          "name": "driver_id",
          "in": "path",
          "required": true,
          "type": "string"
        }],
        "responses": {
          "200": { "description": "Latest driver location", "schema": { "$ref": "#/definitions/DriverLocation" } },
          "404": { "description": "Driver location not found" },
          "500": { "description": "Internal server error" }
        }
      }
    },
    "/driver-locations": {
      "post": {
        "summary": "Create or Update Driver Locations",
        "description": "Processes a batch of driver locations for creation or updating",
        "parameters": [{
          "in": "body",
          "name": "locations",
          "schema": { "type": "array", "items": { "$ref": "#/definitions/DriverLocation" } }
        }],
        "responses": {
          "200": { "description": "Locations processed successfully" },
          "400": { "description": "Failed to decode request body" },
          "500": { "description": "Failed to process locations" }
        }
      },
      "put": {
        "summary": "Update Driver Locations",
        "description": "Processes a batch of driver locations for updating",
        "parameters": [{
          "in": "body",
          "name": "locations",
          "schema": { "type": "array", "items": { "$ref": "#/definitions/DriverLocation" } }
        }],
        "responses": {
          "200": { "description": "Locations processed successfully" },
          "400": { "description": "Failed to decode request body" },
          "500": { "description": "Failed to process locations" }
        }
      }
    },
    "/find-driver-within-radius": {
      "post": {
        "summary": "Find Drivers Within Radius",
        "description": "Finds drivers within a radius from a specified location",
        "parameters": [{
          "in": "body",
          "name": "request",
          "schema": { "$ref": "#/definitions/DriverRequest" }
        }],
        "responses": {
          "200": { "description": "List of drivers within radius", "schema": { "type": "array", "items": { "$ref": "#/definitions/DriverLocation" } } },
          "400": { "description": "Invalid request payload or validation error" },
          "500": { "description": "Internal server error" }
        }
      }
    }
  }
}`
