package app

import "strings"

const GoModTemPlate = `module template

go 1.20
require (
	github.com/gin-gonic/gin v1.9.0
	gorm.io/driver/mysql v1.5.0
	gorm.io/gorm v1.25.0
)`

func createGoMod(nameEn string) []byte {
	mod := strings.ReplaceAll(GoModTemPlate, "template", nameEn)
	return []byte(mod)
}
