package main

import (
	"flag"
	"fmt"
	"DockerGoNginx/app/models"
	db "DockerGoNginx/config/databases/mysql"
)

var (
	h bool
	m string
)

var tables []interface{} = []interface{}{
	// 輸入要建立的 model struct
	models.User{},
	models.Token{},
	models.ApiLog{},
}

func init() {
	flag.BoolVar(&h, "h", false, "migration help")
	flag.StringVar(&m, "m", "default", "migrate command: auto/drop/create/refresh")
}

func auto() {
	db.GormDB.AutoMigrate(tables...)
}
func drop() {
	db.GormDB.DropTable(tables...)
}
func refresh() {
	drop()
	auto()
}

func migrate() {
	fmt.Println("--- Start Migrate ---")
	switch m {
	case "auto":
		fmt.Println("--- AutoMigrate ---")
		auto()
	case "drop":
		fmt.Println("--- DropTable ---")
		drop()
	case "refresh":
		fmt.Println("--- FreshTable ---")
		refresh()
	case "create":
		fmt.Println("create (not yet)")
	default:
		fmt.Println("Please input auto/drop/create/refresh")
	}
	fmt.Println("--- End Migrate ---")
}
func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	migrate()
}
