
package main

import (
  "fmt"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"

  // "github.com/go-sql-driver/mysql"
)

type Product struct {
  gorm.Model
  Name  string
  Price uint
}


func  ConnSqlite() (*gorm.DB, error){
  return gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
}

// func  ConnMysql() (*gorm.DB, error){
//   // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
//   dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
//   return gorm.Open(mysql.Open(dsn), &gorm.Config{})
// }


func TestGorm() {

  db, err := ConnSqlite()
  
  if err != nil {
    panic("failed to connect database")
  }

  //GORM支持Migration特性，支持根据Go Struct结构自动生成对应的表结构。
  db.AutoMigrate(&Product{})

  // 插入数据
  db.Create(&Product{Name: "苹果12", Price: 5200})
  db.Create(&Product{Name: "苹果11", Price: 4600})
  fmt.Println("插入两条数据")

  // 读取数据
  var product Product
  db.First(&product, 2) // find product with integer primary key
  fmt.Println(fmt.Sprintf("查找第二个产品，名称:%s，价格:%d",product.Name, product.Price))

  // 寻找id=2的产品
  db.First(&product, "id = ?", "2") 
  fmt.Println(fmt.Sprintf("查找第二个产品，名称:%s，价格:%d",product.Name, product.Price))


  // Update - update product's price to 200
  db.Model(&product).Update("Price", 4800)
  fmt.Println(fmt.Sprintf("更新第二个产品，名称:%s，更新后价格:%d",product.Name, product.Price))

  // Update - update multiple fields
  db.Model(&product).Updates(Product{Price: 4666, Name: "苹果11-美版"}) // non-zero fields
  fmt.Println(fmt.Sprintf("第二次更新第二个产品，更新后名称:%s，更新后价格:%d",product.Name, product.Price))

  db.Model(&product).Updates(map[string]interface{}{"Price": 3600, "Name": "苹果11-美版-二手"})
  fmt.Println(fmt.Sprintf("第三次更新第二个产品，更新后名称:%s，更新后价格:%d",product.Name, product.Price))

  // 删除
  db.Delete(&product, 2)
  fmt.Println("删除第二个产品成功")


}
