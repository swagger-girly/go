// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/swagger-girly/go"
	"github.com/swagger-girly/go/internal/testutil"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/shared"
)

func TestPlacementNewWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Placements.New(context.TODO(), helloworldtestingggg.PlacementNewParams{
		ApplicationID: "applicationId",
		InitialEvent: helloworldtestingggg.PlacementEventUnionParam{
			OfPlacementEventPlacementTransferEvent: &helloworldtestingggg.PlacementEventPlacementTransferEventParam{
				ID: "id",
				Leg: helloworldtestingggg.TransferLegParam{
					Location: shared.AddressParam{
						City: helloworldtestingggg.String("Palo Alto"),
						Geo: shared.AddressGeoParam{
							Latitude:  37.4443,
							Longitude: -122.1598,
						},
						State:  helloworldtestingggg.String("CA"),
						Street: helloworldtestingggg.String("437 Lytton"),
						Zip:    helloworldtestingggg.String("94301"),
					},
					Contact: helloworldtestingggg.TransferLegContactParam{
						Name:  helloworldtestingggg.String("name"),
						Phone: helloworldtestingggg.String("phone"),
					},
					Cost: shared.MoneyParam{
						Amount:   2500,
						Currency: "USD",
					},
					Window: helloworldtestingggg.TransferLegWindowParam{
						End:   helloworldtestingggg.Time(time.Now()),
						Start: helloworldtestingggg.Time(time.Now()),
					},
				},
				OccurredAt: time.Now(),
				Type:       "transfer",
				Note:       helloworldtestingggg.String("note"),
			},
		},
		Logistics: helloworldtestingggg.PlacementNewParamsLogistics{
			Delivery: helloworldtestingggg.TransferLegParam{
				Location: shared.AddressParam{
					City: helloworldtestingggg.String("Palo Alto"),
					Geo: shared.AddressGeoParam{
						Latitude:  37.4443,
						Longitude: -122.1598,
					},
					State:  helloworldtestingggg.String("CA"),
					Street: helloworldtestingggg.String("437 Lytton"),
					Zip:    helloworldtestingggg.String("94301"),
				},
				Contact: helloworldtestingggg.TransferLegContactParam{
					Name:  helloworldtestingggg.String("name"),
					Phone: helloworldtestingggg.String("phone"),
				},
				Cost: shared.MoneyParam{
					Amount:   2500,
					Currency: "USD",
				},
				Window: helloworldtestingggg.TransferLegWindowParam{
					End:   helloworldtestingggg.Time(time.Now()),
					Start: helloworldtestingggg.Time(time.Now()),
				},
			},
			Handler: helloworldtestingggg.PlacementNewParamsLogisticsHandler{
				Certification: helloworldtestingggg.PlacementNewParamsLogisticsHandlerCertification{
					Level: "standard",
					Audit: helloworldtestingggg.PlacementNewParamsLogisticsHandlerCertificationAudit{
						RecordedAt: time.Now(),
						RecordedBy: helloworldtestingggg.String("recordedBy"),
						Source:     "manual",
					},
					IssuedBy: helloworldtestingggg.String("issuedBy"),
				},
				Name: helloworldtestingggg.String("name"),
			},
			Pickup: helloworldtestingggg.TransferLegParam{
				Location: shared.AddressParam{
					City: helloworldtestingggg.String("Palo Alto"),
					Geo: shared.AddressGeoParam{
						Latitude:  37.4443,
						Longitude: -122.1598,
					},
					State:  helloworldtestingggg.String("CA"),
					Street: helloworldtestingggg.String("437 Lytton"),
					Zip:    helloworldtestingggg.String("94301"),
				},
				Contact: helloworldtestingggg.TransferLegContactParam{
					Name:  helloworldtestingggg.String("name"),
					Phone: helloworldtestingggg.String("phone"),
				},
				Cost: shared.MoneyParam{
					Amount:   2500,
					Currency: "USD",
				},
				Window: helloworldtestingggg.TransferLegWindowParam{
					End:   helloworldtestingggg.Time(time.Now()),
					Start: helloworldtestingggg.Time(time.Now()),
				},
			},
		},
		ScheduledFor: helloworldtestingggg.Time(time.Now()),
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlacementGet(t *testing.T) {
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
	)
	_, err := client.Placements.Get(context.TODO(), "placementId")
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlacementListWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Placements.List(context.TODO(), helloworldtestingggg.PlacementListParams{
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

func TestPlacementRecordEventWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Placements.RecordEvent(
		context.TODO(),
		"placementId",
		helloworldtestingggg.PlacementRecordEventParams{
			PlacementEvent: helloworldtestingggg.PlacementEventUnionParam{
				OfPlacementEventPlacementTransferEvent: &helloworldtestingggg.PlacementEventPlacementTransferEventParam{
					ID: "id",
					Leg: helloworldtestingggg.TransferLegParam{
						Location: shared.AddressParam{
							City: helloworldtestingggg.String("Palo Alto"),
							Geo: shared.AddressGeoParam{
								Latitude:  37.4443,
								Longitude: -122.1598,
							},
							State:  helloworldtestingggg.String("CA"),
							Street: helloworldtestingggg.String("437 Lytton"),
							Zip:    helloworldtestingggg.String("94301"),
						},
						Contact: helloworldtestingggg.TransferLegContactParam{
							Name:  helloworldtestingggg.String("name"),
							Phone: helloworldtestingggg.String("phone"),
						},
						Cost: shared.MoneyParam{
							Amount:   2500,
							Currency: "USD",
						},
						Window: helloworldtestingggg.TransferLegWindowParam{
							End:   helloworldtestingggg.Time(time.Now()),
							Start: helloworldtestingggg.Time(time.Now()),
						},
					},
					OccurredAt: time.Now(),
					Type:       "transfer",
					Note:       helloworldtestingggg.String("note"),
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
