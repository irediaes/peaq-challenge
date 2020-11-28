## 3. Write an application for an analytics of cryptocurrency exchange rate. Your application must get from any market exchange rate for pairs:

# Stack
1. Golang
2. gRPC
3. MySQL
3. Protobuf

## Services

1. Exchange
2. Analytics

Build the services with `docker-composer build`
Run the services with `docker-composer up`

## Exchange Service 
Handles the core business logic.
Data Manipulation
Communication with external APIs

## Analytics Service
The Client interface between the Exchange service and the users.
Handles API calls and exposes endpoint for users to retrieve data

## Endpoint:
## base: 
`127.0.0.1:50052/export/analytics`
## params: 
  ```
    from=2020-11-27T23%3A50%3A01
    to=2020-12-30T07%3A34%3A00
    format=json
  ```

## Example
```
# CURL
'127.0.0.1:50052/export/analytics?from=2020-11-01T00%3A01%3A00&to=2020-12-01T23%3A59%3A00&format=json' \
```
```
# code block
[
  {
    "from": "2020-11-27T07:29:53Z",
    "to": "2020-11-27T07:33:04Z",
    "market_data": [
      {
        "market_pair": "BTC-ADA",
        "growth_data": {
          "volume_growth": -0.0081,
          "high_growth": 0,
          "low_growth": 0
        }
      },
      {
        "market_pair": "ETH-ADA",
        "growth_data": {
          "volume_growth": -0.0617,
          "high_growth": 0,
          "low_growth": 0
        }
      }
    ]
  }
]

```

