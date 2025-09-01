package frontend

import "github.com/gin-gonic/gin"

func FrontendRoutes(router *gin.Engine) {
  router.Static("/static", "./frontend/static")
  router.LoadHTMLGlob("./frontend/templates/*")
  
  router.GET("/", func(c *gin.Context) {
    c.HTML(200, "index.html", nil)
  })

  // invalid route
  router.NoRoute(func(c *gin.Context) {
    c.HTML(404, "404.html", nil)
  })
}