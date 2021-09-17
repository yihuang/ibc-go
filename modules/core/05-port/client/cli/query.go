package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmos/ibc-go/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/modules/core/24-host"
)

// GetCmdQueryPort defines the command to query an IBC application associated with a port
func GetCmdQueryPort() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "port [port-id]",
		Short:   "Query an IBC application port",
		Long:    "Query an IBC application associated with a port, by providing it's port identifier",
		Example: fmt.Sprintf("%s query %s %s port [port-id]", version.AppName, host.ModuleName, types.SubModuleName),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryPortRequest{
				PortId: args[0],
			}

			res, err := queryClient.Port(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
