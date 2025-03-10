package database

import (
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"
)

// Migration 数据库迁移记录
type Migration struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"uniqueIndex"`
	AppliedAt time.Time
}

// Migrator 数据库迁移管理器
type Migrator struct {
	db         *gorm.DB
	migrations []MigrationScript
}

type MigrationScript struct {
	Name     string
	Up       string
	Down     string
	CheckSum string
}

func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{
		db:         db,
		migrations: make([]MigrationScript, 0),
	}
}

// LoadMigrations 加载迁移脚本
func (m *Migrator) LoadMigrations(path string) error {
	// 确保migration表存在
	if err := m.db.AutoMigrate(&Migration{}); err != nil {
		return err
	}

	// 读取并解析迁移文件
	files, err := filepath.Glob(filepath.Join(path, "*.sql"))
	if err != nil {
		return err
	}

	sort.Strings(files)

	for _, file := range files {
		script, err := parseMigrationFile(file)
		if err != nil {
			return err
		}
		m.migrations = append(m.migrations, script)
	}

	return nil
}

// Migrate 执行迁移
func (m *Migrator) Migrate() error {
	return m.db.Transaction(func(tx *gorm.DB) error {
		for _, script := range m.migrations {
			var migration Migration
			result := tx.Where("name = ?", script.Name).First(&migration)

			if result.Error == gorm.ErrRecordNotFound {
				// 执行迁移
				if err := tx.Exec(script.Up).Error; err != nil {
					return err
				}

				// 记录迁移
				migration = Migration{
					Name:      script.Name,
					AppliedAt: time.Now(),
				}
				if err := tx.Create(&migration).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}
