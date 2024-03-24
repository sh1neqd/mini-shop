package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	category2 "testAssignment/internal/domain/category"
	"testAssignment/internal/domain/item"
)

type ItemRepo struct {
	db *sqlx.DB
}

func NewItemRepo(db *sqlx.DB) *ItemRepo {
	return &ItemRepo{db: db}
}

func (r ItemRepo) Create(dto item.CreateItemDTO) (int, error) {
	var id int
	q := `INSERT INTO public.item (name, price) VALUES ($1, $2) RETURNING id`
	logrus.Println(q)
	row := r.db.QueryRow(q, dto.Name, dto.Price)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	logrus.Println(q)
	return id, nil
}

func (r ItemRepo) GetAll() ([]item.GetItemsDto, error) {
	var items []item.GetItemsDto
	q := `SELECT i.name, i.price, c.name FROM public.item AS i
	      JOIN item_category AS ic ON i.id = ic.item_id
	      JOIN public.category AS c ON ic.category_id = c.id`
	fmt.Println(q)
	rows, err := r.db.Query(q)
	if err != nil {
		logrus.Println()
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		i := item.GetItemsDto{}
		category := category2.Category{}
		err = rows.Scan(&i.Name, &i.Price, &category.Name)
		if err != nil {
			return nil, err
		}
		i.Categories = append(i.Categories, category.Name)
		items = append(items, i)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r ItemRepo) GetById(id int) (item.GetItemsDto, error) {
	var i item.GetItemsDto
	q := `SELECT name, price FROM public.item WHERE id = $1`
	row := r.db.QueryRow(q, id)
	err := row.Scan(&i.Name, &i.Price)
	if err != nil {
		fmt.Println(err.Error())
		return item.GetItemsDto{}, err
	}

	i.Categories, err = r.getCategoriesByUserId(id)
	if err != nil {
		logrus.Error(err.Error())
	}
	logrus.Println(q)

	return i, err
}

func (r ItemRepo) getCategoriesByUserId(id int) ([]string, error) {
	var categories []string
	q := `SELECT public.category.name FROM public.item 
      JOIN item_category ON item.id = item_category.item_id 
      JOIN public.category ON item_category.category_id = category.id 
      WHERE public.item.id=$1`
	fmt.Println(q)
	err := r.db.Select(&categories, q, id)
	if err != nil {
		logrus.Printf("Failed to get Categories By User id. error: %v", err)
		return nil, err
	}
	return categories, nil
}

func (r ItemRepo) Delete(id int) error {
	q1 := `DELETE FROM public.item_category WHERE item_id=$1`
	_, err := r.db.Exec(q1, id)
	logrus.Println(q1)

	q2 := `DELETE FROM public.item WHERE id=$1`
	logrus.Println(q2)
	_, err = r.db.Exec(q2, id)
	if err != nil {
		return err
	}
	return nil
}

func (r ItemRepo) Update(id int, dto item.UpdateItemDTO) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if dto.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *dto.Name)
		argId++
	}

	if dto.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *dto.Price)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	var err error
	q := fmt.Sprintf(`UPDATE public.item SET %s WHERE id=$%d`, setQuery, argId)
	logrus.Println(q)
	args = append(args, id)
	_, err = r.db.Exec(q, args...)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r ItemRepo) AddCategoryForItem(itemId, categoryId int) error {
	q := `INSERT INTO public.item_category(item_id, category_id) VALUES ($1, $2)`
	fmt.Println(q)
	_, err := r.db.Query(q, itemId, categoryId)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
