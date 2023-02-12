package dao

type Comment struct {
	CommentId int64  `gorm:colum:comment_id`
	FloorId   int32  `gorm:colum:floor_id`
	ArticleId int64  `gorm:colum:article_id`
	UserId    int64  `gorm:colum:user_id`
	UserName  string `gorm:colum:user_name`
	Content   string `gorm:colum:content`
}
