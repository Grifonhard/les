package main

import (
	"fmt"
	"slog"
)

func main() {
	fmt.Println(slog.info("hello", "count", 3))
}