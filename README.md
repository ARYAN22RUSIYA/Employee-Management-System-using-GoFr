# 📂 Student_Record_Maintainence_System
 
 This is a mini for a Student record management system which is implemented using GOfr and go language.

 The project comprises of the CRUD(create, read, update, delete) operations functionality along with MYSQL docker image integration for accessing database.

 Project implementation explanation :
  1) configs/.env - 
     It contains the configuration regulated through environment variables which are required for th interaction with our SQL database.

  2) model/model.go -
     We’ve chosen Student as our model. The Student model defines four fields: ID, Name, Age, and Class, each tagged with JSON annotations for easy serialization and deserialization when working with RESTful endpoints.

  3) datastore/interface.go -
     To interact with the Student model in your database, we need to define an interface that outlines the basic operations we’ll perform.

  4) datastore/datastore.go -
     The data access layer using GORM and GoFr. This layer will encapsulate the logic for interacting with the database, performing CRUD operations on our student data.This datastore.go file forms the bridge between your REST API and the database, providing a clean and organized way to perform data operations.

  5) handler/handler.go -
     Handlers are responsible for receiving HTTP requests, processing data, and generating responses. They act as the bridge between the application logic and the external world, orchestrating the interactions with the data access layer to fulfill API requests.

  6) main.go -
     This file takes center stage. This is where we tie all the elements together and set up the server to handle incoming requests.

 Functionalities -
  1)  datastore.go - 
 
      Sure, let's summarize the functionalities of each of the four functions:

      GetByID Function:

      Purpose: Retrieves a student from the database based on their ID.
      Parameters:
      ctx: Context for the database operation.
      id: ID of the student to retrieve.
      Return:
      Returns a pointer to a model.Student struct if the student is found.
      Returns an error, including a custom errors.EntityNotFound error if the student with the specified ID is not found.
      Create Function:
 
      Purpose: Inserts a new student into the database.
      Parameters:
      ctx: Context for the database operation.
      student: Pointer to a model.Student struct representing the student to be created.
      Return:
      Returns a pointer to a model.Student struct representing the newly created student.
      Returns an error, including a custom errors.DB error if there is an issue with the database operation.
      Update Function:

      Purpose: Updates an existing student in the database based on their ID.
      Parameters:
      ctx: Context for the database operation.
      student: Pointer to a model.Student struct containing updated information.
      Return:
      Returns a pointer to a model.Student struct representing the updated student.
      Returns an error, including a custom errors.DB error if there is an issue with the database update.
      Delete Function:

      Purpose: Deletes a student from the database based on their ID.
      Parameters:
      ctx: Context for the database operation.
      id: ID of the student to be deleted.
      Return:
      Returns nil if the deletion is successful.
      Returns an error, including a custom errors.DB error if there is an issue with the database deletion.
      These functions collectively provide basic CRUD operations (Create, Read, Update, Delete) for managing student records in a database using Go and a SQL-based database. The error handling in each function allows for proper handling of potential issues that may occur during database interactions.

  2)  handler.go -
      New Function:

      Purpose: Constructor function to create a new handler instance.
      Parameters:
      s: An instance of the datastore.Student interface.
      Return:
      Returns a new handler instance.
      GetByID Function:

      Purpose: Retrieves a student by ID from the datastore.
      Parameters:
      ctx: Context for the HTTP request.
      Return:
      Returns the retrieved student (model.Student) if successful.
      Returns an error if the student is not found or if there are validation issues.
      Create Function:

      Purpose: Creates a new student based on the data from the HTTP request.
      Parameters:
      ctx: Context for the HTTP request.
      Return:
      Returns the newly created student (model.Student) if successful.
      Returns an error if there are issues with data binding or if the database operation fails.
      Update Function:

      Purpose: Updates an existing student based on the data from the HTTP request.
      Parameters:
      ctx: Context for the HTTP request.
      Return:
      Returns the updated student (model.Student) if successful.
      Returns an error if there are issues with parameters, data binding, or if the update operation fails.
      Delete Function:

      Purpose: Deletes a student by ID.
      Parameters:
      ctx: Context for the HTTP request.
      Return:
      Returns a success message if the deletion is successful.
      Returns an error if there are issues with parameters or if the deletion operation fails.
      validateID Function:

      Purpose: Validates and converts a string ID to an integer.
      Parameters:
      id: The string representation of the ID.
      Return:
      Returns the integer ID if successful.
      Returns an error if the conversion fails.
      These functions collectively handle CRUD operations for a student entity in a web server, interacting with a datastore through the datastore.Student interface. The error handling ensures proper responses in case of validation issues or errors during database operations.

  3)  main.go -
      Creating a gofr Application:

      app := gofr.New(): Initializes a new instance of the gofr application. The gofr framework is likely used for building web applications in Go.
      Creating a Datastore and Handler:

      s := datastore.New(): Creates a new instance of the datastore, which is assumed to implement the datastore.Student interface. This datastore is responsible for handling CRUD operations on student data.
      h := handler.New(s): Initializes a new instance of the handler type, passing the created datastore instance as a parameter. The handler is designed to handle HTTP requests and interact with the datastore for student-related operations.
      Defining Routes and Linking to Handler Methods:

      app.GET("/students/{id}", h.GetByID): Defines a route for handling HTTP GET requests at the path /students/{id}. This route is linked to the GetByID method of the handler, which retrieves a student by ID.
      app.POST("/students", h.Create): Defines a route for handling HTTP POST requests at the path /students. This route is linked to the Create method of the handler, which creates a new student.
      app.PUT("/students/{id}", h.Update): Defines a route for handling HTTP PUT requests at the path /students/{id}. This route is linked to the Update method of the handler, which updates a student by ID.
      app.DELETE("/students/{id}", h.Delete): Defines a route for handling HTTP DELETE requests at the path /students/{id}. This route is linked to the Delete method of the handler, which deletes a student by ID.
      Configuring the Server Port:

      app.Server.HTTP.Port = 9092: Configures the HTTP server to listen on port 9092. This is the port on which the server will accept incoming HTTP requests.
      Starting the Server:

      app.Start(): Initiates the web server, causing it to start listening for incoming HTTP requests on the specified port (9092 in this case).
      In summary, this main function sets up a web server using the gofr framework, defines routes for CRUD operations on student records, and associates these routes with corresponding methods in the handler type. The handler, in turn, interacts with a datastore to perform the necessary operations on student data. The server starts listening for incoming requests after configuration.
    
## Prerequisites

Before running the project, ensure you have the following installed:

- [Go](https://golang.org/dl/) (Version 1.16 or later)
- [Git](https://git-scm.com/downloads) 
- [GoFr](https://gofr.dev/docs)
- Make sure to install it using the following command:
  ```
  go get gofr.dev
- [Docker](https://hub.docker.com/_/mysql)
- To simplify the setup and management of the database for your REST API, consider using a Docker image for SQL.

  
## Running the Project

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Pratishtha33/Go.git

2. change the directory
   ```bash
   cd .\Go

4. Install Dependencies:
   ```bash
   go get -u ./...

5. Build and Run the Application:
   ```bash
   go run main.go

# 🧪 Unit Testing

1.  Test Datastore: 
    ```bash
     go test ./datastore

2. Test Handler:
   ```bash
     go test ./handler


