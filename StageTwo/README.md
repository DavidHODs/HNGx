# HNGx Stage 2 Task

A simple REST API capable of CRUD operations on a "person" resource, interfacing with any database of your choice. The API should dynamically handle parameters, such as adding or retrieving a person by name. Accompany the development with UML diagrams to represent your system's design and database structure.  Host your entire project on GitHub, and provide a well-structured documentation in the repository that outlines request/response formats, setup instructions, and sample API usage.

## Objectives

1. REST API Development: Develop an API with endpoints for:
CREATE: Adding a new person.  =>/api
READ: Fetching details of a person.  => /api/user_id
UPDATE: Modifying details of an existing person => /api/user_id
DELETE: Removing a person => /api/user_id
Ensure all interactions with the database are secure and free from common vulnerabilities (e.g., SQL injections)

2. Database Modelling: (Bonus)
UML Diagram: Design and present a UML (Unified Modeling Language) diagram that represents the structure and relationships of your API's classes and models.

3. Testing:
Using tools like Postman or (scripts written in Python using the requests library) that tests each CRUD operation in your API.
This  should:
Add a new person (e.g., "Mark Essien").
Fetch details of a person
Modify the details of an existing person.
Remove a person

4. Dynamic Parameter Handling:
Your API should be flexible enough to handle dynamic input. If we provide a name (or other details), your backend should be able to process operations using that name.
Example: If we pass "Mark Essien", we should be able to perform all CRUD operations on "Mark Essien".
Add validation – field should only be strings; integers or any other data type should not be allowed.

5. Hosting
Using the same Server used in the Stage One task (or another server, if possible), modify it accordingly to  host your endpoint.
Test extensively with various testing tools to make sure it is accessible before submitting

## Programming Languages and Technologies Used for this Task

- Programming Language - Golang
- Database - Postgres
- Database Migration - Golang-Migrate
- Automated Shell Commands - Makefile
- UML - Whimsical
- Documentation - Postman
- API Testing - Postman
- Postman Documentation and Test Reports (JSON) to Html - Newmann
- Reverse Proxy Server - Nginx
- DNS - Name Cheap
- Cloud Server - AWS

## System Design

![CRUD](HNGx-CRUD.png)

## API Documentation

A full set of the API documentation showing the live link, test report, request and response examples is deployed on AWS and available on this [link](http://hng-x.davidoluwatobi.me/api/static/CRUD_test_report_and_documentation.html)

The Summary page shows a brief of what the API and project is all about as well as the summary of the tests run on the API.
Click on Total Requests to access the links, request and response body examples. Clicking on `Expand Requests` at the top right corner will bring all neccesary detail into view

## Reproducibility

1. Install Makefile, Golang-Migrate, Postgres, Golang if you don't have them installed

2. Create a new postgres Database

3. Create a .env file the root of your project. The file should contain two parameters ==> "APP_PORT" and "DATABASE_URL". To work optimally with the structure of the code, database URL should be in the form of <postgresql://postgres_user:postgres_password@addr:postgres_port/database_name>

4. Run the command `make migrate` to carry out database migration

5. Run the command `go run main.go` or `./main` to run the program on your local setup.

6. You can test with Postman. A postman requests collection (json format) and environment setup (json format) is also provided for you to import for easy API testing
