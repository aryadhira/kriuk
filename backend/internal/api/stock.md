# Stock API Documentation

The Stock API is a RESTful API that allows users to manage their stock data.

## Endpoints

### GET /stocks

Retrieves a list of all stocks.

#### Request Parameters

None

#### Response

* Status Code: `200 OK`
* Body: JSON array containing stock records

### GET /stocks/{id}

Retrieves a stock record by ID.

#### Request Parameters

* `id` (required): The stock ID

#### Response

* Status Code: `200 OK`
* Body: JSON object containing the stock record
* Status Code: `400 Bad Request`
* Body: Error message if the provided ID is invalid
* Status Code: `404 Not Found`
* Body: Error message if the stock with the provided ID does not exist

### POST /stocks

Creates a new stock record.

#### Request Body

* `id` (optional): A unique identifier for the stock record (defaults to auto-generation)
* `name` (required): The stock name
* `price` (required): The stock price
* `quantity` (required): The stock quantity

#### Response

* Status Code: `200 OK`
* Body: JSON object containing the saved stock record and a success message
* Status Code: `400 Bad Request`
* Body: Error message if the provided request body is invalid

### PUT /stocks

Updates an existing stock record.

#### Request Body

Same as the POST request body, but the `id` field is required.

#### Response

* Status Code: `200 OK`
* Body: JSON object containing the updated stock record and a success message
* Status Code: `400 Bad Request`
* Body: Error message if the provided request body is invalid or the stock does not exist
* Status Code: `404 Not Found`
* Body: Error message if the stock with the provided ID does not exist
