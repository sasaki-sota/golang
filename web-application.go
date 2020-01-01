package main

// io/ioutilはファイルの読み込みなどについて
// ここでインポートすることによって.をつけた後に自動的に候補を表示してくれる
import (
	"fmt"
	"io/ioutil"
)

type Page struct{
	Title string
	Body  []byte
}
// ここの部分で肩を定義してメソッドを作成しているイメージ

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "test", Body: []byte("This is a sampke Page.")}
	p1.save()

	p2, _ := loadPage(p1.Title)
	fmt.Println(p2.Body)
}