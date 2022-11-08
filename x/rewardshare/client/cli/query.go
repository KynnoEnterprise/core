package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/kynnoenterprise/core/x/rewardshare/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	feesQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	feesQueryCmd.AddCommand(
		GetCmdQueryRewardshares(),
		GetCmdQueryRewardshare(),
		GetCmdQueryParams(),
		GetCmdQueryDeployerRewardshares(),
		GetCmdQueryWithdrawerRewardshares(),
	)

	return feesQueryCmd
}

// GetCmdQueryRewardshares implements a command to return all registered contracts
// for fee distribution
func GetCmdQueryRewardshares() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contracts",
		Short: "Query all rewardshares",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryRewardsharesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.Rewardshares(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryRewardshare implements a command to return a registered contract for fee
// distribution
func GetCmdQueryRewardshare() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "contract [contract-address]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query a registered contract for fee distribution by hex address",
		Long:    "Query a registered contract for fee distribution by hex address",
		Example: fmt.Sprintf("%s query rewardshare contract <contract-address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryRewardshareRequest{ContractAddress: args[0]}

			// Query store
			res, err := queryClient.Rewardshare(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryParams implements a command to return the current rewardshare
// parameters.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the current rewardshare module parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryParamsRequest{}

			res, err := queryClient.Params(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryDeployerRewardshares implements a command that returns all contracts
// that a deployer has registered for fee distribution
func GetCmdQueryDeployerRewardshares() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deployer-contracts [deployer-address]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query all contracts that a given deployer has registered for fee distribution",
		Long:    "Query all contracts that a given deployer has registered for fee distribution",
		Example: fmt.Sprintf("%s query rewardshare deployer-contracts <deployer-address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			// Query store
			res, err := queryClient.DeployerRewardshares(context.Background(), &types.QueryDeployerRewardsharesRequest{
				DeployerAddress: args[0],
				Pagination:      pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryWithdrawerRewardshares implements a command that returns all
// contracts that have registered for fee distribution with a given withdraw
// address
func GetCmdQueryWithdrawerRewardshares() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "withdrawer-contracts [withdrawer-address]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query all contracts that have been registered for fee distribution with a given withdrawer address",
		Long:    "Query all contracts that have been registered for fee distribution with a given withdrawer address",
		Example: fmt.Sprintf("%s query rewardshare withdrawer-contracts <withdrawer-address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			// Query store
			res, err := queryClient.WithdrawerRewardshares(context.Background(), &types.QueryWithdrawerRewardsharesRequest{
				WithdrawerAddress: args[0],
				Pagination:        pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
