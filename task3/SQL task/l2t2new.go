//1.新增用户：创建用户并默认开启激活状态
func CreateUser(db *gorm.DB, name, email string) (*User, error) {
// 你的实现
  user := &User{
    Name: name,
    Email: email,
    Status: "active",
    }
  if err := db.Create(user).Error; err != nil{
    return nil,err
    }
  return user, nil
}
//2.模糊查询：根据邮箱模糊查询用户列表（支持分页）
func SearchUsersByEmail(db *gorm.DB, emailPattern string, page, size int) ([]User, error) {
    // 你的实现
    if page <= 0{
      page = 1
    }
    if size <= 0{
      size = 10
    }
    offset := (page - 1) * size
    var users []User
    if err := db,Where("email Like ?", "%"+emailPattern+"%").Offset(offset).Limit(size).Find(&users).Error;
    err != nil{
      return nil, err
    }
    return users,nil
}
//3.批量更新状态：批量更新用户状态
func UpdateUserStatus(db *gorm.DB, ids []uint, status string) error {
    // 你的实现
    if len(ids) == 0{
    return nil
    }
    return db.Model(&User{}).Where("id IN ?", ids).Update("status",status).Error
  }
}
//4.删除过期用户：删除超过 30 天未登录的用户
func DeleteInactiveUsers(db *gorm.DB) error {
    // 你的实现（注意：软删除将在进阶模块讲解）
    llogintime := time.Now().AddDate(0, 0, -30)
    return db.Where("last_login_at IS NULL OR last_login_at < ?",llogintime).Delete(&User{}).Error
}
