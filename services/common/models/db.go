package models

import (
	"fmt"
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
