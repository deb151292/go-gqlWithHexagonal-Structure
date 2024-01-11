package port

type Calculator interface {
	Add(a, b int) (int, error)
}