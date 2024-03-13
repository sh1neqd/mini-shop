package item

type Item struct {
	ID         int32  `json:"id"`
	Name       string `json:"name"`
	CategoryId int32  `json:"category_id"`
}
