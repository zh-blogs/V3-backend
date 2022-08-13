package mock

import "github/zhblogs/backend/provider/types"

type Blog struct {
	ID      string   `json:"id"`
	Idx     int      `json:"idx"`
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Enabled bool     `json:"enabled"`
	Tags    []string `json:"tags"`

	JoinTime   int64 `json:"join_time"`
	UpdateTime int64 `json:"update_time"`

	Arch      string `json:"arch"`
	Feed      string `json:"feed"`
	Recommend bool   `json:"recommend"`
	Repeat    bool   `json:"repeat"`
	SaveWebID string `json:"saveweb_id"`
	Sign      string `json:"sign"`
	SiteMap   string `json:"sitemap"`
	Status    string `json:"status"`
}

func (blog Blog) ToBlogSimple() types.BlogSimple {
	return types.BlogSimple{
		ID:   blog.ID,
		Idx:  blog.Idx,
		Name: blog.Name,
		URL:  blog.URL,
		Tags: append([]string{}, blog.Tags...),
		Sign: blog.Sign,
		Feed: blog.Feed,
	}
}

type Blogs []Blog

func (bs Blogs) Filter(callback func(blog Blog) bool) Blogs {
	res := make(Blogs, 0, len(bs))

	for _, blog := range bs {
		if callback(blog) {
			res = append(res, blog)
		}
	}

	bs = res
	return res
}

func (bs Blogs) ToBlogSimple() []types.BlogSimple {
	res := make([]types.BlogSimple, len(bs))

	for i, blog := range bs {
		res[i] = blog.ToBlogSimple()
	}

	return res
}
