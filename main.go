package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var PAGINATION_LIMIT int8 = 4

type BlogPreview struct {
	Name        string   `firestore:"name"`
	Description string   `firestore:"desc"`
	Keywords    []string `firestore:"keywords"`
	Genre       string   `firestore:"genre"`
	Date        string   `firestore:"dateCreated"`
	TimeToRead  int8     `firestore:"timeToRead"`
	PostId      int8     `firestore:"postId"`
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
	Name        string   `firestore:"name"`
	Description string   `firestore:"description"`
	Features    []string `firestore:"features"`
	Github      string   `firestore:"github"`
	Website     string   `firestore:"website"`
}

type WorkExperience struct {
	Name           string   `firestore:"name"`
	Position       string   `firestore:"position"`
	Responsibities []string `firestore:"responsibilities"`
	FromDate       string   `firestore:"fromDate"`
	ToDate         string   `firestore:"toDate"`
}

func min(a int8, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

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

func getBlogPreviews(aContext *gin.Context, aFireStore *firestore.Client, ctx context.Context) {
	pageNumber, err := strconv.Atoi(aContext.Param("page"))

	lDocuments, err := aFireStore.Collection("blogPreviews").DocumentRefs(ctx).GetAll()
	if err != nil {
		log.Fatalf("An error occured")
	}

	var lBlogPrev BlogPreview = BlogPreview{}
	var lData []BlogPreview = []BlogPreview{}

	for _, lDocument := range lDocuments {
		lDatum, _ := lDocument.Get(ctx)
		if err := lDatum.DataTo(&lBlogPrev); err != nil {
			fmt.Println(err)
		}
		lData = append(lData, lBlogPrev)
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
	if lPaginationEnd > int8(len(lData)+4) {
		var lResponse = ServerResponse{
			Message: "Page does not exist.",
			Data:    []string{},
		}
		aContext.IndentedJSON(http.StatusBadRequest, lResponse)
		return
	}

	var lResponse = ServerResponse{
		Message: "Cool",
		Data:    paginatedPayload(lData, int8(pageNumber)),
	}

	aContext.IndentedJSON(http.StatusOK, lResponse)
}

func getProjects(aContext *gin.Context, aFireStore *firestore.Client, ctx context.Context) {

	lDocuments, err := aFireStore.Collection("projects").DocumentRefs(ctx).GetAll()
	if err != nil {
		log.Fatalf("An error occured") // TODO: Add error handling.
	}

	var lProj Project = Project{}
	var lData []Project = []Project{}

	for _, lDocument := range lDocuments {
		lDatum, _ := lDocument.Get(ctx)
		if err := lDatum.DataTo(&lProj); err != nil {
			fmt.Println(err)
		}
		lData = append(lData, lProj)
	}
	var lResponse ServerResponse = ServerResponse{
		Message: "Cool",
		Data:    lData,
	}

	aContext.IndentedJSON(http.StatusOK, lResponse)
}

func getWorkExperience(aContext *gin.Context, aFireStore *firestore.Client, ctx context.Context) {

	lDocuments, err := aFireStore.Collection("workExperience").DocumentRefs(ctx).GetAll()
	if err != nil {
		log.Fatalf("An error occured") // TODO: Add error handling.
	}

	var lWorkExp WorkExperience = WorkExperience{}
	var lData []WorkExperience = []WorkExperience{}

	for _, lDocument := range lDocuments {
		lDatum, _ := lDocument.Get(ctx)
		if err := lDatum.DataTo(&lWorkExp); err != nil {
			fmt.Println(err)
		}
		lData = append(lData, lWorkExp)
	}
	var lResponse ServerResponse = ServerResponse{
		Message: "Cool",
		Data:    lData,
	}
	aContext.IndentedJSON(http.StatusOK, lResponse)
}

func main() {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	lFirestoreClient, err := app.Firestore(ctx)

	if err != nil {
		log.Fatal("error instantiating a firestore var")
	}

	lRouter := gin.New()
	lRouter.Use(gin.Logger())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	lRouter.Use(cors.New(config))

	lRouter.GET("/work", func(aContext *gin.Context) {
		getWorkExperience(aContext, lFirestoreClient, ctx)
	})

	lRouter.GET("/projects", func(aContext *gin.Context) {
		getProjects(aContext, lFirestoreClient, ctx)
	})

	lRouter.GET("/blogpreviews/:page", func(aContext *gin.Context) {
		getBlogPreviews(aContext, lFirestoreClient, ctx)
	})

	lRouter.Run(":8080")
}
