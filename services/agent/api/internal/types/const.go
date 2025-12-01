package types

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	JWTAgentID = "agent_id"
	JWTPhone   = "phone"
)

func GetAgentID(ctx context.Context) (int64, error) {
	value := ctx.Value(JWTAgentID)
	if value == nil {
		return 0, fmt.Errorf("JWTAgentID not found in context")
	}
	// 尝试断言为 json.Number
	number, ok := value.(json.Number)
	if !ok {
		return 0, fmt.Errorf("JWTAgentID is not a json.Number")
	}
	// 转换为 int64
	id, err := number.Int64()
	if err != nil {
		return 0, fmt.Errorf("failed to convert JWTAgentID to int64: %w", err)
	}
	return id, nil
}
