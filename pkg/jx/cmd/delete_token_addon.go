package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/jenkins-x/jx/pkg/jx/cmd/templates"
	"github.com/jenkins-x/jx/pkg/log"
	"github.com/jenkins-x/jx/pkg/util"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
)

var (
	deleteTokenAddonLong = templates.LongDesc(`
		Deletes one or more API tokens for your addon from your local settings
`)

	deleteTokenAddonExample = templates.Examples(`
		# Deletes an addon user token
		jx delete token addon -n anchore myusername
	`)
)

// DeleteTokenAddonOptions the options for the create spring command
type DeleteTokenAddonOptions struct {
	CreateOptions

	ServerFlags ServerFlags
	Kind        string
}

// NewCmdDeleteTokenAddon defines the command
func NewCmdDeleteTokenAddon(f Factory, in terminal.FileReader, out terminal.FileWriter, errOut io.Writer) *cobra.Command {
	options := &DeleteTokenAddonOptions{
		CreateOptions: CreateOptions{
			CommonOptions: CommonOptions{
				Factory: f,
				In:      in,
				Out:     out,
				Err:     errOut,
			},
		},
	}

	cmd := &cobra.Command{
		Use:     "addon",
		Short:   "Deletes one or more API tokens for a user on an issue addon server",
		Long:    deleteTokenAddonLong,
		Example: deleteTokenAddonExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			CheckErr(err)
		},
	}
	options.ServerFlags.addGitServerFlags(cmd)
	cmd.Flags().StringVarP(&options.Kind, "kind", "k", "", "The kind of addon. Defaults to the addon name if not specified")
	return cmd
}

// Run implements the command
func (o *DeleteTokenAddonOptions) Run() error {
	args := o.Args
	if len(args) == 0 {
		return fmt.Errorf("Missing addon user name")
	}
	authConfigSvc, err := o.createAddonAuthConfigService()
	if err != nil {
		return err
	}
	config := authConfigSvc.Config()

	kind := o.Kind
	if kind == "" {
		kind = o.ServerFlags.ServerName
	}
	if kind == "" {
		kind = "addon"
	}
	server, err := o.findAddonServer(config, &o.ServerFlags, kind)
	if err != nil {
		return err
	}
	for _, username := range args {
		err = server.DeleteUser(username)
		if err != nil {
			return err
		}
	}
	err = authConfigSvc.SaveConfig()
	if err != nil {
		return err
	}
	log.Infof("Deleted API tokens for users: %s for addon server %s at %s from local settings\n",
		util.ColorInfo(strings.Join(args, ", ")), util.ColorInfo(server.Name), util.ColorInfo(server.URL))
	return nil
}
