package server

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/templates"

	"github.com/hoangndst/vision/cmd/util"
	"github.com/hoangndst/vision/domain/constant"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdServer() *cobra.Command {
	var (
		serverShort = i18n.T(`Start vision server`)

		serverLong = i18n.T(`Start vision server.`)

		serverExample = i18n.T(`
		# Start vision server
		vision server --db-host localhost --db-port 5432 --db-user admin --db-password admin --db-name vision`)
	)

	o := NewServerOptions()
	cmd := &cobra.Command{
		Use:     "server",
		Short:   serverShort,
		Long:    templates.LongDesc(serverLong),
		Example: templates.Examples(serverExample),
		RunE: func(_ *cobra.Command, args []string) (err error) {
			defer util.RecoverErr(&err)
			o.Complete(args)
			util.CheckErr(o.Validate())
			util.CheckErr(o.Run())
			return
		},
	}

	o.AddServerFlags(cmd)

	return cmd
}

func (o *ServerOptions) AddServerFlags(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&o.Port, "port", "p", 80,
		i18n.T("Specify the port to listen on"))
	cmd.Flags().StringVarP(&o.LogFilePath, "log-file-path", "", constant.DefaultLogFilePath,
		i18n.T("File path to write logs to. Default to /home/admin/logs/kusion.log"))
	o.Database.AddFlags(cmd.Flags())
}
