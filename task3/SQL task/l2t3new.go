//创建一个 youngUsers scope，筛选年龄在 18-30 岁之间的用户，并实现分页查询。
func youngUsers() func(db *gorm.DB) *gorm.DB{
    return func(db *gorm.DB) *gorm.DB{
      return db.Where("age BETWEEN ? AND ?", 18, 30)
    }
}
func paginate(page ,size int) func(db *gorm.DB) *gorm.DB{
    return func(db *gorm.DB) *gorm.DB{
      if page <= 0 {
        page = 1
      }
      if size <= 0 {
        size = 10
      }
      offset := ( page - 1 ) * size
      return db.Offset(offset).Limit(size)
    }
}
var users []User
db.Scopes(youngUsers(),paginate(1,10)).Find(&users)
