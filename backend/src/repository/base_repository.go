package repository

import (
	"global"
)

func Save(row interface{}) {
	global.DataBase.Save(row)
}

func Create(row interface{}) {
	global.DataBase.Create(row)
}

func Take(dest interface{}) {
	global.DataBase.Take(dest)
}

func Find(obj interface{}, id uint) {
	global.DataBase.Find(obj, id)
}

func First(obj interface{}, example interface{}) {
	global.DataBase.First(obj, example)
}

func Delete(obj interface{}) {
	global.DataBase.Delete(obj)
}

func FindAll(obj interface{}) {
	global.DataBase.Find(obj)
}

func BatchDelete(query interface{}, value interface{}) {
	global.DataBase.Where(query).Delete(value)
}

// 条件查询
func FindBy(by interface{}, list interface{}) {
	global.DataBase.Where(by).Find(list)
}

// 查询单个
func FindByFirst(by interface{}, obj interface{}) {
	global.DataBase.Where(by).First(obj)
}
