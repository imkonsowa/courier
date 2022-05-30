<div align="center">
<h3 align="center">Courier service -  Golang Monorepo with gRPC and docker</h3>
</div>

<details open>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about">About</a>
    </li>
    <li>
      <a href="#dependencies">Dependencies</a>
    </li>
    <li>
      <a href="#database">Database</a>
    </li>
    <li>
      <a href="#setup">Setup</a>
    </li>
    <li>
      <a href="#api-endpoints">API Endpoints</a>
      <ul>
        <li><a href="#send-csv-file-to-process">Send CSV File to process</a></li>
        <li><a href="#generate-cargo-json-file">Generate Cargo JSON file</a></li>
      </ul>
    </li>
    <li><a href="#possible-enhancements">Possible Enhancements</a></li>
  </ol>
</details>

## About

This repo contains a solution for JUMIA hiring process task, It's basically contains two services:

- `CSV Parser` service, responsible for receiving and processing csv files
- `Courier service`, responsible for persisting data in Postgres DB and generating cargos daily reports

## How it works

- User send a CSV file contains rows of parcel details to the `CSV Parser` service.
- The `CSV Parser` service parses the file, reads its data, detecting the parcel country and posting all rows to
  the `Courier` service over gRPC stream as chunks signed with the `day` data received at.
- The `Courier` service receives the streamed chunks and stores it in the `parcels` table in Postgres DB named `courier`
  .
- The `Courier` service has an exposed endpoint where cargo manifest file should be generated through.

## Dependencies

* [Go-Gonic](https://github.com/gin-gonic/gin) The foundational http router and data binding.
* [Gorm.io](https://gorm.io/gorm) The DB ORM.
* [Assert](https://pkg.go.dev/github.com/stretchr/testify/assert) For testing assertions.

## Database

This service uses Postgres as it's datastore.

## Setup

To run the service, make sure you have Docker and Compose installed on your testing machine.

## Setup and Run

```shell
make up
```

> **Warning**: before running the below command please make sure that ports `5432`, `1996`, `1997` and `1998` are not
> allocated.

- This command builds each service into a standalone container
- The SCV Parser service exposed over port `1996` and accessible via `http://localhost:1996`
- The Courier service exposed over port `1998` and accessible via `http://localhost:1998`

## API Endpoints

This section documents how to consume this service APIs, listed below all requests with request/response payloads
examples.

### Send CSV File to process

Used to parse a csv file and persist its rows in the DB.

#### Request URL

`http://localhost:1996/upload`

#### Request Method

`POST`

#### Request example

This endpoint expects a request with multipart content type with one field which is `file` holds the CSV file.

#### Response examples

##### 1- valid file received

```json
{
  "code": 200,
  "message": "Well received, Processing!",
  "success": true
}
```

##### 2- file is missing

```json
{
  "code": 400,
  "message": "File is missing!",
  "success": false
}
```

##### 3- invalid multipart request

```json
{
  "code": 500,
  "message": "Failed to process your request",
  "success": false
}
```

### Generate Cargo JSON file

This request simply generates a cargo json file for the parcels needs to be shipped for a given date.

This request must have a `date` query param with the standard date format e.g `2022-08-05`.  
It also may receive a `country` param to filter the parcels for a specific country only, 
country must be one of `Cameroon, Ethiopia, Morocco, Mozambique, Uganda`

#### request URL

`http://localhost:1998/cargo-file?day=2022-05-30&country=Ethiopia`

#### Request Method

`GET`

#### Response

Response is a json file attached to the response.

### Possible Enhancements

Below list states app possible enhancements I couldn't mount more time to implement.

- Utilize Go concurrency with DB data fetching.
- Increase testing coverage.