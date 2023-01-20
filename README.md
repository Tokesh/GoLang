## About The Project


This project was made for improving understanding in Back-End part using Gin as a main web-framework.
<br>
The idea is to organize a full-fledged platform for an online store
<br>
The entire business logic of the store was thought out. Like from view of list of product until of order a products
<br>
Order tracking,cart, payments were also implemented


### Built With
* [Golang]
* [Gin Gonic Framework]
* [PGX]
* [Redis]
* [Docker]
* [PostgreSQL]
* [JWT Authorization]


### How to run project with Docker?
--> Open path to project in terminal<br>
--> Write "docker-compose build"<br>
--> After loading write "docker-compose up"<br>


### Layout
```tree
project
├── app
│   ├── controller
│   │   ├── brandController.go
│   │   ├── cartController.go
│   │   ├── deliveryController.go
│   │   ├── inventoryController.go
│   │   ├── orderController.go
│   │   ├── paymentController.go
│   │   ├── productController.go
│   │   └── userController.go
│   │
│   └── server.go
│
├── src
│   ├── app
│   │   └── service
│   │       ├── brandService.go
│   │       ├── cartService.go
│   │       ├── deliveryService.go
│   │       ├── orderService.go
│   │       ├── paymentService.go
│   │       ├── productService.go
│   │       ├── storeInventoryService.go
│   │       └── userService.go
│   │
│   ├── domain
│   │   └── model
│   │       ├── brand.go
│   │       ├── cart.go
│   │       ├── config.go
│   │       ├── delivery.go
│   │       ├── order.go
│   │       ├── payment.go
│   │       ├── product.go
│   │       ├── storeInventory.go
│   │       └── user.go
│   │
│   └── infrastructure
│       ├── postgresql
│       │   └── postgresql.go 
│       │
│       └── repositories
│           ├── brandRepository.go
│           ├── cartRepository.go
│           ├── deliveryRepository.go
│           ├── inventoryRepository.go
│           ├── orderRepo.go
│           ├── paymentRepository.go
│           ├── productRepository.go
│           └── userRepository.go
│
├── .gitignore
├── .env.example
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── GolangProject_postman.json
├── wait-for-postgres.sh
└── README.md
```


## Contact
Torekeldi Niyazbek(tokesh.niyazbek@gmail.com)
