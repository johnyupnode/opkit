package indexer

import (
	"cosmossdk.io/log"
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Api interface {
	GetDomainInfo(ctx sdk.Context, domain string) (DomainIndexer, error)
	GetDomainOwner(ctx sdk.Context, address string) ([]DomainIndexer, error)
	GetDomainStringRecord(ctx sdk.Context, key, value string) ([]DomainIndexer, error)
}

type apiHandle struct {
	baseUrl string
	logger  log.Logger
}

func NewApiHandle(baseUrl string, logger log.Logger) Api {
	return &apiHandle{
		baseUrl: baseUrl,
		logger:  logger,
	}
}

func (a *apiHandle) GetDomainInfo(ctx sdk.Context, domain string) (DomainIndexer, error) {
	time := ctx.BlockTime().UTC().Unix()
	url := a.baseUrl + "/domain/" + domain + "?timestamp=" + strconv.FormatInt(time, 10)
	data, err := a.doHttp(url)
	if err != nil {
		return DomainIndexer{}, err
	}
	return a.parseDomain(data)
}

func (a *apiHandle) GetDomainOwner(ctx sdk.Context, address string) ([]DomainIndexer, error) {
	time := ctx.BlockTime().UTC().Unix()
	url := a.baseUrl + "/domains/" + address + "?timestamp=" + strconv.FormatInt(time, 10)
	data, err := a.doHttp(url)
	if err != nil {
		a.logger.Error("[GetDomainOwner] do http", "err", err.Error())
		return nil, err
	}

	return a.parseDomains(data)
}

func (a *apiHandle) GetDomainStringRecord(ctx sdk.Context, key, value string) ([]DomainIndexer, error) {
	time := ctx.BlockTime().UTC().Unix()
	url := a.baseUrl + "/string-records/" + key + "/" + value + "?timestamp=" + strconv.FormatInt(time, 10)
	data, err := a.doHttp(url)
	if err != nil {
		a.logger.Error("[GetDomainStringRecord] do http", "err", err.Error())
		return nil, err
	}

	return a.parseDomains(data)
}

func (a *apiHandle) doHttp(url string) ([]byte, error) {
	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		a.logger.Error("Error new http request", "err", err.Error())
		return nil, err
	}

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		a.logger.Error("Error do http request", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	return ioutil.ReadAll(resp.Body)
}

func (a *apiHandle) parseDomains(data []byte) ([]DomainIndexer, error) {
	var domains []DomainIndexer
	err := json.Unmarshal(data, &domains)
	if err != nil {
		a.logger.Error("Error parsing JSON", "err", err.Error())
		return nil, err
	}

	return domains, nil
}

func (a *apiHandle) parseDomain(data []byte) (DomainIndexer, error) {
	var domain DomainIndexer
	err := json.Unmarshal(data, &domain)
	if err != nil {
		a.logger.Error("Error parsing JSON", "err", err.Error())
		return DomainIndexer{}, err
	}

	return domain, nil
}
