package main

//  AULA - PADRÃO DE CONCORRÊNCIA: MULTIPLEXAR

// Multiplexar é unir 2 fontes de dados em um único canal,
// ou seja, unindo dois channels em um só

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

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

func main() {
	c := juntar(
		titulo("https://www.cod3r.com.br", "https://www.google.com"),
		titulo("https://github.com/", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
