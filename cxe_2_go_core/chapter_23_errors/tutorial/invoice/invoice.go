package invoice

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/Rhymond/go-money"
)

// ReadFromFile mở tệp được chỉ định và phân tích dữ liệu CSV bên trong
func ReadFromFile(filename string) (*money.Money, *money.Money, error) {
    record, err := readCSV(filename)
    if err != nil {
        return nil, nil, err
    }
    
    vatExcConverted, err := strconv.Atoi(record[2])
    if err != nil {
        return nil, nil, err
    }
    vatExc := money.New(int64(vatExcConverted), "USD")

    vatConverted, err := strconv.Atoi(record[3])
    if err != nil {
        return nil, nil, err
    }
    vat := money.New(int64(vatConverted), "USD")

    return vatExc, vat, nil
}

func readCSV(filename string) ([]string, error) {
    invoiceCSV, err := os.Open(filename)
    defer invoiceCSV.Close()
    if err != nil {
        return nil, err
    }
    reader := csv.NewReader(invoiceCSV)
    record, err := reader.Read()
    if err == io.EOF {
        return nil, errors.New("file is empty")
    }
    if err != nil {
        return nil, err
    }
    return record, nil
}