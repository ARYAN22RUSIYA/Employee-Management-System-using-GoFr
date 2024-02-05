
# Employee Management System API using Gofr

This an API project that helps you with Management of your employees records.It uses database as SQL.It performs CURD operations to do Creation,updation,Read and Deleting of a Record.


## To install Gofr


```bash
  go get gofr.dev
```
      go mod init employee_manager

      go mod tidy

      go run main.go

## Connection to MYSQL Database

```bash
  docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=employees_db -p 3306:3306 -d mysql:8.0.30

  docker exec -it gofr-mysql mysql -uroot -proot123 employees_db -e "CREATE TABLE employees (eid int(255) PRIMARY KEY, empname VARCHAR(255) NOT NULL,salary int(50),email VARCHAR(255));"
```
    Project implementation explanation :

configs/.env - It contains the configuration regulated through environment variables which are required for th interaction with our SQL database.

model/model.go - We’ve chosen Student as our model. The Student model defines four fields: ID, Name, Age, and Class, each tagged with JSON annotations for easy serialization and deserialization when working with RESTful endpoints.

datastore/interface.go - To interact with the Student model in your database, we need to define an interface that outlines the basic operations we’ll perform.

datastore/datastore.go - The data access layer using GORM and GoFr. This layer will encapsulate the logic for interacting with the database, performing CRUD operations on our student data.This datastore.go file forms the bridge between your REST API and the database, providing a clean and organized way to perform data operations.

handler/handler.go - Handlers are responsible for receiving HTTP requests, processing data, and generating responses. They act as the bridge between the application logic and the external world, orchestrating the interactions with the data access layer to fulfill API requests.

main.go - This file takes center stage. This is where we tie all the elements together and set up the server to handle incoming requests.

Functionalities -

datastore.go -

Sure, let's summarize the functionalities of each of the four functions:

GetByID Function:

Purpose: Retrieves a student from the database based on their ID. Parameters: ctx: Context for the database operation. id: ID of the student to retrieve. Return: Returns a pointer to a model.Student struct if the student is found. Returns an error, including a custom errors.EntityNotFound error if the student with the specified ID is not found.

## Create Function:

Purpose: Inserts a new student into the database. Parameters: ctx: Context for the database operation. student: Pointer to a model.Student struct representing the student to be created. Return: Returns a pointer to a model.Student struct representing the newly created student. Returns an error, including a custom errors.DB error if there is an issue with the database operation.

## Update Function:

Purpose: Updates an existing student in the database based on their ID. Parameters: ctx: Context for the database operation. student: Pointer to a model.Student struct containing updated information. Return: Returns a pointer to a model.Student struct representing the updated student. Returns an error, including a custom errors.DB error if there is an issue with the database update. 

## Delete Function:

Purpose: Deletes a student from the database based on their ID. Parameters: ctx: Context for the database operation. id: ID of the student to be deleted. Return: Returns nil if the deletion is successful. Returns an error, including a custom errors.DB error if there is an issue with the database deletion. These functions collectively provide basic CRUD operations (Create, Read, Update, Delete) for managing student records in a database using Go and a SQL-based database. The error handling in each function allows for proper handling of potential issues that may occur during database interactions.

## handler.go - New Function:

Purpose: Constructor function to create a new handler instance. Parameters: s: An instance of the datastore.Student interface. Return: Returns a new handler instance. GetByID Function:

Purpose: Retrieves a student by ID from the datastore. Parameters: ctx: Context for the HTTP request. Return: Returns the retrieved student (model.Student) if successful. Returns an error if the student is not found or if there are validation issues. 

## Create Function:

Purpose: Creates a new student based on the data from the HTTP request. Parameters: ctx: Context for the HTTP request. Return: Returns the newly created student (model.Student) if successful. Returns an error if there are issues with data binding or if the database operation fails. 

## Update Function:

Purpose: Updates an existing student based on the data from the HTTP request. Parameters: ctx: Context for the HTTP request. Return: Returns the updated student (model.Student) if successful. Returns an error if there are issues with parameters, data binding, or if the update operation fails. 

