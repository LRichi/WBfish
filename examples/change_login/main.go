//
// SPDX-License-Identifier: BSD-3-Clause
//
package main

import (
	"github.com/LRichi/WBfish"
)

func main() {
	// Create a new instance of gofish client, ignoring self-signed certs
	username := "my-username"
	config := wbfish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: username,
		Password: "my-password",
		Insecure: true,
	}
	c, err := wbfish.Connect(config)
	if err != nil {
		panic(err)
	}
	defer c.Logout()

	// Retrieve the service root
	service := c.Service

	// Query the AccountService using the session token
	accountService, err := service.AccountService()
	if err != nil {
		panic(err)
	}
	// Get list of accounts
	accounts, err := accountService.Accounts()
	if err != nil {
		panic(err)
	}
	// Iterate over accounts to find the current user
	for _, account := range accounts {
		if account.UserName == username {
			account.UserName = "new-username"
			// New password must follow the rules set in AccountService :
			// MinPasswordLength and MaxPasswordLength
			account.Password = "new-password"
			err := account.Update()
			if err != nil {
				panic(err)
			}
		}
	}
}
