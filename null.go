package null

import (
	"bytes"
	"encoding/json"
)

const (
	hasValue int = iota // 0: 有值
	noValue             // 1: 无值
)

// Null
// 1 has 2 nil -> 有值 Value
// 1 nil 2 has -> 值为 Null
// 1 nil 2 nil -> 未定义值 Undefined
type Null[T any] [2]*T

func New[T any](val T) Null[T] {
	return Null[T]{&val, nil}
}

// Nullable 空值
func Nullable[T any]() Null[T] {
	return Null[T]{nil, new(T)}
}

// Undefined 未定义
func Undefined[T any]() Null[T] {
	return Null[T]{nil, nil}
}

func (t *Null[T]) Set(val T) {
	t[hasValue] = &val
}

func (t *Null[T]) Get() (T, bool) {
	if t[hasValue] == nil {
		var empty T
		if t[noValue] != nil {
			return empty, false
		} else {
			return empty, false
		}
	}
	return *t[hasValue], true
}

// MustGet .
func (t *Null[T]) MustGet() T {
	if t[hasValue] == nil {
		panic("no value available")
	}
	return *t[hasValue]
}

func (t *Null[T]) SetNull() {
	*t = Null[T]{nil, new(T)}
}

func (t *Null[T]) IsNull() bool {
	return t[hasValue] == nil && t[noValue] != nil
}

func (t *Null[T]) SetUndefined() {
	*t = Null[T]{nil, nil}
}

func (t *Null[T]) IsUndefined() bool {
	return t[hasValue] == nil && t[noValue] == nil
}

func (t Null[T]) MarshalJSON() ([]byte, error) {
	if t[hasValue] == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(t[hasValue])
}

func (t *Null[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		t.SetNull()
		return nil
	}
	if err := json.Unmarshal(data, &t[hasValue]); err == nil {
		return nil
	} else {
		return err
	}
}
