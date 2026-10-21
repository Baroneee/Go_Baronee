package config

import (
	"fmt"
	"os"
)

const fileHeader = "APPCONF"

func Load() (string, error) {
    data, err := os.ReadFile("/tmp/myHotelApp/config.txt")
    if err != nil {
        return "", err
    }
    conf := string(data)
    if conf[0:8] != fileHeader {
        return "", fmt.Errorf("the config file do not begin by accepted header")
    }
    return conf, nil
}