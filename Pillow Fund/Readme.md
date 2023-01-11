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

<img width="368" alt="Screenshot 2023-01-11 at 1 12 02 PM" src="https://user-images.githubusercontent.com/99721005/211746614-a4fbd3e6-2a6b-430b-a1bf-e4dfe816d178.png">


Expected output :-


Before reaching X-Rate-Limit :


<img width="1440" alt="Screenshot 2023-01-11 at 1 12 40 PM" src="https://user-images.githubusercontent.com/99721005/211746766-3a94d171-6ab8-4902-a995-6cc471c2d663.png">




After Reaching X-Rate-Limit :


<img width="1440" alt="Screenshot 2023-01-11 at 1 12 55 PM" src="https://user-images.githubusercontent.com/99721005/211746812-c52ee9d4-5692-426f-9de7-b1d64302a0f8.png">



Important Files :-

schema.resolvers.go - comprises of the main logic of the solution

schema.resolvers_test.go - contains the test cases for checking the sanity of the main functions'

config.ini - it is the config file which contains configurable X-Rate-Limit and X-Wait-Till values
