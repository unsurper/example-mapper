package main

import (
	"plcmapper/driver"
	"plcmapper/mapper-sdk-go/pkg/service"
)

// main Virtual device program entry
func main() {
	vd := &driver.VirtualDevice{}
	// protocolName must be consistent with the protocol name defined by CRD
	service.Bootstrap("virtualProtocol", vd)
}
