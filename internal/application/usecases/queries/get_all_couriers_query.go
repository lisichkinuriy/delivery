package queries

type GetAllCouriersQuery struct {
	isSet bool
}

func NewGetAllCouriersQuery() (GetAllCouriersQuery, error) {
	return GetAllCouriersQuery{
		isSet: true,
	}, nil
}

func (q GetAllCouriersQuery) IsEmpty() bool {
	return !q.isSet
}
