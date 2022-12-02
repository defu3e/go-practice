package filter

import (
	"encoding/json"
	"io"
)

func (fc *FilterCity) UnmarshalFilt(body *io.ReadCloser) error {
	err := json.NewDecoder(*body).Decode(&fc)
	if err != nil {
		return err
	}
	defer (*body).Close()

	return nil
}

type FilterCity struct {
	Region     string   `json:"region,omitempty"`
	District   string   `json:"district,omitempty"`
	Foundation []uint64 `json:"foundation,omitempty"`
	Population []uint64 `json:"population,omitempty"`
}
