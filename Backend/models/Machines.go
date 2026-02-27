package models

type Machine struct {
	Id     int    `json:"id" example:"1"`
	Name   string `json:"name" example:"Ubuntu VM/LXC"`
	Status string `json:"status" example:"running"`
	Cores  int    `json:"cores" example:"2"`
	Ram    int    `json:"ram" example:"4096"`
}

type VM struct {
	Machine
	Usbs []string `json:"usbs" example:""`
	Pcie []string `json:"pcie" example:""`
}

type LXC struct {
	Machine
	Unprivileged int `json:"unprivileged" example:"1"`
}
