package main

import (
	"fmt"
	"github.com/turret-io/go-menu/menu"
	"github.com/xackery/eqcleanup/core"
	"github.com/xackery/eqcleanup/era"
	"github.com/xackery/eqcleanup/misc"
	"github.com/xackery/eqcleanup/tools/eqdb"
	"github.com/xackery/eqcleanup/zone"
)

var isDBConnected bool
var isConfigLoaded bool

func main() {
	var err error
	fmt.Println("                         EQCleanup, v0.15")
	fmt.Println("                       Written by Shin Noir")
	fmt.Println("--------------------------------------------------------------------------")
	var option string
	_, err = eqdb.Load()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Press q then enter to quit.")
		fmt.Scanf(option)
		return
	}
	fmt.Println("* !!!NOTE!!! Backup your database before using this tool.")

	//var err error
	mainMenu := []menu.CommandOption{
		{"era", "Out of era edits, ex. removing items or npcs", era.Menu},
		{"zone", "Zone migration tools, ex. newer to classic zone", zone.Menu},
		{"core", "System maintainence options, ex. removing all character data", core.Menu},
		{"misc", "Misc edits, ex. disabling rainfall", misc.Menu},
	}
	menuOptions := menu.NewMenuOptions("'menu' for help []> ", 0)

	menu := menu.NewMenu(mainMenu, menuOptions)
	menu.Start()
}
