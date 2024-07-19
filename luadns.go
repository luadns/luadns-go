package luadns

const (
	baseURL = "https://api.luadns.com/v1"
)

// ListMeta stores parsed headers used for result pagination.
type ListMeta struct {
	Page       uint64
	Limit      uint64
	TotalCount uint64
	PagesCount uint64
}
