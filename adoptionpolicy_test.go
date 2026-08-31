// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/swagger-girly/go"
	"github.com/swagger-girly/go/internal/testutil"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/shared"
)

func TestAdoptionPolicyNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := helloworldtestingggg.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBasicAuthUsername("My Basic Auth Username"),
		option.WithBasicAuthPassword("My Basic Auth Password"),
	)
	_, err := client.Adoptions.Policies.New(context.TODO(), helloworldtestingggg.AdoptionPolicyNewParams{
		Name: "name",
		Rules: []helloworldtestingggg.AdoptionRuleParam{{
			Attribute: helloworldtestingggg.RuleAttributePetStatus,
			Operation: helloworldtestingggg.AdoptionRuleOperationIsOneOf,
			Value: helloworldtestingggg.ConditionValueUnionParam{
				OfString: helloworldtestingggg.String("string"),
			},
		}},
		Escalation: helloworldtestingggg.AdoptionPolicyNewParamsEscalation{
			Operator: "ALL",
			Groups:   []any{},
			Rules: []helloworldtestingggg.AdoptionRuleParam{{
				Attribute: helloworldtestingggg.RuleAttributePetStatus,
				Operation: helloworldtestingggg.AdoptionRuleOperationIsOneOf,
				Value: helloworldtestingggg.ConditionValueUnionParam{
					OfString: helloworldtestingggg.String("string"),
				},
			}},
		},
		FeeCap: shared.MoneyParam{
			Amount:   2500,
			Currency: "USD",
			RelatedAddress: shared.AddressParam{
				City: helloworldtestingggg.String("Palo Alto"),
				Geo: shared.AddressGeoParam{
					Latitude:  37.4443,
					Longitude: -122.1598,
				},
				RelatedCategory: map[string]any{},
				RelatedCustomer: map[string]any{},
				RelatedOrder:    map[string]any{},
				RelatedPet: helloworldtestingggg.PetParam{
					Name:               "doggie",
					PhotoURLs:          []string{"string"},
					ID:                 helloworldtestingggg.Int(10),
					AcquisitionChannel: helloworldtestingggg.PetAcquisitionChannelBreeder,
					Category:           map[string]any{},
					MicrochipID: helloworldtestingggg.PetMicrochipIDUnionParam{
						OfString: helloworldtestingggg.String("string"),
					},
					RelatedCategory: map[string]any{},
					RelatedCustomer: map[string]any{},
					RelatedOrder:    map[string]any{},
					RelatedShelter:  map[string]any{},
					RelatedUser:     map[string]any{},
					Status:          helloworldtestingggg.PetStatusAvailable,
					Tags:            []any{map[string]any{}},
				},
				RelatedShelter: map[string]any{},
				RelatedTag:     map[string]any{},
				RelatedUser:    map[string]any{},
				State:          helloworldtestingggg.String("CA"),
				Street:         helloworldtestingggg.String("437 Lytton"),
				Zip:            helloworldtestingggg.String("94301"),
			},
			RelatedCategory: map[string]any{},
			RelatedCustomer: map[string]any{},
			RelatedOrder:    map[string]any{},
			RelatedPet: helloworldtestingggg.PetParam{
				Name:               "doggie",
				PhotoURLs:          []string{"string"},
				ID:                 helloworldtestingggg.Int(10),
				AcquisitionChannel: helloworldtestingggg.PetAcquisitionChannelBreeder,
				Category:           map[string]any{},
				MicrochipID: helloworldtestingggg.PetMicrochipIDUnionParam{
					OfString: helloworldtestingggg.String("string"),
				},
				RelatedCategory: map[string]any{},
				RelatedCustomer: map[string]any{},
				RelatedOrder:    map[string]any{},
				RelatedShelter:  map[string]any{},
				RelatedUser:     map[string]any{},
				Status:          helloworldtestingggg.PetStatusAvailable,
				Tags:            []any{map[string]any{}},
			},
			RelatedShelter: map[string]any{},
			RelatedTag:     map[string]any{},
			RelatedUser:    map[string]any{},
		},
		NotificationTarget: helloworldtestingggg.AdoptionPolicyNewParamsNotificationTargetUnion{
			OfString: helloworldtestingggg.String("string"),
		},
		RegionalOverrides: map[string]helloworldtestingggg.ConditionValueUnionParam{
			"foo": {
				OfString: helloworldtestingggg.String("string"),
			},
		},
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdoptionPolicyGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := helloworldtestingggg.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBasicAuthUsername("My Basic Auth Username"),
		option.WithBasicAuthPassword("My Basic Auth Password"),
	)
	_, err := client.Adoptions.Policies.Get(context.TODO(), "policyId")
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdoptionPolicyUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := helloworldtestingggg.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBasicAuthUsername("My Basic Auth Username"),
		option.WithBasicAuthPassword("My Basic Auth Password"),
	)
	_, err := client.Adoptions.Policies.Update(
		context.TODO(),
		"policyId",
		helloworldtestingggg.AdoptionPolicyUpdateParams{
			OfAdoptionPolicyDetailsUpdate: &helloworldtestingggg.AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdate{
				Escalation: helloworldtestingggg.AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalation{
					Operator: "ALL",
					Groups:   []any{},
					Rules: []helloworldtestingggg.AdoptionRuleParam{{
						Attribute: helloworldtestingggg.RuleAttributePetStatus,
						Operation: helloworldtestingggg.AdoptionRuleOperationIsOneOf,
						Value: helloworldtestingggg.ConditionValueUnionParam{
							OfString: helloworldtestingggg.String("string"),
						},
					}},
				},
				Name: helloworldtestingggg.String("name"),
				NotificationTarget: helloworldtestingggg.AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetUnion{
					OfString: helloworldtestingggg.String("string"),
				},
			},
		},
	)
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdoptionPolicyListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := helloworldtestingggg.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithBasicAuthUsername("My Basic Auth Username"),
		option.WithBasicAuthPassword("My Basic Auth Password"),
	)
	_, err := client.Adoptions.Policies.List(context.TODO(), helloworldtestingggg.AdoptionPolicyListParams{
		Cursor: helloworldtestingggg.String("cursor"),
		Limit:  helloworldtestingggg.Int(0),
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
