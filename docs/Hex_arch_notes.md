# Hexagonal arch

Hexagonal arch is not very complicated.

* Separate the code in 3 large areas
  * User side
  * Business side: isolated from both sites
  * Server side
* Connected via port, adapters
* Dependencies go inside the center

![Hex architecture](img/ss_hex.png)

*Port*: is an interface

On the next use case diagram is shown what we're going to build:

![App use_case_diagram](img/arch.png)



* User or test sending the request
* All ports are defined as interfaces
  * use loose coupling
* Can use mock implementation instead of server
* Can use business logic cut from the rest of the world

For more information see [this article](https://www.qwan.eu/2020/08/20/hexagonal-architecture.html).

