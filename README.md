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

   

## Authentication

* On the main folder open a terminal and run to start the API:

```shell
$ cd banking-auth/
$ ./start.sh
```



### Postman

* Check the `banking/users` database to check the correct users & passwords, please note that we are using  `usr` and `pass` instead as an example
* To check the API, open Postman and use the POST method with the following address:

```json
http://localhost:8181/auth/login
```

* Add the username & password in the *Body* of the request.

  ```json
  {
      "username": "usr",
      "password": "pass"
  }
  ```

* **Copy the response** as it represents the bearer token, which is necessary  for the `banking` App requests.



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

* Go to *Auth* and select *Bearer token*, then proceed to paste the token obtained from the request previously made with `banking-auth` 

### Routes

These are the following routes you can use based on the bearer token obtained and the role selected (either admin or user):

| **Rou**tes      | Method | Direction                                      |
| --------------- | ------ | ---------------------------------------------- |
| GetAllCustomers | GET    | `/customer`                                    |
| GetCustomer     | GET    | `/customer/(customer_id)`                      |
| NewAccount      | POST   | `/customer/(customer_id)/account`              |
| NewTransaction  | POST   | `/customer/(customer_id)/account/{account_id}` |

- Admin role can request via all 4 routes.

- User role can only use 2 methods:

  - Get Customer by Id (**GetCustomer**)
  - Make a transaction (**NewTransaction**)

- Format: 

  ```
  {
  	"admin": ["GetAllCustomers", "GetCustomer", "NewAccount", "NewTransaction"]
  	"user": ["GetCustomer", "NewTransaction"]
  }
  ```

## Testing

* Some tests require us to generate mock files, so we should run the following command first:

  ```shell
  $ cd banking
  $ go generate ./...
  ```

* You can test various files and see the current coverage by going to the folder of the test and  running:

  ```shell
  $ go test -v -cover
  ```

  

## Additional Notes

* [Requirements on assignment 1](./docs/Assignment_1.md)
* [Requirements on assignment 2](./docs/Assignment_2.md)

Please review the [Course notes](./docs/Notes.md) for more information. 

