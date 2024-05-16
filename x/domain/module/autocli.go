package domain

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "opkit/api/opkit/domain"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "DomainAll",
					Use:       "list-domain",
					Short:     "List all domain",
				},
				{
					RpcMethod:      "Domain",
					Use:            "show-domain [id]",
					Short:          "Shows a domain by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "ListDomainOpkit",
					Use:            "list-domain-opkit [address]",
					Short:          "Query list-domain-opkit",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},

				{
					RpcMethod:      "ListDomainEvm",
					Use:            "list-domain-evm [address]",
					Short:          "Query list-domain-evm",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},

				{
					RpcMethod:      "ListDomainByString",
					Use:            "list-domain-by-string [key] [value]",
					Short:          "Query list-domain-by-string",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "key"}, {ProtoField: "value"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateDomain",
					Use:            "create-domain [domain] [owner] [timestamp] [txhash]",
					Short:          "Create domain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "domain"}, {ProtoField: "owner"}, {ProtoField: "timestamp"}, {ProtoField: "txhash"}},
				},
				{
					RpcMethod:      "UpdateDomain",
					Use:            "update-domain [id] [domain] [owner] [timestamp] [txhash]",
					Short:          "Update domain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "domain"}, {ProtoField: "owner"}, {ProtoField: "timestamp"}, {ProtoField: "txhash"}},
				},
				{
					RpcMethod:      "DeleteDomain",
					Use:            "delete-domain [id]",
					Short:          "Delete domain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
