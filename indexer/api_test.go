package indexer

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	jsonStr := `[{"domain":"test1.opkit","owner":"0xf01Dd015Bc442d872275A79b9caE84A6ff9B2A27","stringRecords":{"opkit":"opkit14f2t5kwrl9fsw9dy58egusm3wplynyv8wgvrvt","twitter":"Chomtana", "github":"Chomtana"}}]`

	var records []DomainIndexer
	err := json.Unmarshal([]byte(jsonStr), &records)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	for _, record := range records {
		fmt.Printf("Domain: %s\nOwner: %s\n", record.Domain, record.Owner)
		fmt.Printf("String Records:\n")
		for k, v := range record.StringRecords {
			fmt.Printf("\t%s: %s\n", k, v)
		}
		fmt.Println()
	}
}
