package workflow

import (
	"context"
	"strings"
	"sync"
)

// StageParallel 并行执行任务的工作流
// 支持任务优先级调度和最大并行度控制
type StageParallel struct {
	// limit 最大并行任务数，当为 0 时表示默认同时执行所有任务
	limit int
	// tasks 存储待执行的任务列表
	tasks []TaskInterface
}

// NewStageParallel 创建一个并行工作流实例
// 默认并行任务数为 0，表示默认执行所有任务
func NewStageParallel() *StageParallel {
	return &StageParallel{
		limit: 0,
		tasks: make([]TaskInterface, 0),
	}
}

// SetLimit 设置并行工作流的最大并行任务数
// limit <= 0 表示不限制并行数量
func (s *StageParallel) SetLimit(limit int) {
	s.limit = limit
}

// Add 向并行工作流添加一个或多个任务
func (w *StageParallel) Add(task ...TaskInterface) {
	w.tasks = append(w.tasks, task...)
}

// Run 执行所有任务，按优先级调度，并行执行
// 如果 limit > 0，则并行执行的任务数量不超过 limit
// 任务出错会立即取消其他任务，返回第一个错误
func (w *StageParallel) Run(ctx context.Context) error {
	if len(w.tasks) == 0 {
		return nil
	}
	if w.limit <= 0 {
		w.limit = len(w.tasks)
	}
	var (
		wg     sync.WaitGroup
		errCh  = make(chan *StageParallelErr, len(w.tasks))
		taskCh = make(chan TaskInterface, len(w.tasks))
	)
	for _, task := range w.tasks {
		taskCh <- task
	}
	close(taskCh)
	// 启动工作协程
	for i := 0; i < w.limit; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					// 如果外部上下文被取消，当前协程退出
					return
				case task, ok := <-taskCh:
					if !ok {
						// 任务通道已关闭，没有更多任务了
						return
					}
					// 执行任务前回调
					beforeCtx := task.OnBefore(ctx)
					// 执行任务逻辑
					nextCtx, err := task.OnAction(beforeCtx)
					// 执行任务完成回调
					task.OnFinish(nextCtx, err)
					if err != nil {
						// 将错误发送到通道
						errCh <- &StageParallelErr{Stage: w, Task: task, Err: err}
					}
				}
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(errCh)
	}()
	var errs StageParallelErrs
	for err := range errCh {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// StageParallelErr 单个任务的执行错误（包含上下文信息）
type StageParallelErr struct {
	Stage *StageParallel // 所属的并行工作流实例
	Task  TaskInterface  // 出错的任务实例
	Err   error          // 具体的错误信息
}

// AssertStageParallelErrNum 判断错误是否为 StageParallelErrs 类型，并返回错误数量
// 参数：
//
//	err: 待判断的错误实例
//
// 返回值：
//
//	int: 错误数量（若为 StageParallelErrs 类型），否则为 0
//	bool: 是否为 StageParallelErrs 类型（true=是，false=否）
func AssertStageParallelErrNum(err error) (int, bool) {
	if errs, ok := err.(StageParallelErrs); ok {
		return len(errs), true
	}
	return 0, false
}

// StageParallelErrs 多个任务的错误集合（实现 error 接口）
type StageParallelErrs []*StageParallelErr

// Error 实现 error 接口，返回所有错误的汇总信息
// 格式：错误1; 错误2; ...（多个错误用分号分隔）
func (sps StageParallelErrs) Error() string {
	var parts []string
	for _, e := range sps {
		parts = append(parts, e.Err.Error())
	}
	return strings.Join(parts, "; ")
}
