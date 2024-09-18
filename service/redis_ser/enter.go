package redis_ser

const (
	articleLookPrefix    = "article_look"
	articleCommentPrefix = "article_comment_count"
	articleDiggPrefix    = "article_digg"
	commentDiggPrefix    = "comment_digg"
)

func NewDigg() CountDB {
	return CountDB{
		Index: articleDiggPrefix,
	}
}
func NewArticleLook() CountDB {
	return CountDB{
		Index: articleLookPrefix,
	}
}
func NewCommentCount() CountDB {
	return CountDB{
		Index: articleCommentPrefix,
	}
}
func NewCommentDigg() CountDB {
	return CountDB{
		Index: commentDiggPrefix,
	}
}
