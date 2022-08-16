package getprojects

import (
	"net/http"

	response "github.com/abdullahkhan9375/portfolio-blog-api/model"

	"github.com/gin-gonic/gin"
)

type Tlink struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type Project struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Features    []string `json:"features"`
	TechStack   []string `json:"techStack"`
	Links       *[]Tlink `json:"links"`
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

func GetProjects(aContext *gin.Context) {
	var lResponse response.ServerResponse = response.ServerResponse{
		Message: "Cool",
		Data:    lProjects,
	}
	aContext.IndentedJSON(http.StatusOK, lResponse)
}
