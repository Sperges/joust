package data

import (
	"context"

	"gorm.io/gorm"
)

func RandomFromTable(ctx context.Context, db *gorm.DB, table string) (string, error) {
	var name string
	res := db.WithContext(ctx).Table(table).Select("name").Order("random()").Limit(1).Scan(&name)
	return name, res.Error
}
