package rpchandlers

import (
	"github.com/Nexellia-Network/nexelliad/app/appmessage"
	"github.com/Nexellia-Network/nexelliad/app/rpc/rpccontext"
	"github.com/Nexellia-Network/nexelliad/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
