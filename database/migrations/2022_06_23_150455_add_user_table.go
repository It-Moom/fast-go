package migrations

import (
	"database/sql"
	"fast-go/app/model/entity/user"
	"fast-go/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	//type User struct {
	//	entity.BaseEntity
	//
	//	Name     string `gorm:"type:varchar(255);not null;index"`
	//	Email    string `gorm:"type:varchar(255);index;default:null"`
	//	Phone    string `gorm:"type:varchar(20);index;default:null"`
	//	Password string `gorm:"type:varchar(255)"`
	//
	//	entity.CommonTimestampsField
	//}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&user.User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&user.User{})
	}

	migrate.Add("2022_06_23_150455_add_user_table", up, down)
}
