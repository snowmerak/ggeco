package imgmnger

import "errors"

type ClientBeanKey string

var errClientNotFound error = errors.New("client not found")

func PushClient(beanContainer Container, value *ImageManager) {
	beanContainer.Set(ClientBeanKey("Clientkey"), value)
}

func GetClient(beanContainer Container) (value *ImageManager, err error) {
	maybe, ok := beanContainer.Get(ClientBeanKey("Clientkey"))
	if !ok {
		return nil, errClientNotFound
	}
	value, ok = maybe.(*ImageManager)
	if !ok {
		return nil, errClientNotFound
	}
	return value, nil
}

func IsErrClientNotFound(err error) (ok bool) {
	return errors.Is(err, errClientNotFound)
}
