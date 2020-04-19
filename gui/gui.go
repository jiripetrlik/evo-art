package gui

import (
	"html/template"
	"image"
	"image/color"
	"image/png"
	"net/http"

	"github.com/jiripetrlik/evo-art/cgp"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ModelType struct {
	Chromosomes []string
}

var cgpModel = cgp.NewCgp(2, 25, 3, 0.2)

func choiceEndpoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	var model ModelType
	model.Chromosomes = make([]string, 4)

	chromosomeParam, found := request.URL.Query()["chromosome"]
	if found == true {
		chromosome := cgp.LoadChromosome(chromosomeParam[0])
		model.Chromosomes[0] = chromosome.ToString()

		for i := 1; i < 4; i++ {
			childChromosome := cgpModel.MutateChromosome(chromosome)
			model.Chromosomes[i] = childChromosome.ToString()
		}
	} else {
		for i := 0; i < 4; i++ {
			chromosome := cgpModel.GenerateChromosome()
			model.Chromosomes[i] = chromosome.ToString()
		}
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
	chromosomeParam, found := request.URL.Query()["chromosome"]
	if !found {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - query parameter 'chromosome' is missing"))
	}

	chromosome := cgp.LoadChromosome(chromosomeParam[0])
	r := image.Rect(0, 0, width, height)
	img := image.NewRGBA(r)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			inputs := []int{x, y}
			outputs := cgpModel.Evaluate(inputs, chromosome)

			c := color.RGBA{uint8((*outputs)[0]), uint8((*outputs)[1]),
				uint8((*outputs)[2]), 255}
			img.SetRGBA(x, y, c)
		}
	}
	png.Encode(writer, img)
}

func GuiServer() {
	http.HandleFunc("/", choiceEndpoint)
	http.HandleFunc("/style.css", cssEndpoint)
	http.HandleFunc("/image.png", imageEndpoint)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
