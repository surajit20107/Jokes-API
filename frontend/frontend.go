package frontend

import "github.com/gin-gonic/gin"

func FrontendRoutes(router *gin.Engine) {
  router.Static("/static", "./frontend/static")
  router.LoadHTMLGlob("./frontend/templates/*")
  
  router.GET("/", func(c *gin.Context) {
    c.HTML(200, "login.html", nil)
  })

  router.GET("/register", func(c *gin.Context) {
    c.HTML(200, "register.html", nil)
  })

  router.GET("/home", func(c *gin.Context) {
    c.HTML(200, "home.html", nil)
  })

  router.GET("/profile", func(c *gin.Context) {
    c.HTML(200, "profile.html", nil)
  })

  router.GET("/create", func(c *gin.Context) {
    c.HTML(200, "createJokes.html", nil)
  })

  // invalid route
  router.NoRoute(func(c *gin.Context) {
    c.HTML(404, "404.html", nil)
  })
}