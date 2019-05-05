package main

import (
	"os"
	"strconv"

	"github.com/logrusorgru/aurora"
	"github.com/olekukonko/tablewriter"
)

func getInstanceColor(state *string) aurora.Value {
	switch *state {
	case "pending":
		return aurora.BrightYellow(*state)
	case "running":
		return aurora.Green(*state)
	case "stopping":
		return aurora.BrightRed(*state)
	case "shutting-down":
		return aurora.BrightRed(*state)
	case "stopped":
		return aurora.Red(*state)
	case "terminated":
		return aurora.Red(*state)
	default:
		return aurora.White(*state)
	}
}

func showInstances(instances []*Instance) {
	total := len(instances)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "IP", "State", "Type", "LaunchTime", "StackName", "Auto scaling group"})
	table.SetFooter([]string{"", "", "", "", "", "", "Total", strconv.Itoa(total)})

	table.SetBorder(false)
	table.SetRowSeparator("-")
	table.SetCenterSeparator("|")

	for i := 0; i < total; i++ {
		instance := instances[i]

		table.Append([]string{
			*instance.ID,
			*instance.Name,
			aurora.Cyan(*instance.IP).String(),
			getInstanceColor(instance.State).String(),
			*instance.Type,
			aurora.Magenta(instance.LaunchTime.Format("2006-01-02T15:04:05")).String(),
			*instance.StackName,
			*instance.AutoScalingGroup,
		})
	}

	table.Render()
}
