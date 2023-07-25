package admin

import (
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/admin/structs"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"gorm.io/gorm"
	"time"
)

type adminRepository struct {
	db *gorm.DB
}

func (a *adminRepository) InsertAccessToken(tx *gorm.DB, adminID uint, expiredAt time.Time) (result *entity.AdminAccessToken, err error) {
	result = &entity.AdminAccessToken{
		AdminID:   adminID,
		ExpiredAt: expiredAt,
	}

	if err = tx.Table(result.TableName()).Create(result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) DeleteAccessToken(tx *gorm.DB, id uint) (err error) {
	var _accessToken = &entity.AdminAccessToken{ID: id}

	if err = tx.Table(_accessToken.TableName()).Where("id = ?", id).Delete(_accessToken).Error; err != nil {
		return err
	}

	return nil
}

func (a *adminRepository) InsertApiLog(tx *gorm.DB, log *entity.AdminApiLog) (*entity.AdminApiLog, error) {

	if err := tx.Table(log.TableName()).Create(log).Error; err != nil {
		return nil, err
	}

	return log, nil
}

func (a *adminRepository) FindAllApiLog(tx *gorm.DB, limit, offset int) (result []*entity.AdminApiLog, err error) {
	var _apiLog *entity.AdminApiLog
	if err = tx.Table(_apiLog.TableName()).Offset(offset).Limit(limit).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindApiLogByAdminID(tx *gorm.DB, adminID uint, limit, offset int) (result []*entity.AdminApiLog, err error) {
	var _table *entity.AdminApiLog

	if err = tx.Table(_table.TableName()).Where("admin_id = ?", adminID).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindApiLogByTimeCond(tx *gorm.DB, start, end time.Time, limit, offset int) (result []*entity.AdminApiLog, err error) {
	var _apiLog *entity.AdminApiLog

	if err = tx.Table(_apiLog.TableName()).Where("created_at BETWEEN ? AND ?", start, end).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindApiLogByMethodCond(tx *gorm.DB, method entity.ApiMethod, limit, offset int) ([]*entity.AdminApiLog, error) {
	var apiLogs []*entity.AdminApiLog
	var _apiLog = &entity.AdminApiLog{Method: method}

	if err := tx.Table(_apiLog.TableName()).Where(_apiLog).Limit(limit).Offset(offset).Find(&apiLogs).Error; err != nil {
		return nil, err
	}
	return apiLogs, nil
}

func (a *adminRepository) FindApiLogByEndpoint(tx *gorm.DB, endpoint string, limit, offset int) ([]*entity.AdminApiLog, error) {
	var apiLogs []*entity.AdminApiLog
	var _apiLog = &entity.AdminApiLog{Endpoint: endpoint}

	if err := tx.Table(_apiLog.TableName()).Where(_apiLog).Limit(limit).Offset(offset).Find(&apiLogs).Error; err != nil {
		return nil, err
	}

	return apiLogs, nil
}

func (a *adminRepository) SearchApiLog(tx *gorm.DB, input *structs.SearchApiLogInput) ([]*entity.AdminApiLog, error) {
	var _apiLog *entity.AdminApiLog
	var result []*entity.AdminApiLog

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

func (a *adminRepository) SearchDeleteLog(tx *gorm.DB, input *structs.SearchDeleteLogInput) ([]*entity.AdminDeleteLog, error) {
	var _deleteLog *entity.AdminDeleteLog
	var result []*entity.AdminDeleteLog

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

func (a *adminRepository) FindAllDeleteLog(tx *gorm.DB, limit, offset int) ([]*entity.AdminDeleteLog, error) {
	var _deleteLog = &entity.AdminDeleteLog{}
	var result []*entity.AdminDeleteLog

	if err := tx.Table(_deleteLog.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindDeleteLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*entity.AdminDeleteLog, error) {
	var _deleteLog = &entity.AdminDeleteLog{ExecutorID: executor}
	var result []*entity.AdminDeleteLog

	if err := tx.Table(_deleteLog.TableName()).Where(_deleteLog).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindDeleteLogOfTarget(tx *gorm.DB, target uint) (*entity.AdminDeleteLog, error) {
	var _deleteLog = &entity.AdminDeleteLog{TargetID: target}

	if err := tx.Table(_deleteLog.TableName()).Where(_deleteLog).First(&_deleteLog).Error; err != nil {
		return nil, err
	}

	return _deleteLog, nil
}

func (a *adminRepository) InsertDeleteLog(tx *gorm.DB, executor, target uint) (*entity.AdminDeleteLog, error) {
	var _deleteLog = &entity.AdminDeleteLog{ExecutorID: executor, TargetID: target}

	if err := tx.Table(_deleteLog.TableName()).Create(_deleteLog).Error; err != nil {
		return nil, err
	}

	return _deleteLog, nil
}

func (a *adminRepository) SearchGradeUpdateLog(tx *gorm.DB, input *structs.SearchGradeUpdateLogInput) ([]*entity.AdminGradeUpdateLog, error) {
	var _gradeUpdateLog *entity.AdminGradeUpdateLog
	var result []*entity.AdminGradeUpdateLog
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

func (a *adminRepository) InsertGradeUpdateLog(tx *gorm.DB, executor, target uint, prev, curr entity.GradeType) (*entity.AdminGradeUpdateLog, error) {
	var _gradeUpdateLog = &entity.AdminGradeUpdateLog{
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

func (a *adminRepository) FindAllGradeUpdateLog(tx *gorm.DB, limit, offset int) ([]*entity.AdminGradeUpdateLog, error) {
	var _gradeUpdateLog *entity.AdminGradeUpdateLog
	var result []*entity.AdminGradeUpdateLog

	if err := tx.Table(_gradeUpdateLog.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindGradeUpdateLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*entity.AdminGradeUpdateLog, error) {
	var _gradeUpdateLog *entity.AdminGradeUpdateLog
	var result []*entity.AdminGradeUpdateLog
	if err := tx.Table(_gradeUpdateLog.TableName()).Where("executor_id = ?", executor).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindGradeUpdateLogOfTarget(tx *gorm.DB, target uint, limit, offset int) ([]*entity.AdminGradeUpdateLog, error) {
	var _gradeUpdateLog *entity.AdminGradeUpdateLog
	var result []*entity.AdminGradeUpdateLog

	if err := tx.Table(_gradeUpdateLog.TableName()).Where("target_id = ?", target).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindLoginFailedLogByAdminID(tx *gorm.DB, adminID uint, limit, offset int) ([]*entity.AdminLoginFailedLog, error) {
	var _loginFailedLog = &entity.AdminLoginFailedLog{AdminID: adminID}
	var result []*entity.AdminLoginFailedLog

	if err := tx.Table(_loginFailedLog.TableName()).Where(_loginFailedLog).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) InsertLoginFailedLog(tx *gorm.DB, adminID uint) (*entity.AdminLoginFailedLog, error) {
	var _loginFailedLog = &entity.AdminLoginFailedLog{AdminID: adminID}

	if err := tx.Table(_loginFailedLog.TableName()).Create(_loginFailedLog).Error; err != nil {
		return nil, err
	}

	return _loginFailedLog, nil
}

func (a *adminRepository) FindLoginHistoryByAdminID(tx *gorm.DB, adminID uint, limit, offset int) ([]*entity.AdminLoginHistory, error) {
	var _loginHistory = &entity.AdminLoginHistory{AdminID: adminID}
	var result []*entity.AdminLoginHistory

	if err := tx.Table(_loginHistory.TableName()).Where(_loginHistory).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) InsertLoginHistory(tx *gorm.DB, adminID uint) (*entity.AdminLoginHistory, error) {
	var _loginHistory = &entity.AdminLoginHistory{AdminID: adminID}

	if err := tx.Table(_loginHistory.TableName()).Create(_loginHistory).Error; err != nil {
		return nil, err
	}

	return _loginHistory, nil
}

func (a *adminRepository) SearchPasswordUpdateLog(tx *gorm.DB, input *structs.SearchPasswordUpdateLogInput) ([]*entity.AdminPasswordLog, error) {
	var _passwordLog *entity.AdminPasswordLog
	var result []*entity.AdminPasswordLog

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

func (a *adminRepository) FindAllPasswordUpdateLog(tx *gorm.DB, limit, offset int) ([]*entity.AdminPasswordLog, error) {
	var _passwordLog = &entity.AdminPasswordLog{}
	var result []*entity.AdminPasswordLog

	if err := tx.Table(_passwordLog.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindPasswordUpdateLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*entity.AdminPasswordLog, error) {
	var _passwordLog = &entity.AdminPasswordLog{ExecutorID: executor}
	var result []*entity.AdminPasswordLog

	if err := tx.Table(_passwordLog.TableName()).Where(_passwordLog).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindPasswordUpdateLogOfTarget(tx *gorm.DB, target uint, limit, offset int) ([]*entity.AdminPasswordLog, error) {
	var _passwordLog = &entity.AdminPasswordLog{TargetID: target}
	var result []*entity.AdminPasswordLog

	if err := tx.Table(_passwordLog.TableName()).Where(_passwordLog).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) InsertPasswordUpdateLog(tx *gorm.DB, executor, target uint) (*entity.AdminPasswordLog, error) {
	var _passwordLog = &entity.AdminPasswordLog{TargetID: target, ExecutorID: executor}

	if err := tx.Table(_passwordLog.TableName()).Create(_passwordLog).Error; err != nil {
		return nil, err
	}

	return _passwordLog, nil
}

func (a *adminRepository) FindAdminByID(tx *gorm.DB, id uint) (*entity.Admin, error) {
	var result *entity.Admin
	var err error

	if err = tx.Table(result.TableName()).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindAdminByEmail(tx *gorm.DB, email string) (*entity.Admin, error) {
	var result *entity.Admin
	var err error

	if err = tx.Table(result.TableName()).Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindAdminCredentialsByID(tx *gorm.DB, id uint) (*entity.Admin, error) {
	var _admin *entity.Admin
	if err := tx.Table(_admin.TableName()).Select("password").Where("id = ?", id).First(&_admin).Error; err != nil {
		return nil, err
	}
	return _admin, nil
}

func (a *adminRepository) FindAccessToken(tx *gorm.DB, limit, offset int) ([]*entity.AdminAccessToken, error) {
	var _accessToken *entity.AdminAccessToken
	var result []*entity.AdminAccessToken

	if err := tx.Table(_accessToken.TableName()).Where("expired_at >= current_timestamp()").Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindAdminByGrade(tx *gorm.DB, grade entity.GradeType, limit, offset int) ([]*entity.Admin, error) {
	var result []*entity.Admin
	var _admin *entity.Admin
	var err error

	if err = tx.Table(_admin.TableName()).Where("grade = ?", grade).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindAllAdmin(tx *gorm.DB, limit, offset int) ([]*entity.Admin, error) {
	var result []*entity.Admin
	var _admin *entity.Admin
	var err error

	if err = tx.Table(_admin.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) InsertAdmin(tx *gorm.DB, email, password string) (*entity.Admin, error) {
	var err error
	var input *entity.Admin

	input = &entity.Admin{
		Email:    email,
		Password: password,
		Grade:    entity.SUPPORT,
	}

	if err = tx.Table(input.TableName()).Create(input).Error; err != nil {
		return nil, err
	}

	return input, nil
}

func (a *adminRepository) UpdateAdminByID(tx *gorm.DB, id uint, admin *entity.Admin) error {
	var err error
	if err = tx.Table(admin.TableName()).Where("id = ?", id).Updates(admin).Error; err != nil {
		return err
	}

	return nil
}

func (a *adminRepository) DeleteAdminByID(tx *gorm.DB, id uint) error {
	var err error
	var _admin = &entity.Admin{ID: id}
	if err = tx.Table(_admin.TableName()).Where("id = ?", id).Delete(_admin).Error; err != nil {
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
