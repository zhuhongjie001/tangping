package controllers

import (
	"test.com.zhj/myblog/models"
)

/**
 * 相册控制器
 */
type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		err.Error()
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
	//this.TplName = "album.html"
}
