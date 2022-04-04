package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/defund-labs/defund/v1/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgInvest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgInvest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgInvest{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgInvest{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
