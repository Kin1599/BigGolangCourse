package storage

import (
	"strconv"

	"go.uber.org/zap"
)

type Value struct {
	s    string
	d    int
	a    any
	b    bool
	Kind string
}

type Storage struct {
	innerString map[string]Value
	Logger      *zap.Logger
}

func NewStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}

	defer logger.Sync()
	logger.Info("created new storage")

	return Storage{
		innerString: make(map[string]Value),
		Logger:      logger,
	}, nil
}

func (r Storage) Set(key, value string) {

	var v Value
	if intValue, err := strconv.Atoi(value); err == nil {
		v = Value{d: intValue, Kind: "D"}
	} else {
		v = Value{s: value, Kind: "S"}
	}

	r.innerString[key] = v
	r.Logger.Info("Установлен ключ", zap.String("key", key), zap.String("value", value), zap.String("kind", v.Kind))
	r.Logger.Sync()
}

func (r Storage) Get(key string) *string {
	res, ok := r.innerString[key]
	if !ok {
		r.Logger.Info("Значение для ключа не найдено", zap.String("key", key))
		return nil
	}
	var result string
	if res.Kind == "D" {
		result = strconv.Itoa(res.d)
	} else {
		result = res.s
	}
	r.Logger.Info("Значение получено", zap.String("key", key), zap.String("value", res.s))
	return &result
}

func (r Storage) GetKind(key string) string {
	res, ok := r.innerString[key]
	if !ok {
		r.Logger.Info("Тип значения для ключа не найден", zap.String("key", key))
		return ""
	}
	r.Logger.Info("Тип значения не получен", zap.String("key", key), zap.String("kind", res.Kind))
	return res.Kind
}

func Sum[T int64 | uint64](x, y T) T {
	return x + y
}
