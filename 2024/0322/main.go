package main

import (
	"fmt"
	"log"
	"golang.org/x/sys/windows/svc/mgr"
)

func main() {
	const serviceName = "MyService"
	const displayName = "My Service"
	const binaryPathName = `C:\path\to\myservice.exe`

	m, err := mgr.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer m.Disconnect()

	// Check if the service already exists
	s, err := m.OpenService(serviceName)
	if err == nil {
		s.Close()
		fmt.Println("Service already exists.")
		return
	}

	// Create a new service
	//c := &svc.Config{
	//	ServiceName: serviceName,
	//	DisplayName: displayName,
	//	BinaryPathName: binaryPathName,
	//}
	//
	//s, err = m.CreateService(c)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer s.Close()

	fmt.Printf("Service %s created.\n", serviceName)
}