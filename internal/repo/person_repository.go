package repo

import (
	"context"
	"fmt"
	"time"

	"log/slog"

	"github.com/jmoiron/sqlx"
)

type PersonRepository interface {
	Create(ctx context.Context, person *Person) (*Person, error)
	GetByID(ctx context.Context, id int64) (*Person, error)
	GetByFilters(ctx context.Context, filter *Filter) ([]*Person, error)
	DeleteByID(ctx context.Context, id int64) error
	Update(ctx context.Context, p *Person) error
}

type personRepository struct {
	db *sqlx.DB
}

func NewPersonRepository(db *sqlx.DB) PersonRepository {
	return &personRepository{
		db: db,
	}
}

type Person struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Surname     string    `db:"surname" json:"surname"`
	Age         int       `db:"age" json:"age"`
	Gender      string    `db:"gender" json:"gender"`
	Nationality string    `db:"nationality" json:"nationality"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

func (rp *personRepository) Create(ctx context.Context, person *Person) (*Person, error) {
	query := `
		INSERT INTO people(name, surname, age, gender, nationality)
		VALUES(:name, :surname, :age, :gender, :nationality)`

	rows, err := rp.db.NamedQueryContext(ctx, query, person)
	if err != nil {
		return nil, fmt.Errorf("failed to create person: %w", err)
	}

	if rows.Next() {
		err = rows.Scan(person.ID, person.CreatedAt, person.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan person: %w", err)
		}
	}

	return person, nil
}

func (rp *personRepository) GetByID(ctx context.Context, id int64) (*Person, error) {
	query := `
		SELECT id, name, surname, age, gender, nationality, created_at, updated_at
		FROM people
		WHERE id=$1;`

	var p Person
	err := rp.db.GetContext(ctx, &p, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch person by id: %w", err)
	}

	return &p, nil
}

func (rp *personRepository) GetByFilters(ctx context.Context, filter *Filter) ([]*Person, error) {
	query := `
		SELECT id, name, surname, age, gender, nationality, created_at, updated_at 
		FROM people;`
	query += filter.String()
	slog.Info(query)

	var p []*Person
	err := rp.db.SelectContext(ctx, &p, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch people by filters: %w", err)
	}

	return p, nil
}

func (rp *personRepository) DeleteByID(ctx context.Context, id int64) error {
	query := `
		DELETE FROM people WHERE id=$1`

	_, err := rp.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete person by id: %w", err)
	}

	return nil
}

func (rp *personRepository) Update(ctx context.Context, p *Person) error {
	slog.Debug("person on update", "person", *p)
	query := `
        UPDATE people
        SET name = :name, surname = :surname, age = :age, nationality = :nationality,
            gender = :gender, updated_at = :updated_at
        WHERE id = :id`

	_, err := rp.db.NamedExecContext(ctx, query, p)
	if err != nil {
		return fmt.Errorf("failed to update person by id: %w", err)
	}
	return nil
}
