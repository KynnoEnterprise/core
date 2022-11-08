package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidCollection = sdkerrors.Register(ModuleName, 20, "invalid nft collection")
	ErrUnknownCollection = sdkerrors.Register(ModuleName, 21, "unknown nft collection")
	ErrInvalidNFT        = sdkerrors.Register(ModuleName, 22, "invalid nft")
	ErrNFTAlreadyExists  = sdkerrors.Register(ModuleName, 23, "nft already exists")
	ErrUnknownNFT        = sdkerrors.Register(ModuleName, 24, "unknown nft")
	ErrEmptyTokenData    = sdkerrors.Register(ModuleName, 25, "nft data can't be empty")
	ErrUnauthorized      = sdkerrors.Register(ModuleName, 26, "unauthorized address")
	ErrInvalidDenom      = sdkerrors.Register(ModuleName, 27, "invalid denom")
	ErrInvalidTokenID    = sdkerrors.Register(ModuleName, 28, "invalid nft id")
	ErrInvalidTokenURI   = sdkerrors.Register(ModuleName, 29, "invalid nft uri")
)
