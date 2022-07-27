![status](https://github.com/ekomanurung/gin-todolist-app/actions/workflows/go.yml/badge.svg?branch=master)
# gin-todolist-app

sample todolist microservice using go and mysql

#### prerequisite
- docker desktop installed in your local

#### how to run the application
- go to project directory
- execute command `docker-compose up -d`
- execute `docker exec -it todolist-mysql-db` to go inside mysql container
- after entering the container, run `mysql -u root -p`
- input password (default should be: root)
- after that, execute query from `migration/insert_database.sql`
- execute query from `migration/insert_table_todo.sql`
- done, you can start to use api on `localhost:8080/v1/todos`