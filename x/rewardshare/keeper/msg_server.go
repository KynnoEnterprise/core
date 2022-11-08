package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/kynnoenterprise/core/x/rewardshare/types"
)

var _ types.MsgServer = &Keeper{}

// RegisterRewardshare registers a contract to receive transaction fees
func (k Keeper) RegisterRewardshare(
	goCtx context.Context,
	msg *types.MsgRegisterRewardshare,
) (*types.MsgRegisterRewardshareResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if !params.EnableRewardshare {
		return nil, types.ErrRewardshareDisabled
	}

	contract := common.HexToAddress(msg.ContractAddress)

	if k.IsRewardshareRegistered(ctx, contract) {
		return nil, sdkerrors.Wrapf(
			types.ErrRewardshareAlreadyRegistered,
			"contract is already registered %s", contract,
		)
	}

	deployer := sdk.MustAccAddressFromBech32(msg.DeployerAddress)
	deployerAccount := k.evmKeeper.GetAccountWithoutBalance(ctx, common.BytesToAddress(deployer))
	if deployerAccount == nil {
		return nil, sdkerrors.Wrapf(
			errortypes.ErrNotFound,
			"deployer account not found %s", msg.DeployerAddress,
		)
	}

	if deployerAccount.IsContract() {
		return nil, sdkerrors.Wrapf(
			types.ErrRewardshareDeployerIsNotEOA,
			"deployer cannot be a contract %s", msg.DeployerAddress,
		)
	}

	// contract must already be deployed, to avoid spam registrations
	contractAccount := k.evmKeeper.GetAccountWithoutBalance(ctx, contract)

	if contractAccount == nil || !contractAccount.IsContract() {
		return nil, sdkerrors.Wrapf(
			types.ErrRewardshareNoContractDeployed,
			"no contract code found at address %s", msg.ContractAddress,
		)
	}

	var withdrawer sdk.AccAddress
	if msg.WithdrawerAddress != "" && msg.WithdrawerAddress != msg.DeployerAddress {
		withdrawer = sdk.MustAccAddressFromBech32(msg.WithdrawerAddress)
	}

	derivedContract := common.BytesToAddress(deployer)

	// the contract can be directly deployed by an EOA or created through one
	// or more factory contracts. If it was deployed by an EOA account, then
	// msg.Nonces contains the EOA nonce for the deployment transaction.
	// If it was deployed by one or more factories, msg.Nonces contains the EOA
	// nonce for the origin factory contract, then the nonce of the factory
	// for the creation of the next factory/contract.
	for _, nonce := range msg.Nonces {
		ctx.GasMeter().ConsumeGas(
			params.AddrDerivationCostCreate,
			"rewardshare registration: address derivation CREATE opcode",
		)

		derivedContract = crypto.CreateAddress(derivedContract, nonce)
	}

	if contract != derivedContract {
		return nil, sdkerrors.Wrapf(
			errortypes.ErrorInvalidSigner,
			"not contract deployer or wrong nonce: expected %s instead of %s",
			derivedContract, msg.ContractAddress,
		)
	}

	// prevent storing the same address for deployer and withdrawer
	rewardshare := types.NewRewardshare(contract, deployer, withdrawer)
	k.SetRewardshare(ctx, rewardshare)
	k.SetDeployerMap(ctx, deployer, contract)

	// The effective withdrawer is the withdraw address that is stored after the
	// rewardshare registration is completed. It defaults to the deployer address if
	// the withdraw address in the msg is omitted. When omitted, the withdraw map
	// dosn't need to be set.
	effectiveWithdrawer := msg.DeployerAddress

	if len(withdrawer) != 0 {
		k.SetWithdrawerMap(ctx, withdrawer, contract)
		effectiveWithdrawer = msg.WithdrawerAddress
	}

	k.Logger(ctx).Debug(
		"registering contract for transaction fees",
		"contract", msg.ContractAddress, "deployer", msg.DeployerAddress,
		"withdraw", effectiveWithdrawer,
	)

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeRegisterRewardshare,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.DeployerAddress),
				sdk.NewAttribute(types.AttributeKeyContract, msg.ContractAddress),
				sdk.NewAttribute(types.AttributeKeyWithdrawerAddress, effectiveWithdrawer),
			),
		},
	)

	return &types.MsgRegisterRewardshareResponse{}, nil
}

