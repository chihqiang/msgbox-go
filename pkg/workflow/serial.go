package workflow

import (
	"context"
)

// StageSerial 表示一个串行执行任务的工作流
type StageSerial struct {
	// tasks 串行执行的任务列表
	tasks []TaskInterface
}

// NewStageSerial 创建一个新的串行工作流实例
func NewStageSerial() *StageSerial {
	// 初始化 StageSerial 并创建空任务列表
	return &StageSerial{tasks: make([]TaskInterface, 0)}
}

// Add 向串行工作流中添加任务
func (w *StageSerial) Add(task ...TaskInterface) {
	// 将任务追加到任务列表末尾
	w.tasks = append(w.tasks, task...)
}

// Run 按顺序执行所有任务，遇到错误或 context 取消则停止
func (w *StageSerial) Run(ctx context.Context) error {
	// 当前上下文
	currentCtx := ctx
	// 遍历所有任务
	for _, t := range w.tasks {
		// 检查上下文是否被取消
		select {
		case <-currentCtx.Done():
			// 上下文取消，返回错误
			return currentCtx.Err()
		default:
		}
		// 执行任务前回调
		beforeCtx := t.OnBefore(currentCtx)
		// 执行任务逻辑
		nextCtx, err := t.OnAction(beforeCtx)
		// 执行任务完成回调
		t.OnFinish(nextCtx, err)
		// 如果任务执行出错，则终止后续任务
		if err != nil {
			return err
		}
		// 更新上下文供下一个任务使用
		currentCtx = nextCtx
	}
	// 所有任务执行成功
	return nil
}
