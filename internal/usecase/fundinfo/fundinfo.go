package fundinfo

type FundInfoFetcher interface {
	GetFundInfo(ticker string) (string, error)
}

type FundInfoUseCase struct {
	fetcher FundInfoFetcher
}

func NewFundInfoUseCase(fetcher FundInfoFetcher) *FundInfoUseCase {
	return &FundInfoUseCase{fetcher: fetcher}
}

func (uc *FundInfoUseCase) Execute(ticker string) (string, error) {
	return uc.fetcher.GetFundInfo(ticker)
}
