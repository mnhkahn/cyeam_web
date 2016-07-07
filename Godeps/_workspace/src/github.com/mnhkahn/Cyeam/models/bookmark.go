package models

// "github.com/go-xorm/xorm"
// _ "github.com/mattn/go-sqlite3"

type Bookmark struct {
	Id    int `xorm:"pk"`
	Cat   string
	Title string
	Link  string
}

type BookmarkView struct {
	Cat   string
	Items []*BookmarkItemView
}

type BookmarkItemView struct {
	Title string
	Link  string
}

// func AddBookmark(b *Bookmark) (err error) {
// 	if b != nil {
// 		_, err = engine.Insert(b)
// 	}
// 	return
// }

// func GetAllBookmarks() []*BookmarkView {
// 	view := make([]*BookmarkView, 0)

// 	bookmarks := make([]Bookmark, 0)
// 	err := engine.Cols("cat").GroupBy("cat").Find(&bookmarks)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, bookmark := range bookmarks {
// 		bookmarkView := new(BookmarkView)
// 		bookmarkView.Cat = bookmark.Cat

// 		bks := make([]Bookmark, 0)
// 		err = engine.Where("cat=?", bookmark.Cat).Find(&bks)
// 		if err != nil {
// 			panic(err)
// 		}
// 		for _, bk := range bks {
// 			item := new(BookmarkItemView)
// 			item.Title = bk.Title
// 			item.Link = bk.Link
// 			bookmarkView.Items = append(bookmarkView.Items, item)
// 		}
// 		view = append(view, bookmarkView)
// 	}

// 	return view
// }

// var engine *xorm.Engine

// func init() {
// 	var err error
// 	engine, err = xorm.NewEngine("sqlite3", "./cyeam.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	engine.ShowSQL = true
// 	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

// 	// err = AddBookmark(&Bookmark{Id: 1, Cat: "重要", Title: "官方网站", Link: "http://cyeam.com"})
// 	GetAllBookmarks()
// }
