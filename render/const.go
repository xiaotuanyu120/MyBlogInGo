package render

const (
	ChromaCssRelPath = "/static/css/chroma.css"
	DateFormatLayout = "02 Jan 2006"
)

const (
	// TPLHtmlStart : html<start>
	TPLHtmlStart = `<!DOCTYPE html>
<html lang="zh-cmn">
`
	// TPLHead : html > head
	TPLHead = `
<head>
    <title>XTY Blog | Linux Ops Docs | SRE | DEVOPS</title>
    <meta charset="utf-8"/>
    <meta content="width=device-width, initial-scale=1" name="viewport"/>
    <link rel="stylesheet" href="/static/css/chroma.css">
    <link rel="stylesheet" href="/static/css/main.css">
</head>
`
	// TPLBlogTitle : html > blog-title
	TPLBlogTitle = `
<div class="blog-title">
	<div class="container">
		<div class="row">
			<div class="col-lg-12">
				<div>
					<a class="main-title" href="/">XTY的小站</a>
                </div>
                <div>
                    <a class="small-title" href="/">记录技术笔记和技术博客</a>
                </div>
			</div>
		</div>
	</div>
</div>
`
	// TPLBodyStart : html > body<start>
	TPLBodyStart = `
<body>
  <div class="container">
`
	// TPLSidebarStart : html > body > sidebar<start>
	TPLSidebarStart = `
    <div class="col-lg-4 col-lg-offset-1 col-md-4 col-md-offset-1 col-sm-4 col-sm-offset-1">
	  <div id="sidebar">
`
	// TPLSidebarNewPageStart : html > body > sidebar > new-page<start>
	TPLSidebarNewPageStart = `
		<h3>最新文章</h3>
          <ul>
`

	// TPLSidebarNewPageEnd : html > body > sidebar > new-page<end>
	TPLSidebarNewPageEnd = `
          </ul>
`

	// TPLSidebarCategoriesStart : html > body > sidebar > categories<start>
	TPLSidebarCategoriesStart = `
		<h3>文章分类</h3>
		  <ul>
`

	// TPLSidebarCategoriesEnd : html > body > sidebar > categories<end>
	TPLSidebarCategoriesEnd = `
          </ul>
`

	// TPLSidebarEnd : html > body > sidebar<end>
	TPLSidebarEnd = `
      </div>
    </div>
`

	// TPLArticleStart : html > body > article<start>
	TPLArticleStart = `
    <div class="col-lg-7 col-md-7 col-sm-7">
`

	// TPLArticleEnd : html > body > article<end>
	TPLArticleEnd = `
    </div>
`

	// TPLArticleSummaryStart : html > body > article > summary<start>
	TPLArticleSummaryStart = `
        <div class="article-summary">
`

	// TPLArticleSummaryEnd : html > body > article > summary<end>
	TPLArticleSummaryEnd = `
        </div>
`

	// TPLCatIndexHeaderStart : html > body > category index header<start>
	TPLCatIndexHeaderStart = `
    <div class="cat-header">
      <h2>
        <span>文章类别：</span>
`

	// TPLCatIndexHeaderEnd : html > body > category index header<end>
	TPLCatIndexHeaderEnd = `
      </h2>
    </div>
`

	// TPLBodyEnd : html > body<end>
	TPLBodyEnd = `
  </div>
</body>
`

	// TPLFooter : html > footer
	TPLFooter = `
<footer>
    <div class="container">
        <div class="row footer-links">
            <div class="col-lg-2 col-sm-2">
                <h3>友情链接</h3>
                <ul>
                    <li><a href="">友链位招租</a></li>
                    <li><a href="">友链位招租</a></li>
                </ul>
            </div>
            <div class="col-lg-2 col-sm-2">
                <h3>没想好</h3>
                <ul>
                    <li><a href="">我爸没想好</a></li>
                    <li><a href="">我哥说我爸没想好</a></li>
                </ul>
            </div>
            <div class="col-lg-2 col-sm-2">
                <h3>Hooray</h3>
                <ul>
                    <li><a href="">Hooray</a></li>
                    <li><a href="">What are we Hooray For?</a></li>
                </ul>
            </div>
            <div class="col-lg-2 col-sm-2">
                <h3>前面的footer太浪了</h3>
                <ul>
                    <li><a href="">就是就是</a></li>
                    <li><a href="">偷偷的表示羡慕</a></li>
                </ul>
            </div>
            <div class="col-lg-4 col-sm-4">
                <h3>网站信息</h3>
                <a class="" href="" target="_blank"></a>
                <a class="" href="" target="_blank"></a>
                <a class="" href="" target="_blank"></a>
                <a class="" href="" target="_blank"></a>
                <div class="fine-print">
                    <p>网战由以下技术支撑</p>
                    <ul>
                        <li>Markdown Processor: <a href="https://github.com/russross/blackfriday/tree/v2">Blackfriday V2</a></li>
                        <li>Renderer Engine: <a href="https://github.com/Depado/bfchroma/">bfchroma</a></li>
                        <li>Syntax Highlighter: <a href="https://github.com/alecthomas/chroma">Chroma</a></li>
                        <li>Coding Language: <a href="https://go.dev/">Golang</a></li>
                        <li>Others: Markdown, HTML, CSS</li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</footer>
`

	// TPLHtmlEnd : html<end>
	TPLHtmlEnd = `
</html>`

	TPLHR = `<hr style="border: 0; border-top: 1px dashed #a2a9b6">`
)

const (
	CatLiStartFormatStr = `
            <li>
              <a href="/%s/index.html">%s</a>`

	SubCatULStartFormatStr = `
              <ul>`

	SubCatLiStartFormatStr = `
                <li>
                  <a href="/%s/%s/index.html">%s</a>`

	PageULStartFormatStr = `
                  <ul>`

	PageLiFormatStr = `
                    <li><a href="/%s">%s</a></li>`

	PageULEndFormatStr = `
                  </ul>`

	SubCatLiEndFormatStr = `
                </li>`

	SubCatULEndFormatStr = `
              </ul>`

	CatLiEndFormatStr = `
            </li>`

	CatIndexHeaderFormatStr = `        [ %s ]`

	SubCatIndexHeaderFormatStr = `        [ %s ] > [ %s ]`

	LatestLiFormatStr = `
            <li>
              <a href="/%s">%s</a>
            </li>`

	// ArticleTitleFormatStr is between TPLArticleStart and real article content
	ArticleTitleFormatStr = `      <h2>%s</h2>
      <div>
        <hr style="border: 0; border-top: 1px dashed #a2a9b6">
      </div>
      <div class="postDate">
        <p>%s</p>
      </div>
      <div>
        <hr style="border: 0; border-bottom: 1px dashed #a2a9b6">
      </div>
`

	ArticleSummaryTitleFormatStr = `        <div class="article-title">
          <a href="%s">%s</a>
        </div>
        <div>
          <hr style="border: 0; border-top: 1px dashed #a2a9b6">
        </div>
        <div class="postDate">
          <p>%s</p>
        </div>
        <div>
          <hr style="border: 0; border-bottom: 1px dashed #a2a9b6">
        </div>
`
)
