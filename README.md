# Employee Manager

## Overview

Employee Manager is a simple Go application for managing employee records. It provides CRUD (Create, Read, Update, Delete) operations via a RESTful API.

## Features

- **CRUD Operations:** Perform Create, Read, Update, and Delete operations on employee records.
- **Pagination:** Support for listing employee records with pagination.
- **Concurrency Safety:** Ensure safe concurrent access to the employee database.
- **In-memory Store:** Store employee records in an in-memory database.

## Installation

1. Clone the repository:

    ```
    git clone https://github.com/chetan1196/employee-manager.git
    ```

2. Navigate to the project directory:

    ```
    cd employee-manager
    ```

3. Build the application:

    ```
    go build
    ```

4. Run the application:

    ```
    ./employee-manager
    ```

## Usage

Interact with the application using HTTP requests. Example requests:

- **List Employees:** GET /v1/employees?page=1&pageSize=10
- **Create Employee:** `POST /v1/employees` (with JSON body containing employee details)
  - Example Payload:
    ```json
    {
        "name": "Chetan Tiwari",
        "position": "Software Engineer",
        "salary": 50000
    }
    ```
- **Get Employee by ID:** GET /v1/employees/{id}
- **Update Employee:** PUT /v1/employees/{id} (with JSON body containing updated employee details)
- **Delete Employee:** DELETE /v1/employees/{id}

## Configuration

The application listens on port 7075 by default. You can change this in the `main.go` file.

## Dependencies

- Gorilla Mux: Go package for HTTP routing and URL matching.
- testify: Testing utilities for Go, used for assertions in unit tests.

