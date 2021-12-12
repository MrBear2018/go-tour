package model

type Tag struct {
	*Model        // 这里通过
	Name   string `json:"name"`
	State  uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
