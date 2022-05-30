package service

import (
	"context"
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func Upload(filename string) error {
	localFile := "D:\\simple-demo-main\\simple-demo-main\\public\\" + filename
	key := filename

	putPolicy := storage.PutPolicy{
		Scope: config.QiNiuCDN.Bucket,
	}
	mac := qbox.NewMac(config.QiNiuCDN.AccessKey, config.QiNiuCDN.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
