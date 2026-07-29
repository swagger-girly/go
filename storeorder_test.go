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
		ID:       helloworldtestingggg.F(int64(10)),
		Complete: helloworldtestingggg.F(true),
		PetID:    helloworldtestingggg.F(int64(198772)),
		Quantity: helloworldtestingggg.F(int64(7)),
		RelatedAddress: helloworldtestingggg.F(shared.AddressParam{
			City: helloworldtestingggg.F("Palo Alto"),
			Geo: helloworldtestingggg.F(shared.AddressGeoParam{
				Latitude:  helloworldtestingggg.F(37.444300),
				Longitude: helloworldtestingggg.F(-122.159800),
			}),
			RelatedCategory: helloworldtestingggg.F[any](map[string]interface{}{}),
			RelatedCustomer: helloworldtestingggg.F[any](map[string]interface{}{}),
			RelatedMoney: helloworldtestingggg.F(shared.MoneyParam{
				Amount:          helloworldtestingggg.F(int64(2500)),
				Currency:        helloworldtestingggg.F("USD"),
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
			RelatedOrder: helloworldtestingggg.F[any](map[string]interface{}{}),
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
		RelatedMoney: helloworldtestingggg.F(shared.MoneyParam{
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
		RelatedPet: helloworldtestingggg.F(helloworldtestingggg.PetParam{
			Name:               helloworldtestingggg.F("doggie"),
			PhotoURLs:          helloworldtestingggg.F([]string{"string"}),
			ID:                 helloworldtestingggg.F(int64(10)),
			AcquisitionChannel: helloworldtestingggg.F(helloworldtestingggg.PetAcquisitionChannelBreeder),
			Category:           helloworldtestingggg.F[any](map[string]interface{}{}),
			MicrochipID:        helloworldtestingggg.F[helloworldtestingggg.PetMicrochipIDUnionParam](shared.UnionString("string")),
			RelatedAddress: helloworldtestingggg.F(shared.AddressParam{
				City: helloworldtestingggg.F("Palo Alto"),
				Geo: helloworldtestingggg.F(shared.AddressGeoParam{
					Latitude:  helloworldtestingggg.F(37.444300),
					Longitude: helloworldtestingggg.F(-122.159800),
				}),
				RelatedCategory: helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedCustomer: helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedMoney: helloworldtestingggg.F(shared.MoneyParam{
					Amount:          helloworldtestingggg.F(int64(2500)),
					Currency:        helloworldtestingggg.F("USD"),
					RelatedCategory: helloworldtestingggg.F[any](map[string]interface{}{}),
					RelatedCustomer: helloworldtestingggg.F[any](map[string]interface{}{}),
					RelatedOrder:    helloworldtestingggg.F[any](map[string]interface{}{}),
					RelatedShelter:  helloworldtestingggg.F[any](map[string]interface{}{}),
					RelatedTag:      helloworldtestingggg.F[any](map[string]interface{}{}),
					RelatedUser:     helloworldtestingggg.F[any](map[string]interface{}{}),
				}),
				RelatedOrder:   helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedShelter: helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedTag:     helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedUser:    helloworldtestingggg.F[any](map[string]interface{}{}),
				State:          helloworldtestingggg.F("CA"),
				Street:         helloworldtestingggg.F("437 Lytton"),
				Zip:            helloworldtestingggg.F("94301"),
			}),
			RelatedCategory: helloworldtestingggg.F[any](map[string]interface{}{}),
			RelatedCustomer: helloworldtestingggg.F[any](map[string]interface{}{}),
			RelatedMoney: helloworldtestingggg.F(shared.MoneyParam{
				Amount:          helloworldtestingggg.F(int64(2500)),
				Currency:        helloworldtestingggg.F("USD"),
				RelatedCategory: helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedCustomer: helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedOrder:    helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedShelter:  helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedTag:      helloworldtestingggg.F[any](map[string]interface{}{}),
				RelatedUser:     helloworldtestingggg.F[any](map[string]interface{}{}),
			}),
			RelatedOrder:   helloworldtestingggg.F[any](map[string]interface{}{}),
			RelatedShelter: helloworldtestingggg.F[any](map[string]interface{}{}),
			RelatedUser:    helloworldtestingggg.F[any](map[string]interface{}{}),
			Status:         helloworldtestingggg.F(helloworldtestingggg.PetStatusAvailable),
			Tags:           helloworldtestingggg.F([]interface{}{map[string]interface{}{}}),
		}),
		RelatedShelter: helloworldtestingggg.F[any](map[string]interface{}{}),
		RelatedTag:     helloworldtestingggg.F[any](map[string]interface{}{}),
		RelatedUser:    helloworldtestingggg.F[any](map[string]interface{}{}),
		ShipDate:       helloworldtestingggg.F(time.Now()),
		Status:         helloworldtestingggg.F(helloworldtestingggg.StoreOrderNewParamsStatusApproved),
		Total: helloworldtestingggg.F(shared.MoneyParam{
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
	_, err := client.Store.Order.Get(context.TODO(), int64(0))
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
	err := client.Store.Order.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
