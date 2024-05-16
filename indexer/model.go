package indexer

type DomainIndexer struct {
	Domain        string            `json:"domain"`
	Owner         string            `json:"owner"`
	StringRecords map[string]string `json:"stringRecords"`
}
