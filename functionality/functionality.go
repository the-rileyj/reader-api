package functionality

import (
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/the-rileyj/KeyMan/keymanaging"
)

type rjCredentials struct {
	local bool
}

func (rjc *rjCredentials) Retrieve() (creds credentials.Value, err error) {
	if rjc.local {
		return credentials.Value{AccessKeyID: os.Getenv("id"), SecretAccessKey: os.Getenv("key")}, nil
	}
	awsIdRequest, err := http.NewRequest(
		http.MethodGet,
		"https://keys.therileyjohnson.com/key/aws-polly-id",
		nil,
	)

	if err != nil {
		return creds, err
	}

	awsKeyRequest, err := http.NewRequest(
		http.MethodGet,
		"https://keys.therileyjohnson.com/key/aws-polly-id",
		nil,
	)

	var (
		idResponseJSON, keyResponseJSON keymanaging.Response
	)

	if err != nil {
		return creds, err
	}

	return credentials.Value{AccessKeyID: os.Getenv("id"), SecretAccessKey: os.Getenv("key")}, nil
}

func (rjc *rjCredentials) IsExpired() bool { return false }

type basicEpisodeInfo struct {
	Description string `json:"description"`
	IframeLink  string `json:"iframeLink"`
	InfoLink    string `json:"infoLink"`
	Title       string `json:"title"`
}

type EpisodeVideoInfo struct {
	URL      string `json:"url"`
	Quality  string `json:"quality"`
	Filetype string `json:"filetype"`
}

type EpisodeUpdatedInfo struct {
	EpisodeVideoSources []EpisodeVideoInfo `json:"episodeVideoInfo"`
	NewDescription      string             `json:"newDescription"`
	NewTitle            string             `json:"newTitle"`
}

func init() {
	var err error

	if err != nil {
		panic(err)
	}
}
