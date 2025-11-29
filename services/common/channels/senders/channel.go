package senders

var (
	senders = []ISender{
		&DingTalkSender{},
		&WorkWxSender{},
	}
)

func Get() []ISender {
	return senders
}

type IMessage interface {
	GetReceiver() string
	GetSignature() string
	GetVendorCode() string
	GetTitle() string
	GetContent() string
	GetVariables() map[string]any
	GetExtra() map[string]any
}

type ISender interface {
	GetName() string
	GetLabel() string
	SetConfig(config map[string]any) error
	Send(message IMessage) (map[string]any, error)
}
