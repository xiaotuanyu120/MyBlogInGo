package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/Depado/bfchroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/russross/blackfriday/v2"

	"github.com/xiaotuanyu120/MyBlogInGo/utils"
)

type GenParam struct {
	Latest      []byte
	Index       PageIndex
	DstDir      string
	ChromaStyle string
}

/*
GenerateHomePage

generate blog's home page
*/
func GenerateHomePage(mds Markdowns, param *GenParam) error {
	// gen page summary
	article, err := genHomePage(mds, param.ChromaStyle)
	if err != nil {
		return err
	}
	// render HTML with templates
	sidebar := renderIndex(
		&renderIndexParam{
			index:   param.Index,
			startBy: StartByCat,
			stopBy:  StopByCat,
		})
	page := renderPage(&renderPageParam{
		article: article,
		sidebar: sidebar,
		latest:  param.Latest,
	})

	// write Page
	// create not existed destination directory
	err = utils.CreateNonExistDir(param.DstDir)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(param.DstDir, "/index.html"), page, 0644)
	if err != nil {
		return err
	}
	log.Println("Generate HomePage finish")

	return nil
}

/*
genHomePage

get latest 10 summary of article, and generate these html
*/
func genHomePage(mds Markdowns, chromaStyle string) ([]byte, error) {
	var homePage []byte
	for _, md := range mds.subSlice(0, 10) {
		// convert content to html
		articleSummary, _, err := convertMDToStaticPage(md, chromaStyle)
		if err != nil {
			return nil, err
		}
		// render html with template
		homePage = append(homePage, renderHomePage(articleSummary, md)...)
	}
	return homePage, nil
}

/*
GeneratePage

1. convert page from markdown to html with template
2. write page to file
*/
func GeneratePage(md *Markdown, param *GenParam) error {
	// convert markdown to html and css
	article, CSS, err := generateArticle(md, param.ChromaStyle)
	if err != nil {
		return err
	}

	// render HTML with templates
	sidebar := renderIndex(
		&renderIndexParam{
			index:       param.Index,
			category:    md.MDInfo.Header.Categories.Category,
			subCategory: md.MDInfo.Header.Categories.SubCategory,
			startBy:     StartByCat,
			stopBy:      StopByUri,
		})

	page := renderPage(&renderPageParam{
		article: article,
		sidebar: sidebar,
		latest:  param.Latest,
	})

	// write HTML
	// create not existed destination directory
	dir, _ := path.Split(md.HTMLInfo.Path)
	err = utils.CreateNonExistDir(dir)
	if err != nil {
		return err
	}
	err = os.WriteFile(md.HTMLInfo.Path, page, 0644)
	if err != nil {
		return err
	}

	// write CSS
	CSSPath := filepath.Join(param.DstDir, ChromaCssRelPath)
	dir, _ = path.Split(CSSPath)
	err = utils.CreateNonExistDir(dir)
	if err != nil {
		return errors.New("create directory failed, message: [" + err.Error() + "]")
	}
	err = os.WriteFile(CSSPath, CSS, 0644)
	if err != nil {
		return errors.New("Append to " + CSSPath + " failed, message: [" + err.Error() + "]")
	}

	return nil
}

/*
generateArticle
*/
func generateArticle(md *Markdown, chromaStyle string) (article []byte, CSS []byte, err error) {
	article, CSS, err = convertMDToStaticPage(md, chromaStyle)
	if err != nil {
		return nil, nil, err
	}

	article = append(
		[]byte(fmt.Sprintf(
			ArticleTitleFormatStr,
			md.MDInfo.Header.Title,
			md.MDInfo.Header.Date.Format(DateFormatLayout))),
		article...)

	return article, CSS, nil
}

/*
GenerateCatIndexPage

generate ALL category index page

For myWritingDesk
- "/linux/index.html"(linux is category)
- "/python/index.html"
- "/golang/index.html"
etc
*/
func GenerateIndexPages(param *GenParam) {
	for _, cat := range param.Index {
		// generate category's index page
		err := genCatIndexPage(param, cat.name)
		if err != nil {
			log.Printf("Generate failed [Category]: %s/index.html, error message: %s", cat.name, err)
		}

		// generate sub category's index page
		for _, subCat := range cat.subCats {
			err := genSubCatIndexPage(param, cat.name, subCat.name)
			if err != nil {
				log.Printf("Generate failed [SubCategory]: %s/%s/index.html, error message: %s", cat.name, subCat.name, err)
			}
		}
	}

	log.Println("Generate Index finish")
}

/*
genCatIndexPage

generate single category index page
*/
func genCatIndexPage(param *GenParam, category string) error {
	sidebar := renderIndex(
		&renderIndexParam{
			index:    param.Index,
			category: category,
			startBy:  StartByCat,
			stopBy:   StopBySubCat,
		})

	article := renderIndex(
		&renderIndexParam{
			index:         param.Index,
			category:      category,
			startBy:       StartBySubCat,
			stopBy:        StopByUri,
			renderFullUri: true,
		})

	page := renderPage(&renderPageParam{
		article:  article,
		sidebar:  sidebar,
		latest:   param.Latest,
		category: category,
	})

	// write HTML
	dstPath := filepath.Join(param.DstDir, category, "index.html")

	dir, _ := path.Split(dstPath)
	err := utils.CreateNonExistDir(dir)
	if err != nil {
		return err
	}

	err = os.WriteFile(dstPath, page, 0644)
	if err != nil {
		return err
	}

	return nil
}

