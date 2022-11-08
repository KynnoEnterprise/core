package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethermint "github.com/evmos/ethermint/types"
)

// NewRewardshare returns an instance of Rewardshare. If the provided withdrawer
// address is empty, it sets the value to an empty string.
func NewRewardshare(contract common.Address, deployer, withdrawer sdk.AccAddress) Rewardshare {
	withdrawerAddr := ""
	if len(withdrawer) > 0 {
		withdrawerAddr = withdrawer.String()
	}

	return Rewardshare{
		ContractAddress:   contract.String(),
		DeployerAddress:   deployer.String(),
		WithdrawerAddress: withdrawerAddr,
	}
}

// GetContractAddr returns the contract address
func (fs Rewardshare) GetContractAddr() common.Address {
	return common.HexToAddress(fs.ContractAddress)
}

// GetDeployerAddr returns the contract deployer address
func (fs Rewardshare) GetDeployerAddr() sdk.AccAddress {
	return sdk.MustAccAddressFromBech32(fs.DeployerAddress)
}

// GetWithdrawerAddr returns the account address to where the funds proceeding
// from the fees will be received. If the withdraw address is not defined, it
// defaults to the deployer address.
func (fs Rewardshare) GetWithdrawerAddr() sdk.AccAddress {
	if fs.WithdrawerAddress == "" {
		return nil
	}

	return sdk.MustAccAddressFromBech32(fs.WithdrawerAddress)
}

// Validate performs a stateless validation of a Rewardshare
func (fs Rewardshare) Validate() error {
	if err := ethermint.ValidateNonZeroAddress(fs.ContractAddress); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(fs.DeployerAddress); err != nil {
		return err
	}

	if fs.WithdrawerAddress != "" {
		if _, err := sdk.AccAddressFromBech32(fs.WithdrawerAddress); err != nil {
			return err
		}
	}

	return nil
}
