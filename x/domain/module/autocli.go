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
					Use:            "get-primary-domain [domain]",
					Short:          "Shows a domain by domain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "domain"}},
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

				{
					RpcMethod:      "Info",
					Use:            "info [domain]",
					Short:          "Query information of specific domain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "domain"}},
				},

				{
					RpcMethod:      "CheckOwner",
					Use:            "check-owner [address]",
					Short:          "Query check-owner",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},

				{
					RpcMethod:      "Reward",
					Use:            "reward [domain]",
					Short:          "Query reward",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "domain"}},
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
					RpcMethod:      "SetPrimaryDomain",
					Use:            "set-primary-domain [domain]",
					Short:          "Send a set-primary-domain tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "domain"}},
				},
				{
					RpcMethod:      "ClaimReward",
					Use:            "claim-reward [domain]",
					Short:          "Send a claim-reward tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "domain"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
