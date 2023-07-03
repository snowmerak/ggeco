package storage

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"net/url"
	"strings"
	"time"
)

// jetti:bean Client
type Client struct {
	baseClient  *azblob.Client
	accountName string
	accountKey  string
	host        string
}

func New(accountName string, accountKey string) (*Client, error) {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return nil, err
	}
	host := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)
	client, err := azblob.NewClientWithSharedKeyCredential(host, credential, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient:  client,
		accountName: accountName,
		accountKey:  accountKey,
		host:        host,
	}, nil
}

func GetSASURL(container Container, blobContainer string, fileName string) (urlValue string, err error) {
	client, err := GetClient(container)
	if err != nil {
		return "", err
	}

	rawUrl, err := client.baseClient.ServiceClient().GetSASURL(sas.AccountResourceTypes{
		Container: true,
		Object:    true,
	}, sas.AccountPermissions{
		Read:  true,
		Write: true,
	}, time.Now().Add(1*time.Hour), nil)

	if err != nil {
		return
	}

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return
	}

	parsedUrl.Path = fmt.Sprintf("%s/%s/%s", strings.TrimSuffix(parsedUrl.Path, "/"), blobContainer, fileName)

	return parsedUrl.String(), nil
}
