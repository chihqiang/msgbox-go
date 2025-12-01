package senders

import (
	"chihqiang/msgbox-go/pkg/htmlx"
	"fmt"
	"sync"
)

var (
	_sr = NewSenderRegistry()
)

func init() {
	_ = Register("dingtalk", "钉钉机器人", &DingTalkSender{})
	_ = Register("workwx", "企业微信机器人", &WorkWxSender{})
}
func Register(name, label string, sender ISender) error {
	return _sr.Register(name, label, sender)
}

func Get(name string) (*NameSender, bool) {
	return _sr.Get(name)
}

func List() []*NameSender {
	return _sr.List()
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
	SetConfig(config map[string]any) error
	Send(message IMessage) (map[string]any, error)
}

type NameSender struct {
	Name   string
	Label  string
	Sender ISender
}

func (s *NameSender) SenderFormFields() []htmlx.FormField {
	return htmlx.ToFormFields(s.Sender)
}

type SenderRegistry struct {
	mu      sync.RWMutex
	senders map[string]*NameSender
	order   []string
}

func NewSenderRegistry() *SenderRegistry {
	return &SenderRegistry{
		senders: make(map[string]*NameSender),
		order:   []string{},
	}
}
func (r *SenderRegistry) Get(name string) (*NameSender, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.senders[name]
	return s, ok
}

func (r *SenderRegistry) Register(name, label string, sender ISender) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.senders[name]; exists {
		return fmt.Errorf("sender %s already registered", name)
	}
	r.senders[name] = &NameSender{
		Name:   name,
		Label:  label,
		Sender: sender,
	}
	r.order = append(r.order, name)
	return nil
}

func (r *SenderRegistry) List() []*NameSender {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]*NameSender, 0, len(r.senders))
	for _, name := range r.order {
		list = append(list, r.senders[name])
	}
	return list
}
