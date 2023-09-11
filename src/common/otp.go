package common

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
)

var cfg = config.GetConfig()

func GenerateOtp() string {
	rand.Seed(time.Now().UnixNano())

	min := int(math.Pow10(6 - 1))
	max := int(math.Pow10(6) - 1)
	var num = rand.Intn(max-min) + min
	return strconv.Itoa(num)
}
