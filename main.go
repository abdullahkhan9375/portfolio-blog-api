package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var PAGINATION_LIMIT int8 = 4

type BlogPreview struct {
	Id          int      `json:"blogId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	Genre       string   `json:"genre"`
	Date        string   `json:"date"`
	TimeToRead  int8     `json:"timeToRead"`
	BlogId      int8     `json:"blogPostId"`
}

type Tlink struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type ServerResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Project struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Features    []string `json:"features"`
	TechStack   []string `json:"techStack"`
	Links       *[]Tlink `json:"links"`
}

type WorkExperience struct {
	WorkId         int
	CompanyName    string   `json:"company"`
	Position       string   `json:"position"`
	Responsibities []string `json:"responsibility"`
	FromDate       string   `json:"fromDate"`
	ToDate         string   `json:"toDate"`
}

func min(a int8, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

// TODO: Find a good way to paginate the last few entries.
func paginatedPayload(aPayload []BlogPreview, pageNumber int8) []BlogPreview {
	lPaginatedBlogPreviews := make([]BlogPreview, PAGINATION_LIMIT)
	var lEndEntry int8 = PAGINATION_LIMIT * pageNumber //12
	var lStartEntry int8 = lEndEntry - PAGINATION_LIMIT
	var lCopyIndex int8 = 0
	var lEndIndex = min(lEndEntry, int8(len(aPayload)))
	for lIndex := lStartEntry; lIndex < lEndIndex; lIndex++ {
		lPaginatedBlogPreviews[lCopyIndex] = aPayload[lIndex]
		lCopyIndex++
	}
	return lPaginatedBlogPreviews
}

func getBlogPreviews(aContext *gin.Context, aDB *sql.DB) {
	pageNumber, err := strconv.Atoi(aContext.Param("page"))

	lRows, err := aDB.Query("SELECT * FROM blogPreview")
	if err != nil {
		log.Fatalf("An error occured") // TODO: Add error handling.
	}

	defer lRows.Close()

	lBlogPreviews := make([]BlogPreview, 0)

	for lRows.Next() {
		var blogId int
		var blogName string
		var blogDesc string
		var blogKeyword1 string
		var blogKeyword2 string
		var blogKeyword3 string
		var blogGenre string
		var blogDateCreated string
		var blogTimeToRead int
		var blogPostId int

		lRows.Scan(
			&blogId,
			&blogName,
			&blogDesc,
			&blogKeyword1,
			&blogKeyword2,
			&blogKeyword3,
			&blogGenre,
			&blogDateCreated,
			&blogTimeToRead,
			&blogPostId,
		)
		lBlogPreviews = append(lBlogPreviews, BlogPreview{
			Id:          blogId,
			Name:        blogName,
			Description: blogDesc,
			Keywords:    []string{blogKeyword1, blogKeyword2, blogKeyword3},
			Genre:       blogGenre,
			Date:        blogDateCreated,
			TimeToRead:  int8(blogTimeToRead),
			BlogId:      int8(blogPostId),
		})
	}

	if err != nil {
		var lResponse = ServerResponse{
			Message: "Bad Request.",
			Data:    []string{},
		}
		aContext.IndentedJSON(http.StatusBadRequest, lResponse)
		return
	}

	var lPaginationEnd = PAGINATION_LIMIT * int8(pageNumber)
	if lPaginationEnd > int8(len(lBlogPreviews)+4) {
		var lResponse = ServerResponse{
			Message: "Page does not exist.",
			Data:    []string{},
		}
		aContext.IndentedJSON(http.StatusBadRequest, lResponse)
		return
	}

	var lResponse = ServerResponse{
		Message: "Cool",
		Data:    paginatedPayload(lBlogPreviews, int8(pageNumber)),
	}

	aContext.IndentedJSON(http.StatusOK, lResponse)
}

func getProjects(aContext *gin.Context, aDB *sql.DB) {
	lRows, err := aDB.Query("SELECT * FROM project")
	if err != nil {
		log.Fatalf("An error occured") // TODO: Add error handling.
	}

	defer lRows.Close()

	lProjects := make([]Project, 0)

	for lRows.Next() {
		var projectId int
		var projectName string
		var projectDescription string
		var projectFeature1 string
		var projectFeature2 string
		var projectFeature3 string
		var projectGithub string
		var projectWebsite string

		lRows.Scan(
			&projectId,
			&projectName,
			&projectDescription,
			&projectFeature1,
			&projectFeature2,
			&projectFeature3,
			&projectGithub,
			&projectWebsite,
		)
		lProjects = append(lProjects, Project{
			Id:          projectId,
			Name:        projectName,
			Description: projectDescription,
			Features:    []string{projectFeature1, projectFeature2, projectFeature3},
			Links: &[]Tlink{{
				Name: "Github",
				Link: projectGithub,
			},
				{
					Name: "Website",
					Link: projectWebsite,
				}},
		})
	}
	var lResponse ServerResponse = ServerResponse{
		Message: "Cool",
		Data:    lProjects,
	}

	aContext.IndentedJSON(http.StatusOK, lResponse)
}

func getWorkExperience(aContext *gin.Context, aDB *sql.DB) {
	lRows, err := aDB.Query("SELECT * FROM workexperience")
	if err != nil {
		log.Fatalf("An error occured") // TODO: Add error handling.
	}

	defer lRows.Close()

	lWorkExperiences := make([]WorkExperience, 0)

	for lRows.Next() {
		var workId int
		var companyName string
		var position string
		var responsibility1 string
		var responsibility2 string
		var responsibility3 string
		var fromDate string
		var toDate string

		lRows.Scan(
			&workId,
			&companyName,
			&position,
			&responsibility1,
			&responsibility2,
			&responsibility3,
			&fromDate,
			&toDate,
		)
		fmt.Println(workId, companyName, position)
		lWorkExperiences = append(lWorkExperiences, WorkExperience{
			WorkId:         workId,
			CompanyName:    companyName,
			Position:       position,
			Responsibities: []string{responsibility1, responsibility2, responsibility3},
			FromDate:       fromDate,
			ToDate:         toDate,
		})
	}
	var lResponse ServerResponse = ServerResponse{
		Message: "Cool",
		Data:    lWorkExperiences,
	}
	aContext.IndentedJSON(http.StatusOK, lResponse)
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	lPostGresUserName := os.Getenv("POSTGRES_USERNAME")
	lPostGresPassword := os.Getenv("POSTGRES_PASSWORD")
	lPostGresPORT, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	lPostGresHost := os.Getenv("POSTGRES_HOST")
	lPostGresDB := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		lPostGresHost, lPostGresPORT, lPostGresUserName, lPostGresPassword, lPostGresDB)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	lRouter := gin.New()
	lRouter.Use(gin.Logger())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	lRouter.Use(cors.New(config))

	lRouter.GET("/work", func(aContext *gin.Context) {
		getWorkExperience(aContext, db)
	})

	lRouter.GET("/projects", func(aContext *gin.Context) {
		getProjects(aContext, db)
	})

	// GET Paginated blog previews.
	lRouter.GET("/blogpreviews/:page", func(aContext *gin.Context) {
		getBlogPreviews(aContext, db)
	})

	lRouter.Run(":8080")
}
