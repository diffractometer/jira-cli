package issue

import (
	"github.com/spf13/cobra"

	"github.com/ankitpokhrel/jira-cli/internal/cmd/issue/create"
	"github.com/ankitpokhrel/jira-cli/internal/cmd/issue/list"
	"github.com/ankitpokhrel/jira-cli/internal/cmd/issue/move"
	"github.com/ankitpokhrel/jira-cli/internal/cmd/issue/view"
)

const helpText = `Issue manage issues in a given project. See available commands below.`

// NewCmdIssue is an issue command.
func NewCmdIssue() *cobra.Command {
	cmd := cobra.Command{
		Use:         "issue",
		Short:       "Issue manage issues in a project",
		Long:        helpText,
		Aliases:     []string{"issues"},
		Annotations: map[string]string{"cmd:main": "true"},
		RunE:        issue,
	}

	lc := list.NewCmdList()
	cc := create.NewCmdCreate()

	cmd.AddCommand(lc, cc, move.NewCmdMove(), view.NewCmdView())

	list.SetFlags(lc)
	create.SetFlags(cc)

	return &cmd
}

func issue(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}