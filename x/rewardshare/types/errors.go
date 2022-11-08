package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// errors
var (
	ErrInternalRewardshare              = sdkerrors.Register(ModuleName, 2, "internal rewardshare error")
	ErrRewardshareDisabled              = sdkerrors.Register(ModuleName, 3, "rewardshare module is disabled by governance")
	ErrRewardshareAlreadyRegistered     = sdkerrors.Register(ModuleName, 4, "rewardshare already exists for given contract")
	ErrRewardshareNoContractDeployed    = sdkerrors.Register(ModuleName, 5, "no contract deployed")
	ErrRewardshareContractNotRegistered = sdkerrors.Register(ModuleName, 6, "no rewardshare registered for contract")
	ErrRewardshareDeployerIsNotEOA      = sdkerrors.Register(ModuleName, 7, "no rewardshare registered for contract")
)
