package main

import (
	"fmt"

	"qlip/pkg/download"
	"qlip/pkg/uploadhandler"

	"github.com/spf13/cobra"
)

var (
	file string
	upload string
	downloadString string
	qlipbord bool
)

var rootCmd = &cobra.Command{
	Use: "qlipcli",
	Short: "A service that can upload files or download files",
	Run: func(cmd *cobra.Command, args []string) {

		if downloadString != "" {
			download.DownloadHandler(downloadString)
		} 

		if upload != "" {
			uploadhandler.UploadToService(upload)
		}

		if qlipbord {
			fmt.Println("To doo")
		}
	},
}

func main() {
	rootCmd.Execute()
}

func init(){
	rootCmd.Flags().StringVarP(&downloadString, "download", "d", "", "Use this to download file")
	rootCmd.Flags().StringVarP(&upload, "upload", "u", "", "Use this to upload file")
	rootCmd.Flags().BoolVarP(&qlipbord, "qlipbord", "q", false, "Use this to get it in qlipbord")
}

