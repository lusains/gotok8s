package script

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ScriptExecutor 脚本执行器
type ScriptExecutor struct {
	scriptPath string
	env        map[string]string
}

func NewScriptExecutor(scriptPath string) *ScriptExecutor {
	return &ScriptExecutor{
		scriptPath: scriptPath,
		env:        make(map[string]string),
	}
}

// SetEnv 设置环境变量
func (e *ScriptExecutor) SetEnv(key, value string) {
	e.env[key] = value
}

// ExecuteSQL 执行SQL脚本
func (e *ScriptExecutor) ExecuteSQL(db *gorm.DB, scriptName string) error {
	scriptPath := filepath.Join(e.scriptPath, scriptName)

	content, err := os.ReadFile(scriptPath)
	if err != nil {
		return err
	}

	// 按分号分割多个SQL语句
	statements := strings.Split(string(content), ";")

	return db.Transaction(func(tx *gorm.DB) error {
		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			if err := tx.Exec(stmt).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// ExecuteShell 执行Shell脚本
func (e *ScriptExecutor) ExecuteShell(scriptName string) error {
	scriptPath := filepath.Join(e.scriptPath, scriptName)

	cmd := exec.Command("bash", scriptPath)

	// 设置环境变量
	env := os.Environ()
	for k, v := range e.env {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	cmd.Env = env

	// 获取标准输出和错误输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		return err
	}

	// 实时打印输出
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// 等待命令完成
	return cmd.Wait()
}
