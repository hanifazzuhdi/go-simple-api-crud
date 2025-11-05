package postgres

import (
	"context"
	"errors"
	"simple-api/internal/entity"
	"simple-api/internal/repository"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type customerRepository struct {
	db *pgx.Conn
	sq sq.StatementBuilderType
}

func NewCustomerRepository(db *pgx.Conn) repository.CustomerRepository {
	return &customerRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (repository customerRepository) GetAll(ctx context.Context) ([]entity.Customer, error) {
	builder := repository.sq.Select("cst_id", "cst_name", "cst_dob", "cst_phonenum", "cst_email", "customers.nationality_id", "nationality_name", "nationality_code").
		Join("nationalities ON customers.nationality_id = nationalities.nationality_id").
		From("customers")

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := repository.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []entity.Customer
	for rows.Next() {
		var customer entity.Customer
		err = rows.Scan(&customer.CstId, &customer.CstName, &customer.CstDob, &customer.CstPhonenum, &customer.CstEmail, &customer.NationalityId, &customer.Nationality.NationalityName, &customer.Nationality.NationalityCode)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (repository customerRepository) GetById(ctx context.Context, id int) (entity.Customer, error) {
	customer := entity.Customer{}

	builder := repository.sq.Select("cst_id", "cst_name", "cst_dob", "cst_phonenum", "cst_email", "customers.nationality_id", "nationality_name", "nationality_code").
		From("customers").
		Join("nationalities ON customers.nationality_id = nationalities.nationality_id").
		Where(sq.Eq{"cst_id": id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return customer, err
	}

	row := repository.db.QueryRow(ctx, sql, args...)
	err = row.Scan(&customer.CstId, &customer.CstName, &customer.CstDob, &customer.CstPhonenum, &customer.CstEmail, &customer.NationalityId, &customer.Nationality.NationalityName, &customer.Nationality.NationalityCode)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return customer, pgx.ErrNoRows
		}

		return customer, err
	}

	return customer, nil
}

func (repository customerRepository) Create(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	builder := repository.sq.Insert("customers").Columns("cst_name", "cst_dob", "cst_phonenum", "cst_email", "nationality_id").
		Values(customer.CstName, customer.CstDob, customer.CstPhonenum, customer.CstEmail, customer.NationalityId).
		Suffix("RETURNING cst_id")

	sql, args, err := builder.ToSql()
	if err != nil {
		return customer, err
	}

	var lastInsertId int
	row := repository.db.QueryRow(ctx, sql, args...)
	err = row.Scan(&lastInsertId)
	if err != nil {
		return customer, err
	}

	customer.CstId = lastInsertId
	return customer, nil
}

func (repository customerRepository) Update(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	builder := repository.sq.Update("customers").
		Set("cst_name", customer.CstName).
		Set("cst_dob", customer.CstDob).
		Set("cst_phonenum", customer.CstPhonenum).
		Set("cst_email", customer.CstEmail).
		Set("nationality_id", customer.NationalityId).
		Where(sq.Eq{"cst_id": customer.CstId})

	sql, args, err := builder.ToSql()
	if err != nil {
		return customer, err
	}

	exec, err := repository.db.Exec(ctx, sql, args...)
	if err != nil {
		return customer, err
	}

	if exec.RowsAffected() == 0 {
		return customer, pgx.ErrNoRows
	}

	return customer, nil
}

func (repository customerRepository) Delete(ctx context.Context, id int) error {
	builder := repository.sq.Delete("customers").Where(sq.Eq{"cst_id": id})
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
