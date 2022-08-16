package getblogpreviews

import (
	"net/http"
	"strconv"

	response "pb-api/model"

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

func GetBlogPreviews(aContext *gin.Context) {
	pageNumber, err := strconv.Atoi(aContext.Param("page"))

	if err != nil {
		var lResponse = response.ServerResponse{
			Message: "Bad Request.",
			Data:    []string{},
		}
		aContext.IndentedJSON(http.StatusBadRequest, lResponse)
		return
	}

	var lPaginationEnd = PAGINATION_LIMIT * int8(pageNumber)
	if lPaginationEnd > int8(len(lBlogPreviews)+4) {
		var lResponse = response.ServerResponse{
			Message: "Page does not exist.",
			Data:    []string{},
		}
		aContext.IndentedJSON(http.StatusBadRequest, lResponse)
		return
	}

	var lResponse = response.ServerResponse{
		Message: "Cool",
		Data:    paginatedPayload(lBlogPreviews, int8(pageNumber)),
	}

	aContext.IndentedJSON(http.StatusOK, lResponse)
}
