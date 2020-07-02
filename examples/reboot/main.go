package main

import (
	"fmt"

	"github.com/LRichi/WBfish"
	"github.com/LRichi/WBfish/redfish"
)

func main() {
	// Create a new instance of gofish client, ignoring self-signed certs
	config := wbfish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	}

	c, err := wbfish.Connect(config)
	if err != nil {
		panic(err)
	}
	defer c.Logout()

	// Attached the client to service root
	service := c.Service

	// Query the computer systems
	ss, err := service.Systems()
	if err != nil {
		panic(err)
	}

	// Creates a boot override to pxe once
	bootOverride := redfish.Boot{
		BootSourceOverrideTarget:  redfish.PxeBootSourceOverrideTarget,
		BootSourceOverrideEnabled: redfish.OnceBootSourceOverrideEnabled,
	}

	for _, system := range ss {
		fmt.Printf("System: %#v\n\n", system)
		err := system.SetBoot(bootOverride)
		if err != nil {
			panic(err)
		}
		err = system.Reset(redfish.ForceRestartResetType)
		if err != nil {
			panic(err)
		}
	}
}
