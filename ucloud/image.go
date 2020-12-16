package ucloud

import (
	"fmt"
	"ucloudmanager/log"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
)

// CreateImage 创建镜像
func CreateImage(uhostClient *uhost.UHostClient, imageName *string, uhostID *string) error {

	req := uhostClient.NewCreateCustomImageRequest()
	req.ImageName = imageName
	req.UHostId = uhostID

	resp, err := uhostClient.CreateCustomImage(req)
	if err != nil {
		log.Infoln("[ERROR]", err)
		return err
	}

	log.Infoln("[RESPONSE]", resp)
	return nil

}

// DeleteImage 删除镜像
func DeleteImage(uhostClient *uhost.UHostClient, imageID *string) error {

	req := uhostClient.NewTerminateCustomImageRequest()
	req.ImageId = imageID

	resp, err := uhostClient.TerminateCustomImage(req)
	if err != nil {
		log.Infoln("[ERROR]", err)
		return err
	}

	log.Infoln("[RESPONSE]", resp)
	return nil

}

// GetImages 获取镜像列表
func GetImages(uhostClient *uhost.UHostClient) []uhost.UHostImageSet {

	req := uhostClient.NewDescribeImageRequest()

	resp, err := uhostClient.DescribeImage(req)
	if err != nil {
		log.Infoln("[ERROR]", err)
		return nil
	}

	for i := 0; i < len(resp.ImageSet); i++ {
		fmt.Printf("ImageName: %s ImageID: %s", resp.ImageSet[i].ImageId, resp.ImageSet[i].ImageName)
	}

	return resp.ImageSet
}
