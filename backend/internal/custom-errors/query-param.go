package customerror

// Error for invalid query parameter.
type QueryParamError struct {
	Message string
}

func (q *QueryParamError) Error() string {
	return q.Message
}
