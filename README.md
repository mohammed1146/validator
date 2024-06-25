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
```


### Run unit tests:

```
go test -mod=vendor -race ./...
```








