package savedata

import (
	"fmt"

	"../dbconnection"
)

// save data to mysql
func SaveData(title string, description string, linkPath string, linkImage string, link string) {
	_, err := dbconnection.Connect.Exec("insert products set title= ?, description = ?, link_path = ?, link_image = ?, link = ?", title, linkPath, linkImage, link)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
