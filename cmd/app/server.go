package app

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/fundoplicatedFundus/toy-server/pkg/server"
)

// NewServerCommand creates a cobra command that serves as the entrypoint into
// our server.
func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "toy-server",
		Version: "v0.1.0",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Infof("version: %+v", cmd.Version)

			server, stop := server.Init()
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}

			<-stop
			log.Info("app stopped")
			return nil
		},
	}

	return cmd
}
