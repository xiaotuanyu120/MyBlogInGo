package render

import (
	"io/fs"
	"sort"
	"time"
)

/*
DATA FOR PAGES
*/
type Markdown struct {
	Content  string
	MDInfo   *MDInfo
	HTMLInfo *HTMLInfo
}

type HTMLInfo struct {
	Path string
	URI  string
}

type MDInfo struct {
	FileInfo *FileInfo
	Header   *Header

	// Analysed: default is false, set it to true if header analysis is passed
	Analysed bool
}

type FileInfo struct {
	Path string
	Info fs.FileInfo
}

type Header struct {
	Title      string
	Date       time.Time
	Categories *Categories
	Tags       []string
}

type Categories struct {
	Category    string
	SubCategory string
}

func InitMarkdown(path string, info fs.FileInfo) *Markdown {
	fileInfo := &FileInfo{
		Path: path,
		Info: info,
	}

	mdInfo := &MDInfo{
		FileInfo: fileInfo,
		Header:   &Header{},
	}

	return &Markdown{
		MDInfo:   mdInfo,
		HTMLInfo: &HTMLInfo{},
	}
}

type Markdowns []*Markdown

func (mds Markdowns) subSlice(start, step int) Markdowns {
	sort.SliceStable(mds, func(i, j int) bool { return mds[j].MDInfo.Header.Date.Before(mds[i].MDInfo.Header.Date) })

	if len(mds) < step {
		step = len(mds)
	}

	return mds[start : start+step]
}

/*
DATA FOR PAGE INDEX
*/
type page struct {
	uri   string
	title string
	date  time.Time
}

type pages []*page

/*
search

determine a page(uri string) exist in a pages or not
*/
func (idx pages) search(uri string) (bool, int) {
	i := sort.Search(len(idx), func(i int) bool { return idx[i].uri == uri })

	if i < len(idx) && idx[i].uri == uri {
		return true, i
	} else {
		return false, -1
	}
}

type subCat struct {
	name  string
	pages pages
}

type subCats []*subCat

/*
search

determine a subCat(name string) exist in a subCats or not
*/
func (idx subCats) search(subCat string) (bool, int) {
	i := sort.Search(len(idx), func(i int) bool { return idx[i].name == subCat })

	if i < len(idx) && idx[i].name == subCat {
		return true, i
	} else {
		return false, -1
	}
}

type cat struct {
	name    string
	subCats subCats
}

type PageIndex []*cat

/*
search

determine a cat(name string) exist in a PageIndex or not
*/
func (idx PageIndex) search(cat string) (bool, int) {
	i := sort.Search(len(idx), func(i int) bool { return idx[i].name == cat })

	if i < len(idx) && idx[i].name == cat {
		return true, i
	} else {
		return false, -1
	}
}

/*
Add

add new cat into PageIndex
*/
func (idx PageIndex) Add(category, subCategory, uri, title string, date time.Time) PageIndex {
	ret := idx
	// create cat if it's not exist
	okCat, iCat := ret.search(category)
	if !okCat {
		ret = append(ret, &cat{
			name:    category,
			subCats: newSubCats(subCategory, newPages(uri, title, date)),
		})
	} else {
		// create subCat if it's not exist
		okSubCat, iSubCat := ret[iCat].subCats.search(subCategory)
		if !okSubCat {
			ret[iCat].subCats = append(ret[iCat].subCats, &subCat{
				name:  subCategory,
				pages: newPages(uri, title, date),
			})
		} else {
			// create page if it's not exist
			okPage, iPage := ret[iCat].subCats[iSubCat].pages.search(uri)
			if !okPage {
				ret[iCat].subCats[iSubCat].pages = append(ret[iCat].subCats[iSubCat].pages, &page{
					uri:   uri,
					title: title,
					date:  date,
				})
			} else {
				// update page if it's exist
				ret[iCat].subCats[iSubCat].pages[iPage] = &page{
					uri:   uri,
					title: title,
					date:  date,
				}
			}
		}
	}

	return ret
}

func newPages(uri, title string, date time.Time) pages {
	return pages{
		&page{
			uri:   uri,
			title: title,
			date:  date,
		},
	}
}

func newSubCats(subCategory string, pages pages) subCats {
	return subCats{
		&subCat{
			name:  subCategory,
			pages: pages,
		},
	}
}
