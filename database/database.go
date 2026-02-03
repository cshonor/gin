package database

import (
	"fmt"
	"log"

	"gin/config"
	"gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase(cfg *config.Config) error {
	var err error
	var dialector gorm.Dialector

	// 根据配置选择数据库驱动
	if cfg.DatabaseURL != "" {
		// 如果提供了数据库 URL，尝试解析
		if isPostgresURL(cfg.DatabaseURL) {
			dialector = postgres.Open(cfg.DatabaseURL)
		} else if isMySQLURL(cfg.DatabaseURL) {
			dialector = mysql.Open(cfg.DatabaseURL)
		} else {
			// 默认使用 SQLite
			dialector = sqlite.Open(cfg.DatabaseURL)
		}
	} else {
		// 默认使用 SQLite（开发环境）
		dialector = sqlite.Open("gin.db")
	}

	// 配置 GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	if cfg.Environment == "production" {
		gormConfig.Logger = logger.Default.LogMode(logger.Error)
	}

	// 连接数据库
	DB, err = gorm.Open(dialector, gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	// 自动迁移
	err = AutoMigrate()
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("数据库连接成功")
	return nil
}

// AutoMigrate 自动迁移数据库表
func AutoMigrate() error {
	return DB.AutoMigrate(
		&models.User{},
		// 在这里添加其他模型
	)
}

// CloseDatabase 关闭数据库连接
func CloseDatabase() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// isPostgresURL 检查是否为 PostgreSQL URL
func isPostgresURL(url string) bool {
	if len(url) >= 11 {
		return url[:10] == "postgres://" || url[:11] == "postgresql://"
	}
	return false
}

// isMySQLURL 检查是否为 MySQL URL
func isMySQLURL(url string) bool {
	return len(url) >= 8 && url[:8] == "mysql://"
}

