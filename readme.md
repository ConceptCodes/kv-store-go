# Key Value Store

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A [Microservice](https://microservices.io/) for Key-Value Storage

## Features
- Multi Tenancy
- Authorization
- REST API

## Prerequisites

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/doc/install) (version 1.21 or higher)

## Installation

1. Clone this repository:

   ```sh
   git clone https://github.com/conceptcodes/kv-store-go.git
   cd kv-store-go
   ```

## Usage (Api)

1. Run the server

  ```sh
  gow run cmd/server/main.go
  ```

2. Verify that the service is running

  ```sh
  curl http://localhost:8080/health/alive
  ```
  ```json
  {
    "message": "Service is alive",
    "data": null,
    "error_code": ""
  }
  ```

3. Onboard a new tenant. This action will add an authorization header to the response. Utilize this header for all subsequent requests.

  ```sh
  curl --location 'http://localhost:8080/api/tenant/onboard'
  ```
  ```json
  {
    "message": "Tenant onboarded successfully",
    "data": {
        "tenant_id": "sample_tenant_id",
        "tenant_secret": "sample_tenant_secret"
    },
    "error_code": ""
  }
  ```

4. Create a new key-value pair

  ```sh
  curl --location 'http://localhost:8080/api/records' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer <sample_auth_token>' \
  --data '{
      "key": "long",
      "value": "word",
      "ttl": 3600
  }'
  ```
  ```json
  {
    "message": "Record created successfully",
    "data": {
        "key": "long",
        "value": "word",
    },
    "error_code": ""
  }
  ```

5. Get a key-value pair

  ```sh
  curl --location 'http://localhost:8080/api/records/long' \
  --header 'Authorization: Bearer <sample_auth_token>'
  ```
  ```json
  {
    "message": "Record Found Successfully",
    "data": {
        "key": "long",
        "value": "world"
    },
    "error_code": ""
  }
  ```


## Roadmap

- [ ] CLI wrapper over the api
- [ ] Add known errors to the api
- [ ] Standard Log Formatting (with log levels)