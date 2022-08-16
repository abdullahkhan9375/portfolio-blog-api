package main

import (
	getblogpreviews "github.com/abdullahkhan9375/portfolio-blog-api/blog"

	getprojects "github.com/abdullahkhan9375/portfolio-blog-api/projects"

	getworkexperience "github.com/abdullahkhan9375/portfolio-blog-api/work"
	"github.com/gin-gonic/gin"
)

func main() {

	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }

	lRouter := gin.New()
	lRouter.Use(gin.Logger())

	lRouter.GET("/work", getworkexperience.GetWorkExperience)
	lRouter.GET("/projects", getprojects.GetProjects)

	// GET Paginated blog previews.
	lRouter.GET("/blogpreviews/:page", getblogpreviews.GetBlogPreviews)

	lRouter.Run(":8080")
}
