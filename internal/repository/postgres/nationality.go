package postgres

import (
	"context"
	"simple-api/internal/entity"
	"simple-api/internal/repository"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type nationalityRepository struct {
	db *pgx.Conn
	sq sq.StatementBuilderType
}

func NewNationalityRepository(db *pgx.Conn) repository.NationalityRepository {
	return &nationalityRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (repository nationalityRepository) GetAll(ctx context.Context) ([]entity.Nationality, error) {
	var nationalities []entity.Nationality

	builder := repository.sq.Select("nationality_id", "nationality_name", "nationality_code").From("nationalities")
	sql, args, err := builder.ToSql()

	rows, err := repository.db.Query(ctx, sql, args...)
	if err != nil {
		return nationalities, err
	}

	defer rows.Close()

	for rows.Next() {
		nationality := entity.Nationality{}
		err = rows.Scan(&nationality.NationalityId, &nationality.NationalityName, &nationality.NationalityCode)
		if err != nil {
			return nationalities, err
		}

		nationalities = append(nationalities, nationality)
	}

	return nationalities, nil
}

func (repository nationalityRepository) GetById(ctx context.Context, id int) (entity.Nationality, error) {
	nationality := entity.Nationality{}

	builder := repository.sq.Select("nationality_id", "nationality_name", "nationality_code").
		From("nationalities").
		Where(sq.Eq{"nationality_id": id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return nationality, err
	}

	rows, err := repository.db.Query(ctx, sql, args...)
	if err != nil {
		return nationality, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&nationality.NationalityId, &nationality.NationalityName, &nationality.NationalityCode)
		if err != nil {
			return nationality, err
		}

		return nationality, nil
	}

	return nationality, pgx.ErrNoRows
}

func (repository nationalityRepository) Create(ctx context.Context, nationality entity.Nationality) (entity.Nationality, error) {
	builder := repository.sq.Insert("nationalities").
		Columns("nationality_name", "nationality_code").
		Values(nationality.NationalityName, nationality.NationalityCode).
		Suffix("RETURNING nationality_id")

	sql, args, err := builder.ToSql()
	if err != nil {
		return nationality, err
	}

	var lastInsertId int
	err = repository.db.QueryRow(ctx, sql, args...).Scan(&lastInsertId)
	if err != nil {
		return nationality, err
	}

	nationality.NationalityId = lastInsertId

	return nationality, nil
}

func (repository nationalityRepository) Update(ctx context.Context, nationality entity.Nationality) (entity.Nationality, error) {
	builder := repository.sq.Update("nationalities").
		Set("nationality_name", nationality.NationalityName).
		Set("nationality_code", nationality.NationalityCode).
		Where(sq.Eq{"nationality_id": nationality.NationalityId})

	sql, args, err := builder.ToSql()
	if err != nil {
		return nationality, err
	}

	exec, err := repository.db.Exec(ctx, sql, args...)
	if err != nil {
		return nationality, err
	}

	if exec.RowsAffected() == 0 {
		return nationality, pgx.ErrNoRows
	}

	return nationality, nil
}

func (repository nationalityRepository) Delete(ctx context.Context, id int) error {
	builder := repository.sq.Delete("nationalities").Where(sq.Eq{"nationality_id": id})
	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	exec, err := repository.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	if exec.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
