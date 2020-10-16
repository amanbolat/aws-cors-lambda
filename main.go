package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
	"log"
	"net/http/httptest"
	"strings"
)

type config struct {
	CorsOrigins []string `split_words:"true" required:"true"`
	CorsHeaders []string `split_words:"true" required:"true"`
}

var conf *config

func init() {
	conf = &config{}
	err := envconfig.Process("", conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	for i, v := range conf.CorsOrigins {
		conf.CorsOrigins[i] = strings.TrimSpace(v)
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   conf.CorsOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "ANY"},
		AllowedHeaders:   conf.CorsHeaders,
		AllowCredentials: true,
		Debug:            true,
	})

	rw := httptest.NewRecorder()
	converter := core.RequestAccessor{}
	r, err := converter.ProxyEventToHTTPRequest(req)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	corsHandler.HandlerFunc(rw, r)

	headers := make(map[string]string)

	for key := range rw.Header() {
		headers[key] = rw.Header().Get(key)
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         headers,
		Body:            "",
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
