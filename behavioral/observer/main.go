package main

import (
	"fmt"
)

// Subject/Publisher interface
type AlertManager interface {
	subscribe(observer)
	unsubscribe(observer)
	notifyAll()
}

// concrete subject or publisher
type ServerDownAlert struct {
	observerList []observer
	isFired      bool
}

func newServerDownAlert() *ServerDownAlert {
	return &ServerDownAlert{
		isFired: true,
	}
}

func (a *ServerDownAlert) fire() {
	fmt.Println("Server is down! Alert!")
	a.notifyAll()
}

func (a *ServerDownAlert) subscribe(o ...observer) {
	a.observerList = append(a.observerList, o...)
}

func (a *ServerDownAlert) unsubscribe(o observer) {
	a.observerList = removeFromslice(a.observerList, o)
}

func (a *ServerDownAlert) notifyAll() {
	for _, o := range a.observerList {
		o.update("server is down")
	}
}

func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getEmail() == observer.getEmail() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

type observer interface {
	update(string)
	getEmail() string
}

// concrete observer
type clusterAdmin struct {
	email string
}

func (c *clusterAdmin) update(alertDetails string) {
	fmt.Printf("Sending email to %s abour alert: %s\n", c.email, alertDetails)
}

func (c *clusterAdmin) getEmail() string {
	return c.email
}

func main() {
	fmt.Println("\t Design Patterns in Golang(3)")
	fmt.Printf("\t \t Observer Pattern\n\n")

	alert := newServerDownAlert()

	sysAdmin := &clusterAdmin{email: "sysadmin@example.com"}
	serverAdmin := &clusterAdmin{"svradmin@example.com"}

	alert.subscribe(sysAdmin, serverAdmin)
	alert.fire()

	alert.unsubscribe(serverAdmin)
	alert.fire()
}

// output:
//     Design Patterns in Golang(3)
//           Observer Pattern

// Server is down! Alert!
// Sending email to sysadmin@example.com abour alert: server is down
// Sending email to svradmin@example.com abour alert: server is down
// Server is down! Alert!
// Sending email to sysadmin@example.com abour alert: server is down
