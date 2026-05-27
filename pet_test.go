// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/hello-world-testingggg-go"
	"github.com/stainless-sdks/hello-world-testingggg-go/internal/testutil"
	"github.com/stainless-sdks/hello-world-testingggg-go/option"
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
	)
	_, err := client.Pet.New(context.TODO(), helloworldtestingggg.PetNewParams{
		Pet: helloworldtestingggg.PetParam{
			Name:      helloworldtestingggg.F("doggie"),
			PhotoURLs: helloworldtestingggg.F([]string{"string"}),
			ID:        helloworldtestingggg.F(int64(10)),
			Category: helloworldtestingggg.F(helloworldtestingggg.PetCategoryParam{
				ID:   helloworldtestingggg.F(int64(1)),
				Name: helloworldtestingggg.F("Dogs"),
			}),
			Status: helloworldtestingggg.F(helloworldtestingggg.PetStatusAvailable),
			Tags: helloworldtestingggg.F([]helloworldtestingggg.PetTagParam{{
				ID:   helloworldtestingggg.F(int64(0)),
				Name: helloworldtestingggg.F("name"),
			}}),
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
	)
	_, err := client.Pet.Get(context.TODO(), int64(0))
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
	)
	_, err := client.Pet.Update(context.TODO(), helloworldtestingggg.PetUpdateParams{
		Pet: helloworldtestingggg.PetParam{
			Name:      helloworldtestingggg.F("doggie"),
			PhotoURLs: helloworldtestingggg.F([]string{"string"}),
			ID:        helloworldtestingggg.F(int64(10)),
			Category: helloworldtestingggg.F(helloworldtestingggg.PetCategoryParam{
				ID:   helloworldtestingggg.F(int64(1)),
				Name: helloworldtestingggg.F("Dogs"),
			}),
			Status: helloworldtestingggg.F(helloworldtestingggg.PetStatusAvailable),
			Tags: helloworldtestingggg.F([]helloworldtestingggg.PetTagParam{{
				ID:   helloworldtestingggg.F(int64(0)),
				Name: helloworldtestingggg.F("name"),
			}}),
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
	)
	err := client.Pet.Delete(context.TODO(), int64(0))
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
	)
	_, err := client.Pet.FindByStatus(context.TODO(), helloworldtestingggg.PetFindByStatusParams{
		Status: helloworldtestingggg.F(helloworldtestingggg.PetFindByStatusParamsStatusAvailable),
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
	)
	_, err := client.Pet.FindByTags(context.TODO(), helloworldtestingggg.PetFindByTagsParams{
		Tags: helloworldtestingggg.F([]string{"string"}),
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
	)
	err := client.Pet.UpdateWithForm(
		context.TODO(),
		int64(0),
		helloworldtestingggg.PetUpdateWithFormParams{
			Name:   helloworldtestingggg.F("name"),
			Status: helloworldtestingggg.F("status"),
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
	)
	_, err := client.Pet.UploadImage(
		context.TODO(),
		int64(0),
		io.Reader(bytes.NewBuffer([]byte("Example data"))),
		helloworldtestingggg.PetUploadImageParams{
			AdditionalMetadata: helloworldtestingggg.F("additionalMetadata"),
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
