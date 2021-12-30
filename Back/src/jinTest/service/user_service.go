package user

import (
  "github.com/gin-gonic/gin"
  "jinTest/db"
  "jinTest/entity"
)

type Service struct{}

type User entity.User

type Parameter struct {
  sei string
  mei string
}


// 検索
func (s Service) Search(sei string, mei string) ([]User, error) {

  // DB接続
  db := db.GetDB()

  // 本モデルから作成
  var user []User

  // パラメータセット
  p := Parameter{sei, mei}

  // DB接続確認
  if err := db.Take(&user).Error; err != nil {
    return nil, err
  }

  // 本検索クエリ
  tx := db
  tx = tx.Find(&user)

  // タイトル
  if p.sei != "" {
    tx = tx.Where("name_sei = ?", p.sei).Find(&user)
  }

  // カテゴリ
  if p.mei != "" {
    tx = tx.Where("name_mei = ?", p.mei).Find(&user)
  }

  return user, nil
}

//更新 update
func (s Service) UpdateByID(id int, c *gin.Context) (User, error) {
  db := db.GetDB()
  var u User
  if err := db.Where("id = ?", id).First(&u).Error; err != nil {
    return u, err
  }
  if err := c.BindJSON(&u); err != nil {
    return u, err
  }
  u.Id = uint(id)
  db.Save(&u)

  return u, nil
}
