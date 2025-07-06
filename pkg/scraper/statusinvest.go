package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type StatusInvestScraper struct{}

func NewStatusInvestScraper() *StatusInvestScraper {
	return &StatusInvestScraper{}
}

func (s *StatusInvestScraper) GetStockInfo(ticker string) (string, error) {
	url := fmt.Sprintf("https://statusinvest.com.br/acoes/%s", ticker)

	log.Println("Bunscando:", url)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	// Headers comuns de navegador
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "pt-BR,pt;q=0.9")
	req.Header.Set("Referer", "https://www.google.com/")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println("Status:", resp.StatusCode)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	header := doc.Find(".top-info").First()

	title := doc.Find("h1").First().Text()
	price := header.Find(".value").First().Text()
	change := doc.Find(".sub-value b").First().Text()

	log.Println("Preço: R$", price)
	log.Println("Variação:", change)

	if price == "" || change == "" {
		return "", fmt.Errorf("não foi possível encontrar preço ou variação")
	}

	return fmt.Sprintf("%s\nPreço: R$ %s\nVariação: %s", title, price, change), nil
}

func (s *StatusInvestScraper) GetFundInfo(ticker string) (string, error) {
	url := fmt.Sprintf("https://statusinvest.com.br/fundos-imobiliarios/%s", ticker)

	log.Println("Bunscando:", url)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	// Headers comuns de navegador
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "pt-BR,pt;q=0.9")
	req.Header.Set("Referer", "https://www.google.com/")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println("Status:", resp.StatusCode)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	header := doc.Find(".top-info").First()

	title := doc.Find("h1").First().Text()
	price := header.Find(".value").First().Text()
	change := doc.Find(".sub-value b").First().Text()

	log.Println("Preço:", price)
	log.Println("Variação:", change)

	if price == "" || change == "" {
		return "", fmt.Errorf("não foi possível encontrar preço ou variação")
	}

	return fmt.Sprintf("%s\nPreço: R$ %s\nVariação: %s", title, price, change), nil
}

func (s *StatusInvestScraper) GetStockPrice(ticker string) (float64, error) {
	url := fmt.Sprintf("https://statusinvest.com.br/acoes/%s", strings.ToLower(ticker))

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("erro na requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("status %d recebido", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("erro ao ler HTML: %w", err)
	}

	// Procura por algo como <strong class="value">R$ 10,99</strong>
	priceStr := doc.Find("strong[class='value']").First().Text()
	priceStr = strings.ReplaceAll(priceStr, "R$", "")
	priceStr = strings.ReplaceAll(priceStr, ".", "")
	priceStr = strings.ReplaceAll(priceStr, ",", ".")
	priceStr = strings.TrimSpace(priceStr)

	price, err := parseFloat(priceStr)
	if err != nil {
		return 0, fmt.Errorf("erro ao converter preço: %w", err)
	}

	return price, nil
}

func (s *StatusInvestScraper) GetFundPrice(ticker string) (float64, error) {

	url := fmt.Sprintf("https://statusinvest.com.br/fundos-imobiliarios/%s", strings.ToLower(ticker))

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("erro na requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("status %d recebido", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("erro ao ler HTML: %w", err)
	}

	// Procura por algo como <strong class="value">R$ 10,99</strong>
	priceStr := doc.Find("strong[class='value']").First().Text()
	priceStr = strings.ReplaceAll(priceStr, "R$", "")
	priceStr = strings.ReplaceAll(priceStr, ".", "")
	priceStr = strings.ReplaceAll(priceStr, ",", ".")
	priceStr = strings.TrimSpace(priceStr)

	price, err := parseFloat(priceStr)
	if err != nil {
		return 0, fmt.Errorf("erro ao converter preço: %w", err)
	}

	return price, nil
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
