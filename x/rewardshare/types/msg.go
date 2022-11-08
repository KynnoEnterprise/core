package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	ethermint "github.com/evmos/ethermint/types"
)

var (
	_ sdk.Msg = &MsgRegisterRewardshare{}
	_ sdk.Msg = &MsgCancelRewardshare{}
	_ sdk.Msg = &MsgUpdateRewardshare{}
)

const (
	TypeMsgRegisterRewardshare = "register_rewardshare"
	TypeMsgCancelRewardshare   = "cancel_rewardshare"
	TypeMsgUpdateRewardshare   = "update_rewardshare"
)

// NewMsgRegisterRewardshare creates new instance of MsgRegisterRewardshare
func NewMsgRegisterRewardshare(
	contract common.Address,
	deployer,
	withdrawer sdk.AccAddress,
	nonces []uint64,
) *MsgRegisterRewardshare {
	withdrawerAddress := ""
	if withdrawer != nil {
		withdrawerAddress = withdrawer.String()
	}

	return &MsgRegisterRewardshare{
		ContractAddress:   contract.String(),
		DeployerAddress:   deployer.String(),
		WithdrawerAddress: withdrawerAddress,
		Nonces:            nonces,
	}
}

// Route returns the name of the module
func (msg MsgRegisterRewardshare) Route() string { return RouterKey }

// Type returns the the action
func (msg MsgRegisterRewardshare) Type() string { return TypeMsgRegisterRewardshare }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterRewardshare) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.DeployerAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid deployer address %s", msg.DeployerAddress)
	}

	if err := ethermint.ValidateNonZeroAddress(msg.ContractAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid contract address %s", msg.ContractAddress)
	}

	if msg.WithdrawerAddress != "" {
		if _, err := sdk.AccAddressFromBech32(msg.WithdrawerAddress); err != nil {
			return sdkerrors.Wrapf(err, "invalid withdraw address %s", msg.WithdrawerAddress)
		}
	}

	if len(msg.Nonces) < 1 {
		return sdkerrors.Wrapf(errortypes.ErrInvalidRequest, "invalid nonces - empty array")
	}

	if len(msg.Nonces) > 20 {
		return sdkerrors.Wrapf(errortypes.ErrInvalidRequest, "invalid nonces - array length must be less than 20")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgRegisterRewardshare) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterRewardshare) GetSigners() []sdk.AccAddress {
	from := sdk.MustAccAddressFromBech32(msg.DeployerAddress)
	return []sdk.AccAddress{from}
}

// NewMsgCancelRewardshare creates new instance of MsgCancelRewardshare.
func NewMsgCancelRewardshare(
	contract common.Address,
	deployer sdk.AccAddress,
) *MsgCancelRewardshare {
	return &MsgCancelRewardshare{
		ContractAddress: contract.String(),
		DeployerAddress: deployer.String(),
	}
}

// Route returns the message route for a MsgCancelRewardshare.
func (msg MsgCancelRewardshare) Route() string { return RouterKey }

// Type returns the message type for a MsgCancelRewardshare.
func (msg MsgCancelRewardshare) Type() string { return TypeMsgCancelRewardshare }

// ValidateBasic runs stateless checks on the message
func (msg MsgCancelRewardshare) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.DeployerAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid deployer address %s", msg.DeployerAddress)
	}

	if err := ethermint.ValidateNonZeroAddress(msg.ContractAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid contract address %s", msg.ContractAddress)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCancelRewardshare) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCancelRewardshare) GetSigners() []sdk.AccAddress {
	funder := sdk.MustAccAddressFromBech32(msg.DeployerAddress)
	return []sdk.AccAddress{funder}
}

// NewMsgUpdateRewardshare creates new instance of MsgUpdateRewardshare
func NewMsgUpdateRewardshare(
	contract common.Address,
	deployer,
	withdraw sdk.AccAddress,
) *MsgUpdateRewardshare {
	return &MsgUpdateRewardshare{
		ContractAddress:   contract.String(),
		DeployerAddress:   deployer.String(),
		WithdrawerAddress: withdraw.String(),
	}
}

// Route returns the name of the module
func (msg MsgUpdateRewardshare) Route() string { return RouterKey }

// Type returns the the action
func (msg MsgUpdateRewardshare) Type() string { return TypeMsgUpdateRewardshare }

// ValidateBasic runs stateless checks on the message
func (msg MsgUpdateRewardshare) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.DeployerAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid deployer address %s", msg.DeployerAddress)
	}

	if err := ethermint.ValidateNonZeroAddress(msg.ContractAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid contract address %s", msg.ContractAddress)
	}

	if _, err := sdk.AccAddressFromBech32(msg.WithdrawerAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid withdraw address %s", msg.WithdrawerAddress)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgUpdateRewardshare) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpdateRewardshare) GetSigners() []sdk.AccAddress {
	from := sdk.MustAccAddressFromBech32(msg.DeployerAddress)
	return []sdk.AccAddress{from}
}
