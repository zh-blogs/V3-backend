package types

import "fmt"

type BlogStatus int

const (
	BlogStatusDisabled  BlogStatus = -1
	BlogStatusAll       BlogStatus = 0
	BlogStatusEnabled   BlogStatus = 1
	BlogStatusRecommend BlogStatus = 2
)

func (s BlogStatus) String() string {
	switch s {
	case BlogStatusDisabled:
		return "BlogStatusDisabled"
	case BlogStatusAll:
		return "BlogStatusAll"
	case BlogStatusEnabled:
		return "BlogStatusEnabled"
	case BlogStatusRecommend:
		return "BlogStatusRecommend"
	default:
		return fmt.Sprintf("UnkownBlogStatus(%d)", s)
	}
}

func (s BlogStatus) Vaild() bool {
	switch s {
	case BlogStatusAll, BlogStatusDisabled, BlogStatusEnabled, BlogStatusRecommend:
		return true
	default:
		return false
	}
}
