package cli

import (
	"coslms/x/std/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/cosmos/cosmos-sdk/x/staking/types"
)

// default values
// var (
// 	DefaultTokens                  = sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction)
// 	defaultAmount                  = DefaultTokens.String() + sdk.DefaultBondDenom
// 	defaultCommissionRate          = "0.1"
// 	defaultCommissionMaxRate       = "0.2"
// 	defaultCommissionMaxChangeRate = "0.01"
// 	defaultMinSelfDelegation       = "1"
// )

// NewTxCmd returns a root CLI command handler for all x/staking transaction commands.

func NewTxCmd() *cobra.Command {
	stdTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Leave management system",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	stdTxCmd.AddCommand(
		AddStudentCmd(),
		RegisterAdminCmd(),
		ApplyLeaveReqCmd(),
		AcceptLeaveReqCmd(),
	)

	return stdTxCmd
}
func RegisterAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registerAdmin",
		Short: "registerAdmin -name- -address-",
		Long:  `to register admin`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			admin := types.RegisterAdminRequest{
				Address: args[0],
				Name:    args[1],
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &admin)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func AddStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "addstud",
		Short: "addstud -address-  -Name- -id-",
		Long:  `adds student`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			adminAddress := args[0]
			students := []*types.Student{}
			for i := 0; i < (len(args)-1)/3; i++ {
				student := &types.Student{
					Address: args[3*i+1],
					Name:    args[3*i+2],
					Id:      args[3*i+3],
				}
				students = append(students, student)
			}
			AddStudents := types.AddStudentRequest{
				Admin:    adminAddress,
				Students: students,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &AddStudents)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func ApplyLeaveReqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "applyleave",
		Short: "applyleave -Address- -Reason-",
		Long:  "To apply for a leave",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			Address := args[0]
			Reason := args[1]
			msg := types.ApplyLeaveReq(sdk.AccAddress(Address), Reason)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func AcceptLeaveReqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acceptleave ",
		Short: "acceptleave -Admin- -LeaveId- -Status-",
		Long:  "To accept a leave request",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCxt, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			Admin := args[0]
			LeaveId := args[1]
			Status := args[2]
			msg := types.NewAcceptLeaveReq(sdk.AccAddress(Admin), LeaveId, Status)
			return tx.GenerateOrBroadcastTxCLI(clientCxt, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

/*
 address: cosmos1k353axl2gjevx6u6rc8e0gd3sg2uhhyymnhj9r
  name: validator-key
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A1I7Jz0SiL1qxxyBLD3TlirQmyak9XkZ8T/DyXIxyGeD"}'
  type: local

  address: cosmos17kpv6tgf2kh0x2rqg07h6tn9tavamy6ehn5gv4
  name: mykey1
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A/xW9Ml1eUYLsmTKYMIZtoGo5mGrM2dcghgLeVytAuT0"}'
  type: local

  address: cosmos13ehauvmvq07hrejgasl6gqjp8ze9525yhp4r8c
  name: mykey2
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A5cd1KttR1UNjbT30mJag9txFK+IZL9/ETtE0BSkB7eF"}'
  type: local

*/
