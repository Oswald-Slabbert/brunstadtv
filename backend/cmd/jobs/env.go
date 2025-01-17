package main

import (
	"github.com/bcc-code/brunstadtv/backend/auth0"
	"github.com/bcc-code/brunstadtv/backend/crowdin"
	"github.com/bcc-code/brunstadtv/backend/members"
	"github.com/bcc-code/brunstadtv/backend/search"
	"github.com/bcc-code/brunstadtv/backend/utils"
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

type awsConfig struct {
	Region                string
	PackagingGroupARN     string
	MediapackageRoleARN   string
	MediapackageSourceARN string
	IngestBucket          string
	StorageBucket         string
}

type postgres struct {
	ConnectionString string
}

type directusConfig struct {
	BaseURL string
	Key     string
}

type firebase struct {
	ProjectID string
}

type cloudTasks struct {
	QueueID string
}

type envConfig struct {
	AWS               awsConfig
	Directus          directusConfig
	Port              string
	DeleteIngestFiles bool
	DB                postgres
	Algolia           search.Config
	Crowdin           crowdin.Config
	Firebase          firebase
	ImageCDNDomain    string
	Tracing           utils.TracingConfig
	CloudTasks        cloudTasks
	ServiceUrl        string
	Redis             utils.RedisConfig
	Auth0             auth0.Config
	Members           members.Config
}

func getEnvConfig() envConfig {

	deleteIngestFilesString := os.Getenv("DELETE_INGEST_FILES")
	// Error is intentionally ignored, if not set default to FALSE
	deleteIngestFilesStringBool, _ := strconv.ParseBool(deleteIngestFilesString)

	crowdinProjectIDs := lo.Map(strings.Split(os.Getenv("CROWDIN_PROJECT_IDS"), ","),
		func(s string, _ int) int {
			r, _ := strconv.ParseInt(s, 10, 64)
			return int(r)
		})

	return envConfig{
		Port:              os.Getenv("PORT"),
		DeleteIngestFiles: deleteIngestFilesStringBool,
		ImageCDNDomain:    os.Getenv("IMAGE_CDN_DOMAIN"),
		AWS: awsConfig{
			Region:                os.Getenv("AWS_DEFAULT_REGION"),
			PackagingGroupARN:     os.Getenv("AWS_PACKAGING_GROUP"),
			MediapackageRoleARN:   os.Getenv("AWS_MEDIAPACKAGE_ROLE"),
			MediapackageSourceARN: os.Getenv("AWS_MEDIAPACKAGE_SOURCE"),
			IngestBucket:          os.Getenv("AWS_INGEST_BUCKET"),
			StorageBucket:         os.Getenv("AWS_STORAGE_BUCKET"),
		},
		Directus: directusConfig{
			BaseURL: os.Getenv("DIRECTUS_URL"),
			Key:     os.Getenv("DIRECTUS_KEY"),
		},
		DB: postgres{
			ConnectionString: os.Getenv("DB_CONNECTION_STRING"),
		},
		Algolia: search.Config{
			AppID:  os.Getenv("ALGOLIA_APP_ID"),
			APIKey: os.Getenv("ALGOLIA_API_KEY"),
		},
		Crowdin: crowdin.Config{
			Token:      os.Getenv("CROWDIN_TOKEN"),
			ProjectIDs: crowdinProjectIDs,
		},
		Firebase: firebase{
			ProjectID: os.Getenv("FIREBASE_PROJECT_ID"),
		},
		Tracing: utils.TracingConfig{
			UptraceDSN:        os.Getenv("UPTRACE_DSN"),
			SamplingFrequency: os.Getenv("TRACE_SAMPLING_FREQUENCY"),
			TracePrettyPrint:  os.Getenv("TRACE_PRETTY"),
		},
		CloudTasks: cloudTasks{
			QueueID: os.Getenv("CLOUD_TASKS_QUEUE_ID"),
		},
		Redis: utils.RedisConfig{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Username: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
			Database: utils.AsInt(os.Getenv("REDIS_DATABASE")),
		},
		ServiceUrl: os.Getenv("SERVICE_URL"),
		Auth0: auth0.Config{
			ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
			ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
			Domain:       os.Getenv("AUTH0_DOMAIN"),
		},
		Members: members.Config{
			Domain: os.Getenv("MEMBERS_API_DOMAIN"),
		},
	}
}
