package booking

import (
	"fmt"
)

const vatRate = 20.0

func PrintVatRate() {
	fmt.Printf("VAT rate: %.2f%%\n", vatRate)
}