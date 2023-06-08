package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "strings"
import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

func main() {
	db, err := gorm.Open(sqlite.Open("car.db"), &gorm.Config{})
	if err!= nil {
        fmt.Println(err)
        os.Exit(1)
    }
	db.AutoMigrate(&Car{})
    fmt.Println("Car storage")
	fmt.Println("1. Add a car")
	fmt.Println("2. Get cars")
	fmt.Println("3. Delete a car")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your choice: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	choiceInt, _ := strconv.Atoi(choice)
	switch choiceInt {
		case 1:
            addCar(db)
        case 2:
            getCars(db)
        case 3:
            deleteCar(db)
		default:
			fmt.Println("Invalid choice")
		}
}

func addCar(db *gorm.DB) {
	fmt.Print("Enter the brand of the car: ")
    name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    name = strings.TrimSpace(name)
	fmt.Print("Enter the model of the car: ")
	model, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	model = strings.TrimSpace(model)
	fmt.Print("Enter the year of the car: ")
	year, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	year = strings.TrimSpace(year)
	yearInt, _ := strconv.ParseInt(year,10,64)
    car := Car{Brand: name, CarModel: model, Year: yearInt}
    db.Create(&car)
	main()
}

func getCars(db *gorm.DB) {
	var cars []Car
    db.Find(&cars)
    for _, car := range cars {
        fmt.Println(car)
    }
	main()
}

func deleteCar(db *gorm.DB) {
	fmt.Print("Enter the id of the car to delete: ")
    id, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    id = strings.TrimSpace(id)
    idInt, _ := strconv.ParseInt(id,10,64)
    db.Delete(&Car{}, idInt)
	main()
}


type Car struct {
	gorm.Model
	Brand string
	CarModel string
	Year int64
}