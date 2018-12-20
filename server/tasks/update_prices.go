package tasks

import (
	"fmt"
	"time"
)

//TaskUpdatePrices updates prices using ESI API
func TaskUpdatePrices() error {

	fmt.Println("TaskUpdatePrices started", time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("TaskUpdatePrices finished", time.Now().Format("2006-01-02 15:04:05"))

	return nil
}
