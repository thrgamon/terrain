package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type Earth struct {
	depth int
	left  *Earth
	right *Earth
}

type Terrain struct {
	width     int
	depth     int
	heightmap []*Earth
	seed      rand.Source
}

func newTerrain(depth int, width int, seed rand.Source) *Terrain {
	terrain := Terrain{depth: depth, width: width, seed: seed}

	for i := 0; i < width; i++ {
		e := &Earth{depth: depth}
		terrain.heightmap = append(terrain.heightmap, e)
	}

	for i := 0; i < width; i++ {
		e := terrain.heightmap[i]
		if i-1 > 0 {
			e.left = terrain.heightmap[i-1]
		}
		if i+1 < width {
			e.right = terrain.heightmap[i+1]
		}
	}

	return &terrain
}

func (terr *Terrain) toSlice() (slizzard []int) {
	for _, earth := range terr.heightmap {
		slizzard = append(slizzard, earth.depth)
	}
	return slizzard
}

func main() {
	http.HandleFunc("/", chartHandler)
	http.ListenAndServe(":8081", nil)
}

func chartHandler(w http.ResponseWriter, _ *http.Request) {
	seed := rand.NewSource(time.Now().UnixNano())
	terr := newTerrain(100, 1000, seed)
	rain := 10000
	data := Foo(terr, rain)

	err := renderTemplate(w, data.toSlice())

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

func Foo(terr *Terrain, rain int) *Terrain {
	rng := rand.New(terr.seed)

	for i := 0; i < rain; i++ {
		spot := rng.Intn(terr.width)
		trickleDown := false
		e := terr.heightmap[spot]
		if e.right != nil && e.right.depth < e.depth {
			e.right.depth = e.right.depth - 1
			trickleDown = true
		}

		if e.left != nil && e.left.depth < e.depth {
			e.left.depth = e.left.depth - 1
			trickleDown = true
		}

		if !trickleDown {
			e.depth = e.depth - 1
		}
	}


	for _, earth := range terr.heightmap {
  	slopify(earth)
	}

	return terr
}

func slopify(e *Earth) {
	if e.right != nil && e.right.depth-e.depth > 1 {
		e.right.depth = e.depth + 1
		slopify(e.right)
	}

	if e.left != nil && e.left.depth-e.depth > 1 {
		e.left.depth = e.depth + 1
		slopify(e.left)
	}

  if e.right == nil {
    e.depth = e.left.depth + 1
    return
  }

  if e.left == nil {
    e.depth = e.right.depth + 1
    return
  }

	return
}
