Problem Statement :- Considering a single node environment where the server is deployed, Implement an API
which when called returns the number of requests made till now and rate limits any
requests after X requests with 60 seconds. Return a 429 HTTP error response when the
user hits the rate limit X. The response header should also contain the X-WAIT-TILL
(time when the next API request to succeed) and X-RATE-LIMIT (the rate limit X
enforced) details.

Tech Stack Used : GraphQL, GoLang

[Note] : The API has been registered onto the GQL server

How to Run the Solution :- 
[In Your Terminal]
1.) Clone the Repo
2.) Change Directory to Pillow Fund
3.) Run go mod vendor
4.) Run go mod tidy
5.) Run the following command :- go run server.go
6.) Go to the url prompted in the terminal
7.) Paste the following query on the left hand side and run it :- 
query{
    RequestCounter{
        Extensions {
        Status
        Message
        X_Wait_Till
        X_Rate_Limit
        }
        Data{
        NumberOfRequests
        }
    
    }
}

Expected output :-

Before reaching X-Rate-Limit :






After Reaching X-Rate-Limit :




Important Files :-

schema.resolvers.go - comprises of the main logic of the solution
schema.resolvers_test.go - contains the test cases for checking the sanity of the main functions'
config.ini - it is the config file which contains configurable X-Rate-Limit and X-Wait-Till values
