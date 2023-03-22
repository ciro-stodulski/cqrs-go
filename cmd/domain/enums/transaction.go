package enums

type TransactionKind string

const (
	Pix  TransactionKind = "PIX_PAYMENT"
	Bill TransactionKind = "BILL_PAYMENT"
)

type TransactionStatus string

const (
	Successful TransactionStatus = "SUCCESSFUL"
	Fail       TransactionStatus = "FAIL"
)
