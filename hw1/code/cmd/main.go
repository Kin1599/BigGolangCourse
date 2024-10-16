package main

import (
	"fmt"
	"livecode/internal/pkg/storage"

	"go.uber.org/zap"
)

func main() {
	stor, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	stor.Set("banana", "banana")
	stor.Set("avocado", "69")

	value1 := stor.Get("banana")
	if value1 != nil {
		stor.Logger.Info("Значение получено", zap.Any("value", value1))
		fmt.Println("Value for banana:", *value1)
	} else {
		stor.Logger.Info("Значение равно nil")
	}

	value2 := stor.Get("avocado")
	if value2 != nil {
		stor.Logger.Info("Значение получено", zap.Any("value", value2))
		fmt.Println("Value for avocado:", *value2)
	} else {
		stor.Logger.Info("Значение равно nil")
	}

	kind1 := stor.GetKind("banana")
	fmt.Println("Kind for banana:", kind1)
	kind2 := stor.GetKind("avocado")
	fmt.Println("Kind for avocado:", kind2)

	sumResult := storage.Sum[int64](10, 20)
	fmt.Println("Сумма:", sumResult)
}
