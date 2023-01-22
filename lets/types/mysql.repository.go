package types

import "gorm.io/gorm"

type IMySQLRepository interface {
	SetDriver(*gorm.DB)
}
