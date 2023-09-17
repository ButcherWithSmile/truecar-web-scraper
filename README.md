# Web Scraping with Go (Golang), Colly, GORM, and MySQL

## Overview

This project demonstrates how to perform web scraping using the Go programming language (Golang) with the Colly library, and how to store the scraped data in a MySQL database using GORM. Specifically, the project scrapes information about used cars for sale from the "www.truecar.com" website. It extracts the brand name, model name, price, and mileage of the first 20 cars listed in the used car for sale section and saves this data to a MySQL database.

## Requirements

Before you can run this project, you'll need the following:

- Go (Golang): Make sure you have Go installed on your machine. You can download it from the official website: https://golang.org/dl/

- Colly Library: Colly is a popular web scraping library for Go. You can install it using `go get`:
  ```
  go get -u github.com/gocolly/colly/v2
  ```

- GORM: GORM is a Go Object-Relational Mapping (ORM) library. You can install it using `go get`:
  ```
  go get -u gorm.io/gorm
  ```

- MySQL Database: You need to have a MySQL database installed and running. You'll also need to create a database and configure the connection details in the project's code.

## Usage

1. Clone this repository to your local machine:
   ```
   git clone https://github.com/your-username/truecar-web-scraper.git
   ```

2. Navigate to the project directory:
   ```
   cd truecar-web-scraper
   ```

3. Configure the MySQL database connection in the `main.go` file. Update the `dsn` variable with your MySQL connection details.

4. Build and run the project:
   ```
   go run main.go
   ```

5. The program will scrape data from "www.truecar.com," specifically the used car for sale section, and save the brand name, model name, price, and mileage of the first 20 cars into the MySQL database.

## Code Structure

- `main.go`: The main application file. It contains the scraping logic using Colly and the database interaction using GORM.

- `models/car.go`: Defines the data structure for the Car entity, which corresponds to the MySQL table schema.

- `database/database.go`: Handles database initialization and connection setup using GORM.

## Contributing

Feel free to contribute to this project by submitting issues, proposing new features, or sending pull requests. Your contributions are greatly appreciated!

## License

This project is licensed under the GPL-3.0 license. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Colly](http://go-colly.org/) - A powerful scraping framework for Go.
- [GORM](https://gorm.io/) - The fantastic ORM library for Go.
- [MySQL](https://www.mysql.com/) - An open-source relational database management system.
- [TrueCar](https://www.truecar.com/) - The website from which the data is being scraped.

---

Happy scraping with Go! 🚗💨