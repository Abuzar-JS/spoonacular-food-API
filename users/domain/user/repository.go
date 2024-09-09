package user

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
}

type WriteRepository interface {
}
