package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	url := "https://solucoes.sje.ifmg.edu.br/CardapioUan/cardapio.php"

	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Erro: Status code %d recebido da URL", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("Erro ao carregar o documento HTML: %v", err)
	}

	currentTime := time.Now().Format("02/01/2006")

	re := regexp.MustCompile(`\s+`)

	var find bool = false

	text := "Cardápio - " + currentTime

	doc.Find("table thead").Each(func(i int, s *goquery.Selection) {
		txt := s.Text()
		cleanText := re.ReplaceAllString(strings.TrimSpace(txt), " ")

		if find {
			save(text)
		}

		if strings.Contains(cleanText, currentTime) {
			find = true
			nextTbody := s.NextFiltered("tbody")
			if nextTbody.Length() > 0 {
				nextTbody.Each(func(j int, tbody *goquery.Selection) {
					text += tbody.Text()
					text = formatText(text)
				})
			}
		}
	})

	if !find {
		text = "Cardápio do dia " + currentTime + " não encontrado."
	}

	save(text)

}

func formatText(text string) string {
	text = strings.ReplaceAll(text, "                    ", "")
	return strings.TrimSpace(text)
}

func save(text string) error {
	err := os.WriteFile("cardapio-"+time.Now().Format("02-01-2006")+".txt", []byte(text), 0644)
	if err != nil {
		log.Fatalf("Erro ao salvar o arquivo")
	}
	return err
}
