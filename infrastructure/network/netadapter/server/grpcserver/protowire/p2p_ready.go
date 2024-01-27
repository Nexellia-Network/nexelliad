package protowire

import (
	"github.com/Nexellia-Network/nexelliad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *NexelliadMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NexelliadMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *NexelliadMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
