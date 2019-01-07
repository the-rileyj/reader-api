package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/the-rileyj/reader-api/functionality"
)

func getAwsCredentialsFromKeeper() (aws.Provider, error) {
}

// TODO:
// - Have context api handle timeout for requests to keyman -- Sorta soon
// - Have tests for functionality -- ASAP
// - Have routes use global aws polly struct

func main() {
	var (
		id, key string
		router  *gin.Engine
	)

	debug := flag.Bool("debug", false, "Set debug mode, use environmental variables to get AWS credentials")

	flag.Parse()

	if *debug {
		router = gin.Default()

		id = os.Getenv("id")
		key = os.Getenv("key")
	} else {
		router = gin.New()

		id, key, err := getAwsCredentialsFromKeeper()
	}

	router.GET("/all-episodes-basic-info", functionality.HandleAllEpisodeInfo)

	router.GET("/get-episode-basic-info/:episode", functionality.HandleGetEpisodeInfo)

	router.GET("/advanced-episode-info/:episode", functionality.HandleAdvancedEpisodeInfo)

	router.Run(":80")
}
