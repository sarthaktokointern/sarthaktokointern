// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type QueryConfig struct {
	Query *string `json:"query"`
}

type YouTubeResults struct {
	VideoTitle       *string `json:"VideoTitle"`
	VideoDescription *string `json:"VideoDescription"`
	DateOfPublishing *string `json:"DateOfPublishing"`
	ThumbnailURL     *string `json:"ThumbnailURL"`
}