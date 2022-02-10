package render

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/xiaotuanyu120/MyBlogInGo/utils"
)

/*
MDInfoCollect

collect all markdown file's info
*/
func MDInfoCollect(SrcDir, DstDir, baseDir string) (Markdowns, PageIndex, error) {
	var markdowns Markdowns
	var index PageIndex

	// READ all markdown files
	err := filepath.WalkDir(SrcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}

			if filepath.Ext(info.Name()) == ".md" {
				markdown := InitMarkdown(path, info)
				markdowns = append(markdowns, markdown)
			}
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	// COLLECT information of markdown file
	var needRemovedIDX []int
	for i, md := range markdowns {

		err := collectHeaderContent(md, DstDir, baseDir)
		if err != nil {
			needRemovedIDX = append(needRemovedIDX, i)
			log.Printf("COLLECT HEADER ERROR: %s", err)
			continue
		}

		// collect index info
		index = index.Add(
			md.MDInfo.Header.Categories.Category,
			md.MDInfo.Header.Categories.SubCategory,
			md.HTMLInfo.URI,
			md.MDInfo.Header.Title,
			md.MDInfo.Header.Date,
		)
	}
	markdowns = removeMDItem(markdowns, needRemovedIDX)

	return markdowns, index, nil
}

/*
collectHeaderContent

read header and content of markdown file
*/
func collectHeaderContent(md *Markdown, DstDir, baseDir string) error {
	mdLines, err := utils.ReadFile(md.MDInfo.FileInfo.Path)
	if err != nil {
		return err
	}

	var title string
	var date time.Time
	var categories []string
	var tags []string

	mdRelPath, _ := filepath.Rel(baseDir, md.MDInfo.FileInfo.Path)

	// analyse header of markdown file
	nonEmptyNo := -1
	contentStartNo := 0
	for number, mdLine := range mdLines {
		// TYPE: PREPARE
		// record actual content start No. whether analysis header success or failed
		contentStartNo = number + 1
		// skip empty line and record nonEmptyNo
		mdLine = strings.TrimSpace(mdLine)
		if mdLine == "" {
			continue
		} else {
			nonEmptyNo += 1
		}

		if nonEmptyNo == 0 && mdLine == "---" {
			continue
		} else if nonEmptyNo > 0 && nonEmptyNo < 5 && strings.HasPrefix(mdLine, "title:") {
			title = strings.SplitN(mdLine, "title:", 2)[1]
			title = strings.TrimSpace(title)

		} else if nonEmptyNo > 0 && nonEmptyNo < 5 && strings.HasPrefix(mdLine, "date:") {
			dateStr := strings.SplitN(mdLine, "date:", 2)[1]
			dateStr = strings.TrimSpace(dateStr)
			dateStrSlice := strings.Split(dateStr, " ")
			if len(dateStrSlice) != 2 {
				return errors.New(
					fmt.Sprintf(
						"[date format error]: %s [%s]",
						dateStr,
						md.MDInfo.FileInfo.Path,
					),
				)
			}
			dateStr = dateStrSlice[0] + "T" + dateStrSlice[1] + "+08:00"
			date, err = time.Parse(time.RFC3339, dateStr)
			if err != nil {
				return errors.New(
					fmt.Sprintf(
						"[date format error]: %s [%s]",
						err.Error(),
						mdRelPath,
					),
				)
			}

		} else if nonEmptyNo > 0 && nonEmptyNo < 5 && strings.HasPrefix(mdLine, "categories:") {
			categoriesStr := strings.SplitN(mdLine, "categories:", 2)[1]
			categoriesStr = strings.TrimSpace(categoriesStr)
			categories = strings.Split(categoriesStr, "/")
			if len(categories) != 2 {
				return errors.New(
					fmt.Sprintf(
						"[categories format error]: %s [%s]",
						categoriesStr,
						mdRelPath,
					),
				)
			}
			// TYPE: SUCCESS

		} else if nonEmptyNo > 0 && nonEmptyNo < 5 && strings.HasPrefix(mdLine, "tags:") {
			tagsStr := strings.SplitN(mdLine, "tags:", 2)[1]
			tagsStr = strings.TrimSpace(tagsStr)
			if strings.HasPrefix(tagsStr, "[") {
				tagsStr = strings.SplitN(tagsStr, "[", 2)[1]
			} else {
				return errors.New(
					fmt.Sprintf(
						"[tags format error]: %s('tags: [tags1, tags2...]') [%s]",
						mdLine,
						mdRelPath,
					),
				)
			}

			tagsStr = strings.TrimSpace(tagsStr)
			if strings.HasSuffix(tagsStr, "]") {
				tagsStr = strings.SplitN(tagsStr, "]", 2)[0]
			} else {
				return errors.New(
					fmt.Sprintf(
						"[tags format error]: %s('tags: [tags1, tags2...]') [%s]",
						mdLine,
						mdRelPath,
					),
				)
			}

			tags = strings.Split(tagsStr, ",")
			for n, t := range tags {
				tags[n] = strings.TrimSpace(t)
			}

		} else if nonEmptyNo == 5 && mdLine == "---" {
			break
		} else {
			return errors.New(fmt.Sprintf("[header not existed] [%s]", mdRelPath))
		}
	}

	categoriesObj := &Categories{
		Category:    categories[0],
		SubCategory: categories[1],
	}
	md.MDInfo.Header = &Header{
		Title:      title,
		Date:       date,
		Categories: categoriesObj,
		Tags:       tags,
	}
	md.MDInfo.Analysed = true

	// collect html path after the success of analyse header of markdown
	collectHTMLPath(md, DstDir)
	// TODO: FULL AMOUNT MODE & INCREMENT MODE LOGIC HERE
	md.Content = strings.Join(mdLines[contentStartNo:], "\n")

	return nil
}

/*
collectHTMLPath

collect html's storage path

	Example:
	Input
		mdf > *MarkdownFile{Info: {Name: page1, ...},Header: {categories: [baseCat, subCat], ...}, ...}
		DstDir > /path/to/resource/html
	Output
		/path/to/resource/html/baseCat/subCat/page1.html
*/
func collectHTMLPath(md *Markdown, DstDir string) {
	HTMLFileName := strings.TrimSuffix(md.MDInfo.FileInfo.Info.Name(), ".md") + ".html"

	md.HTMLInfo.URI = filepath.Join(
		md.MDInfo.Header.Categories.Category,
		md.MDInfo.Header.Categories.SubCategory,
		HTMLFileName,
	)

	md.HTMLInfo.Path = filepath.Join(
		DstDir,
		md.HTMLInfo.URI,
	)
}

/*
removeMDItem

remove error collected information item from Markdowns of markdown
*/
func removeMDItem(mds Markdowns, idxs []int) Markdowns {
	if len(idxs) == 0 {
		return mds
	}

	mdsLen := len(mds)
	sort.Ints(idxs)
	ret := make(Markdowns, 0)
	left := 0
	right := mdsLen
	for i, si := range idxs {
		switch {
		case si < 0:
			continue
		case si >= 0 && si < mdsLen:
			right = si
		case si >= mdsLen:
			right = mdsLen
		}

		if left <= right {
			ret = append(ret, mds[left:right]...)
		}

		if i == len(idxs)-1 && left < mdsLen {
			ret = append(ret, mds[si+1:]...)
		}
		left = si + 1
	}
	return ret
}
