package graph

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/fampay/graph/model"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const (
	username = "root"
	password = "fampay123"
	hostname = "127.0.0.1:3306"
	dbName   = "YouTube"
)

type RequestConfig struct {
	apiKey          string
	requestInterval time.Duration
	Parameters      []string
}

var (
	cfg                        RequestConfig
	YouTubeSearchQueryResponse []*model.YouTubeResults
)

func SetConfig() {
	cfg.apiKey = "AIzaSyAJcJWkEVBmVCWTmKAg5yR_TAg5HGmp74Y"
	cfg.requestInterval = 10
	cfg.Parameters = []string{"id", "snippet"}
}

// SearchOnYouTube is the resolver for the SearchOnYouTube field.
func (r *queryResolver) SearchOnYouTube(ctx context.Context, input *model.QueryConfig) ([]*model.YouTubeResults, error) {
	//Setting the Request Config
	SetConfig()

	//Creating New YouTube Service
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(cfg.apiKey))

	if err != nil {

		return []*model.YouTubeResults{}, err
	}

	ticker := time.NewTicker(cfg.requestInterval * time.Second)

	tickerChan := make(chan bool)

	//Section for calling the YouTube API asynchronously after every 10 seconds
	go func() {
		for {
			select {
			case <-tickerChan:
				return
			case Time := <-ticker.C:
				fmt.Println("YouTube Videos fetched at: ", Time)
				fetchYouTubeVideos(youtubeService, input)

			}
		}
	}()

	time.Sleep(10 * time.Second)

	ticker.Stop()

	tickerChan <- true

	//Updating the SQL Database with the results obtained [WIP hence commented]
	//UpdateSQLDataBase()

	return YouTubeSearchQueryResponse, nil

}

func UpdateSQLDataBase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
	// Open the connection
	db, er := sql.Open("mysql", dsn)
	if er != nil {
		log.Fatalf("impossible to create the connection: %v", er)
	}
	sqlQuery := `CREATE TABLE IF NOT EXISTS YouTube(Video_Title text, Video_Description text, 
        Thumbnail_URL text, Date_Published text)`

	_, er = db.ExecContext(context.Background(), sqlQuery)

	if er != nil {
		log.Fatalf("Error %v when creating product table", er)

	}

	for _, resp := range YouTubeSearchQueryResponse {
		sqlQuery = "INSERT INTO YouTube(Video_Title, Video_Description,Thumbnail_URL,Date_Published) VALUES (?,?,?,?)"
		stmt, er := db.PrepareContext(context.Background(), sqlQuery)
		if er != nil {
			log.Fatalf("Error %v when preparing SQL statement", er)

		}

		_, er = stmt.ExecContext(context.Background(), resp.VideoTitle, resp.VideoDescription, resp.ThumbnailURL, resp.DateOfPublishing)
		if er != nil {
			log.Fatalf("Error %v when inserting row into products table", er)

		}

		fmt.Println("DataBase Updated Successfully!!")

	}
	defer db.Close()
}

func fetchYouTubeVideos(service *youtube.Service, input *model.QueryConfig) {

	//Calling the Search API of YouTube
	APICall := service.Search.List(cfg.Parameters).Q(*input.Query).Order("date").Type("video")

	response, err := APICall.Do()

	if err != nil {
		return
	}

	for _, video := range response.Items {

		//Populating YouTubeSearchQueryResponse with the data fetched
		YouTubeSearchQueryResponse = append(YouTubeSearchQueryResponse, &model.YouTubeResults{
			VideoTitle:       &video.Snippet.Title,
			VideoDescription: &video.Snippet.Description,
			ThumbnailURL:     &video.Snippet.Thumbnails.Default.Url,
			DateOfPublishing: &video.Snippet.PublishedAt,
		})

	}

}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
