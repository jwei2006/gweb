package service

import (
	"errors"
	"github.com/jwei2006/gweb/config"
	"github.com/jwei2006/gweb/sys"
	"math/rand"
	"strconv"
	"time"
)

func SendCode(ctx sys.Context, phone string) error {
	if phone == ""{
		return errors.New(config.GetErrorTitle(config.ParamError))
	}
	return nil
}
func CreateCode() string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	intn := r.Intn(899999) + 100000
	return strconv.Itoa(intn)
}
