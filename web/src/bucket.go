package src

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"cloud.google.com/go/storage"
)

type Image struct {
	Bucket  string
	Object  string
	Content string
}

func NewProfileImage(uid, content string) Image {
	return Image{
		Bucket:  GetCloudBucketImage(),
		Object:  fmt.Sprintf("profile/%s.png", uid),
		Content: content,
	}
}

func (i Image) Upload(ctx context.Context) error {
	data, err := base64.StdEncoding.DecodeString(i.Content)
	if err != nil {
		return err
	}
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	bucket := client.Bucket(i.Bucket)
	wc := bucket.Object(i.Object).NewWriter(ctx)
	wc.ContentType = "image/png"
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	if _, err := wc.Write(data); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

func (i Image) Download(ctx context.Context) ([]byte, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	reader, err := client.Bucket(i.Bucket).Object(i.Object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(reader); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
