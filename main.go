package main

import (
	"math/rand"
	"time"
  "net/http"
  "html/template"
)

func main()  {

  http.HandleFunc("/", chartHandler)
	http.ListenAndServe(":8081", nil)
}

func chartHandler(w http.ResponseWriter, _ *http.Request) {
  seed := rand.NewSource(time.Now().UnixNano())
  depth := 100
  width := 1000
  rain := 10000
  data := Foo(seed, depth, width, rain)

  err := renderTemplate(w, data)

  if err != nil {
    handleError(w, err)
  }
}


func handleError(w http.ResponseWriter, err error) {
  http.Error(w, err.Error(), http.StatusInternalServerError)
}

func renderTemplate(w http.ResponseWriter, data []int) error {
  var t = template.Must(template.ParseFiles("page.html"))
  err := t.Execute(w, data)

  return err
}

func Foo(seed rand.Source, depth int, width int, rain int) []int{
  rng := rand.New(seed)

  line := []int{}
  for i := 0; i < width; i++ {
   line = append(line, depth) 
  }
  for i := 0; i < rain; i++ {
    spot := rng.Intn(width)
    println(spot)
    line[spot] = line[spot] - 1
  }

  return line
}
