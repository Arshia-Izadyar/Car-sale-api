package service_errors

type ServiceError struct {
	EndUserMessage string `json:"end_user_error"`
	Err            error  `json:"error"`
}

func (s *ServiceError) Error() string {
	return s.EndUserMessage
}