## Delete Function:

Purpose: Deletes a student by ID. Parameters: ctx: Context for the HTTP request. Return: Returns a success message if the deletion is successful. Returns an error if there are issues with parameters or if the deletion operation fails. 

## validateID Function:

Purpose: Validates and converts a string ID to an integer. Parameters: id: The string representation of the ID. Return: Returns the integer ID if successful. Returns an error if the conversion fails. These functions collectively handle CRUD operations for a student entity in a web server, interacting with a datastore through the datastore.Student interface. The error handling ensures proper responses in case of validation issues or errors during database operations.

## main.go - Creating a gofr Application:

app := gofr.New(): Initializes a new instance of the gofr application. The gofr framework is likely used for building web applications in Go. Creating a Datastore and Handler:

s := datastore.New(): Creates a new instance of the datastore, which is assumed to implement the datastore.Student interface. This datastore is responsible for handling CRUD operations on student data. h := handler.New(s): Initializes a new instance of the handler type, passing the created datastore instance as a parameter. The handler is designed to handle HTTP requests and interact with the datastore for student-related operations. Defining Routes and Linking to Handler Methods:

## app.GET("/students/{id}", h.GetByID):
Defines a route for handling HTTP GET requests at the path /students/{id}. This route is linked to the GetByID method of the handler, which retrieves a student by ID. app.POST("/students", h.Create): Defines a route for handling HTTP POST requests at the path /students. This route is linked to the Create method of the handler, which creates a new student. app.PUT("/students/{id}", h.Update): Defines a route for handling HTTP PUT requests at the path /students/{id}. This route is linked to the Update method of the handler, which updates a student by ID. app.DELETE("/students/{id}", h.Delete): Defines a route for handling HTTP DELETE requests at the path /students/{id}. This route is linked to the Delete method of the handler, which deletes a student by ID. Configuring the Server Port:

app.Server.HTTP.Port = 9092: Configures the HTTP server to listen on port 9092. This is the port on which the server will accept incoming HTTP requests. Starting the Server:

## app.Start(): 
Initiates the web server, causing it to start listening for incoming HTTP requests on the specified port (9092 in this case). In summary, this main function sets up a web server using the gofr framework, defines routes for CRUD operations on student records, and associates these routes with corresponding methods in the handler type. The handler, in turn, interacts with a datastore to perform the necessary operations on student data. The server starts listening for incoming requests after configuration.

## Prerequisites
Before running the project, ensure you have the following installed:

Go (Version 1.16 or later)
Git
GoFr
Make sure to install it using the following command:
go get gofr.dev
Docker
To simplify the setup and management of the database for your REST API, consider using a Docker image for SQL.
Running the Project
Clone the repository:

change the directory

cd .\Go
Install Dependencies:

go get -u ./...
Build and Run the Application:

go run main.go

    
## Various API requests over Postman

### Get API



![GET](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/ad3767cd-2274-4cbc-81e0-34dbd48d7900)

### Post
![POST](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/a04ee2fd-71d4-4260-bb69-c3d39de449d8)
### PUT
![PUT](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/272452b5-08b2-4522-b429-87f8c0a7f94b)
### DELETE
![DELETE](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/f1b4170e-16f0-46dd-af7c-cccea4cf078a)
## UML Diagram for the project
![UML](https://github.com/saksham5701/gofr-mini-project-zopsmart/assets/95173447/afd4512d-0b43-4b0d-8439-4e50f24cad7a)
 ## References
 - [Gofr Documentation ](https://gofr.dev/)
 - [Article over medium](https://medium.com/@mundhraumang.02/sample-rest-api-using-go-gorm-and-gofr-0ea41eaa6c62#:~:text=In%20the%20journey%20to%20create,server%20to%20handle%20incoming%20requests.&text=GoFr%20excels%20in%20simplifying%20the%20process%20of%20setting%20up%20a%20web%20server.) 








