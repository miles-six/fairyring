package keeper

import (
	"context"
	"strconv"

	"fairyring/x/fairblock/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitEncryptedTx(goCtx context.Context, msg *types.MsgSubmitEncryptedTx) (*types.MsgSubmitEncryptedTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.TargetBlockHeight <= uint64(ctx.BlockHeight()) {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.EncryptedTxRevertedEventType,
				sdk.NewAttribute(types.EncryptedTxRevertedEventCreator, msg.Creator),
				sdk.NewAttribute(types.EncryptedTxRevertedEventHeight, strconv.FormatUint(msg.TargetBlockHeight, 10)),
				sdk.NewAttribute(types.EncryptedTxRevertedEventData, msg.Data),
				sdk.NewAttribute(types.EncryptedTxRevertedEventIndex, "0"),
			),
		)

		return nil, types.ErrInvalidTargetBlockHeight
	}

	encryptedTx := types.EncryptedTx{
		TargetHeight: msg.TargetBlockHeight,
		Data:         msg.Data,
		Creator:      msg.Creator,
	}

	txIndex := k.AppendEncryptedTx(ctx, encryptedTx)

	// Emit event after appended ?
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.SubmittedEncryptedTxEventType,
			sdk.NewAttribute(types.SubmittedEncryptedTxEventCreator, msg.Creator),
			sdk.NewAttribute(types.SubmittedEncryptedTxEventTargetHeight, strconv.FormatUint(msg.TargetBlockHeight, 10)),
			sdk.NewAttribute(types.SubmittedEncryptedTxEventData, msg.Data),
			sdk.NewAttribute(types.SubmittedEncryptedTxEventIndex, strconv.FormatUint(txIndex, 10)),
		),
	)

	return &types.MsgSubmitEncryptedTxResponse{}, nil
}
