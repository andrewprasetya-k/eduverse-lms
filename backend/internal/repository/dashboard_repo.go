package repository

import (
	"time"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	// Student
	GetPendingAssignmentsCount(userID string) (int, error)
	GetUpcomingDeadlines(userID string, limit int) ([]map[string]interface{}, error)
	GetAverageScore(userID string) (float64, error)
	GetMaterialProgress(userID string) (completed int, total int, err error)

	// Teacher
	GetPendingReviewsCount(teacherID string) (int, error)
	GetTotalStudentsByTeacher(teacherID string) (int, error)
	GetSubmissionRateByTeacher(teacherID string) (float64, error)
	GetClassPerformance(teacherID string) ([]map[string]interface{}, error)

	// Admin
	GetSchoolStatistics(schoolID string) (map[string]int, error)
	GetEnrollmentTrends(schoolID string) ([]map[string]interface{}, error)
	GetRecentActivities(schoolID string, limit int) ([]map[string]interface{}, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db: db}
}

// Student Dashboard
func (r *dashboardRepository) GetPendingAssignmentsCount(userID string) (int, error) {
	var count int64
	err := r.db.Table("assignments a").
		Joins("JOIN subject_classes sc ON a.asg_scl_id = sc.scl_id").
		Joins("JOIN enrollments e ON sc.scl_cls_id = e.enr_cls_id").
		Joins("JOIN school_users su ON e.enr_scu_id = su.scu_id").
		Where("su.scu_usr_id = ? AND a.asg_deadline > ? AND a.deleted_at IS NULL", userID, time.Now()).
		Where("NOT EXISTS (SELECT 1 FROM submissions s WHERE s.sbm_asg_id = a.asg_id AND s.sbm_usr_id = ? AND s.deleted_at IS NULL)", userID).
		Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) GetUpcomingDeadlines(userID string, limit int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Raw(`
		SELECT 
			a.asg_id as assignment_id,
			a.asg_title as assignment_title,
			sub.sub_name as subject_name,
			a.asg_deadline as deadline,
			EXISTS(SELECT 1 FROM submissions s WHERE s.sbm_asg_id = a.asg_id AND s.sbm_usr_id = ? AND s.deleted_at IS NULL) as is_submitted
		FROM assignments a
		JOIN subject_classes sc ON a.asg_scl_id = sc.scl_id
		JOIN subjects sub ON sc.scl_sub_id = sub.sub_id
		JOIN enrollments e ON sc.scl_cls_id = e.enr_cls_id
		JOIN school_users su ON e.enr_scu_id = su.scu_id
		WHERE su.scu_usr_id = ? 
			AND a.asg_deadline > ?
			AND a.deleted_at IS NULL
		ORDER BY a.asg_deadline ASC
		LIMIT ?
	`, userID, userID, time.Now(), limit).Scan(&results).Error
	return results, err
}

func (r *dashboardRepository) GetAverageScore(userID string) (float64, error) {
	var avg float64
	err := r.db.Raw(`
		SELECT COALESCE(AVG(asm.asm_score), 0) as average
		FROM assessments asm
		JOIN submissions s ON asm.asm_sbm_id = s.sbm_id
		WHERE s.sbm_usr_id = ? AND s.deleted_at IS NULL
	`, userID).Scan(&avg).Error
	return avg, err
}

func (r *dashboardRepository) GetMaterialProgress(userID string) (completed int, total int, err error) {
	err = r.db.Raw(`
		SELECT 
			COUNT(CASE WHEN mp.map_status = 'completed' THEN 1 END) as completed,
			COUNT(*) as total
		FROM materials m
		JOIN subject_classes sc ON m.mat_scl_id = sc.scl_id
		JOIN enrollments e ON sc.scl_cls_id = e.enr_cls_id
		JOIN school_users su ON e.enr_scu_id = su.scu_id
		LEFT JOIN material_progress mp ON m.mat_id = mp.map_mat_id AND mp.map_usr_id = ?
		WHERE su.scu_usr_id = ? AND m.deleted_at IS NULL
	`, userID, userID).Row().Scan(&completed, &total)
	return
}

// Teacher Dashboard
func (r *dashboardRepository) GetPendingReviewsCount(teacherID string) (int, error) {
	var count int64
	err := r.db.Table("submissions s").
		Joins("JOIN assignments a ON s.sbm_asg_id = a.asg_id").
		Joins("JOIN subject_classes sc ON a.asg_scl_id = sc.scl_id").
		Where("sc.scl_scu_id = ? AND s.deleted_at IS NULL", teacherID).
		Where("NOT EXISTS (SELECT 1 FROM assessments asm WHERE asm.asm_sbm_id = s.sbm_id)").
		Count(&count).Error
	return int(count), err
}

func (r *dashboardRepository) GetTotalStudentsByTeacher(teacherID string) (int, error) {
	var count int64
	err := r.db.Raw(`
		SELECT COUNT(DISTINCT e.enr_scu_id)
		FROM enrollments e
		JOIN subject_classes sc ON e.enr_cls_id = sc.scl_cls_id
		WHERE sc.scl_scu_id = ? AND e.enr_role = 'student'
	`, teacherID).Scan(&count).Error
	return int(count), err
}

