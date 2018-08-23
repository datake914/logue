package util

import "github.com/jinzhu/copier"

// Copy copy things
func Copy(src interface{}, dest interface{}) {
	copier.Copy(dest, src)
}
