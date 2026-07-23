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
		option.WithBasicAuthUsername("My Basic Auth Username"),
		option.WithBasicAuthPassword("My Basic Auth Password"),
	)
	_, err := client.Placements.New(context.TODO(), helloworldtestingggg.PlacementNewParams{
		ApplicationID: helloworldtestingggg.F("applicationId"),
		InitialEvent: helloworldtestingggg.F[helloworldtestingggg.PlacementEventUnionParam](helloworldtestingggg.PlacementEventPlacementTransferEventParam{
			ID: helloworldtestingggg.F("id"),
			Leg: helloworldtestingggg.F(helloworldtestingggg.TransferLegParam{
				Location: helloworldtestingggg.F(shared.AddressParam{
					City: helloworldtestingggg.F("Palo Alto"),
					Geo: helloworldtestingggg.F(shared.AddressGeoParam{
						Latitude:  helloworldtestingggg.F(37.444300),
						Longitude: helloworldtestingggg.F(-122.159800),
					}),
					State:  helloworldtestingggg.F("CA"),
					Street: helloworldtestingggg.F("437 Lytton"),
					Zip:    helloworldtestingggg.F("94301"),
				}),
				Contact: helloworldtestingggg.F(helloworldtestingggg.TransferLegContactParam{
					Name:  helloworldtestingggg.F("name"),
					Phone: helloworldtestingggg.F("phone"),
				}),
				Cost: helloworldtestingggg.F(shared.MoneyParam{
					Amount:   helloworldtestingggg.F(int64(2500)),
					Currency: helloworldtestingggg.F("USD"),
				}),
				Window: helloworldtestingggg.F(helloworldtestingggg.TransferLegWindowParam{
					End:   helloworldtestingggg.F(time.Now()),
					Start: helloworldtestingggg.F(time.Now()),
				}),
			}),
			OccurredAt: helloworldtestingggg.F(time.Now()),
			Type:       helloworldtestingggg.F(helloworldtestingggg.PlacementEventPlacementTransferEventTypeTransfer),
			Note:       helloworldtestingggg.F("note"),
		}),
		Logistics: helloworldtestingggg.F(helloworldtestingggg.PlacementNewParamsLogistics{
			Delivery: helloworldtestingggg.F(helloworldtestingggg.TransferLegParam{
				Location: helloworldtestingggg.F(shared.AddressParam{
					City: helloworldtestingggg.F("Palo Alto"),
					Geo: helloworldtestingggg.F(shared.AddressGeoParam{
						Latitude:  helloworldtestingggg.F(37.444300),
						Longitude: helloworldtestingggg.F(-122.159800),
					}),
					State:  helloworldtestingggg.F("CA"),
					Street: helloworldtestingggg.F("437 Lytton"),
					Zip:    helloworldtestingggg.F("94301"),
				}),
				Contact: helloworldtestingggg.F(helloworldtestingggg.TransferLegContactParam{
					Name:  helloworldtestingggg.F("name"),
					Phone: helloworldtestingggg.F("phone"),
				}),
				Cost: helloworldtestingggg.F(shared.MoneyParam{
					Amount:   helloworldtestingggg.F(int64(2500)),
					Currency: helloworldtestingggg.F("USD"),
				}),
				Window: helloworldtestingggg.F(helloworldtestingggg.TransferLegWindowParam{
					End:   helloworldtestingggg.F(time.Now()),
					Start: helloworldtestingggg.F(time.Now()),
				}),
			}),
			Handler: helloworldtestingggg.F(helloworldtestingggg.PlacementNewParamsLogisticsHandler{
				Certification: helloworldtestingggg.F(helloworldtestingggg.PlacementNewParamsLogisticsHandlerCertification{
					Level: helloworldtestingggg.F(helloworldtestingggg.PlacementNewParamsLogisticsHandlerCertificationLevelStandard),
					Audit: helloworldtestingggg.F(helloworldtestingggg.PlacementNewParamsLogisticsHandlerCertificationAudit{
						RecordedAt: helloworldtestingggg.F(time.Now()),
						RecordedBy: helloworldtestingggg.F("recordedBy"),
						Source:     helloworldtestingggg.F(helloworldtestingggg.PlacementNewParamsLogisticsHandlerCertificationAuditSourceManual),
					}),
					IssuedBy: helloworldtestingggg.F("issuedBy"),
				}),
				Name: helloworldtestingggg.F("name"),
			}),
			Pickup: helloworldtestingggg.F(helloworldtestingggg.TransferLegParam{
				Location: helloworldtestingggg.F(shared.AddressParam{
					City: helloworldtestingggg.F("Palo Alto"),
					Geo: helloworldtestingggg.F(shared.AddressGeoParam{
						Latitude:  helloworldtestingggg.F(37.444300),
						Longitude: helloworldtestingggg.F(-122.159800),
					}),
					State:  helloworldtestingggg.F("CA"),
					Street: helloworldtestingggg.F("437 Lytton"),
					Zip:    helloworldtestingggg.F("94301"),
				}),
				Contact: helloworldtestingggg.F(helloworldtestingggg.TransferLegContactParam{
					Name:  helloworldtestingggg.F("name"),
					Phone: helloworldtestingggg.F("phone"),
				}),
				Cost: helloworldtestingggg.F(shared.MoneyParam{
					Amount:   helloworldtestingggg.F(int64(2500)),
					Currency: helloworldtestingggg.F("USD"),
				}),
				Window: helloworldtestingggg.F(helloworldtestingggg.TransferLegWindowParam{
					End:   helloworldtestingggg.F(time.Now()),
					Start: helloworldtestingggg.F(time.Now()),
				}),
			}),
		}),
		ScheduledFor: helloworldtestingggg.F(time.Now()),
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
		option.WithBasicAuthUsername("My Basic Auth Username"),
		option.WithBasicAuthPassword("My Basic Auth Password"),
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
		option.WithBasicAuthUsername("My Basic Auth Username"),
		option.WithBasicAuthPassword("My Basic Auth Password"),
	)
	_, err := client.Placements.List(context.TODO(), helloworldtestingggg.PlacementListParams{
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
		option.WithBasicAuthUsername("My Basic Auth Username"),
		option.WithBasicAuthPassword("My Basic Auth Password"),
	)
	_, err := client.Placements.RecordEvent(
		context.TODO(),
		"placementId",
		helloworldtestingggg.PlacementRecordEventParams{
			PlacementEvent: helloworldtestingggg.PlacementEventPlacementTransferEventParam{
				ID: helloworldtestingggg.F("id"),
				Leg: helloworldtestingggg.F(helloworldtestingggg.TransferLegParam{
					Location: helloworldtestingggg.F(shared.AddressParam{
						City: helloworldtestingggg.F("Palo Alto"),
						Geo: helloworldtestingggg.F(shared.AddressGeoParam{
							Latitude:  helloworldtestingggg.F(37.444300),
							Longitude: helloworldtestingggg.F(-122.159800),
						}),
						State:  helloworldtestingggg.F("CA"),
						Street: helloworldtestingggg.F("437 Lytton"),
						Zip:    helloworldtestingggg.F("94301"),
					}),
					Contact: helloworldtestingggg.F(helloworldtestingggg.TransferLegContactParam{
						Name:  helloworldtestingggg.F("name"),
						Phone: helloworldtestingggg.F("phone"),
					}),
					Cost: helloworldtestingggg.F(shared.MoneyParam{
						Amount:   helloworldtestingggg.F(int64(2500)),
						Currency: helloworldtestingggg.F("USD"),
					}),
					Window: helloworldtestingggg.F(helloworldtestingggg.TransferLegWindowParam{
						End:   helloworldtestingggg.F(time.Now()),
						Start: helloworldtestingggg.F(time.Now()),
					}),
				}),
				OccurredAt: helloworldtestingggg.F(time.Now()),
				Type:       helloworldtestingggg.F(helloworldtestingggg.PlacementEventPlacementTransferEventTypeTransfer),
				Note:       helloworldtestingggg.F("note"),
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
