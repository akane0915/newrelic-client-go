package cloud

import (
	"log"
	"os"

	"github.com/newrelic/newrelic-client-go/pkg/config"
	mock "github.com/newrelic/newrelic-client-go/pkg/testhelpers"
)

func Example_cloudAccounts() {
	// Initialize the client configuration.  A Personal API key is required to
	// communicate with the backend API.
	cfg := config.New()
	cfg.PersonalAPIKey = os.Getenv("NEW_RELIC_API_KEY")

	// Initialize the client.
	client := New(cfg)

	// Get the linked cloud accounts
	linkedAccounts, err := client.GetLinkedAccounts("aws")
	if err != nil {
		log.Fatal("error retrieving linked accounts:", err)
	}

	log.Printf("linked accounts count: %d", len(*linkedAccounts))

	// Link a cloud account
	linkResponse, err := client.CloudLinkAccount(mock.TestAccountID, CloudLinkCloudAccountsInput{
		Aws: []CloudAwsLinkAccountInput{
			{
				Arn:  "arn:aws:iam::12345678:role/MyAWSARN",
				Name: "My Linked AWS Account",
			},
		},
	})
	if err != nil || len(linkResponse.LinkedAccounts) != 1 {
		log.Fatal("error linking cloud account:", err)
	}

	linkedAccountID := linkResponse.LinkedAccounts[0].ID

	// Rename a linked account
	_, err = client.CloudRenameAccount(mock.TestAccountID, []CloudRenameAccountsInput{
		{
			LinkedAccountId: linkedAccountID,
			Name:            "My Renamed Linked AWS Account",
		},
	})
	if err != nil {
		log.Fatal("error renaming linked cloud account:", err)
	}

	// Unlink a linked account
	_, err = client.CloudUnlinkAccount(mock.TestAccountID, []CloudUnlinkAccountsInput{{linkedAccountID}})
	if err != nil {
		log.Fatal("error unlinking linked cloud account:", err)
	}
}
