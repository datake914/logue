package service

// BaseService represents the struct that exexutes service.
type BaseService struct {
}

// NewBaseService returns a pointer to the BaseService struct.
func NewBaseService(opts ...Option) *BaseService {
	s := &BaseService{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
