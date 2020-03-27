package skeleton

import (
	"time"
)

// Skeleton model
type Skeleton struct {
	SkeletonID           int       `db:"id" json:"s_id"`
	SkeletonNip          string    `db:"nip" json:"s_nip"`
	SkeletonNama         string    `db:"nama_lengkap" json:"s_nama"`
	SkeletonTanggalLahir time.Time `db:"tanggal_lahir" json:"s_tanggalLahir"`
	SkeletonJabatan      string    `db:"jabatan" json:"s_jabatan"`
	SkeletonEmail        string    `db:"email" json:"s_email"`
}
