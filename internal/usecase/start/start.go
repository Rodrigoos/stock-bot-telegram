package start

type StartUseCase struct{}

func NewStartUseCase() *StartUseCase {
  return &StartUseCase{}
}

func (uc *StartUseCase) Execute() string {
  return "Olá! Bem vindo ao meu bot do mercado fincanceiro \n" +
    "Envie /stock <TICKER> para ver a cotação de ações. \n" +
    "Envie /fund <TICKER> para ver a cotação de fundos.\n" +
    "Envie /portfolio <NOME> para ver sua carteira.\n" +
    "Envie /carteira <NOME> para ver sua carteira.\n" +
    "Envie /start para ver esta mensagem novamente."
}
