package cmd

import (
	"fmt"
	"log"
	"path"

	"github.com/xiaotuanyu120/MyBlogInGo/render"

	"github.com/xiaotuanyu120/MyBlogInGo/config"

	"github.com/spf13/cobra"
)

var fullAmountMode bool

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate static page from markdown files of blog",
	Long:  `Generate static page from markdown files of blog`,
	Run: func(cmd *cobra.Command, args []string) {
		// STEP 1. LOAD CONFIGURATION
		// 1.1 fetch app configuration
		var conf, err = config.GetConfig()
		if err != nil {
			log.Panicf("Load App Config error: %s", err)
		}

		SrcDir := path.Join(conf.BaseDir, conf.SrcDir)
		DstDir := path.Join(conf.BaseDir, conf.DstDir)

		// 1.2 load notebook mode structure cache file
		// TODO: load notebook mode structure cache file to object

		// STEP 2. Generate Blog Page from Markdown
		markdowns, index, err := render.MDInfoCollect(SrcDir, DstDir, conf.BaseDir)
		if err != nil {
			log.Panicf("Collect markdown file's info failed, error message: [%s]", err)
		}

		// select generate mode: [ increment mode | full amount mode]
		if !fullAmountMode {
			fmt.Println("its full amount mode")
			// todo: filter MDFileInfos
		}

		latest := render.RenderIndexLatest(markdowns)
		// generate page
		for _, md := range markdowns {
			// generate page only when header analysed success
			if md.MDInfo.Analysed {
				// generate page
				err = render.GeneratePage(md, &render.GenParam{
					Latest:      latest,
					Index:       index,
					DstDir:      DstDir,
					ChromaStyle: conf.ChromaStyle,
				})
				if err != nil {
					log.Printf("Generate Page failed, error message: [%s]", err)
				}
			}
		}
		log.Println("Generate Page finish")

		// STEP 3. UPDATE INDEX OF NOTEBOOK MODE
		// errors handled inside render.GenerateCatIndex
		render.GenerateIndexPages(&render.GenParam{
			Latest: latest,
			Index:  index,
			DstDir: DstDir,
		})

		err = render.GenerateHomePage(markdowns, &render.GenParam{
			Latest:      latest,
			Index:       index,
			DstDir:      DstDir,
			ChromaStyle: conf.ChromaStyle,
		})
		if err != nil {
			log.Printf("Generate HomePage failed, error message: [%s]", err)
		}

		// STEP 4. WRITE BACK NOTEBOOK MODE STRUCTURE CACHE FILE
	},
}

func init() {
	genCmd.Flags().BoolVarP(
		&fullAmountMode,
		"full-amount-mode",
		"f",
		false,
		"use full amount mode; default is false which means use increment mode")
	rootCmd.AddCommand(genCmd)
}
