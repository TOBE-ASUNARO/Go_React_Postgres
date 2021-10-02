package user

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  user "jinTest/service"
)

type Controller struct{}

// 検索 GET /books
func (pc Controller) Index(c *gin.Context) {
  // パラメータ取得
  sei := c.Query("sei")
  mei := c.Query("mei")

  // 検索処理
  var s user.Service
  p, err := s.Search(sei, mei)

  // 検索結果を返す
  if err != nil {
    c.AbortWithStatus(http.StatusNotFound)
    fmt.Println(err)
  } else {
    c.JSON(http.StatusOK, p)
  }
}
