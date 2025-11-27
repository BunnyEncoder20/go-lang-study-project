package api

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    int
	Balance int
}

type ErrorResponse struct {
	Code    int
	Message string
}
