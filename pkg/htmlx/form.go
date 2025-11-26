package htmlx

import (
	"fmt"
	"reflect"
	"strings"
)

// FormField 表示前端表单字段的结构
type FormField struct {
	Type        string `json:"type"`        // 表单类型，如 text、select
	Name        string `json:"name"`        // 字段标识（一般对应 JSON key）
	Label       string `json:"label"`       // 前端显示名称
	Required    bool   `json:"required"`    // 是否必填
	Placeholder string `json:"placeholder"` // 输入框提示内容
	Default     string `json:"default"`     // 默认值
}

// ToFormFields 使用反射将 struct 转换成 Form slice
//
// 说明：
// 1. struct 字段必须包含 `json` tag 用作 Name。
// 2. struct 字段可选 `ui` tag 用于表单元数据，例如 label、type、required、placeholder、default。
// 3. 默认类型为 text
//
// ui tag 示例：
//
//	text 输入框:
//	 `ui:"label=Webhook地址;type=text;required;placeholder=请输入 Webhook;default=https://oapi.dingtalk.com/robot/send"`
//
// key 说明:
//   - label: 表单显示名称
//   - type: 表单类型 text
//   - required: 必填字段（只要出现即表示必填）
//   - placeholder: 占位符
//   - default: 默认值
func ToFormFields(data interface{}) []FormField {
	v := reflect.ValueOf(data)
	if !v.IsValid() {
		return nil
	}
	// 如果是指针，解引用
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	// 不是 struct，直接返回空 slice
	if v.Kind() != reflect.Struct {
		return nil
	}
	t := v.Type()
	var forms []FormField
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// 获取 JSON tag 作为 Name
		jsonTag := field.Tag.Get("json")
		name := strings.Split(jsonTag, ",")[0]

		// 初始化 Form，默认 type=text
		form := FormField{
			Name:        name,
			Type:        "text",
			Label:       name,
			Required:    false,
			Placeholder: "",
			Default:     "",
		}

		// 解析 ui tag
		uiTag := field.Tag.Get("ui")
		for _, part := range strings.Split(uiTag, ";") {
			if part == "" {
				continue
			}
			kv := strings.SplitN(part, "=", 2)
			key := kv[0]
			// 防止越界
			var val string
			if len(kv) > 1 {
				val = kv[1]
			}
			switch key {
			case "type":
				form.Type = val
			case "label":
				form.Label = val
			case "required":
				form.Required = true
			case "placeholder":
				form.Placeholder = val
			case "default":
				form.Default = val
			}
		}
		// struct 字段有实际值时覆盖默认值
		if value.Kind() == reflect.String && value.String() != "" {
			form.Default = value.String()
		}

		forms = append(forms, form)
	}

	return forms
}

// MapSet 将 map[string]any 中的值设置到任意 struct 上。
// 支持字段类型：string、int、float、bool。
// 对于 string 类型字段，如果 map 中没有提供值且字段原始值为空，
// 会读取字段 ui tag 中的 default 作为默认值。
func MapSet(v any, maps map[string]any) error {
	// 1. 检查 v 是否为非空指针
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("v must be a non-nil pointer to struct")
	}

	// 2. 获取指针指向的 struct 值
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("v must be a pointer to struct")
	}

	rt := rv.Type()
	// 3. 遍历 struct 的每一个字段
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)    // 字段类型信息
		fieldVal := rv.Field(i) // 字段的值

		// 4. 使用 json tag 作为 map 的 key
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			// 没有 json tag，可以使用字段名作为 key
			jsonTag = field.Name
		} else {
			// 处理可能带 omitempty 或其它选项
			jsonTag = strings.Split(jsonTag, ",")[0]
		}
		hasValue := false // 标记字段是否被 map 覆盖
		// 5. 如果 map 中有对应 key 且非 nil，则尝试设置字段
		if val, ok := maps[jsonTag]; ok && val != nil {
			if !fieldVal.CanSet() {
				continue // 字段不可设置，跳过
			}

			// 6. 根据字段类型进行类型断言和赋值
			switch fieldVal.Kind() {
			case reflect.String:
				if s, ok := val.(string); ok && s != "" {
					fieldVal.SetString(s)
					hasValue = true
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if f, ok := val.(float64); ok { // JSON 数字默认 float64
					fieldVal.SetInt(int64(f))
					hasValue = true
				} else if i, ok := val.(int); ok {
					fieldVal.SetInt(int64(i))
					hasValue = true
				}
			case reflect.Bool:
				if b, ok := val.(bool); ok {
					fieldVal.SetBool(b)
					hasValue = true
				}
			case reflect.Float32, reflect.Float64:
				if f, ok := val.(float64); ok {
					fieldVal.SetFloat(f)
					hasValue = true
				}
			}
		}

		// 7. 对于 string 类型字段，如果 map 没有提供值且字段原始值为空
		//    尝试读取 ui tag 中的 default 作为默认值
		if fieldVal.Kind() == reflect.String && !hasValue && fieldVal.String() == "" {
			uiTag := field.Tag.Get("ui")
			for _, part := range strings.Split(uiTag, ";") {
				kv := strings.SplitN(part, "=", 2)
				if len(kv) == 2 && kv[0] == "default" {
					fieldVal.SetString(kv[1])
					break
				}
			}
		}
	}

	return nil
}
