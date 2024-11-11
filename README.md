# Synapsis Library Management Server

Scalable and modular backend system built with Go (Golang), designed to manage a Library Management System using microservices architecture

## Table of Contents

- [Technologies Used](#technologies-used)
- [Setup](#setup)
- [API Collections](#API-collections)

## Technologies Used
- Go version 1.22.x
- Docker version 4.28
- Postmann version 10.24.x

## Setup

1. Clone this repository:

   ```
   git clone https://github.com/mch-fauzy/synapsis-library-management-server.git
   ```

2. Navigate to the project directory:

   ```
   cd synapsis-library-management-server
   ```

3. To start the application, run the following command in the project root folder:

   ```
   docker-compose up --build
   ```

## API Collections

To simplify testing of the API endpoints, a Postman collections is provided. Follow the steps below to import and use it:

1. Use the Postman collection JSON file [**synapsis-library-management-server.postman_collection.json**](docs/api/synapsis-library-management-server.postman_collection.json) in this project directory

2. Open Postman

3. Click on the "Import" button located at the top left corner of the Postman interface

4. Select the JSON file

5. Once imported, you will see a new collection named "synapsis-library-management-server" in your Postman collections

6. You can now use this collection to test the API endpoints by sending requests to your running API server
