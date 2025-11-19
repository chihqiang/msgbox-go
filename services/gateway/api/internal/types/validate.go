package types

import "chihqiang/msgbox-go/services/common/errors"

func (r SendRequest) Validate() error {
	if r.TemplateCode == "" {
		return errors.ErrParamInvalid
	}
	return nil
}
