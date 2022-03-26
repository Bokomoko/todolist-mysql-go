# Simple todolist to ilustrate how to create a REST API in Golang

## Requirements : Docker, Go 1.17.x and a text editor

### Install MySql server

```bash
docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root mysql
docker exec -it mysql mysql -uroot -proot -e 'CREATE DATABASE todolist'
```

### Install Myqsql drivers for Golang

```bash
go get -u github.com/go-sql-driver/mysql
```

### Install ORM library to access datamodels in go

```bash
go get -u github.com/jinzhu/gorm
go get -u github.com/jinzhu/gorm/dialects/mysql
```
