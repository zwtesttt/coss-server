package table

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"im/services/user/domain/entity"
)

func (d InitDatabase) AddTableUser() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202401031200",
		Migrate: func(tx *gorm.DB) error {
			// 执行迁移操作，例如创建表
			return tx.AutoMigrate(&entity.User{})
		},
	}
}

func (d InitDatabase) AddTableGroup() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202401031200",
		Migrate: func(tx *gorm.DB) error {
			// 执行迁移操作，例如创建表
			return tx.AutoMigrate(&entity.Group{})
		},
	}
}

func (d InitDatabase) AddTableUserRelation() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202401031200",
		Migrate: func(tx *gorm.DB) error {
			// 执行迁移操作，例如创建表
			return tx.AutoMigrate(&entity.UserRelation{})
		},
	}
}
func (d InitDatabase) AddTableUserGroup() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202401031200",
		Migrate: func(tx *gorm.DB) error {
			// 执行迁移操作，例如创建表
			return tx.AutoMigrate(&entity.UserGroup{})
		},
	}
}
