package maps

type Container interface {
	Get(key any) (value any, ok bool)
	Set(key any, value any)
	Delete(key any)
	Keys() (keys []any)
}

func CopyBeanContainer(dst Container, src Container) {
	for _, key := range src.Keys() {
		value, ok := src.Get(key)
		if !ok {
			continue
		}
		dst.Set(key, value)
	}
}
