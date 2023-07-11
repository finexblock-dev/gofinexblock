package admin

import (
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/admin/dto"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
	"gorm.io/gorm"
	"time"
)

var table = &admin.Admin{}

type adminRepository struct {
	db *gorm.DB
}

func (a *adminRepository) InsertAccessToken(tx *gorm.DB, adminID uint, expiredAt time.Time) (*admin.AdminAccessToken, error) {
	var err error

	var _accessToken = &admin.AdminAccessToken{
		AdminID:   adminID,
		ExpiredAt: expiredAt,
	}

	if err = tx.Table(_accessToken.TableName()).Create(_accessToken).Error; err != nil {
		return nil, err
	}

	return _accessToken, nil
}

func (a *adminRepository) DeleteAccessToken(tx *gorm.DB, id uint) error {
	var _accessToken = &admin.AdminAccessToken{ID: id}

	if err := tx.Table(_accessToken.TableName()).Where("id = ?", id).Delete(_accessToken).Error; err != nil {
		return err
	}

	return nil
}

func (a *adminRepository) InsertApiLog(tx *gorm.DB, log *admin.AdminApiLog) (*admin.AdminApiLog, error) {

	if err := tx.Table(log.TableName()).Create(log).Error; err != nil {
		return nil, err
	}

	return log, nil
}

