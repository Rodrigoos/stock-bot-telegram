package usecase

type CriptoInfoFetcher interface {
  GetCriptoInfo(cripto string) (string, error)
}

type CriptoInfoUseCase struct {
  fetcher CriptoInfoFetcher
}

func NewCriptoInfoUseCase(fetcher CriptoInfoFetcher) *CriptoInfoUseCase {
  return &CriptoInfoUseCase{fetcher: fetcher}
}

func (uc *CriptoInfoUseCase) Execute(cripto string) (string, error) {
  return uc.fetcher.GetCriptoInfo(cripto)
}
