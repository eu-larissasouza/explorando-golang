package main

// AULA - PADRÃO DE CONCORRÊNCIA: GENERATORS

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Nomenclatura <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				c <- "Erro ao acessar: " + url
				return
			}
			defer resp.Body.Close()
			html, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				c <- "Erro ao ler: " + url
				return
			}

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			matches := r.FindStringSubmatch(string(html))
			if len(matches) > 1 {
				c <- matches[1]
			} else {
				c <- "Título não encontrado: " + url
			}
		}(url)
	}
	return c
}
