package models

//ItemBlog is
type ItemBlog struct {
	AuthorID string `bson:"author_id"`
	Title    string `bson:"title"`
	Content  string `bson:"content"`
}

//CreateUpdateBlog is
func (item *ItemBlog) CreateUpdateBlog(data *ItemBlog) interface{} {
	if item.AuthorID != "" {
		data.AuthorID = item.AuthorID
	}
	if item.Content != "" {
		data.Content = item.Content
	}
	if item.Title != "" {
		data.Title = item.Content
	}

	return data
}
