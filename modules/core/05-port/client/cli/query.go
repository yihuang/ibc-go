package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmos/ibc-go/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/modules/core/24-host"
)

// GetCmdQueryPort defines the command to query...todo
func GetCmdQueryPort() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "channels",
		Short:   "Query all channels",
		Long:    "Query all channels from a chain",
		Example: fmt.Sprintf("%s query %s %s channels", version.AppName, host.ModuleName, types.SubModuleName),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// clientCtx, err := client.GetClientQueryContext(cmd)
			// if err != nil {
			// 	return err
			// }
			// queryClient := types.NewQueryClient(clientCtx)

			// pageReq, err := client.ReadPageRequest(cmd.Flags())
			// if err != nil {
			// 	return err
			// }

			// req := &types.QueryChannelsRequest{
			// 	Pagination: pageReq,
			// }

			// res, err := queryClient.Channels(cmd.Context(), req)
			// if err != nil {
			// 	return err
			// }

			// return clientCtx.PrintProto(res)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
