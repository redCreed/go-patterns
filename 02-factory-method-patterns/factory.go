package main

import "errors"

//工厂方法模式
const (
	PayWechat = 1
	PayAli    = 2
)

type Payment interface {
	Pay(money float64) error
}

type Wechat struct {
	Balance float32
}

type Ali struct {
	Balance float32
}

func (w *Wechat) Pay(money float32) error {
	if w.Balance < 0 || w.Balance < money {
		return errors.New("balance not enough")
	}
	w.Balance -= money
	return nil
}

func (a *Ali) Pay(money float32) error {
	if a.Balance < 0 || a.Balance < money {
		return errors.New("balance not enough")
	}
	a.Balance -= money
	return nil
}

func Pay(payType int, money float32) error {
	switch payType {
	case PayWechat:
		w := &Wechat{
			Balance: 500,
		}
		return w.Pay(100)
	case PayAli:
		a := &Ali{
			Balance: 500,
		}
		return a.Pay(100)
	default:
		return errors.New("not support pay type")
	}
}
