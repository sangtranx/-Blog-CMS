package postmodel

const EntityName = "Post"

type Post struct {
	ID         int    `json:"-" gorm:"primaryKey"`
	Title      string `json:"title" gorm:"type:varchar(255);not null"`
	Content    string `json:"content" gorm:"type:text;not null"`
	AuthorID   int    `json:"author_id" gorm:"not null"`
	CategoryID *int   `json:"category_id" gorm:"default:null"`
	Status     string `json:"status" gorm:"type:enum('draft', 'published');default:'draft'"`
	Views      int    `json:"views" gorm:"default:0"`
	Likes      int    `json:"likes" gorm:"default:0"`
}

func (Post) TableName() string { return "posts" }

type PostCreate struct {
	ID         int    `json:"-" gorm:"primaryKey"`
	Title      string `json:"title" gorm:"type:varchar(255);not null"`
	Content    string `json:"content" gorm:"type:text;not null"`
	AuthorID   int    `json:"author_id" gorm:"not null"`
	CategoryID *int   `json:"category_id" gorm:"default:null"`
	Status     string `json:"status" gorm:"type:enum('draft', 'published');default:'draft'"`
}

func (PostCreate) TableName() string { return Post{}.TableName() }
