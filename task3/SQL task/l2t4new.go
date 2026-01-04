//练习 1：设计博客系统关联关系
//设计以下模型的关联关系：
//User：用户
//Post：文章
//Comment：评论
//Tag：标签
//要求：
//User Has Many Posts
//Post Belongs To User
//Post Has Many Comments
//Comment Belongs To Post
//Post Many to Many Tags
type User struct{
  ID      uint  `gorm: "primaryKey"`
  Name    string `gorm:"size:64;not null"`
  
  Posts []Post
}
type Post struct{
  ID      uint  `gorm: "primaryKey"`
  Title   string `gorm:"size:128;not null"`
  Content string `gorm:"size:1024"`
  
  UserID  uint
  User    User
  
  Comments []Comment
  Tags     []Tag `gorm: "many2many:post_tags"`
}
type Comment struct{
  ID      uint  `gorm: "primaryKey"`
  Content string `gorm:"size:1024;not null"`
  
  PostID uint
  Post Post
}
type Tag struct{
  ID      uint  `gorm: "primaryKey"`
  Name    string `gorm:"size:32;uniqueIndex;not null"`

  Posts    []Post `gorm: "many2many:post_tags"`

//练习 2：实现查询功能
//查询用户最新文章：查询用户发表的最新 10 篇文章（含标签）

func GetUserLatestPosts(db *gorm.DB, userID uint) ([]Post, error) {
    var Posts []Post
  err :=db.
  Preload("Tags").
  Where("user_id = ?",userID).
  Order("created_at desc").
  Limit(10).
  Find(&posts).Error
  return posts,err
    // 你的实现
}
//统计评论数量：使用 Preload + Count 统计每篇文章的评论数量
type PostWithCount struct{
  Post
  CommentCount
}
func GetPostsWithCommentCount(db *gorm.DB) ([]PostWithCount, error) {
    // 你的实现
    var Posts []Post
    err := db.Preload("Comments").Find(&posts).Error
    if err != nil{
      return nil,err
    }
    var result []PostWithCount
    for _, p := range posts{
      rsult := append(result,PostWithCount{
        Post: p,
        CommentCount: int64(len(p.Comments)),
      })
    }
    return result,nil
}
//练习 3：实现事务操作
//在事务中实现文章发布 + 标签绑定：

func PublishPostWithTags(db *gorm.DB, post *Post, tagIDs []uint) error {
    // 在事务中：
    // 1. 创建文章
    // 2. 绑定标签（Many to Many）
    // 3. 更新用户文章数量
    // 你的实现
  return db.Transaction(func(tx *gorm.DB) error{
    if err := tx.Create(post).Error; err != nil{
      return err
    }
    if len(tagIDs) > 0{
      var tags Tag[]
      if err := tx.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil{
        return err
      }
      if err := tx.Model(post).Association("Tags").append(tags); err != nil{
        return err
      }
    )

    if err := tx.Model(&User{}).
      Where("id = ?".post.UserID).
      Update("post_count",gorm.Expr("post_count +  ?",1)).Error; err != nil{
        return err
        }
      return nil
      })
}
//练习 4：实现软删除
//为评论新增软删除，提供"彻底清除"功能：

// 软删除评论
type Comment struct{
  ID        uint           `gorm:"primaryKey"`
  PostID    uint
  Content   string
  DeletedAt gorm.DeletedAt `gorm:" index"`
}
func SoftDeleteComment(db *gorm.DB, commentID uint) error {
    // 你的实现
  return db.Delete(&Comment{}, commentID).Error
}

// 彻底删除评论
func HardDeleteComment(db *gorm.DB, commentID uint) error {
    // 你的实现
  return db.Unscope().Delete(&Comment{}, commentID).Error
}
