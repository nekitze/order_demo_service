# Order demo service

## Features
- Implemented connection and channel subscription in nats-streaming
- The received order data from nats-streaming is written to the database and stored in the cache
- In case of service crash, the cache is restored from the database
- The service gives order data by id from cache by the handle /api/order/id
- A simple interface for displaying the obtained data has been implemented

## How to run:

1. Run service ```docker compose up```

2. Run nats publisher ```go run cmd/publisher```

Service available on http://localhost:8080/ by default
