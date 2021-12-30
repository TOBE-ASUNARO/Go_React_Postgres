package user

import (
  "fmt"
  "github.com/gin-gonic/gin"
  user "jinTest/service"
  "net/http"
  "strconv"
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

// 更新
func (pc Controller) Update(c *gin.Context) {
  id := c.Params.ByName("id")
  idInt, _ := strconv.Atoi(id)
  var s user.Service
  p, err := s.UpdateByID(idInt, c)

  if err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
  } else {
    c.JSON(200, p)
  }
}
