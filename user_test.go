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

func TestUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.User.New(context.TODO(), helloworldtestingggg.UserNewParams{
		ID:        helloworldtestingggg.F(int64(10)),
		Email:     helloworldtestingggg.F("john@email.com"),
		FirstName: helloworldtestingggg.F("John"),
		LastName:  helloworldtestingggg.F("James"),
		Password:  helloworldtestingggg.F("12345"),
		Phone:     helloworldtestingggg.F("12345"),
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
		RelatedOrder: helloworldtestingggg.F[any](map[string]interface{}{}),
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
		Username:       helloworldtestingggg.F("theUser"),
		UserStatus:     helloworldtestingggg.F(int64(1)),
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserGet(t *testing.T) {
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
	_, err := client.User.Get(context.TODO(), "username")
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	err := client.User.Update(
		context.TODO(),
		"username",
		helloworldtestingggg.UserUpdateParams{
			ID:        helloworldtestingggg.F(int64(10)),
			Email:     helloworldtestingggg.F("john@email.com"),
			FirstName: helloworldtestingggg.F("John"),
			LastName:  helloworldtestingggg.F("James"),
			Password:  helloworldtestingggg.F("12345"),
			Phone:     helloworldtestingggg.F("12345"),
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
			RelatedOrder: helloworldtestingggg.F[any](map[string]interface{}{}),
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
			Username:       helloworldtestingggg.F("theUser"),
			UserStatus:     helloworldtestingggg.F(int64(1)),
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

func TestUserDelete(t *testing.T) {
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
	err := client.User.Delete(context.TODO(), "username")
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithListWithOptionalParams(t *testing.T) {
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
	_, err := client.User.NewWithList(context.TODO(), helloworldtestingggg.UserNewWithListParams{
		Body: []helloworldtestingggg.UserNewWithListParamsBody{{
			ID:        helloworldtestingggg.F(int64(10)),
			Email:     helloworldtestingggg.F("john@email.com"),
			FirstName: helloworldtestingggg.F("John"),
			LastName:  helloworldtestingggg.F("James"),
			Password:  helloworldtestingggg.F("12345"),
			Phone:     helloworldtestingggg.F("12345"),
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
			Username:   helloworldtestingggg.F("theUser"),
			UserStatus: helloworldtestingggg.F(int64(1)),
		}},
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Login(context.TODO(), helloworldtestingggg.UserLoginParams{
		Password: helloworldtestingggg.F("password"),
		Username: helloworldtestingggg.F("username"),
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
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
	err := client.User.Logout(context.TODO())
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserVerifyIdentity(t *testing.T) {
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
	_, err := client.User.VerifyIdentity(context.TODO(), "username")
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
