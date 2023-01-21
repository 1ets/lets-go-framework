package types

import "gorm.io/gorm"

type IMySQLRepository interface {
	SetAdapter(*gorm.DB)
}