func (r *dashboardRepository) GetSubmissionRateByTeacher(teacherID string) (float64, error) {
	var rate float64
	err := r.db.Raw(`
		SELECT 
			CASE 
				WHEN COUNT(DISTINCT a.asg_id) = 0 THEN 0
				ELSE (COUNT(DISTINCT s.sbm_id)::float / COUNT(DISTINCT a.asg_id)) * 100
			END as rate
		FROM assignments a
		JOIN subject_classes sc ON a.asg_scl_id = sc.scl_id
		LEFT JOIN submissions s ON a.asg_id = s.sbm_asg_id AND s.deleted_at IS NULL
		WHERE sc.scl_scu_id = ? AND a.deleted_at IS NULL
	`, teacherID).Scan(&rate).Error
	return rate, err
}

func (r *dashboardRepository) GetClassPerformance(teacherID string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Raw(`
		SELECT 
			c.cls_id as class_id,
			c.cls_title as class_name,
			sub.sub_name as subject_name,
			COALESCE(AVG(asm.asm_score), 0) as average_score,
			COUNT(DISTINCT e.enr_scu_id) as total_students,
			CASE 
				WHEN COUNT(DISTINCT a.asg_id) = 0 THEN 0
				ELSE (COUNT(DISTINCT s.sbm_id)::float / COUNT(DISTINCT a.asg_id)) * 100
			END as submission_rate
		FROM subject_classes sc
		JOIN classes c ON sc.scl_cls_id = c.cls_id
		JOIN subjects sub ON sc.scl_sub_id = sub.sub_id
		LEFT JOIN enrollments e ON c.cls_id = e.enr_cls_id AND e.enr_role = 'student'
		LEFT JOIN assignments a ON sc.scl_id = a.asg_scl_id AND a.deleted_at IS NULL
		LEFT JOIN submissions s ON a.asg_id = s.sbm_asg_id AND s.deleted_at IS NULL
		LEFT JOIN assessments asm ON s.sbm_id = asm.asm_sbm_id
		WHERE sc.scl_scu_id = ? AND c.deleted_at IS NULL
		GROUP BY c.cls_id, c.cls_title, sub.sub_name
	`, teacherID).Scan(&results).Error
	return results, err
}

// Admin Dashboard
func (r *dashboardRepository) GetSchoolStatistics(schoolID string) (map[string]int, error) {
	stats := make(map[string]int)
	
	var totalStudents, totalTeachers, totalClasses, activeClasses int64
	
	r.db.Table("school_users su").
		Joins("JOIN enrollments e ON su.scu_id = e.enr_scu_id").
		Where("su.scu_sch_id = ? AND e.enr_role = 'student'", schoolID).
		Count(&totalStudents)
	
	r.db.Table("school_users su").
		Joins("JOIN subject_classes sc ON su.scu_id = sc.scl_scu_id").
		Where("su.scu_sch_id = ?", schoolID).
		Count(&totalTeachers)
	
	r.db.Table("classes").
		Where("cls_sch_id = ? AND deleted_at IS NULL", schoolID).
		Count(&totalClasses)
	
	r.db.Table("classes").
		Where("cls_sch_id = ? AND is_active = true AND deleted_at IS NULL", schoolID).
		Count(&activeClasses)
	
	stats["totalStudents"] = int(totalStudents)
	stats["totalTeachers"] = int(totalTeachers)
	stats["totalClasses"] = int(totalClasses)
	stats["activeClasses"] = int(activeClasses)
	
	return stats, nil
}

func (r *dashboardRepository) GetEnrollmentTrends(schoolID string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Raw(`
		SELECT 
			c.cls_title as class_name,
			COUNT(DISTINCT e.enr_scu_id) as total_enrolled,
			COUNT(DISTINCT CASE WHEN e.enr_role = 'teacher' THEN e.enr_scu_id END) as teachers,
			COUNT(DISTINCT CASE WHEN e.enr_role = 'student' THEN e.enr_scu_id END) as students
		FROM classes c
		LEFT JOIN enrollments e ON c.cls_id = e.enr_cls_id
		WHERE c.cls_sch_id = ? AND c.deleted_at IS NULL
		GROUP BY c.cls_id, c.cls_title
		ORDER BY total_enrolled DESC
	`, schoolID).Scan(&results).Error
	return results, err
}

func (r *dashboardRepository) GetRecentActivities(schoolID string, limit int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.Raw(`
		SELECT 
			u.usr_nama_lengkap as user_name,
			l.log_action as action,
			l.created_at as timestamp
		FROM logs l
		JOIN users u ON l.log_usr_id = u.usr_id
		WHERE l.log_sch_id = ?
		ORDER BY l.created_at DESC
		LIMIT ?
	`, schoolID, limit).Scan(&results).Error
	return results, err
}
