# Snake Validator
This repository is a simple implementation of API service to valid snake game.


## Requirements

* Go 1.22+

## Notes

1. This repo mock the DB layer using Stub repository.
2. Test only cover one layer (handler) just as a show case.
3. Some edge cases new to be considered like checking if snake moves posit side(other direction).
4. Application should be dockerized.
5. Add integration tests.
6. Add swagger for API docs.
7. Add linter.

### Run the app:

```
go mod tidy

go mod vendor

go run cmd/main.go
```


### Run unit tests:

```
go test -mod=vendor -race ./...
```


### Request Sample:

```
Method: POST
url: http://url:8081/validate

body: Json

{
   "state": {
            "gameId": "x0001",
            "width": 10,
            "height": 10,
            "score": 5,
            "fruit": {
                "x": 1,
                "y": 3
            },
        "snake": {
            "x": 10,
            "y": 10,
            "velX": 0,
            "velY": 0
        }
   },
    "tick" : [
        {
            "x": 0,
            "y": 1
        },
        {
            "x": 0,
            "y": 2
        },
         {
            "x": 1,
            "y": 2
        },
         {
            "x": 1,
            "y": 3
        }
    ]
}
```







