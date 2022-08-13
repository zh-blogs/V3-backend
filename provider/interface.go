package provider

import "github/zhblogs/backend/provider/types"

//go:generate mockgen -source=interface.go -destination mocks/mocks.go
type Provider interface {
	GetFeatureBlogs() ([]types.BlogSimple, error)
	GetRandomBlogs(tags []string, count int) ([]types.BlogSimple, error)
	GetBlogs(search string, tags []string, status types.BlogStatus, offset int, limit int) (int, []types.BlogSimple, error)
}