// UpdateRewardshare updates the withdraw address of a given Rewardshare. If the given
// withdraw address is empty or the same as the deployer address, the withdraw
// address is removed.
func (k Keeper) UpdateRewardshare(
	goCtx context.Context,
	msg *types.MsgUpdateRewardshare,
) (*types.MsgUpdateRewardshareResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if !params.EnableRewardshare {
		return nil, types.ErrRewardshareDisabled
	}

	contract := common.HexToAddress(msg.ContractAddress)
	rewardshare, found := k.GetRewardshare(ctx, contract)
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrRewardshareContractNotRegistered,
			"contract %s is not registered", msg.ContractAddress,
		)
	}

	// error if the msg deployer address is not the same as the fee's deployer
	if msg.DeployerAddress != rewardshare.DeployerAddress {
		return nil, sdkerrors.Wrapf(
			errortypes.ErrUnauthorized,
			"%s is not the contract deployer", msg.DeployerAddress,
		)
	}

	// check if updating rewardshare to default withdrawer
	if msg.WithdrawerAddress == rewardshare.DeployerAddress {
		msg.WithdrawerAddress = ""
	}

	// rewardshare with the given withdraw address is already registered
	if msg.WithdrawerAddress == rewardshare.WithdrawerAddress {
		return nil, sdkerrors.Wrapf(
			types.ErrRewardshareAlreadyRegistered,
			"rewardshare with withdraw address %s", msg.WithdrawerAddress,
		)
	}

	// only delete withdrawer map if is not default
	if rewardshare.WithdrawerAddress != "" {
		k.DeleteWithdrawerMap(ctx, sdk.MustAccAddressFromBech32(rewardshare.WithdrawerAddress), contract)
	}

	// only add withdrawer map if new entry is not default
	if msg.WithdrawerAddress != "" {
		k.SetWithdrawerMap(
			ctx,
			sdk.MustAccAddressFromBech32(msg.WithdrawerAddress),
			contract,
		)
	}
	// update rewardshare
	rewardshare.WithdrawerAddress = msg.WithdrawerAddress
	k.SetRewardshare(ctx, rewardshare)

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeUpdateRewardshare,
				sdk.NewAttribute(types.AttributeKeyContract, msg.ContractAddress),
				sdk.NewAttribute(sdk.AttributeKeySender, msg.DeployerAddress),
				sdk.NewAttribute(types.AttributeKeyWithdrawerAddress, msg.WithdrawerAddress),
			),
		},
	)

	return &types.MsgUpdateRewardshareResponse{}, nil
}

// CancelRewardshare deletes the Rewardshare for a given contract
func (k Keeper) CancelRewardshare(
	goCtx context.Context,
	msg *types.MsgCancelRewardshare,
) (*types.MsgCancelRewardshareResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if !params.EnableRewardshare {
		return nil, types.ErrRewardshareDisabled
	}

	contract := common.HexToAddress(msg.ContractAddress)

	fee, found := k.GetRewardshare(ctx, contract)
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrRewardshareContractNotRegistered,
			"contract %s is not registered", msg.ContractAddress,
		)
	}

	if msg.DeployerAddress != fee.DeployerAddress {
		return nil, sdkerrors.Wrapf(
			errortypes.ErrUnauthorized,
			"%s is not the contract deployer", msg.DeployerAddress,
		)
	}

	k.DeleteRewardshare(ctx, fee)
	k.DeleteDeployerMap(
		ctx,
		fee.GetDeployerAddr(),
		contract,
	)

	// delete entry from withdrawer map if not default
	if fee.WithdrawerAddress != "" {
		k.DeleteWithdrawerMap(
			ctx,
			fee.GetWithdrawerAddr(),
			contract,
		)
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeCancelRewardshare,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.DeployerAddress),
				sdk.NewAttribute(types.AttributeKeyContract, msg.ContractAddress),
			),
		},
	)

	return &types.MsgCancelRewardshareResponse{}, nil
}
