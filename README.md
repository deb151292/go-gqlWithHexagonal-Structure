
# Backend with Hexagonal Architecture - Go & Graphql


This repository contains the source code for a Travel Agency Website, organized using a hexagonal architecture. The hexagonal architecture, also known as Ports and Adapters or the Onion Architecture, is a design pattern that promotes the separation of concerns and modularity in software development.

Project Structure
The project follows a hexagonal structure, which consists of the following main components:

### "cmd" Folder
This folder contains the entry point for the application. Specifically, the server file inside the cmd folder serves as the starting point for the code. 


## Generating New Graphql API 
Always create a new " *.graphqls " file with the name of the API inside gql folder
then Run the following Command at root folder

```bash
  go run github.com/99designs/gqlgen generate
```

## Run Locally

Clone the project 

```bash
git clone https://github.com/deb151292/go-gqlWithHexagonal-Structure.git
```

go inside "graph" Folder

```bash
  go run cmd/server.go
```


## Tech Stack

**Server:** Go with Graphql

**Architecture:** Hexagonal


## Author(s)

- [@Debojyoti Chakraborty](https://github.com/deb151292)

