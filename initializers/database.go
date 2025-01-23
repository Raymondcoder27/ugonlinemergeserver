package initializers

import (
	"log"
	"os"

	"github.com/ugonlinemergeserver/models" // Correct import path for your models
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB") // Using the database URL from the environment variable

	if dsn == "" {
		log.Fatal("Database URL is not set") // Exit if DB URL is missing
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Connect to PostgreSQL using the DSN
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err) // Log a fatal error if connection fails
	}
}

func MigrateDB() {
	// Migrate the FloatRequest model
	if err := DB.AutoMigrate(&models.FloatRequest{}); err != nil {
		log.Printf("Error migrating FloatRequest Database: %v", err)
	}

	// Migrate the ServiceRequest model
	if err := DB.AutoMigrate(&models.ServiceRequest{}); err != nil {
		log.Printf("Error migrating ServiceRequest Database: %v", err)
	}

	// Migrate the Post model
	if err := DB.AutoMigrate(&models.Post{}); err != nil {
		log.Printf("Error migrating Post Database: %v", err)
	}

	// Migrate the User model
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Printf("Error migrating User Database: %v", err)
	}

	// Migrate the BranchManager model
	if err := DB.AutoMigrate(&models.BranchManager{}); err != nil {
		log.Printf("Error migrating BranchManager Database: %v", err)
	}

	// Migrate the Agent model
	if err := DB.AutoMigrate(&models.Agent{}); err != nil {
		log.Printf("Error migrating Agent Database: %v", err)
	}

	// Migrate the Transaction model (if applicable to your app)
	if err := DB.AutoMigrate(&models.Transaction{}); err != nil {
		log.Printf("Error migrating Transaction Database: %v", err)
	}
}
