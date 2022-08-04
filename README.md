![status](https://github.com/ekomanurung/gin-todolist-app/actions/workflows/go.yml/badge.svg?branch=master)
# gin-todolist-app

sample todolist microservice using go and mysql

#### prerequisite
- docker desktop installed in your local

#### how to run the application
- go to project directory
- execute command `docker-compose up -d`
- done, you can start to use api on `localhost:8080/v1/todos`

#### swagger goland
- execute this in goland terminal to apply swag command:
   `export PATH=$(go env GOPATH)/bin:$PATH`
- changing request/response means you need to run `swag init -g {path-to-controller}/{your-controller}.go`
  to generate the swagger docs
- access swagger via `localhost:8080/swagger/index.html`