package types

import (
	"chihqiang/msgbox-go/services/common/errs"
)

func (r SendRequest) Validate() error {
	if r.TemplateCode == "" {
		return errs.ErrParamInvalid
	}
	return nil
}
