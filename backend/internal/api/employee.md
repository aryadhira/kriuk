# Employee API Documentation

The Employee API is a RESTful API that allows users to manage their employee data.

## Endpoints

### GET /employees

Retrieves a list of all employees.

#### Request Parameters

None

#### Response

* Status Code: `200 OK`
* Body: JSON array containing employee records

### GET /employees/{id}

Retrieves an employee record by ID.

#### Request Parameters

* `id` (required): The employee ID

#### Response

* Status Code: `200 OK`
* Body: JSON object containing the employee record
* Status Code: `400 Bad Request`
* Body: Error message if the provided ID is invalid
* Status Code: `404 Not Found`
* Body: Error message if the employee with the provided ID does not exist

### POST /employees

Creates a new employee record.

#### Request Body

* `id` (optional): A unique identifier for the employee record (defaults to auto-generation)
* `name` (required): The employee name
* `email` (required): The employee email
* `department` (optional): The employee department
* `position` (optional): The employee position
* `salary` (optional): The employee salary

#### Response

* Status Code: `200 OK`
* Body: JSON object containing the saved employee record and a success message
* Status Code: `400 Bad Request`
* Body: Error message if the provided request body is invalid

### PUT /employees

Updates an existing employee record.

#### Request Body

Same as the POST request body, but the `id` field is required.

#### Response

* Status Code: `200 OK`
* Body: JSON object containing the updated employee record and a success message
* Status Code: `400 Bad Request`
* Body: Error message if the provided request body is invalid or the employee does not exist
* Status Code: `404 Not Found`
* Body: Error message if the employee with the provided ID does not exist
