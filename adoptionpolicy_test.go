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
		Name: helloworldtestingggg.F("name"),
		Rules: helloworldtestingggg.F([]helloworldtestingggg.AdoptionRuleParam{{
			Attribute: helloworldtestingggg.F(helloworldtestingggg.RuleAttributePetStatus),
			Operation: helloworldtestingggg.F(helloworldtestingggg.AdoptionRuleOperationIsOneOf),
			Value:     helloworldtestingggg.F[helloworldtestingggg.ConditionValueUnionParam](shared.UnionString("string")),
		}}),
		Escalation: helloworldtestingggg.F(helloworldtestingggg.AdoptionPolicyNewParamsEscalation{
			Operator: helloworldtestingggg.F(helloworldtestingggg.AdoptionPolicyNewParamsEscalationOperatorAll),
			Groups:   helloworldtestingggg.F([]interface{}{}),
			Rules: helloworldtestingggg.F([]helloworldtestingggg.AdoptionRuleParam{{
				Attribute: helloworldtestingggg.F(helloworldtestingggg.RuleAttributePetStatus),
				Operation: helloworldtestingggg.F(helloworldtestingggg.AdoptionRuleOperationIsOneOf),
				Value:     helloworldtestingggg.F[helloworldtestingggg.ConditionValueUnionParam](shared.UnionString("string")),
			}}),
		}),
		FeeCap: helloworldtestingggg.F(shared.MoneyParam{
			Amount:   helloworldtestingggg.F(int64(2500)),
			Currency: helloworldtestingggg.F("USD"),
			RelatedAddress: helloworldtestingggg.F(shared.AddressParam{
				City: helloworldtestingggg.F("Palo Alto"),
				Geo: helloworldtestingggg.F(shared.AddressGeoParam{
					Latitude:  helloworldtestingggg.F(37.444300),
					Longitude: helloworldtestingggg.F(-122.159800),
				}),
				RelatedCategory: helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedCustomer: helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedOrder:    helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedPet: helloworldtestingggg.F(helloworldtestingggg.PetParam{
					Name:               helloworldtestingggg.F("doggie"),
					PhotoURLs:          helloworldtestingggg.F([]string{"string"}),
					ID:                 helloworldtestingggg.F(int64(10)),
					AcquisitionChannel: helloworldtestingggg.F(helloworldtestingggg.PetAcquisitionChannelBreeder),
					Category:           helloworldtestingggg.F[any](map[string]interface{}{}),
					MicrochipID:        helloworldtestingggg.F[helloworldtestingggg.PetMicrochipIDUnionParam](shared.UnionString("string")),
					RelatedCategory:    helloworldtestingggg.F[any](map[string]interface{}{}),
					RelatedCustomer:    helloworldtestingggg.F[any](map[string]interface{}{}),
					RelatedOrder:       helloworldtestingggg.F[any](map[string]interface{}{}),
					RelatedShelter:     helloworldtestingggg.F[any](map[string]interface{}{}),
					RelatedUser:        helloworldtestingggg.F[any](map[string]interface{}{}),
					Status:             helloworldtestingggg.F(helloworldtestingggg.PetStatusAvailable),
					Tags:               helloworldtestingggg.F([]interface{}{map[string]interface{}{}}),
				}),
				RelatedShelter: helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedTag:     helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedUser:    helloworldtestingggg.F[any](map[string]interface{}{}),
				State:          helloworldtestingggg.F("CA"),
				Street:         helloworldtestingggg.F("437 Lytton"),
				Zip:            helloworldtestingggg.F("94301"),
			}),
			RelatedCategory: helloworldtestingggg.F[any](map[string]interface{}{}),
			RelatedCustomer: helloworldtestingggg.F[any](map[string]interface{}{}),
			RelatedOrder:    helloworldtestingggg.F[any](map[string]interface{}{}),
			RelatedPet: helloworldtestingggg.F(helloworldtestingggg.PetParam{
				Name:               helloworldtestingggg.F("doggie"),
				PhotoURLs:          helloworldtestingggg.F([]string{"string"}),
				ID:                 helloworldtestingggg.F(int64(10)),
				AcquisitionChannel: helloworldtestingggg.F(helloworldtestingggg.PetAcquisitionChannelBreeder),
				Category:           helloworldtestingggg.F[any](map[string]interface{}{}),
				MicrochipID:        helloworldtestingggg.F[helloworldtestingggg.PetMicrochipIDUnionParam](shared.UnionString("string")),
				RelatedCategory:    helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedCustomer:    helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedOrder:       helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedShelter:     helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedUser:        helloworldtestingggg.F[any](map[string]interface{}{}),
				Status:             helloworldtestingggg.F(helloworldtestingggg.PetStatusAvailable),
				Tags:               helloworldtestingggg.F([]interface{}{map[string]interface{}{}}),
			}),
			RelatedShelter: helloworldtestingggg.F[any](map[string]interface{}{}),
			RelatedTag:     helloworldtestingggg.F[any](map[string]interface{}{}),
			RelatedUser:    helloworldtestingggg.F[any](map[string]interface{}{}),
		}),
		NotificationTarget: helloworldtestingggg.F[helloworldtestingggg.AdoptionPolicyNewParamsNotificationTargetUnion](shared.UnionString("string")),
		RegionalOverrides: helloworldtestingggg.F(map[string]helloworldtestingggg.ConditionValueUnionParam{
			"foo": shared.UnionString("string"),
		}),
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
			Body: helloworldtestingggg.AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdate{
				Escalation: helloworldtestingggg.F(helloworldtestingggg.AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalation{
					Operator: helloworldtestingggg.F(helloworldtestingggg.AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateEscalationOperatorAll),
					Groups:   helloworldtestingggg.F([]interface{}{}),
					Rules: helloworldtestingggg.F([]helloworldtestingggg.AdoptionRuleParam{{
						Attribute: helloworldtestingggg.F(helloworldtestingggg.RuleAttributePetStatus),
						Operation: helloworldtestingggg.F(helloworldtestingggg.AdoptionRuleOperationIsOneOf),
						Value:     helloworldtestingggg.F[helloworldtestingggg.ConditionValueUnionParam](shared.UnionString("string")),
					}}),
				}),
				Name:               helloworldtestingggg.F("name"),
				NotificationTarget: helloworldtestingggg.F[helloworldtestingggg.AdoptionPolicyUpdateParamsBodyAdoptionPolicyDetailsUpdateNotificationTargetUnion](shared.UnionString("string")),
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
		Cursor: helloworldtestingggg.F("cursor"),
		Limit:  helloworldtestingggg.F(int64(0)),
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
