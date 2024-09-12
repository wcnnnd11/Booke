package flag

import "GVB_server/models"

func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex()
}
