# Go API for Current Toronto Time with MySQL Database Logging

This project is a Go-based API that provides the current time in Toronto in JSON format. Each request to the API logs the current time into a MySQL database.

## Objectives

- **Create a Go API Endpoint**: The API will return the current time in Toronto in JSON format.
![Screenshot (206)](https://github.com/user-attachments/assets/b5c1fa8d-57de-4173-9e69-293ad41a573c)


- **MySQL Database Integration**: Connect to a MySQL database and store the time data for each API request.

- ![Screenshot (208)](https://github.com/user-attachments/assets/eebdb303-6306-447a-b7cc-25957bebaf12)


- 
- **Time Zone Handling**: Adjust the time returned to match Toronto’s timezone.
- **Error Handling**: Properly handle errors related to database operations and time zone conversions.

## Requirements

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or later)
- [MySQL](https://dev.mysql.com/downloads/installer/) (version 8.0 or later)
- [MySQL Driver for Go](https://github.com/go-sql-driver/mysql)

### Installing Dependencies

1. Install the required Go MySQL driver:
   ```bash
   go get -u github.com/go-sql-driver/mysql
