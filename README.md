# ucloud
ucloud 云平台 uhost 和 镜像管理

# 获取镜像ID
```
ucloudmanager image list
ImageID: uimage-50c2fwwi ImageName: test1
ImageID: uimage-ddjty5cm ImageName: test2
```
# 根据 ImageID 创建uhost
```
ucloudmanager uhost create --zone cn-sh2-01 --image uimage-ddjty5cm --name "test1-20201216"

```
# Usage
```
ucloudmanager --help
Manage ucloud host and image

Usage:
  ucloudmanager [command]

Available Commands:
  help        Help about any command
  image       manage image
  uhost       manage uhost

Flags:
      --config string   config file (default "./ucloud.toml")
  -h, --help            help for ucloudmanager
      --region string   ucloud region (default "cn-sh2")

Use "ucloudmanager [command] --help" for more information about a command.
```