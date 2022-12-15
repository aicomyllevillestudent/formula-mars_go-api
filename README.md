# GO GIN API

A Go API for the Formula Mars application.

[![Deploy API on DigitalOcean](https://github.com/nicolasloontjens/formula-mars_go-api/actions/workflows/main.yml/badge.svg)](https://github.com/nicolasloontjens/formula-mars_go-api/actions/workflows/main.yml)

## Project Description:

This API was built to communicate between all our deliverables except our AI. 

You can find a API-spec to see what it does [here](https://git.ti.howest.be/TI/2022-2023/s5/trending-topics-sd/students/mars07/documentation/-/blob/main/api-spec.yaml).

But the actual deployed API can be reached via [here](https://go-api-lgafo.ondigitalocean.app/api/)

## Project requirements and dependencies:

Requirements:
- Go (https://go.dev/doc/install)
- Gin-goinic (`go get -u github.com/gin-gonic/gin`)
- VS Code (recommended for launching the application locally)

Dependencies:

After importing the project, go to the root file and run `go get -d ./...` to download all the necessary dependencies


## Project structure:

The API is started via the main.go file you find in the root foler  
The rest is split into 2 main parts and 2 smaller parts:

### Main:
- The controllers folder: contains all the controllers that are necessary to talk to models
- The models folder: contains a the models that talk to the Postgresql database. In here you can also find setup.go which takes care of the connection to the database.

### Smaller:
- The middlewares folder: contains one file with all the middlewares we used in the application.
- The utils folder: contains one file with a token for the JWT authentication

## Running locally:
1. Install Go and Gin-gonic
2. Clone the repository
3. Run the command `go get -d ./...` to get all the dependencies and also run `go mod tidy` to make sure you get everything
4. Setup a postgresql database (https://www.postgresql.org/docs/current/tutorial-createdb.html)
5. Create a .env file following the .env.example we added and enter the postgresql credentials
6. run `go run .` and it should be starting


## Project deployment:

The API is available on <a href="https://go-api-lgafo.ondigitalocean.app/api/">this website</a>.  
This project is using GitHub action deploys and is being hosted on DigitalOcean.   
<a href="https://github.com/nicolasloontjens/formula-mars_go-api/">Project link on GitHub</a>

## Disclaimer

I started commiting on this GIT and then copying it to another directory which was linked to my personal Github account. I stopped doing it because I found it useless.

Then we switched our application to a different webhost (from Koyeb to DigitalOcean). But because I couldn't create a DigitalOcean account, I needed to transfer my project to Nicolas. 

Because of this is it impossible to see my previous commits from before the transfer to Nicolas's Github.