# Transaction API Documentation

The Transaction API is a RESTful API that allows users to manage their financial transactions.

## Endpoints

### GET /transactions

Retrieves a list of all transactions.

#### Request Parameters

None

#### Response

* Status Code: `200 OK`
 * Body: JSON array containing transaction records

### GET /transactions/{id}

Retrieves a transaction record by ID.

#### Request Parameters

* `id` (required): The transaction ID

#### Response

* Status Code: `200 OK`
 * Body: JSON object containing the transaction record
* Status Code: `400 Bad Request`
 * Body: Error message if the provided ID is invalid
* Status Code: `404 Not Found`
 * Body: Error message if the transaction with the provided ID does not exist

### POST /transactions

Creates a new transaction record.

#### Request Body

* `id` (optional): A unique identifier for the transaction record (defaults to auto-generation)
* `date` (required): The transaction date in YYYYMMDD format
* `employee_id` (required): The ID of the employee involved in the transaction
* `stock_id` (required): The ID of the stock involved in the transaction
* `quantity` (required): The quantity of the stock bought or sold
* `price` (required): The price of the stock at the time of the transaction
* `type` (required): The type of the transaction ("buy" or "sell")

#### Response

* Status Code: `200 OK`
 * Body: JSON object containing the saved transaction record and a success message
* Status Code: `400 Bad Request`
 * Body: Error message if the provided request body is invalid

### PUT /transactions

Updates an existing transaction record.

#### Request Body

Same as the POST request body, but the `id` field is required.

#### Response

* Status Code: `200 OK`
 * Body: JSON object containing the updated transaction record and a success message
* Status Code: `400 Bad Request`
 * Body: Error message if the provided request body is invalid or the transaction does not exist
* Status Code: `404 Not Found`
 * Body: Error message if the transaction with the provided ID does not exist
