//假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
user := User{ name: "张三", age: 20, grade: "三年级"}
if err := db.Create(&user).Error;err != nil {
  }
//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
var users []User
db.Where("age > ?", 18).Find(&users)
//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"
db.Model(&User{}).Where("name = ?","张三").Updates(map[string]any{"grade": "四年级"})
//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
db.Where("age < ?", 15).Delete(User{})
err := db.First(&User{}, user.ID).Error
if !errors.Is(err, gorm.ErrRecordNotFound){
  }
