package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

type Plugin struct {
	Endpoint string
	AK       string
	SK       string
	Bucket   string
	Region   string
	Key      string
	Source   string
	Delete   bool
	Zone     string
}

type KodoRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
}

func (p *Plugin) Exec() error {

	var err error

	if len(p.AK) == 0 || len(p.SK) == 0 || len(p.Bucket) == 0 {
		return errors.New("Must set access_key, secret_key and bucket")
	}

	// sign upload token
	putPolicy := storage.PutPolicy{
		Scope:      p.Bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)"}`,
	}

	mac := qbox.NewMac(p.AK, p.SK)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          getZone(p.Zone),
		UseHTTPS:      true,
		UseCdnDomains: true,
	}

	if p.Delete {
		bucketManager := storage.NewBucketManager(mac, &cfg)
		// delete ignore not exist error
		bucketManager.Delete(p.Bucket, p.Key)
	}

	formUploader := storage.NewFormUploader(&cfg)
	ret := KodoRet{}
	err = formUploader.PutFile(context.Background(), &ret, upToken, p.Key, p.Source, nil)
	if err != nil {
		return fmt.Errorf("[PutFile] error: %v", err)
	}

	fmt.Printf("ret: %+v\n", ret)
	return nil
}

func getZone(zoneId string) *storage.Zone {
	switch zoneId {
	case "z0":
		return &storage.ZoneHuadong
	case "z1":
		return &storage.ZoneHuabei
	case "z2":
		return &storage.ZoneHuanan
	case "na0":
		return &storage.ZoneBeimei
	default:
		return &storage.ZoneHuadong
	}
}
