package keeper

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	ethermint "github.com/evmos/ethermint/types"

	"github.com/kynnoenterprise/core/x/rewardshare/types"
)

var _ types.QueryServer = Keeper{}

// Rewardshares returns all Rewardshares that have been registered for fee distribution
func (k Keeper) Rewardshares(
	c context.Context,
	req *types.QueryRewardsharesRequest,
) (*types.QueryRewardsharesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var rewardshares []types.Rewardshare
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixRewardshare)

	pageRes, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var rewardshare types.Rewardshare
		if err := k.cdc.Unmarshal(value, &rewardshare); err != nil {
			return err
		}
		rewardshares = append(rewardshares, rewardshare)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryRewardsharesResponse{
		Rewardshares:   rewardshares,
		Pagination: pageRes,
	}, nil
}

// Rewardshare returns the Rewardshare that has been registered for fee distribution for a given
// contract
func (k Keeper) Rewardshare(
	c context.Context,
	req *types.QueryRewardshareRequest,
) (*types.QueryRewardshareResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if strings.TrimSpace(req.ContractAddress) == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"contract address is empty",
		)
	}

	// check if the contract is a non-zero hex address
	if err := ethermint.ValidateNonZeroAddress(req.ContractAddress); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid format for contract %s, should be non-zero hex ('0x...')", req.ContractAddress,
		)
	}

	rewardshare, found := k.GetRewardshare(ctx, common.HexToAddress(req.ContractAddress))
	if !found {
		return nil, status.Errorf(
			codes.NotFound,
			"fees registered contract '%s'",
			req.ContractAddress,
		)
	}

	return &types.QueryRewardshareResponse{Rewardshare: rewardshare}, nil
}

// Params returns the fees module params
func (k Keeper) Params(
	c context.Context,
	_ *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}

// DeployerRewardshares returns all contracts that have been registered for fee
// distribution by a given deployer
func (k Keeper) DeployerRewardshares( // nolint: dupl
	c context.Context,
	req *types.QueryDeployerRewardsharesRequest,
) (*types.QueryDeployerRewardsharesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if strings.TrimSpace(req.DeployerAddress) == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"deployer address is empty",
		)
	}

	deployer, err := sdk.AccAddressFromBech32(req.DeployerAddress)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid format for deployer %s, should be bech32 ('kynno...')", req.DeployerAddress,
		)
	}

	var contracts []string
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey),
		types.GetKeyPrefixDeployer(deployer),
	)

	pageRes, err := query.Paginate(store, req.Pagination, func(key, _ []byte) error {
		contracts = append(contracts, common.BytesToAddress(key).Hex())
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDeployerRewardsharesResponse{
		ContractAddresses: contracts,
		Pagination:        pageRes,
	}, nil
}

// WithdrawerRewardshares returns all fees for a given withdraw address
func (k Keeper) WithdrawerRewardshares( // nolint: dupl
	c context.Context,
	req *types.QueryWithdrawerRewardsharesRequest,
) (*types.QueryWithdrawerRewardsharesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if strings.TrimSpace(req.WithdrawerAddress) == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"withdraw address is empty",
		)
	}

	deployer, err := sdk.AccAddressFromBech32(req.WithdrawerAddress)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid format for withdraw addr %s, should be bech32 ('kynno...')", req.WithdrawerAddress,
		)
	}

	var contracts []string
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey),
		types.GetKeyPrefixWithdrawer(deployer),
	)

	pageRes, err := query.Paginate(store, req.Pagination, func(key, _ []byte) error {
		contracts = append(contracts, common.BytesToAddress(key).Hex())

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryWithdrawerRewardsharesResponse{
		ContractAddresses: contracts,
		Pagination:        pageRes,
	}, nil
}
