package saga

type Option func(*saga)

func WithMaxGo(maxGo int) Option {
	return func(s *saga) {
		s.maxGo = maxGo
	}
}
