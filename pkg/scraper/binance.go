package scraper

import (
	"encoding/json"
  "fmt"
  "io"
  "log"
  "net/http"
)

type BinanceScraper struct{}

func NewBinanceScraper() *BinanceScraper {
  return &BinanceScraper{}
}

func (b *BinanceScraper) GetCriptoInfo(cripto string) (string, error) {

  if cripto == "BITCOIN" {
    cripto = "BTCUSDT"
  }
  url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", cripto)

  log.Println("Bunscando:", url)
  
  client := &http.Client{}
  req, _ := http.NewRequest("GET", url, nil)

  //   // Headers comuns de navegador
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
  log.Println("response", resp)

	body, _ := io.ReadAll(resp.Body)
  if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %w", err)
	}
  
	var result struct {
		Symbol string  `json:"symbol"`
		Price  float64 `json:"price,string"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("erro ao parsear JSON: %w", err)
	}

	return fmt.Sprintf("Preço do %s: R$ %.2f", cripto, result.Price), nil
}