package initpluginmanager

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Command(ctx context.Context) *cobra.Command {
	opts := &options{}
	cmd := &cobra.Command{
		Use:   "init-plugin-manager",
		Short: "Install the plugin manager",
		RunE: func(cmd *cobra.Command, args []string) error {
			script, err := downloadScript(ctx)
			if err != nil {
				return err
			}
			defer script.Close()
			installCmd := exec.Command("sh")
			installCmd.Stdin = script
			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr
			if opts.home != "" {
				installCmd.Env = append(installCmd.Env, "MESHCTL_HOME="+opts.home)
			}

			return installCmd.Run()
		},
	}

	opts.addToFlags(cmd.Flags())
	cmd.SilenceUsage = true
	return cmd
}

type options struct {
	home string
}

func (o *options) addToFlags(flags *pflag.FlagSet) {
	flags.StringVar(&o.home, "gm-home", "", "Gloo Mesh home directory (default: $HOME/.gloo-mesh)")
}

func downloadScript(ctx context.Context) (io.ReadCloser, error) {
	const uri = "https://storage.googleapis.com/gloo-mesh-enterprise/meshctl-plugins/install.sh"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			logrus.WithError(err).Debug("could not read response body")
		} else {
			logrus.Debugf("response body: %s", string(b))
		}
		res.Body.Close()

		return nil, eris.Errorf("could not download script: %d %s", res.StatusCode, res.Status)
	}

	return res.Body, nil
}
