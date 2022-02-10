package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myblog",
	Short: "myblog is my personal blog generator",
	Long: `myblog is my personal blog generator which is rewrite in Go. 
                  And its sourcecode is available at
                  https://github.com/xiaotuanyu120/MyBlogInGo.git`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is myblog, feel free to use it!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
