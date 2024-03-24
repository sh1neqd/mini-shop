package item

type CreateItemDTO struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type UpdateItemDTO struct {
	Name  *string `json:"name"`
	Price *int    `json:"price"`
}

type GetItemsDto struct {
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	Categories []string `json:"categories"`
}

type GetItemWithoutCategories struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type GetCategoryID struct {
	CategoryId int `json:"category_id"`
}
