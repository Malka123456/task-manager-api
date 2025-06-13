package config

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDBgit ingit remote add origin https://github.com/malka123456/task-manager-api.git
() {
    dsn := "host=localhost user=postgres password=Aa09!@#$ dbname=taskdb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Failed to connect to DB:", err)
    }

    log.Println("✅ Connected to database!")
    DB = db
}
