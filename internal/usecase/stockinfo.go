package usecase

type StockInfoFetcher interface {
	GetStockInfo(ticker string) (string, error)
}

type StockInfoUseCase struct {
	fetcher StockInfoFetcher
}

func NewStockInfoUseCase(fetcher StockInfoFetcher) *StockInfoUseCase {
	return &StockInfoUseCase{fetcher: fetcher}
}

func (uc *StockInfoUseCase) Execute(ticker string) (string, error) {
	return uc.fetcher.GetStockInfo(ticker)
}
