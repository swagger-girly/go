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

func TestStoreOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Store.Order.New(context.TODO(), helloworldtestingggg.StoreOrderNewParams{
		ID:       helloworldtestingggg.Int(10),
		Complete: helloworldtestingggg.Bool(true),
		PetID:    helloworldtestingggg.Int(198772),
		Quantity: helloworldtestingggg.Int(7),
		RelatedAddress: shared.AddressParam{
			City: helloworldtestingggg.String("Palo Alto"),
			Geo: shared.AddressGeoParam{
				Latitude:  37.4443,
				Longitude: -122.1598,
			},
			RelatedCategory: map[string]any{},
			RelatedCustomer: map[string]any{},
			RelatedMoney: shared.MoneyParam{
				Amount:          2500,
				Currency:        "USD",
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
			RelatedOrder: map[string]any{},
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
		RelatedMoney: shared.MoneyParam{
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
		RelatedPet: helloworldtestingggg.PetParam{
			Name:               "doggie",
			PhotoURLs:          []string{"string"},
			ID:                 helloworldtestingggg.Int(10),
			AcquisitionChannel: helloworldtestingggg.PetAcquisitionChannelBreeder,
			Category:           map[string]any{},
			MicrochipID: helloworldtestingggg.PetMicrochipIDUnionParam{
				OfString: helloworldtestingggg.String("string"),
			},
			RelatedAddress: shared.AddressParam{
				City: helloworldtestingggg.String("Palo Alto"),
				Geo: shared.AddressGeoParam{
					Latitude:  37.4443,
					Longitude: -122.1598,
				},
				RelatedCategory: map[string]any{},
				RelatedCustomer: map[string]any{},
				RelatedMoney: shared.MoneyParam{
					Amount:          2500,
					Currency:        "USD",
					RelatedCategory: map[string]any{},
					RelatedCustomer: map[string]any{},
					RelatedOrder:    map[string]any{},
					RelatedShelter:  map[string]any{},
					RelatedTag:      map[string]any{},
					RelatedUser:     map[string]any{},
				},
				RelatedOrder:   map[string]any{},
				RelatedShelter: map[string]any{},
				RelatedTag:     map[string]any{},
				RelatedUser:    map[string]any{},
				State:          helloworldtestingggg.String("CA"),
				Street:         helloworldtestingggg.String("437 Lytton"),
				Zip:            helloworldtestingggg.String("94301"),
			},
			RelatedCategory: map[string]any{},
			RelatedCustomer: map[string]any{},
			RelatedMoney: shared.MoneyParam{
				Amount:          2500,
				Currency:        "USD",
				RelatedCategory: map[string]any{},
				RelatedCustomer: map[string]any{},
				RelatedOrder:    map[string]any{},
				RelatedShelter:  map[string]any{},
				RelatedTag:      map[string]any{},
				RelatedUser:     map[string]any{},
			},
			RelatedOrder:   map[string]any{},
			RelatedShelter: map[string]any{},
			RelatedUser:    map[string]any{},
			Status:         helloworldtestingggg.PetStatusAvailable,
			Tags:           []any{map[string]any{}},
		},
		RelatedShelter: map[string]any{},
		RelatedTag:     map[string]any{},
		RelatedUser:    map[string]any{},
		ShipDate:       helloworldtestingggg.Time(time.Now()),
		Status:         helloworldtestingggg.StoreOrderNewParamsStatusApproved,
		Total: shared.MoneyParam{
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
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreOrderGet(t *testing.T) {
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
	_, err := client.Store.Order.Get(context.TODO(), 0)
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreOrderDelete(t *testing.T) {
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
	err := client.Store.Order.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
