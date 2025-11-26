package tasks

type CtxKey string

const (
	CtxModelAgent      CtxKey = "_model_agent"
	CtxModelTemplate   CtxKey = "_model_template"
	CtxModelChannel    CtxKey = "_model_channel"
	CtxModelSendBatch  CtxKey = "_model_send_batch"
	CtxModelSendRecord CtxKey = "_model_send_record"
)
