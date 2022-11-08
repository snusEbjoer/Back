package db

type DataBaseEvents interface {
	InitDB()
	IsConnected() error
	Check(id string) bool
	CreatePost(post Post) error
	GetPost(id string) (Post, error)
	GetUserPosts(userId string) ([]Post, error)
}

type Post struct {
	ICategory    string `json:"category,omitempty" bson:"category,omitempty"`
	ITitle       string `json:"title,omitempty" bson:"title,omitempty"`
	ICount       string `json:"count,omitempty" bson:"count,omitempty"`
	Text         string `json:"text,omitempty" bson:"text,omitempty"`
	PostId       string `json:"id,omitempty" bson:"id,omitempty"`
	UserId       string `json:"authorId,omitempty" bson:"authorId,omitempty"`
	ISOTimestamp string `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
