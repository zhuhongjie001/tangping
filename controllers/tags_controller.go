package controllers

import (
	"fmt"

	"test.com.zhj/myblog/models"
)

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	tags := models.QueryArticleWithParam("tags")

	fmt.Println(models.HandleTagsListData(tags))

	this.Data["Tags"] = models.HandleTagsListData(tags)

	this.TplName = "tags.html"
}
