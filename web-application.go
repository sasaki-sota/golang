package main

import (
	"html/template"
	"log"
	"net/http"
	"io/ioutil"
)



type Page struct{
	Title string
	Body []byte
}

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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

// wに対しての返答が返ってくるような仕組みになっている
func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil{
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// :の前に何も書かなければローカルホストになる
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
// ListenAndServeは第一引数はアドレスで第二引数はハンドラーとなっている
// Fprintはioに関係するような出力の仕方


// goでサーバを立ち上げるときに立ち上がらないときはgolangのディレクトリからpsとうち
// kill -KILL 4757のような番号を入力して削除した再度でバックする必要がある
// またはvsコードのていしボタンを押す