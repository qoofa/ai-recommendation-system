package orderembeddings

type service struct {
	repo Repository
}

func New(r Repository) *service {
	return &service{
		repo: r,
	}
}
