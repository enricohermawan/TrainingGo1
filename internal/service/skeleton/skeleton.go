package skeleton

import (
	"context"
	"fmt"
	sentity "go-skeleton/internal/entity/skeleton"
	"go-skeleton/pkg/errors"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	GetAllUser(ctx context.Context) ([]sentity.Skeleton, error)
	GetAllNip(ctx context.Context, nip string) (sentity.Skeleton, error)
	InsertAllUser(ctx context.Context, sk sentity.Skeleton) error
	UpdateAllUser(ctx context.Context, sk sentity.Skeleton) (sentity.Skeleton, error)
	DeleteAllUser(ctx context.Context, sk sentity.Skeleton) error
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	data Data
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(data Data) Service {
	// Assign variable dari parameter ke object
	return Service{
		data: data,
	}
}

// GetAllUser ...
func (s Service) GetAllUser(ctx context.Context) ([]sentity.Skeleton, error) {
	result, err := s.data.GetAllUser(ctx)

	if err != nil {
		return result, errors.Wrap(err, "[SERVICES][GetAllUser]")
	}

	return result, err
}

// GetAllNip ...
func (s Service) GetAllNip(ctx context.Context, nip string) (sentity.Skeleton, error) {
	result, err := s.data.GetAllNip(ctx, nip)

	if err != nil {
		return result, errors.Wrap(err, "[SERVICE][GetAllNip]")
	}
	return result, err
}

// InsertAllUser ...
func (s Service) InsertAllUser(ctx context.Context, sk sentity.Skeleton) error {
	err := s.data.InsertAllUser(ctx, sk)
	fmt.Println(sk)
	return err
}

// UpdateAllUser ...
func (s Service) UpdateAllUser(ctx context.Context, sk sentity.Skeleton) (sentity.Skeleton, error) {
	// fmt.Println(sk.SkeletonTanggalLahir)

	// temptglLahir := sk.SkeletonTanggalLahir.Format("2006-01-02")
	// fmt.Println(temptglLahir)
	// sk.SkeletonTanggalLahir = temptglLahir.Time
	result, err := s.data.UpdateAllUser(ctx, sk)

	fmt.Println(sk)
	return result, err

}

// DeleteAllUser ...
func (s Service) DeleteAllUser(ctx context.Context, sk sentity.Skeleton) error {
	err := s.data.DeleteAllUser(ctx, sk)
	fmt.Println(sk)
	return err

}
