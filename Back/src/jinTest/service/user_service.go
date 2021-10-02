package user

import (
  "jinTest/db"
  "jinTest/entity"
)

type Service struct{}

type user entity.User

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
  if p.title != "" {
    tx = tx.Where("sei = ?", p.sei).Find(&user)
  }

  // カテゴリ
  if p.category != 0 {
    tx = tx.Where("mei = ?", p.mei).Find(&user)
  }

  return user, nil
}
