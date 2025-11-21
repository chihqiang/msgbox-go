package workflow

import (
	"context"
)

// TaskInterface 定义任务需要实现的接口方法
type TaskInterface interface {
	// OnAction 执行任务逻辑，返回可能修改后的 context 和错误
	OnAction(ctx context.Context) (context.Context, error)
	// OnBefore 在任务执行前调用，可用于初始化或修改 context
	OnBefore(ctx context.Context) context.Context
	// OnFinish 在任务执行后调用，可用于清理或处理错误
	OnFinish(ctx context.Context, err error)
}

// Task 是 TaskInterface 的基础实现，支持 Before、Action、Finish 回调
type Task struct {
	// Action 任务执行逻辑
	Action func(ctx context.Context) (context.Context, error)
	// Before 任务执行前回调，可修改 context
	Before func(ctx context.Context) context.Context
	// Finish 任务完成后回调，可处理错误
	Finish func(ctx context.Context, err error)
}

// OnAction 执行任务的 Action 回调，如果未设置 Action，则直接返回 context
func (t *Task) OnAction(ctx context.Context) (context.Context, error) {
	// 如果 Action 不为空，则执行任务逻辑
	if t.Action != nil {
		return t.Action(ctx)
	}
	// Action 未设置，直接返回原始 context
	return ctx, nil
}

// OnBefore 执行任务开始前回调，如果未设置 Before，则直接返回 context
func (t *Task) OnBefore(ctx context.Context) context.Context {
	// 如果 Before 不为空，则执行回调
	if t.Before != nil {
		return t.Before(ctx)
	}
	return ctx
}

// OnFinish 执行任务完成后的回调，如果未设置 Finish，则不做任何操作
func (t *Task) OnFinish(ctx context.Context, err error) {
	// 如果 Finish 不为空，则执行回调
	if t.Finish != nil {
		t.Finish(ctx, err)
	}
}
