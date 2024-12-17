
# Receipt Processor

This service provides an HTTP API to process receipts, assign points according to defined rules, and retrieve those
points later. The project is implemented in Go and uses in-memory storage for quick setup and testing. No external
databases or persistent storage are required.

## Endpoints

- **POST** `/receipts/process`  
  Takes a JSON receipt and returns a JSON object with an `id` for the stored receipt. Points are computed according to
  the defined rules.

- **GET** `/receipts/{id}/points`  
  Returns a JSON object containing the number of points awarded to the specified receipt.

## Rules for Points Calculation

1. One point for every alphanumeric character in the retailer name.
2. 50 points if the total is a round dollar amount with no cents.
3. 25 points if the total is a multiple of 0.25.
4. 5 points for every two items on the receipt.
5. If the trimmed length of the item description is a multiple of 3, add `ceil(price * 0.2)` points.
6. 6 points if the day in the purchase date is odd.
7. 10 points if the purchase time is after 2:00pm and before 4:00pm.

## Prerequisites

- Go 1.18+ installed (if running locally without Docker).
- Docker and Docker Compose installed (if running via Docker).

## Installation & Running

### Without Docker

1. **Build**
   ```bash
   make build

This will compile the application into bin/ags and copy .env to bin/.

2. Run

```shell
make run
````

The server starts at http://localhost:8080.

Build and Run in One Step

```shell
make build-and-run
````

This command compiles and immediately runs the application.

Running Tests

To run the test suite:

```shell

make test
```

With Docker
You can also run the project using Docker:

```shell
make up
```

This command uses docker compose to build and run the project inside containers. Once up, the service should be
available at http://localhost:8080.

API Examples

Process a Receipt:

```shell

curl -X POST \
-H "Content-Type: application/json" \
-d '{
"retailer": "Target",
"purchaseDate": "2022-01-01",
"purchaseTime": "13:01",
"items": [
{ "shortDescription": "Mountain Dew 12PK", "price": "6.49" },
{ "shortDescription": "Emils Cheese Pizza", "price": "12.25" },
{ "shortDescription": "Knorr Creamy Chicken", "price": "1.26" },
{ "shortDescription": "Doritos Nacho Cheese", "price": "3.35" },
{ "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ", "price": "12.00" }
],
"total": "35.35"
}' http://localhost:8080/receipts/process

```

# Response:

``` json
{ "id": "some-unique-id" }

```

Get Points:

```shell
curl http://localhost:8080/receipts/some-unique-id/points
```

Response:

```json
{
  "points": 28
}
```

Notes
• Data is stored in memory and will be lost upon server restart.
• This service is provided as-is for the assessment exercise.
• Feel free to explore and modify as needed.

