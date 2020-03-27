package models


//ItemBlog is
type ItemBlog struct {
	AuthorID string  `bson:"author_id"`
	Title string `bson:"title"`
	Content string `bson:"content"`
}