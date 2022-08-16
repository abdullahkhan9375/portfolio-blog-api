package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	Company          string   `json:"company"`
	Position         string   `json:"position"`
	Responsibilities []string `json:"description"`
	FromDate         string   `json:"fromDate"`
	ToDate           string   `json:"toDate"`
}

var lWorkExperiences = []WorkExperience{
	{
		Company:  "Proxima Capital",
		Position: "Junior SWE",
		Responsibilities: []string{
			"Worked on PAS.",
			"Enjoyed free lunches.",
			"Had a lot of fun.",
		},
		FromDate: "2022/01/04",
		ToDate:   "Present",
	},
	{
		Company:  "EY",
		Position: "PACE Developer",
		Responsibilities: []string{
			"Worked on Maskforce.",
			"Had a lot of fun.",
			"I was also a team lead.",
		},
		FromDate: "2021/06/01",
		ToDate:   "2021/12/10",
	},
	{
		Company:  "Hearing Power",
		Position: "Data Intern",
		Responsibilities: []string{
			"Worked on Tinnibot.",
			"Had a lot of fun.",
			"Enjoyed working from home.",
		},
		FromDate: "2021/01/01",
		ToDate:   "2021/05/01",
	},
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

func getWorkExperience(aContext *gin.Context) {
	//TODO: use aContext.JSON in prod. IndentedJSON is CPU intensive.

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

	lRouter := gin.New()
	lRouter.Use(gin.Logger())

	lRouter.GET("/work", getWorkExperience)
	lRouter.GET("/projects", getProjects)

	// GET Paginated blog previews.
	lRouter.GET("/blogpreviews/:page", getBlogPreviews)

	lRouter.Run(":8080")
}
