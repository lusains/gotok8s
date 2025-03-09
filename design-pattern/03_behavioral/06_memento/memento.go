package memento

import "time"

// EditorMemento 编辑器备忘录
type EditorMemento struct {
	content   string
	timestamp time.Time
}

func NewEditorMemento(content string) *EditorMemento {
	return &EditorMemento{
		content:   content,
		timestamp: time.Now(),
	}
}

func (e *EditorMemento) GetContent() string {
	return e.content
}

func (e *EditorMemento) GetTimestamp() time.Time {
	return e.timestamp
}

// Editor 文本编辑器（发起人）
type Editor struct {
	content string
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) Type(text string) {
	e.content += text
}

func (e *Editor) GetContent() string {
	return e.content
}

func (e *Editor) Save() *EditorMemento {
	return NewEditorMemento(e.content)
}

func (e *Editor) Restore(memento *EditorMemento) {
	e.content = memento.GetContent()
}

// History 历史记录（管理者）
type History struct {
	mementos []*EditorMemento
	current  int
}

func NewHistory() *History {
	return &History{
		mementos: make([]*EditorMemento, 0),
		current:  -1,
	}
}

func (h *History) Push(memento *EditorMemento) {
	// 如果当前不是最后一个状态，删除当前状态之后的所有状态
	if h.current+1 < len(h.mementos) {
		h.mementos = h.mementos[:h.current+1]
	}
	h.mementos = append(h.mementos, memento)
	h.current++
}

func (h *History) Undo() *EditorMemento {
	if h.current > 0 {
		h.current--
		return h.mementos[h.current]
	}
	return nil
}

func (h *History) Redo() *EditorMemento {
	if h.current+1 < len(h.mementos) {
		h.current++
		return h.mementos[h.current]
	}
	return nil
}

func (h *History) GetHistory() []string {
	history := make([]string, len(h.mementos))
	for i, memento := range h.mementos {
		history[i] = memento.GetContent()
	}
	return history
}
