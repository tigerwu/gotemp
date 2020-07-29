package main

import "fmt"

type AbstractBanker interface {
	DoBusi()
}

type SaveBanker struct {
	//AbstractBanker
}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行了存款");
}

type TransBanker struct {
	//AbstractBanker
}

func (tb *TransBanker) DoBusi() {
	fmt.Println("进行了转账");
}

type PayBanker struct {
	//AbstractBanker
}

func (pb *PayBanker) DoBusi() {
	fmt.Println("进行了支付");
}

func BankerBusiness(banker AbstractBanker) {
	banker.DoBusi()
}

func main() {
	BankerBusiness(&SaveBanker{})
	BankerBusiness(&TransBanker{})
	BankerBusiness(&PayBanker{})
}