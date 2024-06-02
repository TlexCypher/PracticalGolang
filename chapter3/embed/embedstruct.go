package embed

import "fmt"

type Book struct {
	Title string
	ISBN  string
}

func (b Book) GetAmazonURL() string {
	return "https://amazon.co.jp/dp/" + b.ISBN
}

type OreillyBook struct {
	Book
	ISBN13 string
}

func (o OreillyBook) GetOreillyURL() string {
	return "https://www/oreilly.co.jp/books/" + o.ISBN13 + "/"
}

type Database struct {
	Address string
}

type FileStorage struct {
	Address string
}

type WebServiceConifg struct {
	Database
	FileStorage
}

func main() {
	ob := OreillyBook{
		ISBN13: "9784873119038",
		Book: Book{
			Title: "Real World HTTP",
		},
	}
	//本当はob.Book.GetAmazonURL()だけど複数の同じFieldNameを有していない限り大丈夫
	fmt.Println(ob.GetAmazonURL())
	fmt.Println(ob.GetOreillyURL())

	ws := WebServiceConifg{
		Database: Database{
			Address: "Database field",
		},
		FileStorage: FileStorage{
			Address: "FileStorage field",
		},
	}
	//同じフィールド名が存在する時は明示的なアクセスを示す必要がある
	fmt.Println(ws.Database.Address)
}
