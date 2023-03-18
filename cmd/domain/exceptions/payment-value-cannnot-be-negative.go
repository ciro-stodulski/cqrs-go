package domainexceptions

func PaymentCannotBeNegative() *ApplicationException {
	return new("PAYMENT_VALUE_CANNOT_BE_NEGATIVE", "Payment value cannot be negative")
}
