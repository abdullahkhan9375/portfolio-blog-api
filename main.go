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
	Id          string   `json:"blogId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	Genre       string   `json:"genre"`
	Date        string   `json:"date"`
	TimeToRead  int8     `json:"timeToRead"`
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
	Id          string   `json:"id"`
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

var lBlogPreviews = []BlogPreview{
	{
		Id:          "post-1",
		Name:        "Post 1",
		Description: "The first blog post I have ever written wow.",
		Keywords:    []string{"cool", "first blog"},
		Genre:       "Opinion",
		Date:        "2022-01-01",
		TimeToRead:  5,
	},
	{
		Id:          "post-2",
		Name:        "Post 2",
		Description: "The second blog post I have ever written wow.",
		Keywords:    []string{"cool", "second blog"},
		Genre:       "Story",
		Date:        "2022-03-01",
		TimeToRead:  10,
	},
	{
		Id:          "post-3",
		Name:        "Post 3",
		Description: "The third blog post I have ever written wow.",
		Keywords:    []string{"strange", "cool"},
		Genre:       "Article",
		Date:        "2022-05-01",
		TimeToRead:  15,
	},
	{
		Id:          "post-4",
		Name:        "Post 4",
		Description: "The fourth blog post I have ever written wow.",
		Keywords:    []string{"strange", "cool"},
		Genre:       "Article",
		Date:        "2022-05-01",
		TimeToRead:  15,
	},
	{
		Id:          "post-5",
		Name:        "Post 1",
		Description: "The first blog post I have ever written wow.",
		Keywords:    []string{"cool", "first blog"},
		Genre:       "Opinion",
		Date:        "2022-01-01",
		TimeToRead:  5,
	},
	{
		Id:          "post-6",
		Name:        "Post 2",
		Description: "The second blog post I have ever written wow.",
		Keywords:    []string{"cool", "second blog"},
		Genre:       "Story",
		Date:        "2022-03-01",
		TimeToRead:  10,
	},
	{
		Id:          "post-7",
		Name:        "Post 3",
		Description: "The third blog post I have ever written wow.",
		Keywords:    []string{"strange", "cool"},
		Genre:       "Article",
		Date:        "2022-05-01",
		TimeToRead:  15,
	},
	{
		Id:          "post-8",
		Name:        "Post 4",
		Description: "The fourth blog post I have ever written wow.",
		Keywords:    []string{"strange", "cool"},
		Genre:       "Article",
		Date:        "2022-05-01",
		TimeToRead:  15,
	},
	{
		Id:          "post-9",
		Name:        "Post 4",
		Description: "The fourth blog post I have ever written wow.",
		Keywords:    []string{"strange", "cool"},
		Genre:       "Article",
		Date:        "2022-05-01",
		TimeToRead:  15,
	},
}

var lProjects = []Project{
	{
		Id:          "Project-1",
		Name:        "Notemaking App",
		Description: "Self-explanatory title",
		Features: []string{
			"React front-end",
			"Randomized backgrounds from Unsplash",
			"Mobile responsive",
		},
		TechStack: []string{
			"React",
			"Netlify",
			"CSS",
		},
		Links: &[]Tlink{
			{
				Name: "Github",
				Link: "https://github.com/abdullahkhan9375/Notemaker",
			},
		},
	},
	{
		Id:          "Project-2",
		Name:        "Breaking Bad Quiz",
		Description: "Can you make pure meth?",
		Features: []string{
			"React front-end",
			"Fetch API to grab images from an open-source API",
			"Mobile responsive",
		},
		TechStack: []string{
			"React",
			"Netlify",
			"CSS",
		},
		Links: &[]Tlink{
			{
				Name: "Github",
				Link: "https://github.com/abdullahkhan9375/breakingbadquiz",
			},
			{
				Name: "Website",
				Link: "https://breakingbadquiz.netlify.com",
			},
		},
	},
	{
		Id:          "Project-3",
		Name:        "Google Apps vs Apple Apps",
		Description: "An appstore comparison",
		Features: []string{
			"Exploratory Data analysis that tells a story",
			"Cutting edge python visualization libraries.",
			"Hosted on and powered by Kaggle.",
		},
		TechStack: []string{
			"Plotly",
			"Python",
		},
		Links: &[]Tlink{
			{
				Name: "Github",
				Link: "https://github.com/abdullahkhan9375/GooglePlaystore-vs-AppleAppstore",
			},
		},
	},
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
	var lEndIndex = min(lEndEntry, int8(len(lBlogPreviews)))
	for lIndex := lStartEntry; lIndex < lEndIndex; lIndex++ {
		lBlogEntry := aPayload[lIndex]
		if lBlogEntry.Id != "" {
			lPaginatedBlogPreviews[lCopyIndex] = aPayload[lIndex]
		}
		lCopyIndex++
	}
	return lPaginatedBlogPreviews
}

func getBlogPreviews(aContext *gin.Context) {
	pageNumber, err := strconv.Atoi(aContext.Param("page"))

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

func getProjects(aContext *gin.Context) {
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

	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }
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

	lRouter.GET("/projects", getProjects)

	// GET Paginated blog previews.
	lRouter.GET("/blogpreviews/:page", getBlogPreviews)

	lRouter.Run(":8080")
}
