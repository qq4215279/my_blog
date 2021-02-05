/**
 * upload_file_handler
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/2/1 16:33
 */
package handler

import (
	"backend/src/service"
	"backend/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 上传文件
func UploadFile(ctx *gin.Context) {
	mdFile, err := ctx.FormFile("mdFile")
	if err != nil {
		WrapperResponseBody(ctx, service.NewCustomErrorResponseBody("mdFile not found"))
		return
	}

	mdFileName := mdFile.Filename
	fmt.Println("mdFileName: ", mdFileName)
	path := utils.GetUploadFilePath()

	// 上传
	err = ctx.SaveUploadedFile(mdFile, path+"/"+mdFileName)
	if err != nil {
		WrapperResponseBody(ctx, service.NewCustomErrorResponseBody("upload fail"))
		return
	}

	//utils.InitEngine().Markdown(mdFileName, mdFile)

	WrapperResponseBody(ctx, service.NewSimpleSuccessResponseBody())
}

// 下载文件
func DownloadFile(ctx *gin.Context) {
	fileName := ctx.Query("fileName")

	path := utils.GetUploadFilePath() + "/" + fileName

	//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	disposition := fmt.Sprintf("attachment; filename=%s", fileName)
	ctx.Writer.Header().Add("Content-Disposition", disposition)
	ctx.Writer.Header().Add("Content-Type", "application/octet-strea")

	// 下载
	ctx.File(path)
}

// 图片上传
func UploadPicture(ctx *gin.Context) {
	picFile, err := ctx.FormFile("pic")
	if err != nil {
		WrapperResponseBody(ctx, service.NewCustomErrorResponseBody("pic not found"))
		return
	}

	picName := picFile.Filename
	path := utils.GetUploadPicturePath()
	fmt.Println("picName: ", path)

	// 上传
	err = ctx.SaveUploadedFile(picFile, path+"/"+picName)
	if err != nil {
		WrapperResponseBody(ctx, service.NewCustomErrorResponseBody("upload fail"))
		return
	}

	WrapperResponseBody(ctx, service.NewSimpleSuccessResponseBody())
}

// 查看图片
func FindPicture(ctx *gin.Context) {
	pictureName := ctx.Query("picName")

	path := utils.GetUploadPicturePath() + "/" + pictureName

	// 下载
	ctx.File(path)
}
