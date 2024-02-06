package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

type kafka struct{}

func (s kafka) Run() {
	table := tablewriter.NewWriter(os.Stdout)
	table.Append([]string{"Tool", "Kafka"})
	table.Append([]string{"Check", "OK"})
	table.Append([]string{"Brokers", "broker1,broker2,broker3"})
	table.Append([]string{"Status", "Up"})
	table.Render()
}

var BasePlugin kafka
