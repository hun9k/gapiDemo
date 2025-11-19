//go:generate gapi gen curd

package schemas

import (
	"github.com/hun9k/gapi"
)

type Contents struct {
	Subject  string `gorm:"subject"`
	AuthorId uint   `gorm:"author_id"`
	gapi.Model
}
