package src

import "os"

func GetCloudBucketImage() string {
	return os.Getenv("GOOGLE_CLOUD_BUCKET_IMAGE")
}
