package getworkexperience

import (
	"net/http"

	response "github.com/abdullahkhan9375/portfolio-blog-api/model"

	"github.com/gin-gonic/gin"
)

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

func GetWorkExperience(aContext *gin.Context) {
	//TODO: use aContext.JSON in prod. IndentedJSON is CPU intensive.

	var lResponse response.ServerResponse = response.ServerResponse{
		Message: "Cool",
		Data:    lWorkExperiences,
	}

	aContext.IndentedJSON(http.StatusOK, lResponse)
}
