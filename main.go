package main

import (
	//"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"github.com/xackery/eqcleanup/soulbinder"
	"os"
	"strings"
)

func main() {

	for {
		showMenu()
	}

}

func connectDB(config *eqemuconfig.Config) (db *sqlx.DB, err error) {
	//Connect to DB
	db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Db))
	if err != nil {
		fmt.Println("Error connecting to DB:", err.Error())
		os.Exit(1)
	}
	return
}

func showMenu() {
	var err error
	var option string
	fmt.Println("\n\n===EQ Cleanup===")
	fmt.Println("Choose an option:")
	config := menuConfig()
	db := menuDB(&config)
	defer db.Close()
	fmt.Println("3) Take out Soulbinders")
	fmt.Println("Q) Quit")

	fmt.Scan(&option)
	fmt.Println("You chose option:", option)
	option = strings.ToLower(option)
	if option == "q" || option == "exit" {
		fmt.Println("Exiting")
		os.Exit(0)
	}
	if option == "3" { //Clean up Soulbinders
		err = soulbinder.Clean(db)
		if err != nil {
			fmt.Println("Error removing soulbinders:", err.Error())
		}
	}

	return
}

func menuConfig() (config eqemuconfig.Config) {
	config, err := eqemuconfig.LoadConfig()
	status := "Good"
	if err != nil {
		status = fmt.Sprintf("Bad (%s)", err.Error())
	} else {
		status = fmt.Sprintf("Good (%s)", config.Longame)
	}
	fmt.Printf("1) Reload eqemu_config.xml (Status: %s)\n", status)
	return
}

func menuDB(config *eqemuconfig.Config) (db *sqlx.DB) {
	var err error
	db, err = connectDB(config)
	status := "Good"
	if err != nil {
		status = fmt.Sprintf("Bad (%s)", err.Error())
	}
	fmt.Printf("2) Test DB Connection Settings (Status: %s)\n", status)
	return
}