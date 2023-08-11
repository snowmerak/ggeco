package sqlserver

import "encoding/base64"

type UUID []byte

func (u *UUID) String() string {
	return base64.URLEncoding.EncodeToString(*u)
}

func (u *UUID) From(s string) error {
	b, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	*u = b
	return nil
}
