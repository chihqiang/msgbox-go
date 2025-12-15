package senders

import (
	"chihqiang/msgbox-go/pkg/htmlx"
	"fmt"
	"sync"
)

var (
	_senders = NewSenders()
)

func init() {
	_ = Register("dingtalk", "钉钉机器人", &DingTalkSender{})
	_ = Register("workwx", "企业微信机器人", &WorkWxSender{})
}
func Register(name, label string, sender ISender) error {
	return _senders.Register(name, label, sender)
}

func Get(name string) (*SenderForm, bool) {
	return _senders.Get(name)
}

func List() []*SenderForm {
	return _senders.List()
}

type Senders struct {
	mu      sync.RWMutex
	senders map[string]*SenderForm
	order   []string
}

func NewSenders() *Senders {
	return &Senders{senders: make(map[string]*SenderForm), order: make([]string, 0)}
}

func (r *Senders) Register(name, label string, sender ISender) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.senders[name]; exists {
		return fmt.Errorf("sender %s already registered", name)
	}
	r.senders[name] = &SenderForm{
		Name:   name,
		Label:  label,
		Sender: sender,
	}
	r.order = append(r.order, name)
	return nil
}
func (r *Senders) Get(name string) (*SenderForm, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.senders[name]
	return s, ok
}
func (r *Senders) List() []*SenderForm {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]*SenderForm, 0, len(r.senders))
	for _, name := range r.order {
		list = append(list, r.senders[name])
	}
	return list
}

type SenderForm struct {
	Name   string
	Label  string
	Sender ISender
}

func (s *SenderForm) FormFields() []htmlx.FormField {
	return htmlx.ToFormFields(s.Sender)
}
