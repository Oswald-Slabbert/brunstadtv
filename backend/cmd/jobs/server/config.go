package server

import "github.com/aws/aws-sdk-go-v2/aws"

// ConfigData that can be passed to other things
type ConfigData struct {
	IngestBucket       string
	PackagingGroupID   string
	MediapackageRole   string
	MediapackageSource string
}

// GetIngestBucket that contains the new assets
// the MediapackageSource ARN should correspond to this bucket
func (c ConfigData) GetIngestBucket() *string {
	return aws.String(c.IngestBucket)
}

// GetPackagingGroup that the assets should be ingested into
func (c ConfigData) GetPackagingGroup() *string {
	return aws.String(c.PackagingGroupID)
}

// GetMediapackageRole ARN that should be used for ingesting the assets
func (c ConfigData) GetMediapackageRole() *string {
	return aws.String(c.MediapackageRole)
}

// GetMediapackageSource S3 ARN that the MediapackageRole has access to
func (c ConfigData) GetMediapackageSource() *string {
	return aws.String(c.MediapackageSource)
}
