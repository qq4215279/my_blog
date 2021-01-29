/**
 * tag_service
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/28 17:50
 */
package service

import (
	"backend/src/module"
	"backend/src/repository"
	"backend/src/utils"
	"strings"
)

// 添加标签
func AddTag(param module.Tag) ResponseBody {

	if utils.IsBlank(param.Name) {
		return NewParamEmptyResponseBody()
	}

	// 名称是否重复
	tag := new(module.Tag)
	repository.FindByFirst(&module.Tag{Name: param.Name}, tag)
	if tag.Id > 0 {
		return NewCustomErrorResponseBody("不能添加相同标签")
	}

	repository.Create(&param)

	return NewSimpleSuccessResponseBody()
}

// 获取所有标签
func GetAllTag() ResponseBody {
	var allTag []module.Tag
	repository.FindAll(&allTag)
	return NewSuccessResponseBody(allTag)
}

// 修改标签
func UpdateTag(param module.Tag) ResponseBody {
	var tagById module.Tag
	repository.FindById(&tagById, param.Id)
	if tagById.Id < 1 {
		return NewParamErrorResponseBody()
	}
	// 名称是否重复
	var tagByName module.Tag
	repository.FindByFirst(&module.Tag{Name: param.Name}, &tagByName)
	if len(tagByName.Name) > 0 && strings.EqualFold(tagByName.Name, param.Name) {
		return NewCustomErrorResponseBody("不能添加相同标签")
	}

	if !utils.IsBlank(param.Name) {
		tagById.Name = param.Name
	}

	repository.Save(tagById)
	return NewSimpleSuccessResponseBody()
}

// 删除标签
func DeleteTag(id uint) ResponseBody {
	var tag module.Tag
	repository.First(&tag, id)
	if utils.IsBlank(tag.Name) {
		return NewParamErrorResponseBody()
	}

	repository.Delete(tag)
	return NewSimpleSuccessResponseBody()
}
