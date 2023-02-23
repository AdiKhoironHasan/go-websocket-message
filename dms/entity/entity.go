package entity

type NotificationEntity struct {
	ID        int    `json:"id" db:"id"`
	UserID    int    `json:"user_id" db:"user_id"`
	IsRead    bool   `json:"is_read" db:"is_read"`
	Title     string `json:"title" db:"title"`
	Detail    string `json:"detail" db:"detail"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type NotificationParam struct {
	UserID int    `db:"user_id"`
	Title  string `db:"title"`
	Detail string `db:"detail"`
}

func (n NotificationParam) TableName() string {
	return "notifications"
}

func (n NotificationEntity) TableName() string {
	return "notifications"
}
