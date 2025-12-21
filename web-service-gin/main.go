package main

import(
	"os"
	"log"
	"fmt"
	"errors"
	"regexp" 
	"net/http"
	"html/template"
	"github.com/gin-gonic/gin"
)

type album struct{
	ID		string 	`json:"id"`
	Title 	string 	`json:"title"`
	Artist	string 	`json:"artist"`
	Price 	float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album
	// Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context){
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

type Page struct {
    Title string
    Body  []byte
}
func (p *Page) save() error {
    filename := "./Pages/" + p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) (*Page, error) {
    filename := "./Pages/" + title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
// func viewHandler(w http.ResponseWriter, r *http.Request) {
//     title := r.URL.Path[len("/view/"):]
//     p, _ := loadPage(title)
//     fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>\n", p.Title, p.Body)
// }
var templates = template.Must(template.ParseFiles("Pages/edit.html", "Pages/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}
func renderTemplate(w http.ResponseWriter, p *Page,tmpl string) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// func renderTemplate(w http.ResponseWriter, p *Page, tmpl string){
// 	t, err := template.ParseFiles("Pages/" + tmpl + ".html")
// 	if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     err = t.Execute(w, p)
// 	if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//     }
// }

// version_1 with duplicate validation of error
// func viewHandler(w http.ResponseWriter, r *http.Request) {
//    	title, err := getTitle(w, r)
// 	if err != nil {
//         return
//     }
//     p, err := loadPage(title)
// 	if err != nil {
// 		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
// 		return 
//     }
// 	renderTemplate(w, p, "view")
// }

// func editHandler(w http.ResponseWriter, r *http.Request) {
//    	title, err := getTitle(w, r)
//     if err != nil {
//         return
//     }
//     p, err := loadPage(title)
//     if err != nil {
// 		p = &Page{Title: title}
//     }
// 	renderTemplate(w, p, "edit")
// }

// func saveHandler(w http.ResponseWriter, r *http.Request) {
//    	title, err := getTitle(w, r)
//     if err != nil {
//         return
//     }
//     body := r.FormValue("body")
//     p := &Page{Title: title, Body: []byte(body)}
//     err = p.save()
// 	if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     http.Redirect(w, r, "/view/"+title, http.StatusFound)
// }

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return 
    }
	renderTemplate(w, p, "view")
}

func editHandler(w http.ResponseWriter, r *http.Request,title string) {
    p, err := loadPage(title)
    if err != nil {
		p = &Page{Title: title}
    }
	renderTemplate(w, p, "edit")
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
func makeHandler(w http.ResponseWriter, r *http.Request, title string){

}
func main() {
	http.HandleFunc("/view/",viewHandler)
	http.HandleFunc("/edit/",editHandler)
	http.HandleFunc("/save/",saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// p1 := &Page{Title: "TestPage", Body:[]byte("This is a test page!")}
	// err := p1.save()
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// p2, err := loadPage("TestPage")
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(p2.Body))
	// http.HandleFunc("/", handler)
    // log.Fatal(http.ListenAndServe(":8080", nil))
    // router := gin.Default()
    // router.GET("/albums", getAlbums)
	// router.POST("/albums", postAlbums)
	// router.GET("/albums/:id",getAlbumById)
    // router.Run("localhost:8080")
}