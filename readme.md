# wallet

I use existing libs and tools :

 - Ozzo Validation, for input request validation
 - Godotenv, for env loader
 - jmoiron/sqlx for postgres driver
 - redis for caching (in-memory db)
 - postgresql for DB

 # For setup after cloning/unzip the project:
> cd wallet
> go mod tidy
> make changes in the .env file using your postgresql and redis

# for db table :
> in folder db, there is a .sql file with the create table command. I use postgresql for this case. you can run the command in your sql editor page.

# to do a unit test :
> i've made several unit testing but just in usecases layer
> go to the each usecase package you want to testing then run a command "go test"
> you can see the coverage testing in each usecase package by open the project with vscode, choose the testing file, right click then choose "Go:Toogle Test Coverage in Current Package"

# to run the project
after set the .env file with yoyr database and redis credential, then stay still in root directory, then do "go run main.go" in terminal

# collections
for the collections, you can export the postman file in this project (Wallet.postman_collection.json) or you can use this example curl :

## register
curl --location --request POST 'http://localhost:8080/api/user/create_user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"jodyalmaida3"
}'

## login
curl --location --request POST 'http://localhost:8080/api/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"jodyalmaida"
    
}'

## topup
curl --location --request POST 'http://localhost:8080/api/balance/balance_topup' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNSwid2FsbGV0X2lkIjozLCJleHAiOjE3MTY3MTAyMjZ9.0u62khne6kzFLZoBZIsPsG2SCy7DUR9kPuYYLaHDP40' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount":100000
}'

## Balance Read
curl --location --request GET 'http://localhost:8080/api/balance/balance_read' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNSwid2FsbGV0X2lkIjozLCJleHAiOjE3MTY3MTAyMjZ9.0u62khne6kzFLZoBZIsPsG2SCy7DUR9kPuYYLaHDP40' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount":50000
}'

## Transfer 
curl --location --request POST 'http://localhost:8080/api/transaction/transfer' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNSwid2FsbGV0X2lkIjozLCJ1c2VybmFtZSI6ImpvZHlhbG1haWRhIiwiZXhwIjoxNzE2NzE2MzkxfQ.o0fgUyyQ46NkK7IJqa-nEgbsXXgse5OWCcYNoNPWoVk' \
--header 'Content-Type: application/json' \
--data-raw '{
    "to_username":"jodyalmaida",
    "amount":50000
}'

## TopUser
curl --location --request GET 'http://localhost:8080/api/transaction/top_users' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNSwid2FsbGV0X2lkIjozLCJleHAiOjE3MTY3MTM2MzV9.jWEbo0MAL_ymXRPRcIYNCod_wSVIphKsl7Ox0XuyN7Q'

## Top Transactions Per User
curl --location --request GET 'http://localhost:8080/api/transaction/top_transactions_per_user' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNSwid2FsbGV0X2lkIjozLCJleHAiOjE3MTY3MTM2MzV9.jWEbo0MAL_ymXRPRcIYNCod_wSVIphKsl7Ox0XuyN7Q'