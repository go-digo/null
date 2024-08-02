package null

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Null
// 0 has -> 有值 Value
// 0 nil -> 值为 Null
// len 0 -> 未定义值 Undefined
type Null[T any] []*T

func New[T any](val T) Null[T] {
	return Null[T]{&val}
}

// Nullable 空值
func Nullable[T any]() Null[T] {
	return Null[T]{nil}
}

// Undefined 未定义
func Undefined[T any]() Null[T] {
	return make(Null[T], 0)
}

func (t *Null[T]) Set(val T) {
	if *t == nil {
		*t = Null[T]{&val}
	} else {
		(*t)[0] = &val
	}
}

func (t Null[T]) Get() (T, bool) {
	if len(t) == 0 {
		var empty T
		return empty, false
	}
	if t[0] == nil {
		var empty T
		return empty, false
	}
	return *t[0], true
}

// MustGet .
func (t Null[T]) MustGet() T {
	if len(t) == 0 || t[0] == nil {
		panic("no value available")
	}
	return *t[0]
}

func (t *Null[T]) SetNull() {
	*t = Null[T]{nil}
}

// IsNull 只有长度为1
// 值为nil的时候为true
func (t Null[T]) IsNull() bool {
	if len(t) == 0 {
		return false
	}
	return t[0] == nil
}

func (t *Null[T]) SetUndefined() {
	*t = Null[T]{}
}

// IsUndefined 切片长度为0的时候为true
func (t Null[T]) IsUndefined() bool {
	return len(t) == 0
}

func (t Null[T]) MarshalJSON() ([]byte, error) {
	if len(t) == 0 || t[0] == nil {
		return []byte("null"), nil
	}
	return json.Marshal(t[0])
}

func (t *Null[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		t.SetNull()
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err == nil {
		t.Set(v)
		return nil
	} else {
		return err
	}
}

func (t Null[T]) sqlNull() sql.Null[T] {
	v, b := t.Get()
	return sql.Null[T]{
		V:     v,
		Valid: b,
	}
}

func (t *Null[T]) Scan(value any) error {
	null := t.sqlNull()
	return null.Scan(value)
}

//// GormDataType 返回类型
//func (t Null[T]) GormDataType() string {
//	if len(t) == 0 || t[0] == nil {
//		return string(schema.Int)
//	}
//	return
//}

func (t Null[T]) Value() (driver.Value, error) {
	v, b := t.Get()
	if b {
		return v, nil
	}
	// 返回 nil 会尝试结构体第一个对象的类型
	return int64(0), nil
}
