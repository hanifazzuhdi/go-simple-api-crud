package postgres

import (
	"context"
	"errors"
	"simple-api/internal/entity"
	"simple-api/internal/repository"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type familyListRepository struct {
	db *pgx.Conn
	sq sq.StatementBuilderType
}

func NewFamilyRepository(db *pgx.Conn) repository.FamilyListRepository {
	return &familyListRepository{
		db: db,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (repository familyListRepository) GetAll(ctx context.Context) ([]entity.FamilyList, error) {
	var familyLists []entity.FamilyList

	builder := repository.sq.Select("fl_id", "cst_id", "fl_relation", "fl_name", "fl_dob").From("family_list")
	sql, args, err := builder.ToSql()
	if err != nil {
		return familyLists, err
	}

	rows, err := repository.db.Query(ctx, sql, args...)
	if err != nil {
		return familyLists, err
	}

	defer rows.Close()

	for rows.Next() {
		familyList := entity.FamilyList{}
		err = rows.Scan(&familyList.FlId, &familyList.CstId, &familyList.FlRelation, &familyList.FlName, &familyList.FlDob)
		if err != nil {
			return familyLists, err
		}

		familyLists = append(familyLists, familyList)
	}

	return familyLists, nil
}

func (repository familyListRepository) GetAllByCustomerID(ctx context.Context, customerID int) ([]entity.FamilyList, error) {
	var familyLists []entity.FamilyList

	builder := repository.sq.Select("fl_id", "cst_id", "fl_relation", "fl_name", "fl_dob").From("family_list").Where(sq.Eq{"cst_id": customerID})
	sql, args, err := builder.ToSql()
	if err != nil {
		return familyLists, err
	}

	rows, err := repository.db.Query(ctx, sql, args...)
	if err != nil {
		return familyLists, err
	}

	defer rows.Close()

	for rows.Next() {
		familyList := entity.FamilyList{}
		err = rows.Scan(&familyList.FlId, &familyList.CstId, &familyList.FlRelation, &familyList.FlName, &familyList.FlDob)
		if err != nil {
			return familyLists, err
		}

		familyLists = append(familyLists, familyList)
	}

	return familyLists, nil
}

func (repository familyListRepository) GetByID(ctx context.Context, id int) (entity.FamilyList, error) {
	familyList := entity.FamilyList{}

	builder := repository.sq.Select("fl_id", "cst_id", "fl_relation", "fl_name", "fl_dob").From("family_list").Where(sq.Eq{"fl_id": id})
	sql, args, err := builder.ToSql()
	if err != nil {
		return familyList, err
	}

	row := repository.db.QueryRow(ctx, sql, args...)
	err = row.Scan(&familyList.FlId, &familyList.CstId, &familyList.FlRelation, &familyList.FlName, &familyList.FlDob)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return familyList, pgx.ErrNoRows
		}

		return familyList, err
	}

	return familyList, nil
}

func (repository familyListRepository) Create(ctx context.Context, familyList entity.FamilyList) (entity.FamilyList, error) {
	builder := repository.sq.Insert("family_list").Columns("cst_id", "fl_relation", "fl_name", "fl_dob").
		Values(familyList.CstId, familyList.FlRelation, familyList.FlName, familyList.FlDob).
		Suffix("RETURNING fl_id")

	sql, args, err := builder.ToSql()
	if err != nil {
		return familyList, err
	}

	var lastInsertId int
	err = repository.db.QueryRow(ctx, sql, args...).Scan(&lastInsertId)
	if err != nil {
		return familyList, err
	}

	familyList.FlId = lastInsertId
	return familyList, nil
}

func (repository familyListRepository) Update(ctx context.Context, familyList entity.FamilyList) (entity.FamilyList, error) {
	builder := repository.sq.Update("family_list").
		Set("cst_id", familyList.CstId).
		Set("fl_relation", familyList.FlRelation).
		Set("fl_name", familyList.FlName).
		Set("fl_dob", familyList.FlDob).
		Where(sq.Eq{"fl_id": familyList.FlId})

	sql, args, err := builder.ToSql()
	if err != nil {
		return familyList, err
	}

	exec, err := repository.db.Exec(ctx, sql, args...)
	if err != nil {
		return familyList, err
	}

	if exec.RowsAffected() == 0 {
		return familyList, pgx.ErrNoRows
	}

	return familyList, nil
}

func (repository familyListRepository) Delete(ctx context.Context, id int) error {
	builder := repository.sq.Delete("family_list").Where(sq.Eq{"fl_id": id})
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
