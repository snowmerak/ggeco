package sqlserver

type Container interface {
	Get(key any) (value any, ok bool)
	Set(key any, value any)
	Delete(key any)
}
