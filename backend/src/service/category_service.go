/**
 * category_service
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/29 17:05
 */
package service

import (
	"backend/src/module"
	"backend/src/repository"
	"backend/src/utils"
	"reflect"
)

// 添加分类
func AddCategory(param module.Category) ResponseBody {

	if utils.IsBlank(param.Name) || utils.IsBlank(param.Describe) {
		return NewParamEmptyResponseBody()
	}

	var category module.Category
	repository.FindByFirst(&module.Category{Name: param.Name}, &category)
	if category.Id > 0 {
		return NewCustomErrorResponseBody("已存在当前分类")
	}

	repository.Create(&param)

	return NewSimpleSuccessResponseBody()
}

// 获取所有分类
func GetAllCategory() ResponseBody {
	var allCategory []module.Category
	repository.FindAll(&allCategory)
	return NewSuccessResponseBody(allCategory)
}

// 修改分类信息
func UpdateCategory(param module.Category) ResponseBody {

	var category module.Category
	repository.FindById(&category, param.Id)

	if reflect.DeepEqual(category, module.Category{}) {
		return NewParamErrorResponseBody()
	}
	if utils.IsBlank(param.Describe) {
		return NewCustomErrorResponseBody("描述内容不能为空")
	}

	category.Describe = param.Describe
	repository.Save(&category)

	return NewSimpleSuccessResponseBody()
}

// 删除分类
func DeleteCategory(categoryId uint) ResponseBody {

	var category module.Category
	repository.FindById(&category, categoryId)

	if reflect.DeepEqual(category, module.Category{}) {
		return NewParamErrorResponseBody()
	}

	repository.Delete(&category)

	return NewSimpleSuccessResponseBody()
}
