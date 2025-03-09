package command

// Command 命令接口
type Command interface {
	Execute() string
	Undo() string
}

// Editor 文本编辑器（接收者）
type Editor struct {
	text string
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) InsertText(text string) {
	e.text += text
}

func (e *Editor) DeleteText(length int) {
	if len(e.text) >= length {
		e.text = e.text[:len(e.text)-length]
	}
}

func (e *Editor) GetText() string {
	return e.text
}

// InsertCommand 插入命令
type InsertCommand struct {
	editor *Editor
	text   string
}

func NewInsertCommand(editor *Editor, text string) *InsertCommand {
	return &InsertCommand{
		editor: editor,
		text:   text,
	}
}

func (c *InsertCommand) Execute() string {
	c.editor.InsertText(c.text)
	return "插入文本: " + c.text
}

func (c *InsertCommand) Undo() string {
	c.editor.DeleteText(len(c.text))
	return "撤销插入: " + c.text
}

// DeleteCommand 删除命令
type DeleteCommand struct {
	editor *Editor
	text   string
	length int
}

func NewDeleteCommand(editor *Editor, length int) *DeleteCommand {
	return &DeleteCommand{
		editor: editor,
		length: length,
	}
}

func (c *DeleteCommand) Execute() string {
	if len(c.editor.GetText()) >= c.length {
		c.text = c.editor.GetText()[len(c.editor.GetText())-c.length:]
		c.editor.DeleteText(c.length)
		return "删除文本: " + c.text
	}
	return "无法删除：文本长度不足"
}

func (c *DeleteCommand) Undo() string {
	c.editor.InsertText(c.text)
	return "撤销删除: " + c.text
}

// CommandInvoker 命令调用者
type CommandInvoker struct {
	commands  []Command
	undoStack []Command
}

func NewCommandInvoker() *CommandInvoker {
	return &CommandInvoker{
		commands:  make([]Command, 0),
		undoStack: make([]Command, 0),
	}
}

func (i *CommandInvoker) ExecuteCommand(command Command) string {
	result := command.Execute()
	i.commands = append(i.commands, command)
	i.undoStack = append(i.undoStack, command)
	return result
}

func (i *CommandInvoker) Undo() string {
	if len(i.undoStack) > 0 {
		command := i.undoStack[len(i.undoStack)-1]
		i.undoStack = i.undoStack[:len(i.undoStack)-1]
		return command.Undo()
	}
	return "没有可撤销的命令"
}
