package aws

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	. "github.com/mlabouardy/komiser/services/aws"
	. "github.com/mlabouardy/komiser/services/cache"
)

type AWSHandler struct {
	cfg   aws.Config
	cache Cache
	aws   AWS
}

func NewAWSHandler(cfg aws.Config, cache Cache) *AWSHandler {
	awsHandler := AWSHandler{
		cfg:   cfg,
		cache: cache,
		aws:   AWS{},
	}
	return &awsHandler
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
