package organizationsettinghistory

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/theopenlane/core/cmd/cli/cmd"
	"github.com/theopenlane/core/pkg/openlaneclient"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing organizationSettingHistory",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	command.AddCommand(getCmd)
	getCmd.Flags().StringP("id", "i", "", "id to query")
}

// get an existing organizationSettingHistory in the platform
func get(ctx context.Context) error {
	// setup http client
	client, err := cmd.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer cmd.StoreSessionCookies(client)

	// filter options
	id := cmd.Config.String("id")
	if id != "" {
		o, err := client.GetOrganizationSettingHistories(ctx, &openlaneclient.OrganizationSettingHistoryWhereInput{
			Ref: &id,
		})
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	// get all will be filtered for the authorized organization(s)
	o, err := client.GetAllOrganizationSettingHistories(ctx)
	cobra.CheckErr(err)

	return consoleOutput(o)
}