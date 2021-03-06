package shared

// BaseFilter will hold basic query/filter
type BaseFilter struct {
	Limit   int    `json:"limit" default:"10"`
	Page    int    `json:"page" default:"1"`
	Offset  int    `json:"-"`
	Search  string `json:"search,omitempty"`
	OrderBy string `json:"orderBy,omitempty" default:"created_at" lower:"true"`
	Sort    string `json:"sort,omitempty" default:"desc" lower:"true"`
}

// CalculateOffset method
func (f *BaseFilter) CalculateOffset() {
	f.Offset = (f.Page - 1) * f.Limit
}
