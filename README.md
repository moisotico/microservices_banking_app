# Micro-services Banking App

Part of the [REST based micro-services API development in Golang](https://www.udemy.com/course/rest-based-microservices-api-development-in-go-lang/) course

---

## Install

* Install Golang latest version [from here](https://golang.org/doc/install).
* Install Postman [from here](https://www.postman.com/downloads/) to check the API.
* Install `docker-compose` [from here](https://docs.docker.com/compose/install/).

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

   

## Auth

* On the main folder open a terminal and run to start the API:

```shell
$ cd banking-auth/
$ ./start.sh
```



### Postman

* To check the API, open Postman and use the POST method with the following address:

```json
http://localhost:8181/auth/login
```





## App

* On the main folder open a terminal and run to start the API:

```shell
$ cd banking/
$ ./start.sh
```



### Postman

* To check the API, open Postman and use the GET method with the following address:

```json
http://localhost:8282/customers
```

### Routes

These are the following routes to use:

| **Rou**tes      | Methof | Direction                                      |
| --------------- | ------ | ---------------------------------------------- |
| GetAllCustomers | GET    | `/customer`                                    |
| GetCustomer     | GET    | `/customer/(customer_id)`                      |
| NewAccount      | POST   | `/customer/(customer_id)/account`              |
| NewTransaction  | POST   | `/customer/(customer_id)/account/{account_id}` |

- Admin role can request via all 4 routes.

- User role can only use 2 methods:

  - Get Customer by Id
  - Make a transaction

- Format: 

  ```
  {
  	"admin": ["GetAllCustomers", "GetCustomer", "NewAccount", "NewTransaction"]
  	"user": ["GetCustomer", "NewTransaction"]
  }
  ```

  

## Additional Notes

* [Requirements on assignment 1](./docs/Assignment_1.md)
* [Requirements on assignment 2](./docs/Assignment_2.md)

Please review the [Course notes](./docs/Notes.md) for more information. 

