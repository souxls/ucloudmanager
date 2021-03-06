/**
Homepage: https://github.com/ucloud/ucloud-sdk-go
Examples: https://github.com/ucloud/ucloud-sdk-go/tree/master/examples
*/

package ucloud

import (
	"fmt"
	"strings"
	"ucloudmanager/config"
	"ucloudmanager/log"

	"github.com/sethvargo/go-password/password"
	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
)

// CreateHost 创建主机
func CreateHost(name, imageID, zone string) error {

	h := config.Cfg.Host
	pwd, err := password.Generate(16, 5, 0, false, false)
	req := Uclient.NewCreateUHostInstanceRequest()
	req.Zone = ucloud.String(zone)
	req.ImageId = ucloud.String(imageID)
	req.Password = ucloud.String(pwd)
	req.LoginMode = ucloud.String("Password")
	req.Name = ucloud.String(name)
	req.ChargeType = ucloud.String(h.ChargeType)
	req.CPU = ucloud.Int(h.CPU)
	req.Memory = ucloud.Int(h.Memory)
	req.NetCapability = ucloud.String(h.NetCapability)
	req.MachineType = ucloud.String(h.MachineType)
	req.MinimalCpuPlatform = ucloud.String(h.MinimalCpuPlatform)
	req.NetworkInterface = []uhost.CreateUHostInstanceParamNetworkInterface{
		{
			EIP: &uhost.CreateUHostInstanceParamNetworkInterfaceEIP{
				Bandwidth:    ucloud.Int(h.NetworkInterface[0].Bandwidth),
				PayMode:      ucloud.String(h.NetworkInterface[0].PayMode),
				OperatorName: ucloud.String("Bgp"),
				GlobalSSH: &uhost.CreateUHostInstanceParamNetworkInterfaceEIPGlobalSSH{
					Port: ucloud.Int(22),
				},
			},
		},
	}
	req.Disks = []uhost.UHostDisk{
		{
			IsBoot: ucloud.String(h.Disks[0].IsBoot),
			Size:   ucloud.Int(h.Disks[0].Size),
			Type:   ucloud.String(h.Disks[0].Type),
		},
	}

	resp, err := Uclient.CreateUHostInstance(req)
	if err != nil {
		if strings.Contains(err.Error(), "resource not enough") {
			log.Errorln("资源不足", err)
			if h.MinimalCpuPlatform == "Amd/Epyc2" {
				req.MinimalCpuPlatform = ucloud.String("Intel/Auto")
			} else {
				req.MinimalCpuPlatform = ucloud.String("Amd/Epyc2")
			}
			resp, _ := Uclient.CreateUHostInstance(req)
			log.Infoln("重新创建", resp)
			return nil
		}
		return err
	}
	log.Debugln("[RESPONSE]", resp)
	return nil
}

// StartHost 启动 uhost
func StartHost(uHostID *string) error {

	req := Uclient.NewStartUHostInstanceRequest()
	req.UHostId = uHostID

	resp, err := Uclient.StartUHostInstance(req)
	if err != nil {
		log.Errorln("[ERROR]", err)
		return err
	}

	log.Debugln("[RESPONSE]", resp)
	return nil
}

// StopHost 停止 uHost
func StopHost(uhostID *string) error {

	req := Uclient.NewStopUHostInstanceRequest()
	req.UHostId = uhostID

	resp, err := Uclient.StopUHostInstance(req)
	if err != nil {
		log.Errorln("[ERROR]", err)
		return err
	}

	log.Debugln("[RESPONSE]", resp)
	return nil
}

// DeleteHost 删除 uhost
func DeleteHost(uhostID *string) error {

	req := Uclient.NewTerminateUHostInstanceRequest()
	req.UHostId = uhostID
	req.ReleaseEIP = ucloud.Bool(true)

	resp, err := Uclient.TerminateUHostInstance(req)
	if err != nil {
		log.Errorln("[ERROR]", err)
		return err
	}

	log.Debugln("[RESPONSE]", resp)
	return nil
}

// GetHostIDs 获取 所有主机
func GetHostIDs() []uhost.UHostInstanceSet {

	req := Uclient.NewDescribeUHostInstanceRequest()

	resp, err := Uclient.DescribeUHostInstance(req)
	if err != nil {
		log.Errorln("[ERROR]", err)
		return nil
	}

	for i := 0; i < len(resp.UHostSet); i++ {
		hostInfo := resp.UHostSet[i]
		fmt.Printf("HostID: %s HostName: %s WanIP: %s State: %s \n", hostInfo.UHostId, hostInfo.Name, hostInfo.IPSet[1].IP, hostInfo.State)
	}

	return resp.UHostSet

}
