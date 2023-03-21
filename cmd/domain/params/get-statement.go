package params

import entity "cqrs-go/cmd/domain/entities"

type GetStatementQueryParams struct {
	AccountId entity.ID
}
