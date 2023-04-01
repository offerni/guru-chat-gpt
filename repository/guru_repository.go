package repository

type ListCardsOpts struct {
	Username string
	Password string
}

type GuruRepo struct {
	Username string
	Password string
}

type NewGuruRepoOpts struct {
	Username string
	Password string
}

func NewGuruRepository(opts NewGuruRepoOpts) (*GuruRepo, error) {
	return &GuruRepo{
		Username: opts.Username,
		Password: opts.Password,
	}, nil
}
