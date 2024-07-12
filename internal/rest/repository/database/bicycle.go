package database

import (
	"context"
	"database/sql"

	"pr1/internal/rest/model"
)

type DB struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *DB {
	return &DB{db}
}

func (db *DB) Create(ctx context.Context, bicycle model.Bicycle) error {
	const q = `
		insert into bicycles (brand, model, price) values ($1, $2, $3);
		`

	_, err := db.ExecContext(ctx, q, bicycle.Brand, bicycle.Model, bicycle.Price)
	return err
}

func (db *DB) Read(ctx context.Context, id int64) (model.Bicycle, error) {
	const q = `
		select id, brand, model, price from bicycles where id = $1;	
	`
	bicycle := model.Bicycle{}
	return bicycle, db.QueryRowContext(ctx, q, id).Scan(
		&bicycle.ID,
		&bicycle.Brand,
		&bicycle.Model,
		&bicycle.Price,
	)
}

func (db *DB) Update(ctx context.Context, bicycle model.Bicycle) error {
	const q = `
		update bicycles set brand = $1, model = $2, price = $3 where id = $4;
		`
	_, err := db.ExecContext(ctx, q, bicycle.Brand, bicycle.Model, bicycle.Price, bicycle.ID)
	return err
}

func (db *DB) Delete(ctx context.Context, id int64) error {
	const q = `delete from bicycles where id = $1;`
	_, err := db.ExecContext(ctx, q, id)
	return err
}

func (db *DB) List(ctx context.Context) ([]model.Bicycle, error) {
	const q = `select id, brand, model, price from bicycles order by id desc ;`
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	bicycles := make([]model.Bicycle, 0, 16)
	for rows.Next() {
		bicycle := model.Bicycle{}
		if err := rows.Scan(
			&bicycle.ID,
			&bicycle.Brand,
			&bicycle.Model,
			&bicycle.Price,
		); err != nil {
			return nil, err
		}

		bicycles = append(bicycles, bicycle)
	}
	return bicycles, nil
}
