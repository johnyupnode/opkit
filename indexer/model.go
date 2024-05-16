package indexer

type StringRecord struct {
	Opkit   string `json:"opkit"`
	Twitter string `json:"twitter"`
}

type DomainIndexer struct {
	Domain        string       `json:"domain"`
	Owner         string       `json:"owner"`
	StringRecords StringRecord `json:"stringRecords"`
}
