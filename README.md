# Rate Limit

web service that acts as third party rate limiter service.

Accept following arguments during the startup (command line args):

threshold - Max number of requests per URL within a time period (ttl).

ttl - The time period in which URL visits will be counted.

````
threshold by defualt is 10

ttl by defualt is 5
````

## Usage

### Start Project

```bash
go run main.go -ttl 20 -threshold 120

Or

go run main.go
```

### Endpoint
````
curl --location --request POST 'http://localhost:8080/report' \
--header 'Content-Type: application/json' \
--data-raw '{
		"url":"googsle.com"
}'
````
