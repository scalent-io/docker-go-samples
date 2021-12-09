# Golang and MySQL communication example

Running mysql in container, set password as "password" and by setting MYSQL_ROOT_HOST we are allowing all the other IPs can connect to the mysql
```bash
docker run  --name mysqlhost --network mynet -e MYSQL_ROOT_PASSWORD=password -e MYSQL_ROOT_HOST=% -d mysql/mysql-server
```

Once mysql starts running, we need to get inside the running MySQL container, connect to mysql and create a new database
```bash
docker exec -it mysqlhost bash
```

Connect to MySQL
```bash
bash-4.4# mysql -uroot -ppassword
```

Create new daatabase
```bash
mysql>  create database testdb;
```

Either exit from the docker container or run following commands in other terminal 
## Create Docker image from golang code

```bash
docker build . -t go-mysql-docker:goMySQLDocker
```

## Running the docker image 
Run the docker image as a docker container, syntax "docker run -p HOST-PORT:CONTAINER-PORT IMAGE-ID" 
```bash
    docker run --network mynet -p 8080:80 -it go-mysql-docker:goMySQLDocker
```

## Visit the link 
To add record visit [http://localhost:8080/add](http://localhost:8080/add)
Listing of all the records [http://localhost:8080/list](http://localhost:8080/list)