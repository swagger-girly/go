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

func TestWebhookParsed(t *testing.T) {
	client := helloworldtestingggg.NewClient(
		option.WithWebhookSecret("whsec_c2VjcmV0Cg=="),
		option.WithAPIKey("My API Key"),
		option.WithBasicAuthUsername("My Basic Auth Username"),
		option.WithBasicAuthPassword("My Basic Auth Password"),
	)
	payload := []byte(`{"pet":{"name":"doggie","photoUrls":["string"],"id":10,"acquisitionChannel":"breeder","category":{"id":1,"name":"Dogs","subcategories":[]},"microchipId":"string","status":"available","tags":[{"id":0,"name":"name"}]},"type":"pet.created"}`)
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
	_, err = client.Webhooks.Parsed(payload, headers)
	if err != nil {
		t.Fatal("Failed to unwrap webhook:", err)
	}
}
