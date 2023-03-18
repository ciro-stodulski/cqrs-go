package usecase

import entity "cqrs-go/cmd/domain/entities"

type GetBalanceUseCase interface {
	perform(accountId entity.ID) int
}