func (a *adminRepository) FindAllApiLog(tx *gorm.DB, limit, offset int) ([]*admin.AdminApiLog, error) {
	var _apiLog *admin.AdminApiLog
	var result []*admin.AdminApiLog

	if err := tx.Table(_apiLog.TableName()).Offset(offset).Limit(limit).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindApiLogByAdminID(tx *gorm.DB, adminID uint, limit, offset int) ([]*admin.AdminApiLog, error) {
	var _apiLog *admin.AdminApiLog
	var result []*admin.AdminApiLog

	if err := tx.Table(_apiLog.TableName()).Where("admin_id = ?", adminID).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindApiLogByTimeCond(tx *gorm.DB, start, end time.Time, limit, offset int) ([]*admin.AdminApiLog, error) {
	var apiLogs []*admin.AdminApiLog
	var _apiLog *admin.AdminApiLog

	if err := tx.Table(_apiLog.TableName()).Where("created_at BETWEEN ? AND ?", start, end).Limit(limit).Offset(offset).Find(&apiLogs).Error; err != nil {
		return nil, err
	}

	return apiLogs, nil
}

func (a *adminRepository) FindApiLogByMethodCond(tx *gorm.DB, method admin.ApiMethod, limit, offset int) ([]*admin.AdminApiLog, error) {
	var apiLogs []*admin.AdminApiLog
	var _apiLog = &admin.AdminApiLog{Method: method}

	if err := tx.Table(_apiLog.TableName()).Where(_apiLog).Limit(limit).Offset(offset).Find(&apiLogs).Error; err != nil {
		return nil, err
	}
	return apiLogs, nil
}

func (a *adminRepository) FindApiLogByEndpoint(tx *gorm.DB, endpoint string, limit, offset int) ([]*admin.AdminApiLog, error) {
	var apiLogs []*admin.AdminApiLog
	var _apiLog = &admin.AdminApiLog{Endpoint: endpoint}

	if err := tx.Table(_apiLog.TableName()).Where(_apiLog).Limit(limit).Offset(offset).Find(&apiLogs).Error; err != nil {
		return nil, err
	}

	return apiLogs, nil
}

func (a *adminRepository) SearchApiLog(tx *gorm.DB, input *dto.SearchApiLogInput) ([]*admin.AdminApiLog, error) {
	var _apiLog *admin.AdminApiLog
	var result []*admin.AdminApiLog

	query := tx.Table(_apiLog.TableName())

	if input.AdminID != 0 {
		query.Where("admin_id = ?", input.AdminID)
	}
	if input.Method != "" {
		query.Where("method = ?", input.Method)
	}
	if input.Endpoint != "" {
		query.Where("endpoint = ?", input.Endpoint)
	}
	if input.StartTime != "" {
		query.Where("created_at >= ?", input.StartTime)
	}
	if input.EndTime != "" {
		query.Where("created_at <= ?", input.EndTime)
	}
	if input.Limit > 0 {
		query.Limit(input.Limit)
	}
	if input.Offset > 0 {
		query.Offset(input.Offset)
	}

	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) SearchDeleteLog(tx *gorm.DB, input *dto.SearchDeleteLogInput) ([]*admin.AdminDeleteLog, error) {
	var _deleteLog *admin.AdminDeleteLog
	var result []*admin.AdminDeleteLog

	query := tx.Table(_deleteLog.TableName())

	if input.Target != 0 {
		query.Where("target_id = ?", input.Target)
	}
	if input.Executor != 0 {
		query.Where("executor_id = ?", input.Executor)
	}
	if input.StartTime != "" {
		query.Where("created_at >= ?", input.StartTime)
	}
	if input.EndTime != "" {
		query.Where("created_at <= ?", input.EndTime)
	}
	if input.Limit > 0 {
		query.Limit(input.Limit)
	}
	if input.Offset > 0 {
		query.Offset(input.Offset)
	}

	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindAllDeleteLog(tx *gorm.DB, limit, offset int) ([]*admin.AdminDeleteLog, error) {
	var _deleteLog = &admin.AdminDeleteLog{}
	var result []*admin.AdminDeleteLog

	if err := tx.Table(_deleteLog.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindDeleteLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*admin.AdminDeleteLog, error) {
	var _deleteLog = &admin.AdminDeleteLog{ExecutorID: executor}
	var result []*admin.AdminDeleteLog

	if err := tx.Table(_deleteLog.TableName()).Where(_deleteLog).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindDeleteLogOfTarget(tx *gorm.DB, target uint) (*admin.AdminDeleteLog, error) {
	var _deleteLog = &admin.AdminDeleteLog{TargetID: target}

	if err := tx.Table(_deleteLog.TableName()).Where(_deleteLog).First(&_deleteLog).Error; err != nil {
		return nil, err
	}

	return _deleteLog, nil
}

func (a *adminRepository) InsertDeleteLog(tx *gorm.DB, executor, target uint) (*admin.AdminDeleteLog, error) {
	var _deleteLog = &admin.AdminDeleteLog{ExecutorID: executor, TargetID: target}

	if err := tx.Table(_deleteLog.TableName()).Create(_deleteLog).Error; err != nil {
		return nil, err
	}

	return _deleteLog, nil
}

func (a *adminRepository) SearchGradeUpdateLog(tx *gorm.DB, input *dto.SearchGradeUpdateLogInput) ([]*admin.AdminGradeUpdateLog, error) {
	var _gradeUpdateLog *admin.AdminGradeUpdateLog
	var result []*admin.AdminGradeUpdateLog
	var err error

	query := tx.Table(_gradeUpdateLog.TableName())

	if input.Target != 0 {
		query.Where("target_id = ?", input.Target)
	}

	if input.Executor != 0 {
		query.Where("executor_id = ?", input.Executor)
	}

	if input.StartTime != "" {
		query.Where("created_at >= ?", input.StartTime)
	}

	if input.EndTime != "" {
		query.Where("created_at <= ?", input.EndTime)
	}

	if input.Limit != 0 {
		query.Limit(input.Limit)
	}

	if input.Offset != 0 {
		query.Offset(input.Offset)
	}

	if err = query.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) InsertGradeUpdateLog(tx *gorm.DB, executor, target uint, prev, curr string) (*admin.AdminGradeUpdateLog, error) {
	var _gradeUpdateLog = &admin.AdminGradeUpdateLog{
		ExecutorID: executor,
		TargetID:   target,
		PrevGrade:  prev,
		CurrGrade:  curr,
	}

	if err := tx.Table(_gradeUpdateLog.TableName()).Create(_gradeUpdateLog).Error; err != nil {
		return nil, err
	}

	return _gradeUpdateLog, nil
}

func (a *adminRepository) FindAllGradeUpdateLog(tx *gorm.DB, limit, offset int) ([]*admin.AdminGradeUpdateLog, error) {
	var _gradeUpdateLog *admin.AdminGradeUpdateLog
	var result []*admin.AdminGradeUpdateLog

	if err := tx.Table(_gradeUpdateLog.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindGradeUpdateLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*admin.AdminGradeUpdateLog, error) {
	var _gradeUpdateLog *admin.AdminGradeUpdateLog
	var result []*admin.AdminGradeUpdateLog
	if err := tx.Table(_gradeUpdateLog.TableName()).Where("executor_id = ?", executor).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindGradeUpdateLogOfTarget(tx *gorm.DB, target uint, limit, offset int) ([]*admin.AdminGradeUpdateLog, error) {
	var _gradeUpdateLog *admin.AdminGradeUpdateLog
	var result []*admin.AdminGradeUpdateLog

	if err := tx.Table(_gradeUpdateLog.TableName()).Where("target_id = ?", target).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindLoginFailedLogByAdminID(tx *gorm.DB, adminID uint, limit, offset int) ([]*admin.AdminLoginFailedLog, error) {
	var _loginFailedLog = &admin.AdminLoginFailedLog{AdminID: adminID}
	var result []*admin.AdminLoginFailedLog

	if err := tx.Table(_loginFailedLog.TableName()).Where(_loginFailedLog).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) InsertLoginFailedLog(tx *gorm.DB, adminID uint) (*admin.AdminLoginFailedLog, error) {
	var _loginFailedLog = &admin.AdminLoginFailedLog{AdminID: adminID}

	if err := tx.Table(_loginFailedLog.TableName()).Create(_loginFailedLog).Error; err != nil {
		return nil, err
	}

	return _loginFailedLog, nil
}

func (a *adminRepository) FindLoginHistoryByAdminID(tx *gorm.DB, adminID uint, limit, offset int) ([]*admin.AdminLoginHistory, error) {
	var _loginHistory = &admin.AdminLoginHistory{AdminID: adminID}
	var result []*admin.AdminLoginHistory

	if err := tx.Table(_loginHistory.TableName()).Where(_loginHistory).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) InsertLoginHistory(tx *gorm.DB, adminID uint) (*admin.AdminLoginHistory, error) {
	var _loginHistory = &admin.AdminLoginHistory{AdminID: adminID}

	if err := tx.Table(_loginHistory.TableName()).Create(_loginHistory).Error; err != nil {
		return nil, err
	}

	return _loginHistory, nil
}

func (a *adminRepository) SearchPasswordUpdateLog(tx *gorm.DB, input *dto.SearchPasswordUpdateLogInput) ([]*admin.AdminPasswordLog, error) {
	var _passwordLog *admin.AdminPasswordLog
	var result []*admin.AdminPasswordLog

	query := tx.Table(_passwordLog.TableName())

	if input.Target != 0 {
		query.Where("target_id = ?", input.Target)
	}
	if input.Executor != 0 {
		query.Where("executor_id = ?", input.Executor)
	}
	if input.StartTime != "" {
		query.Where("created_at >= ?", input.StartTime)
	}
	if input.EndTime != "" {
		query.Where("created_at <= ?", input.EndTime)
	}
	if input.Limit > 0 {
		query.Limit(input.Limit)
	}
	if input.Offset > 0 {
		query.Offset(input.Offset)
	}

	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindAllPasswordUpdateLog(tx *gorm.DB, limit, offset int) ([]*admin.AdminPasswordLog, error) {
	var _passwordLog = &admin.AdminPasswordLog{}
	var result []*admin.AdminPasswordLog

	if err := tx.Table(_passwordLog.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindPasswordUpdateLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*admin.AdminPasswordLog, error) {
	var _passwordLog = &admin.AdminPasswordLog{ExecutorID: executor}
	var result []*admin.AdminPasswordLog

	if err := tx.Table(_passwordLog.TableName()).Where(_passwordLog).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindPasswordUpdateLogOfTarget(tx *gorm.DB, target uint, limit, offset int) ([]*admin.AdminPasswordLog, error) {
	var _passwordLog = &admin.AdminPasswordLog{TargetID: target}
	var result []*admin.AdminPasswordLog

	if err := tx.Table(_passwordLog.TableName()).Where(_passwordLog).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) InsertPasswordUpdateLog(tx *gorm.DB, executor, target uint) (*admin.AdminPasswordLog, error) {
	var _passwordLog = &admin.AdminPasswordLog{TargetID: target, ExecutorID: executor}

	if err := tx.Table(_passwordLog.TableName()).Create(_passwordLog).Error; err != nil {
		return nil, err
	}

	return _passwordLog, nil
}

func (a *adminRepository) FindAdminByID(tx *gorm.DB, id uint) (*admin.Admin, error) {
	var result *admin.Admin
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindAdminByEmail(tx *gorm.DB, email string) (*admin.Admin, error) {
	var result *admin.Admin
	var err error

	if err = tx.Table(table.TableName()).Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindAdminCredentialsByID(tx *gorm.DB, id uint) (*admin.Admin, error) {
	var _admin *admin.Admin
	if err := tx.Table(_admin.TableName()).Select("password").Where("id = ?", id).First(&_admin).Error; err != nil {
		return nil, err
	}
	return _admin, nil
}

func (a *adminRepository) FindAccessToken(tx *gorm.DB, limit, offset int) ([]*admin.AdminAccessToken, error) {
	var _accessToken *admin.AdminAccessToken
	var result []*admin.AdminAccessToken

	if err := tx.Table(_accessToken.TableName()).Where("expired_at >= current_timestamp()").Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindAdminByGrade(tx *gorm.DB, grade admin.GradeType, limit, offset int) ([]*admin.Admin, error) {
	var result []*admin.Admin
	var err error

	if err = tx.Table(table.TableName()).Where("grade = ?", grade).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindAllAdmin(tx *gorm.DB, limit, offset int) ([]*admin.Admin, error) {
	var result []*admin.Admin
	var err error

	if err = tx.Table(table.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) InsertAdmin(tx *gorm.DB, email, password string) (*admin.Admin, error) {
	var err error
	var input *admin.Admin

	input = &admin.Admin{
		Email:    email,
		Password: password,
	}

	if err = tx.Table(table.TableName()).Create(input).Error; err != nil {
		return nil, err
	}

	return input, nil
}

func (a *adminRepository) UpdateAdminByID(tx *gorm.DB, id uint, admin *admin.Admin) error {
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Updates(admin).Error; err != nil {
		return err
	}

	return nil
}

func (a *adminRepository) DeleteAdminByID(tx *gorm.DB, id uint) error {
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Delete(&admin.Admin{ID: id}).Error; err != nil {
		return err
	}

	return nil
}

func newAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db: db}
}

func (a *adminRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.db.Begin(&sql.TxOptions{Isolation: level})
}

func (a *adminRepository) Conn() *gorm.DB {
	return a.db
}