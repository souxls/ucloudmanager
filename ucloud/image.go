package ucloud

import (
	"fmt"
	"ucloudmanager/log"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// CreateImage 创建镜像
func CreateImage(imageName *string, uhostID *string) error {

	req := Uclient.NewCreateCustomImageRequest()
	req.ImageName = imageName
	req.UHostId = uhostID

	resp, err := Uclient.CreateCustomImage(req)
	if err != nil {
		log.Infoln("[ERROR]", err)
		return err
	}

	log.Infoln("[RESPONSE]", resp)
	return nil

}

// DeleteImage 删除镜像
func DeleteImage(imageID *string) error {

	req := Uclient.NewTerminateCustomImageRequest()
	req.ImageId = imageID

	resp, err := Uclient.TerminateCustomImage(req)
	if err != nil {
		log.Infoln("[ERROR]", err)
		return err
	}

	log.Infoln("[RESPONSE]", resp)
	return nil

}

// GetImages 获取镜像列表
func GetImages() []uhost.UHostImageSet {

	req := Uclient.NewDescribeImageRequest()
	req.ImageType = ucloud.String("Custom")

	resp, err := Uclient.DescribeImage(req)
	if err != nil {
		log.Infoln("[ERROR]", err)
		return nil
	}

	for i := 0; i < len(resp.ImageSet); i++ {
		imageInfo := resp.ImageSet[i]
		fmt.Printf("ImageID: %s ImageName: %s State: %s \n", imageInfo.ImageId, imageInfo.ImageName, imageInfo.State)
	}

	return resp.ImageSet
}
