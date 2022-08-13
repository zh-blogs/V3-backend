package mock

import (
	"github/zhblogs/backend/provider"
	"github/zhblogs/backend/provider/types"
	"github/zhblogs/backend/utils/set"
	"math/rand"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type MockProvider struct {
	Blogs []Blog
}

func New() provider.Provider {
	return &MockProvider{
		Blogs: append([]Blog{}, mockBlogs...),
	}
}

func (p *MockProvider) GetFeatureBlogs() ([]types.BlogSimple, error) {
	res := make([]types.BlogSimple, 0, 10)
	for _, blog := range p.Blogs {
		if blog.Recommend {
			res = append(res, blog.ToBlogSimple())
		}
	}

	return res, nil
}

func (p *MockProvider) GetRandomBlogs(tags []string, count int) ([]types.BlogSimple, error) {
	_, blogs, err := p.GetBlogs("", tags, types.BlogStatusEnabled, 0, -1)
	if err != nil {
		return nil, errors.Wrapf(err, "get blogs error")
	}

	l := len(blogs)
	idxArr := make([]int, l)
	for i := 0; i < l; i++ {
		idxArr[i] = i
	}

	rand.Shuffle(l, func(i, j int) { idxArr[i], idxArr[j] = idxArr[j], idxArr[i] })

	if count > l || count < 0 {
		count = l
	}

	idxArr = idxArr[:count]

	res := make([]types.BlogSimple, 0, count)
	for _, idx := range idxArr {
		res = append(res, blogs[idx])
	}

	return res, nil
}

func (p *MockProvider) GetBlogs(search string, tags []string, status types.BlogStatus, offset int, limit int) (int, []types.BlogSimple, error) {
	blogsFilter := Blogs(append([]Blog{}, p.Blogs...))
	logrus.Debugf(
		"get blogs: search %s tags %v, status %s offset %d limit %d, total %d blogs",
		search, tags, status, offset, limit, len(blogsFilter),
	)

	if search != "" {
		search = strings.ToLower(search)
		blogsFilter = blogsFilter.Filter(func(blog Blog) bool {
			return strings.Contains(strings.ToLower(blog.Name), search) ||
				strings.Contains(strings.ToLower(blog.URL), search)
		})
	}
	logrus.Debugf("after filter search, has %d blogs", len(blogsFilter))

	tagsSet := set.NewMapSet()
	for _, tag := range tags {
		tag = strings.ToLower(strings.TrimSpace(tag))
		if tag != "" {
			tagsSet.Add(tag)
		}
	}
	if tagsSet.Len() > 0 {
		blogsFilter = blogsFilter.Filter(func(blog Blog) bool {
			blogTags := set.NewMapSet()
			for _, tag := range blog.Tags {
				blogTags.Add(strings.ToLower(strings.TrimSpace(tag)))
			}

			return tagsSet.Difference(blogTags).Len() == 0
		})
	}
	logrus.Debugf("after filter tags, has %d blogs", len(blogsFilter))

	blogsFilter.Filter(func(blog Blog) bool {
		switch status {
		case types.BlogStatusEnabled: // enabled
			return blog.Enabled
		case types.BlogStatusDisabled: // not enabled
			return !blog.Enabled
		case types.BlogStatusRecommend: // recommend
			return blog.Recommend
		case types.BlogStatusAll: // all
			return true
		default: // default enabled
			return blog.Enabled
		}
	})
	logrus.Debugf("after filter status, has %d blogs", len(blogsFilter))

	total := len(blogsFilter)

	if offset > total {
		offset = total
	}
	if limit+offset > total || limit < 0 {
		limit = total - offset
	}

	return total, blogsFilter[offset:limit].ToBlogSimple(), nil
}
