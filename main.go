// This package uses the Colly web scraping library and GORM ORM to scrape car data from TrueCar and insert it into a MySQL database

package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// This struct defines the model for a car
type CarInfo struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"column:name"`
	Price   string `gorm:"column:price"`
	Mileage string `gorm:"column:mileage"`
}

func main() {
	// This code prompts the user to enter the car brand and model
	fmt.Println("Please enter the car brand:")
	var carBrand, carModel string
	fmt.Scanln(&carBrand)
	fmt.Println("Please enter the model of car:")
	fmt.Scanln(&carModel)

	// This code creates a new Colly collector and a slice of CarInfo structs
	c := colly.NewCollector()
	var carInfo []CarInfo

	//This code opens a connection to the MySQL database and creates a table for the CarInfo struct if it does not already exist
	dsn := "username:password@tcp(localhost:3306)/db_name"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&CarInfo{})
	if err != nil {
		log.Fatal(err)
	}

	//This code tells Colly to visit the specified URL and extract the car data from the .vehicle-card-body elements
	// If the number of cars in the carInfo slice is less than 20, the code creates a new CarInfo struct from the extracted data and appends it to the slice
	// The code then inserts the new car into the database
	c.OnHTML(".vehicle-card-body", func(h *colly.HTMLElement) {
		if len(carInfo) < 20 {
			price := h.ChildText("span[data-test='vehicleListingPriceAmount']")
			mileage := h.ChildText("div[data-test='vehicleMileage']")
			car := CarInfo{Name: carModel, Price: price, Mileage: mileage}
			carInfo = append(carInfo, car)

			result := db.Create(&car)
			if result.Error != nil {
				log.Printf("Inserting data error: %v", result.Error)
			}
		}
	})

	// This code tells Colly to visit the TrueCar URL for the specified car brand and model
	c.Visit("https://www.truecar.com/used-cars-for-sale/listings/" + carBrand + "/" + carModel + "/")

	// This code prints a message to the user informing them that the scraping and insertion process is complete
	fmt.Println("Scraping and inserting data into your database completed.")
}
