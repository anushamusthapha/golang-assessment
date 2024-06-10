package config

import (
	"golang-assessment/logger"
	"golang-assessment/models"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Dialect  string `yaml:"dialect"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"` // Change the type to int
	DBName   string `yaml:"dbname"`
}

func LoadDatabaseConfig() *DatabaseConfig {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	var config DatabaseConfig
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}

	log.Printf("Loaded config: %+v", config) // Add this line for debugging

	return &config
}

func SetupDatabase() *gorm.DB {
	dbConfig := LoadDatabaseConfig()

	// Constructing the DSN with proper parameter order
	dsn := "host=" + dbConfig.Host +
		" user=" + dbConfig.Username +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.DBName +
		" port=" + strconv.Itoa(dbConfig.Port) + // Convert port to string
		" sslmode=disable" +
		" TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.Employee{})
	logger.Log.Info("Database connected and migrated")

	return db
}
