/**
Homepage: https://github.com/ucloud/ucloud-sdk-go
Examples: https://github.com/ucloud/ucloud-sdk-go/tree/master/examples
*/

package ucloud

import (
	"fmt"
	"ucloudmanager/config"
	"ucloudmanager/log"

	"github.com/sethvargo/go-password/password"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// CreateHost 创建主机
func CreateHost(uhostClient *uhost.UHostClient, name, imageID, zone string) error {

	h := config.Cfg.Host
	pwd, err := password.Generate(16, 5, 5, false, false)
	req := uhostClient.NewCreateUHostInstanceRequest()
	req.Zone = ucloud.String(zone)
	req.ImageId = ucloud.String(imageID)
	req.Password = ucloud.String(pwd)
	req.LoginMode = ucloud.String(h.LoginMode)
	req.Name = ucloud.String(name)
	req.ChargeType = ucloud.String(h.ChargeType)
	req.CPU = ucloud.Int(h.CPU)
	req.Memory = ucloud.Int(h.Memory)
	req.NetCapability = ucloud.String(h.NetCapability)
	req.MachineType = ucloud.String(h.MachineType)
	req.MinimalCpuPlatform = ucloud.String(h.MinimalCPUPlatform)
	req.NetworkInterface = []uhost.CreateUHostInstanceParamNetworkInterface{
		{
			EIP: &uhost.CreateUHostInstanceParamNetworkInterfaceEIP{
				Bandwidth:    ucloud.Int(h.NetworkInterface.Bandwidth),
				PayMode:      ucloud.String(h.NetworkInterface.PayMode),
				OperatorName: ucloud.String(h.NetworkInterface.OperatorName),
				GlobalSSH: &uhost.CreateUHostInstanceParamNetworkInterfaceEIPGlobalSSH{
					Port: ucloud.Int(22),
				},
			},
		},
	}
	req.Disks = []uhost.UHostDisk{
		{
			IsBoot: ucloud.String(h.Disks.IsBoot),
			Size:   ucloud.Int(h.Disks.Size),
			Type:   ucloud.String(h.Disks.Type),
		},
	}

	resp, err := uhostClient.CreateUHostInstance(req)
	if err != nil {
		log.Infoln("[ERROR]", err)
		return err
	}
	log.Infoln("[RESPONSE]", resp)
	return nil
}

// StartHost 启动 uhost
func StartHost(uhostClient *uhost.UHostClient, uHostID *string) error {

	req := uhostClient.NewStartUHostInstanceRequest()
	req.UHostId = uHostID

	resp, err := uhostClient.StartUHostInstance(req)
	if err != nil {
		log.Infoln("[ERROR]", err)
		return err
	}

	log.Infoln("[RESPONSE]", resp)
	return nil
}

// StopHost 停止 uHost
func StopHost(uhostClient *uhost.UHostClient, uhostID *string) error {

	req := uhostClient.NewStopUHostInstanceRequest()
	req.UHostId = uhostID

	resp, err := uhostClient.StopUHostInstance(req)
	if err != nil {
		log.Infoln("[ERROR]", err)
		return err
	}

	log.Infoln("[RESPONSE]", resp)
	return nil
}

// DeleteHost 删除 uhost
func DeleteHost(uhostClient *uhost.UHostClient, uhostID *string) error {

	req := uhostClient.NewTerminateUHostInstanceRequest()
	req.UHostId = uhostID

	resp, err := uhostClient.TerminateUHostInstance(req)
	if err != nil {
		log.Infoln("[ERROR]", err)
		return err
	}

	log.Infoln("[RESPONSE]", resp)
	return nil
}

// GetHostIDs 获取 所有主机
func GetHostIDs(uhostClient *uhost.UHostClient) []uhost.UHostInstanceSet {

	req := uhostClient.NewDescribeUHostInstanceRequest()

	resp, err := uhostClient.DescribeUHostInstance(req)
	if err != nil {
		log.Infoln("[ERROR]", err)
		return nil
	}

	for i := 0; i < len(resp.UHostSet); i++ {
		fmt.Println(resp.UHostSet[i].UHostId)
	}

	return resp.UHostSet

}
