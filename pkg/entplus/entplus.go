package entplus

import (
	"github.com/jinzhu/copier"
)

func MustCopyValue(to, from interface{}) {
	err := copier.Copy(to, from)
	if err != nil {
		panic(err)
	}
}
