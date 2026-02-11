package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type SubjectRepository interface{
	CreateSubject(subject *domain.Subject) error
	GetAllSubjects(schoolID string) ([]*domain.Subject, error) // Tambahkan schoolID
	convertSchoolCodeToSchoolID(subjectID string) (string, error)
	GetSubjectByCode(subjectCode string, schoolID string) (*domain.Subject, error)
	UpdateSubject(subject *domain.Subject) error
	DeleteSubject(subjectID string) error
}

type subjectRepository struct {
	db *gorm.DB
}

//constructor
func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{db: db}
}

func (r *subjectRepository) CreateSubject(subject *domain.Subject) error {
	return r.db.Create(subject).Error
}

func (r *subjectRepository) GetAllSubjects(schoolCode string) ([]*domain.Subject, error) {
    var subjects []*domain.Subject

    // 1. Cara panggil fungsi internal: pakai 'r', bukan nama struct-nya.
    // 2. Ingat: fungsi itu sekarang balikin DUA nilai (ID dan Error).
    schoolID, err := r.convertSchoolCodeToSchoolID(schoolCode)
    
    // 3. WAJIB cek error dari hasil 'convert' tadi. 
    // Kalau sekolahnya gak ada, ya gak usah lanjut cari subjects.
    if err != nil {
        return nil, err
    }

    // 4. Baru deh cari Subjects pakai ID yang sudah valid.
    err = r.db.Where("sub_sch_id = ?", schoolID).Find(&subjects).Error
    
    return subjects, err
}

// 1. Tambahkan (string, error) agar bisa mengembalikan ID DAN pesan error
func (r *subjectRepository) convertSchoolCodeToSchoolID(schoolCode string) (string, error) {
    var schoolID string

    // 2. Gunakan .Scan() jika ingin mengambil satu kolom saja ke variabel biasa
    err := r.db.Model(&domain.School{}).
        Select("sch_id").
        Where("sch_code = ?", schoolCode).
        Scan(&schoolID).Error // Scan lebih cocok untuk single variable daripada First

    if err != nil {
        return "", err
    }

    // 3. Kembalikan ID-nya dan nil (tanda tidak ada error)
    return schoolID, nil
}

func (r *subjectRepository) GetSubjectByCode(subjectCode string, schoolCode string) (*domain.Subject, error) {
	var subject domain.Subject
	schoolID, err := r.convertSchoolCodeToSchoolID(schoolCode)
	if err != nil {
		return nil, err
	}

	err = r.db.Where("sub_sch_id = ? and sub_code = ? ", schoolID, subjectCode).First(&subject).Error
	return &subject, err
}

func (r *subjectRepository) UpdateSubject(subject *domain.Subject) error {
	return r.db.Updates(subject).Error
}

func (r *subjectRepository) DeleteSubject(id string) error {
	return r.db.Delete(&domain.Subject{}, "sub_id = ?", id).Error
}