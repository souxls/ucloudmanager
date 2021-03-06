/**
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"ucloudmanager/ucloud"

	"github.com/spf13/cobra"
)

var imgageName, imageID string

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "manage image",
}

var createImageCmd = &cobra.Command{
	Use:   "create",
	Short: "create image",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ucloud.CreateImage(&imgageName, &HostID); err != nil {
			fmt.Printf("镜像创建失败")

		}
	},
}

var deleteImageCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete image",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ucloud.DeleteImage(&imageID); err != nil {
			fmt.Println("镜像删除失败", err)

		}
	},
}

var listImageCmd = &cobra.Command{
	Use:   "list",
	Short: "list image",
	Run: func(cmd *cobra.Command, args []string) {
		ucloud.GetImages()
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)

	imageCmd.AddCommand(createImageCmd)
	createImageCmd.PersistentFlags().StringVar(&imgageName, "name", "", "Set image name.")
	createImageCmd.PersistentFlags().StringVar(&HostID, "hostID", "", "Create image by hostid.")
	createImageCmd.MarkPersistentFlagRequired("name")
	createImageCmd.MarkPersistentFlagRequired("hostID")

	imageCmd.AddCommand(deleteImageCmd)
	deleteImageCmd.PersistentFlags().StringVar(&imageID, "imageID", "", "delete image by imageid.")
	deleteImageCmd.MarkPersistentFlagRequired("imageID")

	imageCmd.AddCommand(listImageCmd)
}
