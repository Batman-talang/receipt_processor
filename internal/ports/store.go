package ports

type MemoryStore interface {
	Save(point int) string
	Get(id string) (int, bool)
	Clear()
}
