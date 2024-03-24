package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"testAssignment/internal/domain/category"
)

type CategoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r CategoryRepo) Create(dto category.CreateCategoryDTO) (int, error) {
	var id int
	q := `INSERT INTO public.category (name) VALUES ($1) RETURNING id`
	logrus.Println(q)

	row := r.db.QueryRow(q, dto.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	logrus.Println(q)
	return id, nil
}

func (r CategoryRepo) GetAll() ([]category.Category, error) {
	var categories []category.Category
	q := `SELECT name FROM public.category`
	logrus.Println(q)
	err := r.db.Select(&categories, q)
	if err != nil {
		return nil, err
	}
	return categories, err
}

func (r CategoryRepo) GetById(id int) (category.Category, error) {
	var c category.Category
	q := `SELECT id, name FROM public.category WHERE id = $1`
	logrus.Println(q)
	row := r.db.QueryRow(q, id)
	err := row.Scan(&c.ID, &c.Name)
	if err != nil {
		return category.Category{}, err
	}
	return c, err
}

func (r CategoryRepo) Delete(id int) error {
	q1 := `DELETE FROM public.item_category WHERE public.item_category.category_id=$1`
	logrus.Println(q1)
	_, err := r.db.Exec(q1, id)
	q2 := `DELETE FROM public.category WHERE id=$1`
	logrus.Println(q2)
	_, err = r.db.Exec(q2, id)
	if err != nil {
		return err
	}
	return nil
}

func (r CategoryRepo) Update(id int, dto category.CreateCategoryDTO) error {
	q := `UPDATE public.category SET name=$1 WHERE id=$2`
	_, err := r.db.Exec(q, dto.Name, id)
	if err != nil {
		return err
	}
	return nil
}
