package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func Migrate(db *gorm.DB) error {
	return db.Migrator().AutoMigrate(
		&Agent{},
		&Channel{},
		&Template{},
		&SendBatch{},
		&SendRecord{},
	)
}

type Config struct {
	DBType       string `json:",default=mysql"`     // 数据库类型: "mysql", "postgres""
	Username     string `json:",default=root"`      // 数据库用户名
	Password     string `json:",default=123456"`    // 数据库密码
	Host         string `json:",default=127.0.0.1"` // 数据库主机
	Port         int    `json:",default=3306"`      // 数据库端口
	Database     string `json:",default=msgbox"`    // 数据库名
	SSLMode      string `json:",default=disable"`   // PostgreSQL 专用 SSL 模式 ("disable", "require", 等)
	MaxIdleConns int    `json:",default=10"`        // 连接池最大空闲连接数
	MaxOpenConns int    `json:",default=100"`       // 连接池最大打开连接数
}

func Connect(cfg Config, plugin ...gorm.Plugin) (*gorm.DB, error) {
	var dialector gorm.Dialector
	// 根据数据库类型选择对应的 GORM Dialector
	switch cfg.DBType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
		dialector = mysql.Open(dsn)
	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.Database, cfg.Password, cfg.SSLMode)
		dialector = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.DBType)
	}
	// 使用 GORM 打开数据库连接，并配置额外选项
	db, err := gorm.Open(dialector, &gorm.Config{
		QueryFields:                              true,  // 查询时包括未映射的字段
		SkipDefaultTransaction:                   false, // 不跳过默认事务
		DisableForeignKeyConstraintWhenMigrating: true,  // 自动迁移时不创建外键约束
		PrepareStmt:                              true,  // 预编译 SQL 提高性能
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名使用单数，不加 s
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	for _, p := range plugin {
		_ = db.Use(p)
	}
	// 获取底层 *sql.DB 对象，用于连接池配置
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	// 设置连接池参数，可通过 Config 配置
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	// 最大打开连接数
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, nil
}

func MapToDataTypesJSON(v any) datatypes.JSON {
	if v == nil {
		return datatypes.JSON("{}")
	}
	b, err := json.Marshal(v)
	if err != nil || string(b) == "null" {
		return datatypes.JSON("{}")
	}
	return datatypes.JSON(b)
}

func DataTypesToMap(j datatypes.JSON) map[string]interface{} {
	if len(j) == 0 {
		return map[string]interface{}{} // 空 JSON 返回空 map
	}
	var m map[string]interface{}
	if err := json.Unmarshal(j, &m); err != nil {
		return map[string]interface{}{}
	}
	return m
}

type Pagination[T any] struct {
	DB *gorm.DB
}

func NewPagination[T any](db *gorm.DB) *Pagination[T] {
	return &Pagination[T]{DB: db}
}

func (p *Pagination[T]) QueryPage(page, pageSize int, query func(tx *gorm.DB) *gorm.DB) (total int64, data []T, err error) {
	// normalize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	const maxPageSize = 100
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	// Safe model inference
	var entity T
	tx := p.DB.Model(&entity)

	if query != nil {
		tx = query(tx)
	}

	// Count safely
	countTx := tx.Session(&gorm.Session{})
	if err = countTx.Count(&total).Error; err != nil {
		return
	}

	if total == 0 {
		return total, []T{}, nil
	}

	offset := (page - 1) * pageSize

	// Query data
	err = tx.Limit(pageSize).Offset(offset).Find(&data).Error
	return total, data, err
}
