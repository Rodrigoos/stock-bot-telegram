package start

type StartUseCase struct{}

func NewStartUseCase() *StartUseCase {
	return &StartUseCase{}
}

func (uc *StartUseCase) Execute() string {
	return "Olá! Envie /stock <TICKER> para ver a cotação."
}
