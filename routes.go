package main

func initializeRoutes() {

  // Handle the index route
  // router:=getRouter(true)
  router.GET("/", showIndexPage)
  router.GET("/article/view/:article_id", getArticle)
}