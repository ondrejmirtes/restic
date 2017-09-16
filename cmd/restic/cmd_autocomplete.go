package main

import (
	"github.com/spf13/cobra"
)

var cmdAutocomplete = &cobra.Command{
	Use:   "autocomplete",
	Short: "Generate shell autocompletion script",
	Long: `The "autocomplete" command generates a shell autocompletion script.

NOTE: The current version supports Bash only.
      This should work for *nix systems with Bash installed.

By default, the file is written directly to /etc/bash_completion.d
for convenience, and the command may need superuser rights, e.g.:

$ sudo restic autocomplete`,

	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmdRoot.GenBashCompletionFile(autocompleteTarget); err != nil {
			return err
		}
		return nil
	},
}

var autocompleteTarget string

func init() {
	cmdRoot.AddCommand(cmdAutocomplete)

	cmdAutocomplete.Flags().StringVarP(&autocompleteTarget, "completionfile", "", "/usr/share/bash-completion/completions/restic", "autocompletion file")
	// For bash-completion
	cmdAutocomplete.Flags().SetAnnotation("completionfile", cobra.BashCompFilenameExt, []string{})
}
