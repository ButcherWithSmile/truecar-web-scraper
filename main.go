package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CarInfo struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"column:name"`
	Price   string `gorm:"column:price"`
	Mileage string `gorm:"column:mileage"`
}

func main() {
	fmt.Println("Please enter the car brand:")
	var carBrand, carModel string
	fmt.Scanln(&carBrand)
	fmt.Println("Please enter the model of car:")
	fmt.Scanln(&carModel)

	c := colly.NewCollector()

	var carInfo []CarInfo

	dsn := "username:password@tcp(localhost:3306)/db_name"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&CarInfo{})
	if err != nil {
		log.Fatal(err)
	}

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

	c.Visit("https://www.truecar.com/used-cars-for-sale/listings/" + carBrand + "/" + carModel + "/")

	fmt.Println("Scraping and inserting data into your database completed.")
}
