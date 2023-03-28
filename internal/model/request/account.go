package request

type GetAccountByID struct {
	AccountID int64 `json:"account_id" binding:"required,gte=1"`
}

type GetAccountListByUserID struct {
}
