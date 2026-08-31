// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/swagger-girly/go"
	"github.com/swagger-girly/go/internal/testutil"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/shared"
)

func TestPetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.New(context.TODO(), helloworldtestingggg.PetNewParams{
		Pet: helloworldtestingggg.PetParam{
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
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetGet(t *testing.T) {
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
	_, err := client.Pet.Get(context.TODO(), 0)
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.Update(context.TODO(), helloworldtestingggg.PetUpdateParams{
		Pet: helloworldtestingggg.PetParam{
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
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.List(context.TODO(), helloworldtestingggg.PetListParams{
		CreatedAt: helloworldtestingggg.PetListParamsCreatedAt{
			Gt:  helloworldtestingggg.Time(time.Now()),
			Gte: helloworldtestingggg.Time(time.Now()),
			Lt:  helloworldtestingggg.Time(time.Now()),
			Lte: helloworldtestingggg.Time(time.Now()),
		},
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

func TestPetDelete(t *testing.T) {
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
	err := client.Pet.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.FindByStatus(context.TODO(), helloworldtestingggg.PetFindByStatusParams{
		Status: helloworldtestingggg.PetFindByStatusParamsStatusAvailable,
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByTagsWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.FindByTags(context.TODO(), helloworldtestingggg.PetFindByTagsParams{
		Tags: []string{"string"},
	})
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetListFakePage(t *testing.T) {
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
	_, err := client.Pet.ListFakePage(context.TODO())
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetListFakePageInferred(t *testing.T) {
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
	_, err := client.Pet.ListFakePageInferred(context.TODO())
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetListLeaderboard(t *testing.T) {
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
	_, err := client.Pet.ListLeaderboard(context.TODO())
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetListUnpaginatedWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.ListUnpaginated(context.TODO(), helloworldtestingggg.PetListUnpaginatedParams{
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

func TestPetGetPremium(t *testing.T) {
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
	_, err := client.Pet.GetPremium(context.TODO(), 0)
	if err != nil {
		var apierr *helloworldtestingggg.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.Search(context.TODO(), helloworldtestingggg.PetSearchParams{
		Filters: helloworldtestingggg.PetSearchParamsFilters{
			Color: helloworldtestingggg.String("color"),
			Size:  helloworldtestingggg.Int(0),
		},
		MaxResults: helloworldtestingggg.Int(0),
		RawFilter:  map[string]any{},
		TagFilters: []helloworldtestingggg.PetSearchParamsTagFilter{{
			Key:   helloworldtestingggg.String("key"),
			Match: "exact",
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

func TestPetUpdateWithFormWithOptionalParams(t *testing.T) {
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
	err := client.Pet.UpdateWithForm(
		context.TODO(),
		0,
		helloworldtestingggg.PetUpdateWithFormParams{
			Name:   helloworldtestingggg.String("name"),
			Status: helloworldtestingggg.String("status"),
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

func TestPetUploadImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Pet.UploadImage(
		context.TODO(),
		0,
		io.Reader(bytes.NewBuffer([]byte("Example data"))),
		helloworldtestingggg.PetUploadImageParams{
			AdditionalMetadata: helloworldtestingggg.String("additionalMetadata"),
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
