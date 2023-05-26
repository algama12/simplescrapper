package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// Crear un nuevo objeto Collector
	c := colly.NewCollector()

	// Crear un slice para almacenar los enlaces encontrados
	links := make([]string, 0, 50)

	// Visitar una página web
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !strings.HasPrefix(link, "http") {
			link = fmt.Sprintf("%s%s", e.Request.URL.Scheme, link)
		}
		links = append(links, link)
	})

	// Visitar la página web
	c.Visit("https://www.fotocasa.es/es/alquiler/vivienda/malaga-capital/ascensor-amueblado/177657479/d")

	// Imprimir los enlaces encontrados
	for _, link := range links {
		fmt.Println(link)
	}
}
