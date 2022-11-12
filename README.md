
#Go Version go 1.18
API RANDOM USER 
How to Run 
#1
go mod tidy

#2 
Run source mode is debug
go run cmd/main.go

#3
Run binary build
./main


#4 Unit test coverage
```
go test -cover ./...
?   	api-random-user/cmd	[no test files]
?   	api-random-user/config	[no test files]
ok  	api-random-user/src/app/person/delivery	0.028s	coverage: 100.0% of statements
ok  	api-random-user/src/app/person/repository	0.044s	coverage: 100.0% of statements
ok  	api-random-user/src/app/person/usecase	0.049s	coverage: 100.0% of statements
?   	api-random-user/src/domain	[no test files]
?   	api-random-user/src/model/entity	[no test files]
?   	api-random-user/src/shared/mocks	[no test files]
```
