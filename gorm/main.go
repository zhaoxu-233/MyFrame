package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//创建mysql数据库的链接
	//db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/your_dbname"))
	if err != nil {
		panic("failed to connect database")
	}
	//重新声明db，覆盖之前的的，在控制台打印出执行语句
	db = db.Debug()
	// 迁移 schema
	//建表
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // 根据整型主键查找
	//使用这种就行，不考虑第一种使用方法
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	//更新price和code两个字段，相当于sql语句的` SET `price`=200,`code`="F42"
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&product, 1)
}
