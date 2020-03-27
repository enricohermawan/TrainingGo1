package skeleton

import (
	"context"
	"log"

	sentity "go-skeleton/internal/entity/skeleton"
	"go-skeleton/pkg/errors"

	"github.com/jmoiron/sqlx"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

// Tambahkan query di dalam const
const (
	getAllUser  = "GetAllUser"
	qGetAllUser = "SELECT * FROM user_test"

	getAllID  = "GetAllID"
	qGetAllID = "Select *from user_test where id = ?"

	getAllNip  = "GetAllNip"
	qGetAllNip = "Select * from user_test where nip = ?"

	getAllNamaLengkap  = "GetAllNamaLengkap"
	qGetAllNamaLengkap = "Select *from user_test where nama_lengkap = ?"

	getAllTanggalLahir  = "GetAllTanggalLahir"
	qGetAllTanggalLahir = "Select *from user_test where tanggal_lahir = ?"

	getAllJabatan  = "GetAllJabatan"
	qGetAllJabatan = "Select *from user_test where jabatan = ?"

	getAllEmail  = "GetAllEmail"
	qGetAllEmail = "select *from user_test where email = ?"

	insertAllUser  = "InsertAllUser"
	qInsertAllUser = "insert into user_test VALUES(?, ?, ?, ?, ?, ?)"

	updateAllUser  = "UpdateAllUser"
	qUpdateAllUser = "UPDATE user_test set nama_lengkap = ?, tanggal_lahir = ?, jabatan = ?, email = ? where nip = ?"

	deleteAllUser  = "DeleteAllUser"
	qDeleteAllUser = "DELETE from user_test where nip = ?"
)

// Tambahkan query ke dalam key value order agar menjadi prepared statements

var (
	readStmt = []statement{
		{getAllUser, qGetAllUser},
		{getAllID, qGetAllID},
		{getAllNip, qGetAllNip},
		{getAllNamaLengkap, qGetAllNamaLengkap},
		{getAllTanggalLahir, qGetAllTanggalLahir},
		{getAllJabatan, qGetAllJabatan},
		{getAllEmail, qGetAllEmail},
		{insertAllUser, qInsertAllUser},
		{updateAllUser, qUpdateAllUser},
		{deleteAllUser, qDeleteAllUser},
	}
)

// New ...
func New(db *sqlx.DB) Data {
	d := Data{
		db: db,
	}

	d.initStmt()
	return d
}

func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}

// GetAllUser ...
func (d Data) GetAllUser(ctx context.Context) ([]sentity.Skeleton, error) {
	var (
		rows      *sqlx.Rows
		skeleton  sentity.Skeleton
		skeletons []sentity.Skeleton
		err       error
	)

	rows, err = d.stmt[getAllUser].QueryxContext(ctx)

	for rows.Next() {
		if err := rows.StructScan(&skeleton); err != nil {
			return skeletons, errors.Wrap(err, "[DATA][GetAllUser]")
		}
		skeletons = append(skeletons, skeleton)
	}
	return skeletons, err
}

// GetAllNip ...
func (d Data) GetAllNip(ctx context.Context, nip string) (sentity.Skeleton, error) {
	var (
		skeleton sentity.Skeleton
		err      error
	)

	
	err = d.stmt[getAllNip].QueryRowxContext(ctx, nip).StructScan(&skeleton)

	if err != nil {
		return skeleton, errors.Wrap(err, "[DATA][GetAllNip]")
	}
	return skeleton, err
}

// InsertAllUser ...
func (d Data) InsertAllUser(ctx context.Context, sk sentity.Skeleton) error {
	_, err := d.stmt[insertAllUser].ExecContext(ctx,
		sk.SkeletonID,
		sk.SkeletonNip,
		sk.SkeletonNama,
		sk.SkeletonTanggalLahir,
		sk.SkeletonJabatan,
		sk.SkeletonEmail,
	)
	return err
}

// UpdateAllUser ...
func (d Data) UpdateAllUser(ctx context.Context, sk sentity.Skeleton) (sentity.Skeleton, error) {
	_, err := d.stmt[updateAllUser].ExecContext(ctx,
		sk.SkeletonNama,
		sk.SkeletonTanggalLahir,
		sk.SkeletonJabatan,
		sk.SkeletonEmail,
		sk.SkeletonNip,
	)

	return sk, err
}

// DeleteAllUser ...
func (d Data) DeleteAllUser(ctx context.Context, sk sentity.Skeleton) error {
	_, err := d.stmt[deleteAllUser].ExecContext(ctx, sk.SkeletonNip)

	return err
}
