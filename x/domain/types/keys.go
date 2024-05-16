package types

const (
	// ModuleName defines the module name
	ModuleName = "domain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_domain"
)

var (
	ParamsKey = []byte("p_domain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	DomainKey      = "Domain/value/"
	DomainCountKey = "Domain/count/"
	RewardKey      = "Reward/value/"
)
