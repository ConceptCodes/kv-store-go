# Key Value Store

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Key Value Store Microservice written in Go.

## Features
- Multi Tenancy
- Authentication
- Authorization
- CLI
- REST API

## Prerequisites

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/doc/install) (version 1.21 or higher)

## Installation

1. Clone this repository:

   ```sh
   git clone https://github.com/conceptcodes/kv-store-go.git
   cd kv-store-go
   ```

2. Navigate to the CLI directory and build the application:

   ```sh
   # For Windows 64-bit
   GOOS=windows GOARCH=amd64 go build -o password-generator.exe cmd/main.go

   # For macOS 64-bit
   GOOS=darwin GOARCH=amd64 go build -o password-generator cmd/main.go

   # For Linux 64-bit
   GOOS=linux GOARCH=amd64 go build -o password-generator cmd/main.go
   ```

3. Run the CLI:

   ```sh
   # For Windows 64-bit
   ./password-generator.exe --help

   # For macOS/Linux 64-bit
   ./password-generator --help
   ```

If you want to install the CLI globally, you can run the following command:

```sh
# (Mac and Linux)
sudo mv password-generator /usr/local/bin
```

```sh
password-generator --help

```

## Usage

To run the CLI, use the following command:

```sh
password-generator --length 12 --uppercase --lowercase --numbers --symbols

Your password is: SD5a[M8#2lMi
```
