package controllers

import (
	"fmt"
	"proxmanager/backend/models"
)

type Manager struct {
	ListVMs  []models.VM
	ListLXCs []models.LXC
}

func (m Manager) GetVMs() []models.VM {
	return m.ListVMs
}

func (m Manager) GetLXCs() []models.LXC {
	return m.ListLXCs
}

func (m Manager) TurnOffMachine(id int) (bool, error) {
	// Try VM
	for _, v := range m.ListVMs {
		if v.Id == id {
			fmt.Printf("Stopping VM %v", v.Name) // the %v is in value works for int, string etc.
			return true, nil
		}
	}
	// Try LXC
	for _, v := range m.ListLXCs {
		if v.Id == id {
			fmt.Printf("Stopping LXC %v", v.Name) // the %v is in value works for int, string etc.
			return true, nil
		}
	}

	// No VM or LXC Found with that ID
	return false, nil
}
