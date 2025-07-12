# Cash API Documentation

The Cash API is a RESTful API that allows users to manage their cash transactions.

## Endpoints

### GET /cash

Retrieves a list of all cash transactions.

#### Request Parameters

None

#### Response

* Status Code: `200 OK`
 * Body: JSON array containing cash transaction records

### GET /cash?date=YYYYMMDD

Retrieves cash transactions for a specific date.

#### Request Parameters

* `date` (required): The target date in YYYYMMDD format

#### Response

* Status Code: `200 OK`
 * Body: JSON array containing cash transaction records filtered by date
* Status Code: `400 Bad Request`
 * Body: Error message if the provided date format is invalid

### POST /cash

Creates a new cash transaction record.

#### Request Body

* `date` (required): The cash transaction date in YYYYMMDD format
* `type` (required): The cash transaction type (e.g., "income" or "expense")
* `amount` (required): The cash transaction amount
* `note` (optional): An optional note for the cash transaction record

#### Response

* Status Code: `200 OK`
 * Body: JSON object containing the saved cash transaction record and a success message
* Status Code: `400 Bad Request`
 * Body: Error message if the provided request body is invalid

### PUT /cash

Updates an existing cash transaction record.

#### Request Body

Same as the POST request body, but the `id` field is required.

#### Response

* Status Code: `200 OK`
 * Body: JSON object containing the updated cash transaction record and a success message
* Status Code: `400 Bad Request`
 * Body: Error message if the provided request body is invalid or the cash transaction does not exist
* Status Code: `404 Not Found`
 * Body: Error message if the cash transaction with the provided ID does not exist

### DELETE /cash/{id}

Deletes a cash transaction record by ID.

#### Request Parameters

* `id` (required): The cash transaction ID

#### Response

* Status Code: `200 OK`
 * Body: A success message
* Status Code: `400 Bad Request`
 * Body: Error message if the provided ID is invalid
* Status Code: `404 Not Found`
 * Body: Error message if the cash transaction with the provided ID does not exist
