package seeders

import (
	"blog/pkg/console"
	"blog/pkg/database/factories"
	"blog/pkg/logger"
	"blog/pkg/seed"
	"fmt"
	"gorm.io/gorm"
)

func init() {
	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		users := factories.MakeUsers(100)

		result := db.Table("users").Create(&users)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded",
			result.Statement.Table, result.RowsAffected))
	})
}
