# lil-gqland

Welcome to lil-gqland. lil-gqland is a web application using Go and GraphQL, and for database operations, I'm utilizing SQLBoiler ORM to make things smoother.

## Main Features
`Data Fetching`: In the graph/db directory, you'll find queries designed to fetch data from the database. 
For instance, the `issues.go` file contains the `GetIssueByRepoAndNumber` function, which I use to retrieve issues by their repository and number.

## GraphQL Schemas and Resolvers: 
The graph directory is where the GraphQL schemas and resolvers are defined. This setup allows me to process GraphQL queries from the client side and return the appropriate data.

## Testing:
I'm committed to maintaining code quality, and the project includes various tests to ensure everything works as expected.

## Development Environment
Go: This is the primary language I'm using for the project. Dependencies are managed through Go's module management system.

GitHub Actions: The CI/CD pipeline is set up using GitHub Actions, configured in the .github/workflows/go.yml file.

SQLBoiler: As my ORM choice for Go, SQLBoiler simplifies database operations.

## Setup
To set up db, simply run the `setup.sh` script. Go dependencies are managed using the `go.mod` and `go.sum` files.

