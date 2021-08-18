package cmd

import (
	"github.com/mroach/n64-go/formatters"
	"github.com/mroach/n64-go/rom"
	"github.com/spf13/cobra"
)

func init() {
	var outputFormat string
	var columns []string

	var statCmd = &cobra.Command{
		Use:     "stat",
		Aliases: []string{"info"},
		Short:   "Get ROM file information",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]

			info, err := rom.FromPath(path)
			if err != nil {
				return err
			}

			if len(columns) == 0 {
				columns = formatters.DefaultColumns(outputFormat)
			}

			columns, err := validateColumns(columns)
			if err != nil {
				printColumnHelp()
				return err
			}

			err = info.AddHashes()
			if err != nil {
				return err
			}

			return formatters.PrintOne(info, outputFormat, columns)
		},
	}

	statCmd.Flags().StringVarP(&outputFormat, "output", "o", "text", "Output format")
	statCmd.Flags().StringSliceVarP(&columns, "columns", "c", make([]string, 0), "Column selection")

	rootCmd.AddCommand(statCmd)
}
