package cli

import (
	"coslms/x/std/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	lmsQueryCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Querying commands for the lms module",
		Long:  ``,
		RunE:  client.ValidateCmd,
	}

	lmsQueryCmd.AddCommand(
		GetStudent(),
	)
	return lmsQueryCmd
}

func GetStudent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getstudents",
		Short: "| student Address | student name |",
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.QueryGetStudent(cmd.Context(), &types.GetStudentsRequest{})
			if err != nil {
				return err
			}
			// panic("called")

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
