package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

var _ = strconv.Itoa(0)

func CmdGetValidatorConsAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-validator-cons-address [validator-operator-address]",
		Short: "Query getValidatorConsAddress",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqValidatorOperatorAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetValidatorConsAddressRequest{

				ValidatorOperatorAddress: reqValidatorOperatorAddress,
			}

			res, err := queryClient.GetValidatorConsAddress(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
