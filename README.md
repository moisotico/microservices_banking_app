# Micro-services Banking App

Part of the [REST based micro-services API development in Golang](https://www.udemy.com/course/rest-based-microservices-api-development-in-go-lang/) course

---

## Install

* Install Golang latest version [from here](https://golang.org/doc/install).
* Install Postman [from here](https://www.postman.com/downloads/) to check the API.

## Run

### mysql database

You can use any one of the following procedure to make a database instance, and make the changes to your `start.sh` file accordingly

1. `docker-compose.yml` file. This contains the init script to generate and tables and insert the default data. You just need to bring the container up

   To start the docker container, run :

   ```shell
   $ cd resources/docker
   $ docker-compose up
   ```

   **Note**: in Linux you should run ` docker-compose -f docker-compose.yml up`.

   

2. `resources/database.sql` this contains the SQL for generating the tables. In case you don't want to use the docker-compose file you can use this file to generate tables and insert the default data.

   

### App

* On the main folder open a terminal and run to start the API:

```shell
$ cd banking/
$ go run main.go
```



### Postman

* To check the API, open Postman and use the GET method with the following address:

```json
http://localhost:8000/customers
```





