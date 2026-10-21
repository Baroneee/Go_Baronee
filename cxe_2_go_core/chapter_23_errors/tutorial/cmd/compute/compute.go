package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Rhymond/go-money"
	"maximilien-andile.com/errors/application/invoice"
)

const startDate = "2003-07-01"
const endDate = "2020-07-01"
const baseDirPath = "/Users/maximilienandile/Documents/DEV/goBook/errors/application/data"

var totalVat *money.Money
var totalVatExc *money.Money
var totalVatInc *money.Money

func init() {
    totalVat = money.New(0, "USD")
    totalVatExc = money.New(0, "USD")
	totalVatInc = money.New(0, "USD")
}

func main() {
    start, err := time.Parse("2006-01-02", startDate)
    if err != nil {
        log.Fatalf("impossible to parse start date %s", err)
    }
    end, err := time.Parse("2006-01-02", endDate)
    if err != nil {
        log.Fatalf("impossible to parse end date %s", err)
    }

    for d := start; d.Unix() < end.Unix(); d = d.AddDate(0, 1, 0) {
        monthDirPath := fmt.Sprintf("%s/%s-%d", baseDirPath, d.Month(), d.Year())
        dir, err := os.Open(monthDirPath)
        if err != nil {
            log.Fatalf("failed to open directory %s: %s", monthDirPath, err)
        }
        defer dir.Close()

        list, err := dir.Readdirnames(0)
        if err != nil {
            log.Fatalf("failed to read all files in directory %s: %s", monthDirPath, err)
        }
 
        for _, name := range list {
            filePath := fmt.Sprintf("%s/%s", monthDirPath, name)
            vatExc, vat, err := invoice.ReadFromFile(filePath)
            if err != nil {
                log.Printf("failed to parse invoice %s: %s", filePath, err)
    			continue
            }
            totalVat, err = totalVatExc.Add(vat)
            if err != nil {
                log.Printf("impossible to add VAT to counter")
				continue
            }
            totalVatExc, err = totalVatExc.Add(vatExc)
            if err != nil {
                log.Printf("impossible to add VAT to counter")
				continue
            }
        }
    }
	totalVatInc, err = totalVat.Add(totalVatExc)
	if err != nil {
		log.Fatalf("impossible to calculate total VAT Inc: %s", err)
	}

    log.Println("total VAT", totalVat.Display())
    log.Println("total VAT Exc", totalVatExc.Display())
	log.Println("total VAT Inc", totalVatInc.Display())
}