/*
genSubCatIndexPage

generate single subCategory index page
*/
func genSubCatIndexPage(param *GenParam, category, subCategory string) error {
	sidebar := renderIndex(
		&renderIndexParam{
			index:       param.Index,
			category:    category,
			subCategory: subCategory,
			startBy:     StartByCat,
			stopBy:      StopBySubCat,
		})

	article := renderIndex(
		&renderIndexParam{
			index:         param.Index,
			category:      category,
			subCategory:   subCategory,
			startBy:       StartByUri,
			stopBy:        StopByUri,
			renderFullUri: false,
		})

	page := renderPage(&renderPageParam{
		article:     article,
		sidebar:     sidebar,
		latest:      param.Latest,
		category:    category,
		subCategory: subCategory,
	})

	// write HTML
	dstPath := filepath.Join(param.DstDir, category, subCategory, "index.html")

	dir, _ := path.Split(dstPath)
	err := utils.CreateNonExistDir(dir)
	if err != nil {
		return err
	}

	err = os.WriteFile(dstPath, page, 0644)
	if err != nil {
		return err
	}

	return nil
}

/*
renderHomePage

render homepage's body part
*/
func renderHomePage(article []byte, md *Markdown) []byte {
	// 1. BODY
	var articleSummary []byte
	// 1.2 ARTICLE
	articleSummary = append(articleSummary, TPLArticleSummaryStart...)
	articleSummary = append(
		articleSummary,
		fmt.Sprintf(
			ArticleSummaryTitleFormatStr,
			md.HTMLInfo.URI,
			md.MDInfo.Header.Title,
			md.MDInfo.Header.Date.Format(DateFormatLayout))...)
	articleSummary = append(articleSummary, article...)
	articleSummary = append(articleSummary, TPLArticleSummaryEnd...)

	return articleSummary
}

type renderPageParam struct {
	article     []byte
	sidebar     []byte
	latest      []byte
	category    string
	subCategory string
}

/*
renderPage

render HTML with templates

NOTE:
  when subCategory is not empty, this function will generate subCategory header before article
  and category will be ignored.

  when subCategory is empty and article is not empty, this function will generate category header before article
*/
func renderPage(param *renderPageParam) []byte {
	// 0. prepare HTML TEMPLATE string
	pageStart := TPLHtmlStart + TPLHead + TPLBlogTitle
	pageEnd := TPLFooter + TPLHtmlEnd

	// 1. prepare HTML BODY in []byte format
	var pageBody []byte
	pageBody = append(pageBody, TPLBodyStart...)

	// 1.1 SIDEBAR
	pageBody = append(pageBody, TPLSidebarStart...)
	pageBody = append(pageBody, TPLSidebarNewPageStart...)
	pageBody = append(pageBody, param.latest...)
	pageBody = append(pageBody, TPLSidebarNewPageEnd...)
	pageBody = append(pageBody, TPLSidebarCategoriesStart...)
	pageBody = append(pageBody, param.sidebar...)
	pageBody = append(pageBody, TPLSidebarCategoriesEnd...)
	pageBody = append(pageBody, TPLSidebarEnd...)

	// 1.2 ARTICLE BODY
	pageBody = append(pageBody, TPLArticleStart...)

	// 1.2.1 category information
	if param.subCategory != "" {
		pageBody = append(pageBody, TPLCatIndexHeaderStart...)
		pageBody = append(pageBody, fmt.Sprintf(SubCatIndexHeaderFormatStr, param.category, param.subCategory)...)
		pageBody = append(pageBody, TPLCatIndexHeaderEnd...)
	} else {
		if param.category != "" {
			pageBody = append(pageBody, TPLCatIndexHeaderStart...)
			pageBody = append(pageBody, fmt.Sprintf(CatIndexHeaderFormatStr, param.category)...)
			pageBody = append(pageBody, TPLCatIndexHeaderEnd...)
		}
	}

	// 1.2.2 article content
	pageBody = append(pageBody, param.article...)
	pageBody = append(pageBody, TPLArticleEnd...)
	pageBody = append(pageBody, TPLBodyEnd...)

	// combine HTML TEMPLATE and HTML BODY in []byte format
	var page []byte
	page = append(page, pageStart...)
	page = append(page, pageBody...)
	page = append(page, pageEnd...)

	return page
}

/*
convertMDToStaticPage

convert markdown to html and css
*/
func convertMDToStaticPage(md *Markdown, chromaStyle string) (HTML []byte, css []byte, err error) {
	r := bfchroma.NewRenderer(
		bfchroma.Style(chromaStyle),
		bfchroma.ChromaOptions(html.WithClasses(true)),
	)

	b := new(bytes.Buffer)
	if err := r.ChromaCSS(b); err != nil {
		return nil, nil, err
	}
	css = b.Bytes()

	HTML = blackfriday.Run(
		[]byte(md.Content),
		blackfriday.WithRenderer(r),
		// https://github.com/russross/blackfriday/issues/693
		blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.NoEmptyLineBeforeBlock),
	)

	return HTML, css, nil
}
