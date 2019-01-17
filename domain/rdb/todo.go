package rdb

// +-------+--------------+------+-----+---------+----------------+
// | Field | Type         | Null | Key | Default | Extra          |
// +-------+--------------+------+-----+---------+----------------+
// | id    | int(11)      | NO   | PRI | NULL    | auto_increment |
// | name  | varchar(255) | YES  |     | NULL    |                |
// +-------+--------------+------+-----+---------+----------------+

// Todo is a struct for `todos` table
type Todo struct {
	ID   uint64  `db:"id"`
	Name *string `db:"name"`
}
