// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "multipart/form-data"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Swagger Trash",
    "version": "0.1.0"
  },
  "paths": {
    "/": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "default"
        ],
        "summary": "get update",
        "operationId": "default",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/upload": {
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "uploadS3"
        ],
        "summary": "Uploads a file.",
        "operationId": "uploadS3",
        "parameters": [
          {
            "type": "file",
            "description": "The file to upload.",
            "name": "upfile",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "multipart/form-data"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Swagger Trash",
    "version": "0.1.0"
  },
  "paths": {
    "/": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "default"
        ],
        "summary": "get update",
        "operationId": "default",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/upload": {
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "tags": [
          "uploadS3"
        ],
        "summary": "Uploads a file.",
        "operationId": "uploadS3",
        "parameters": [
          {
            "type": "file",
            "description": "The file to upload.",
            "name": "upfile",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    }
  }
}`))
}
