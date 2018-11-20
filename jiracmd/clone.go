package jiracmd

import (
	"github.com/coryb/figtree"
	"github.com/coryb/oreo"
	"gopkg.in/Netflix-Skunkworks/go-jira.v1"
	"gopkg.in/Netflix-Skunkworks/go-jira.v1/jiracli"
	"gopkg.in/alecthomas/kingpin.v2"
)

type CloneOptions struct {
	jiracli.CommonOptions `yaml:",inline" json:",inline" figtree:",inline"`
	jira.IssueOptions     `yaml:",inline" json:",inline" figtree:",inline"`
	Issue                 string `yaml:"issue,omitempty" json:"issue,omitempty"`
	Deep				  bool `yaml:"deep,omitempty" json:deep,omitempty"`
}

func CmdCloneRegistry() *jiracli.CommandRegistryEntry {
	opts := CloneOptions{
		CommonOptions: jiracli.CommonOptions{
			Template: figtree.NewStringOption("clone"),
		},
	}

	return &jiracli.CommandRegistryEntry{
		"Clone issue",
		func(fig *figtree.FigTree, cmd *kingpin.CmdClause) error {
			jiracli.LoadConfigs(cmd, fig, &opts)
			return CmdCloneUsage(cmd, &opts)
		},
		func(o *oreo.Client, globals *jiracli.GlobalOptions) error {
			return CmdClone(o, globals, &opts)
		},
	}
}

func CmdCloneUsage(cmd *kingpin.CmdClause, opts *CloneOptions) error {
	jiracli.BrowseUsage(cmd, &opts.CommonOptions)
	jiracli.TemplateUsage(cmd, &opts.CommonOptions)
	jiracli.GJsonQueryUsage(cmd, &opts.CommonOptions)
	cmd.Flag("deep", "deep clone subtasks and relationships").BoolVar(&opts.Deep)
	cmd.Arg("ISSUE", "issue id to clones").Required().StringVar(&opts.Issue)
	return nil
}

// Clone with duplicate an existing issue and its sub-objects
func CmdClone(o *oreo.Client, globals *jiracli.GlobalOptions, opts *CloneOptions) error {
	data, err := jira.GetIssue(o, globals.Endpoint.Value, opts.Issue, opts)
	if err != nil {
		return err
	}


	/*
	if err := opts.PrintTemplate(data); err != nil {
		return err
	}
	if opts.Browse.Value {
		return CmdBrowse(globals, opts.Issue)
	}
	*/


	return nil
}