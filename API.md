# API Documentation

## Calculate endpoint
The endpoint is available at http://localhost:8080/calculate and the following payload is accepted as an input:
```json
[
    {
        "Origin": "IND",
        "Destination": "EWR"
    },
    {
        "Origin": "SFO",
        "Destination": "ATL"
    },
    {
        "Origin": "GSO",
        "Destination": "IND"
    },
    {
        "Origin": "ATL",
        "Destination": "GSO"
    }
]
```
As you can notice it's a json with an array of paths, each one containing two fields, **Origin** and **Destination**.

A response with the same structure will be sent by the server.

## Response codes
The API has the following response codes:
- **200** - OK, the calculation was successful and the result is available on the response
- **400** - Bad Request, there's an error in the payload
- **500** - Internal Server Error, if something unexpected happened on the server side

## Postman collection

There's a postman collection [here,](./postman_collection.json) so you can download and import into Postman to be able to
check all the requests payloads used to test the API.

