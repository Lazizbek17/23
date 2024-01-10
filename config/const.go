package config

import "fmt"

const (
	Port     = ":8080"
	
	User     = "lazizbek"
	Password = "2744"
	PortDB   = "5432"
	Host     = "localhost"
	DB       = "instagram"
)

var (
	ConnStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		Host, PortDB, User, DB, Password)
)
