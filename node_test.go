package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMyTestErr(t *testing.T) {
	fmt.Println(time.Unix(1595487025, 0).Format(time.RFC3339))

	years, month, day := time.Now().Date()
	fmt.Println(time.Date(years, month, day, 0, 0, 0, 0, time.FixedZone("CST", 8*60*60)).Format(time.RFC3339))
	fmt.Println(time.Date(years, month, day, 0, 0, 0, 0, time.UTC).Format(time.RFC3339))
}
