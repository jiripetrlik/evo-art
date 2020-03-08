package gui

import (
	"html/template"
	"image"
	"image/color"
	"image/png"
	"net/http"

	"github.com/jiripetrlik/evo-art/cgp"
)

type ModelType struct {
	Chromosomes []string
}

var cgpModel = cgp.NewCgp(2, 25, 3, 0.05)

func choiceEndpoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	var model ModelType
	model.Chromosomes = make([]string, 4)
	for i := 0; i < 4; i++ {
		chromosome := cgpModel.GenerateChromosome()
		model.Chromosomes[i] = chromosome.ToString()
	}

	tmpl, err := template.New("choice").Parse(choicePageTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(writer, model)
}

func cssEndpoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/css")

	writer.Write([]byte(css))
}

func imageEndpoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "image/png")

	const width = 256
	const height = 256

	r := image.Rect(0, 0, width, height)
	img := image.NewRGBA(r)
	c := color.RGBA{0, 0, 255, 255}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.SetRGBA(x, y, c)
		}
	}
	png.Encode(writer, img)
}

func GuiServer() {
	http.HandleFunc("/", choiceEndpoint)
	http.HandleFunc("/style.css", cssEndpoint)
	http.HandleFunc("/image.png", imageEndpoint)

	http.ListenAndServe(":8080", nil)
}
