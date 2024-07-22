package luadns

import "time"

type Zone struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name"`
	Records   []*Record `json:"records,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
