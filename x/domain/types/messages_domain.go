package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDomain{}

func NewMsgCreateDomain(creator string, domain string, owner string, timestamp string, txhash string) *MsgCreateDomain {
	return &MsgCreateDomain{
		Creator:   creator,
		Domain:    domain,
		Owner:     owner,
		Timestamp: timestamp,
		Txhash:    txhash,
	}
}

func (msg *MsgCreateDomain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDomain{}

func NewMsgUpdateDomain(creator string, domain string, owner string, timestamp string, txhash string) *MsgUpdateDomain {
	return &MsgUpdateDomain{
		Creator:   creator,
		Domain:    domain,
		Owner:     owner,
		Timestamp: timestamp,
		Txhash:    txhash,
	}
}

func (msg *MsgUpdateDomain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDomain{}

func NewMsgDeleteDomain(creator string, domain string) *MsgDeleteDomain {
	return &MsgDeleteDomain{
		Domain:  domain,
		Creator: creator,
	}
}

func (msg *MsgDeleteDomain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
