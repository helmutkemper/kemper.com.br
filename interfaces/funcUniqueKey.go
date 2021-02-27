package interfaces

type UniqueKey interface {
	Make() (uniqueKey string)
}
