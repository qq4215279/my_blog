/**
 * constants
 * @author liuzhen
 * @Description
 * @version 1.0.0 2021/1/28 16:56
 */
package repository

import (
	"backend/src/global"
)

func Create(row interface{}) {
	global.DataBase.Create(row)
}

func Save(row interface{}) {
	global.DataBase.Save(row)
}

func Take(dest interface{}) {
	global.DataBase.Take(dest)
}

func First(obj interface{}, example interface{}) {
	global.DataBase.First(obj, example)
}

func FindById(obj interface{}, id uint) {
	global.DataBase.Find(obj, id)
}

func FindAll(obj interface{}) {
	global.DataBase.Find(obj)
}

// 条件查询
func FindBy(by interface{}, list interface{}) {
	global.DataBase.Where(by).Find(list)
}

// 查询单个
func FindByFirst(by interface{}, obj interface{}) {
	global.DataBase.Where(by).First(obj)
}

func Delete(obj interface{}) {
	global.DataBase.Delete(obj)
}

func BatchDelete(query interface{}, value interface{}) {
	global.DataBase.Where(query).Delete(value)
}
