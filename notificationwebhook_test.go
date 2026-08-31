// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg_test

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"github.com/swagger-girly/go"
	"github.com/swagger-girly/go/option"
)

func TestNotificationWebhookParsePet(t *testing.T) {
	client := helloworldtestingggg.NewClient(
		option.WithWebhookSecret("whsec_c2VjcmV0Cg=="),
		option.WithAPIKey("My API Key"),
		option.WithBasicAuthUsername("My Basic Auth Username"),
		option.WithBasicAuthPassword("My Basic Auth Password"),
	)
	payload := []byte(`{"pet":{"name":"doggie","photoUrls":["string"],"id":10,"acquisitionChannel":"breeder","category":{},"microchipId":"string","relatedAddress":{"city":"Palo Alto","geo":{"latitude":37.4443,"longitude":-122.1598},"relatedCategory":{},"relatedCustomer":{},"relatedMoney":{"amount":2500,"currency":"USD","relatedCategory":{},"relatedCustomer":{},"relatedOrder":{},"relatedShelter":{},"relatedTag":{},"relatedUser":{}},"relatedOrder":{},"relatedShelter":{},"relatedTag":{},"relatedUser":{},"state":"CA","street":"437 Lytton","zip":"94301"},"relatedCategory":{},"relatedCustomer":{},"relatedMoney":{"amount":2500,"currency":"USD","relatedCategory":{},"relatedCustomer":{},"relatedOrder":{},"relatedShelter":{},"relatedTag":{},"relatedUser":{}},"relatedOrder":{},"relatedShelter":{},"relatedUser":{},"status":"available","tags":[{}]},"type":"pet.created"}`)
	wh, err := standardwebhooks.NewWebhook("whsec_c2VjcmV0Cg==")
	if err != nil {
		t.Fatal("Failed to sign test webhook message", err)
	}
	msgID := "1"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Fatal("Failed to sign test webhook message:", err)
	}
	headers := make(http.Header)
	headers.Set("webhook-signature", sig)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	_, err = client.Notifications.Webhooks.ParsePet(payload, headers)
	if err != nil {
		t.Fatal("Failed to unwrap webhook:", err)
	}
}
