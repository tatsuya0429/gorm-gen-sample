package queries

import "gorm.io/gen"

// ここに書くとエラーになる
type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	FindByID(id int64) (gen.T, error)
	// SELECT * FROM @@table WHERE user_id = @userId
	FindByUserId(userId int64) ([]gen.T, error)
	// SELECT * FROM @@table
	FindAll() ([]gen.T, error)
}
