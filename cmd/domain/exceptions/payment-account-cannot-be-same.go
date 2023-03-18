package domainexceptions

func PaymentAccountCannotBeSame() *ApplicationException {
	return new("PAYMENT_ACCOUNT_CANNOT_BE_SAME", "Payment value cannot be negative")
}
