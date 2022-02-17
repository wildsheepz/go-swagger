// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

func makeMarkdownCmd() *cobra.Command {

	var markdownCmd = &cobra.Command{
		Use:                   "markdown",
		Short:                 "Generate markdown documentation (that you might be reading right now!!)",
		Long:                  "Generate markdown documentation (that you might be reading right now!!)",
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			outDir := viper.GetString("markdown_dir")
			err := os.MkdirAll(outDir, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating directory: %s\n", err)
				os.Exit(1)
			}
			cmd.DisableAutoGenTag = true
			err = doc.GenMarkdownTree(cmd.Root(), outDir)
			if err != nil {
				logDebugf("Error generating markdown: %s\n", err)
				os.Exit(1)
			}
			logDebugf("Generated markdown successfully!")
		},
	}
	return markdownCmd
}
