package main

import (
	"log"
	"your/package/initjob/config"
	"your/package/initjob/database"
	"your/package/initjob/script"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// 初始化数据库连接
	db, err := database.InitDB(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	// 创建迁移管理器
	migrator := database.NewMigrator(db)
	if err := migrator.LoadMigrations("./migrations"); err != nil {
		log.Fatal(err)
	}

	// 执行迁移
	if err := migrator.Migrate(); err != nil {
		log.Fatal(err)
	}

	// 创建脚本执行器
	executor := script.NewScriptExecutor(cfg.Script.Path)

	// 设置环境变量
	executor.SetEnv("APP_ENV", cfg.App.Mode)
	executor.SetEnv("DB_DSN", cfg.Database.DSN)

	// 执行预处理脚本
	for _, scriptName := range cfg.Script.PreScripts {
		if err := executor.ExecuteSQL(db, scriptName); err != nil {
			log.Fatal(err)
		}
	}

	// 执行后处理脚本
	for _, scriptName := range cfg.Script.PostScripts {
		if err := executor.ExecuteShell(scriptName); err != nil {
			log.Fatal(err)
		}
	}
}
