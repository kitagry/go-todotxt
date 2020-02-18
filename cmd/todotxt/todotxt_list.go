package main

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

func todotxtList(c *cli.Context) error {
	tasks, err := getTasks(c.String("file"))
	if err != nil {
		return xerrors.Errorf("Failed to getTasks: %w", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Done", "Priority", "Description", "Creation Date", "Completion Date"})
	table.SetBorder(false)
	table.SetHeaderColor(
		tablewriter.Colors{},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlackColor, tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor, tablewriter.FgBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgYellowColor, tablewriter.FgBlackColor},
	)

	for i, task := range tasks {
		line := make([]string, 6)
		line[0] = strconv.Itoa(i + 1)
		priColor := tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor}
		switch task.Priority() {
		case 'A':
			priColor = tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor}
		case 'B':
			priColor = tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor}
		case 'C':
			priColor = tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor}
		}
		line[2] = string(task.Priority())
		line[3] = task.Description()

		if task.HasCreationDate() {
			line[4] = task.CreationDate.Format("2006-01-02")
		}

		if task.HasCompletionDate() {
			line[5] = task.CompletionDate.Format("2006-01-02")
		}

		line[1] = "☐ "
		color := tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor}
		if task.Completed {
			line[1] = "☑"
			color = tablewriter.Colors{tablewriter.FgBlackColor}
			priColor = tablewriter.Colors{tablewriter.FgBlackColor}
		}

		table.Rich(line, []tablewriter.Colors{color, color, priColor, color, color, color})
	}
	table.Render()

	return nil
}
