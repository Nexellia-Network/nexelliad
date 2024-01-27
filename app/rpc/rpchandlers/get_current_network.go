package rpchandlers

import (
	"github.com/Nexellia-Network/nexelliad/app/appmessage"
	"github.com/Nexellia-Network/nexelliad/app/rpc/rpccontext"
	"github.com/Nexellia-Network/nexelliad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
