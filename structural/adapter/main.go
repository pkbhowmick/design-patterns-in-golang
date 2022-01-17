package main

import (
	"fmt"
)

// Client code
type client struct {
}

func (c *client) payElectricBill(p PaymentService) {
	fmt.Println("Client requests to pay the electric bill.")
	p.payBill()
}

// client interface
type PaymentService interface {
	payBill()
}

// Our owned service
type BankPaymentService struct {
}

func (b *BankPaymentService) payBill() {
	fmt.Println("Bill is paid by Bank payment service.")
}

// Now an unknown service called Mobile Banking Service
// has started their bill pay service.
// which is third party service and we can't adapt this service to out existing service.
type MobileBankingService struct {
}

func (m *MobileBankingService) payBillViaMobile() {
	fmt.Println("Bill is paid by Mobile banking service.")
}

// So, we need an adaptar to adapt to this third party service.
// because we don't want to change our existing client code.
type MobileBankingServiceAdapter struct {
	svc *MobileBankingService
}

func (m *MobileBankingServiceAdapter) payBill() {
	m.svc.payBillViaMobile()
}

func main() {
	fmt.Println("\t Design Patterns in Golang(2)")
	fmt.Printf("\t \t Adapter Pattern\n\n")

	c := &client{}
	bankSvc := &BankPaymentService{}
	c.payElectricBill(bankSvc)

	mobileSvc := &MobileBankingService{}
	adapter := &MobileBankingServiceAdapter{svc: mobileSvc}
	c.payElectricBill(adapter)
}

// output:
//      Design Patterns in Golang(2)
//          Adapter Pattern
//
// Client requests to pay the electric bill.
// Bill is paid by Bank payment service.
// Client requests to pay the electric bill.
// Bill is paid by Mobile banking service.
