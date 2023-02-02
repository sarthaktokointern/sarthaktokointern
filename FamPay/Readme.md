Problem Statement :- To make an API to fetch latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

Tech Stack Used : GraphQL, GoLang, MySQL

[Note] : The API has been registered onto the GQL server

How to Run the Solution :-

[In Your Terminal]

1.) Clone the Repo

2.) Change Directory to FamPay

3.) Run go mod vendor

4.) Run go mod tidy

5.) Run the following command :- go run server.go

6.) Go to the url prompted in the terminal

7.) Paste the following query on the left hand side and run it :-


<img width="665" alt="Screenshot 2023-02-02 at 1 41 58 PM" src="https://user-images.githubusercontent.com/99721005/216267909-08d50c5e-b4e0-43cf-a2d6-662e2fb3d3a4.png">


Expected output :-


<img width="1440" alt="Screenshot 2023-02-02 at 1 41 01 PM" src="https://user-images.githubusercontent.com/99721005/216267696-954132e2-77c4-40ef-89aa-924a5a733a7b.png">




Important Files :-

schema.resolvers.go - comprises of the main logic of the solution

schema.graphqls - contains the schema of our API
