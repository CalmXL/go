package main

/**
  什么是数据库 => 数据表的集合（MySQL）
		电商项目: Database
				用户  Data table
 				产品  Data table
				订饭  Data table

	用户表:
		id    username    password
		1     xulei       123456

	什么是数据（表）模型 => 基于编程语言层面上的描述数据表结构的一种方式
	type User struct {
		id        uint
		username  string
		password  string
	}

	为什么要有数据模型 => 我们有 orm 库需要语言来描述数据表，根据这个描述来对数据表进行数据的增删改查

	数据表 => 数据的增删改查的 => SQL语言的语句

	1. 连接数据库
	2. UPDATE `user` SET username = 'superCoder' WHERE id = 1  => SQL 语言的语句

	为什么需要 orm => Object Relational Mapping 对象关系映射

	type User struct {
		id        uint
		username  string
		password  string
	}

	用户表:
		id    username    password
		1     xulei       123456

	为什么要做这种映射呢?
		因为 SQL 语句比较难以编写与维护

	orm => 封装一套基于编程语言的操作数据库的方法APIs集合 => orm库的作用（Library）

	orm =>  1. 连接数据库
					2. 返回数据库的实例 db => database
					3. 可以通过实例去调用增删改查的方法

	db.model(User{})
		.Where("id = ?", 1)
		.Update("username", "superCoder")

	优点：
		1. 开发方便高效
		2. 代码整洁易读
	缺点：
		1. 性能可能会牺牲
		2. 过渡依赖 ORM, 会导致对 SQL 语句理解困难

	orm库 => Gorm => Go orm库
*/

import (
	"12_gorm/db"
	"12_gorm/model"
	"fmt"
)

// var newLogger = logger.New(
// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 	logger.Config{
// 		SlowThreshold:             time.Second, // Slow SQL threshold
// 		LogLevel:                  logger.Info, // Log level
// 		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
// 		ParameterizedQueries:      true,        // Don't include params in the SQL log
// 		Colorful:                  false,       // Disable color
// 	},
// )

// type GormModel struct {
// 	ID        uint         `json:"id" gorm:"primarykey"`
// 	CreatedAt time.Time    `json:"created_at"`
// 	UpdatedAt time.Time    `json:"updated_at"`
// 	DeletedAt sql.NullTime `json:"deletedAt" gorm:"index"`
// }
//
// type Student struct {
// 	// gorm.Model
// 	GormModel
// 	Name string `json:"name" gorm:"size:10;default:游客;"`
// 	Age  uint   `json:"age" gorm:"not null"`
// }

// type User struct {
// 	GormModel
// 	Username string `json:"username" gorm:"size:10;not null;unique"`
// 	Password string `json:"password" gorm:"size:10;not null;unique"`
// }

// func (s *Student) BeforeSave(db *gorm.DB) (err error) {
// 	fmt.Println("Before Save")
// 	return err
// }
//
// func (s *Student) BeforeCreate(db *gorm.DB) (err error) {
// 	fmt.Println("Before Create")
// 	return err
// }
//
// func (s *Student) AfterCreate(db *gorm.DB) (err error) {
// 	fmt.Println("After Create")
// 	return err
// }
//
// func (s *Student) AfterSave(db *gorm.DB) (err error) {
// 	fmt.Println("After Save")
// 	return err
// }

func main() {
	// 1. 连接数据库
	dbName := "golang"
	// dsn := "root:xL1210...@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: newLogger,
	// })
	//

	db, err := db.Connect(dbName)

	if err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Println(db)

	// 返回前端 { "id": 1, "name": "xulei", "age": 28 }

	// db.AutoMigrate(&Student{}, &User{}) // 同步表功能
	// db.AutoMigrate(&Student{}) // 同步表功能

	// s1 := Student{
	// 	Name: "Mike",
	// 	Age:  18,
	// }
	//
	// s2 := Student{
	// 	Name: "Cyrstal",
	// 	Age:  22,
	// }
	//
	// s3 := Student{
	// 	Name: "Jack",
	// 	Age:  15,
	// }

	// 创建单条数据
	// result := db.Create(&s1)
	// fmt.Println(result)

	// 同时创建多条数据
	// students := []*Student{
	// 	&s1,
	// 	&s2,
	// 	&s3,
	// }
	// result := db.Create(students)
	//
	// fmt.Println(result.RowsAffected)

	// s1 := Student{
	// 	Name: "Mendy",
	// 	Age:  13,
	// }
	//
	// db.Select("age").Create(&s1)

	// db.Omit("Name").Create(&s1)

	// 批量添加
	s1 := model.Student{
		Name: "Mike",
		Age:  18,
	}

	s2 := model.Student{
		Name: "Cyrstal",
		Age:  22,
	}

	s3 := model.Student{
		Name: "Jack",
		Age:  15,
	}
	//
	students := []model.Student{s1, s2, s3}
	// db.Create(&students)
	// db.CreateInBatches(students, 1) // 创建 -> 分条语SQL句

	/**
	MySQL 进行数据操作的时候
			// 开始事务
			BeforeSave
			BeforeCreate
			// 关联前的 save
			// 插入记录至 db
			// 关联后的 save
			AfterCreate
			AfterSave
			// 提交或回滚事务
	*/

	// for _, s := range students {
	// 	fmt.Println(s.ID)
	// }

	// m1 := map[string]any{
	// 	"name": "Menda",
	// 	"age":  21,
	// }

	// db.Model(&model.Student{}).Create(m1)

	db.Create(students)
}
