package render

import (
	"fmt"
)

const (
	StartByCat = iota
	StartBySubCat
	StartByUri
	StopByCat
	StopBySubCat
	StopByUri
)

type renderIndexParam struct {
	index         PageIndex
	category      string
	subCategory   string
	startBy       int
	stopBy        int
	renderFullUri bool
}

/*
renderIndex

render sidebar categories
*/
func renderIndex(param *renderIndexParam) []byte {
	var result []byte
	for _, cat := range param.index {
		if param.startBy <= StartByCat {
			result = append(result, fmt.Sprintf(CatLiStartFormatStr, cat.name, cat.name)...)
		}
		if param.stopBy > StopByCat && cat.name == param.category {
			result = renderIndexSubCat(param, cat.subCats, result)
		}
		if param.startBy <= StartByCat {
			result = append(result, CatLiEndFormatStr...)
		}
	}
	return result
}

/*
renderIndexSubCat

render sidebar subcategories
*/
func renderIndexSubCat(param *renderIndexParam, subCats subCats, result []byte) []byte {
	result = append(result, SubCatULStartFormatStr...)
	for _, subCat := range subCats {
		if param.startBy <= StartBySubCat {
			result = append(result, fmt.Sprintf(SubCatLiStartFormatStr, param.category, subCat.name, subCat.name)...)
		}
		if (param.stopBy > StopBySubCat && subCat.name == param.subCategory) || param.renderFullUri {
			result = renderIndexUri(subCat.pages, result)
		}
		if param.startBy <= StartBySubCat {
			result = append(result, SubCatLiEndFormatStr...)
		}
	}
	result = append(result, SubCatULEndFormatStr...)

	return result
}

/*
renderIndexUri

render sidebar uri
*/
func renderIndexUri(pages pages, result []byte) []byte {
	result = append(result, PageULStartFormatStr...)
	for _, page := range pages {
		result = append(result, fmt.Sprintf(PageLiFormatStr, page.uri, page.title)...)
	}
	result = append(result, PageULEndFormatStr...)

	return result
}

/*
RenderIndexLatest

render sidebar latest article
*/
func RenderIndexLatest(mds Markdowns) []byte {
	var result []byte

	for _, md := range mds.subSlice(0, 5) {
		// render html with template
		result = append(result, fmt.Sprintf(LatestLiFormatStr, md.HTMLInfo.URI, md.MDInfo.Header.Title)...)
	}

	return result
}
