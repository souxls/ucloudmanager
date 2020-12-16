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

var listHost, all bool
var uhostName, imageName, zoneName string

// uhostCmd represents the uhost command
var uhostCmd = &cobra.Command{
	Use:   "uhost",
	Short: "manage uhost",
}

var createUhostCmd = &cobra.Command{
	Use:   "create",
	Short: "create uhost",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ucloud.CreateHost(uhostName, imageName, zoneName); err != nil {
			fmt.Println("Host创建失败", err)
		}
	},
}
var deleteUhostCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete uhost",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ucloud.DeleteHost(&HostID); err != nil {
			fmt.Println("Host删除失败", err)
		}
	},
}

var listUhostCmd = &cobra.Command{
	Use:   "list",
	Short: "list uhost",
	Run: func(cmd *cobra.Command, args []string) {
		ucloud.GetHostIDs()
	},
}

var startUhostCmd = &cobra.Command{
	Use:   "start",
	Short: "start uhost",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ucloud.StartHost(&HostID); err != nil {
			fmt.Println("Host启动失败", err)
		}
	},
}

var stopUhostCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop uhost",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ucloud.StopHost(&HostID); err != nil {
			fmt.Println("Host停止失败", err)

		}
	},
}

func init() {
	rootCmd.AddCommand(uhostCmd)
	uhostCmd.AddCommand(createUhostCmd)
	createUhostCmd.PersistentFlags().StringVar(&uhostName, "name", "", "uhost name.")
	createUhostCmd.PersistentFlags().StringVar(&imageName, "image", "", "image id.")
	createUhostCmd.PersistentFlags().StringVar(&zoneName, "zone", "cn-sh2-01", "zone name.")
	createUhostCmd.MarkPersistentFlagRequired("zone")
	createUhostCmd.MarkPersistentFlagRequired("image")

	uhostCmd.AddCommand(deleteUhostCmd)
	deleteUhostCmd.PersistentFlags().StringVar(&HostID, "hostID", "", "Delete uhost by uhostID.")
	deleteUhostCmd.MarkPersistentFlagRequired("hostID")

	uhostCmd.AddCommand(startUhostCmd)
	startUhostCmd.PersistentFlags().StringVar(&HostID, "hostID", "", "Start uhost by uhostID.")
	startUhostCmd.MarkPersistentFlagRequired("hostID")

	uhostCmd.AddCommand(stopUhostCmd)
	stopUhostCmd.PersistentFlags().StringVar(&HostID, "hostID", "", "Stop uhost by uhostID.")
	stopUhostCmd.MarkPersistentFlagRequired("hostID")

	uhostCmd.AddCommand(listUhostCmd)
	listUhostCmd.PersistentFlags().BoolVar(&listHost, "list", true, "List uhost of region.")
}
