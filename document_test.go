// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessnix_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/paperless-nix-go"
	"github.com/stainless-sdks/paperless-nix-go/internal/testutil"
	"github.com/stainless-sdks/paperless-nix-go/option"
)

func TestDocumentGetWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Get(
		context.TODO(),
		0,
		paperlessnix.DocumentGetParams{
			Fields:    []string{"string"},
			FullPerms: paperlessnix.Bool(true),
		},
	)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Update(
		context.TODO(),
		0,
		paperlessnix.DocumentUpdateParams{
			ArchiveSerialNumber: paperlessnix.Int(0),
			Content:             paperlessnix.String("content"),
			Correspondent:       paperlessnix.Int(0),
			Created:             paperlessnix.Time(time.Now()),
			CreatedDate:         paperlessnix.Time(time.Now()),
			CustomFields: []paperlessnix.CustomFieldInstanceRequestParam{{
				Field: 0,
				Value: paperlessnix.CustomFieldInstanceRequestValueUnionParam{
					OfString: paperlessnix.String("string"),
				},
			}},
			DeletedAt:       paperlessnix.Time(time.Now()),
			DocumentType:    paperlessnix.Int(0),
			Owner:           paperlessnix.Int(0),
			RemoveInboxTags: paperlessnix.Bool(true),
			SetPermissions: paperlessnix.DocumentUpdateParamsSetPermissions{
				Change: paperlessnix.DocumentUpdateParamsSetPermissionsChange{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessnix.DocumentUpdateParamsSetPermissionsView{
					Groups: []int64{0},
					Users:  []int64{0},
				},
			},
			StoragePath: paperlessnix.Int(0),
			Tags:        []int64{0},
			Title:       paperlessnix.String("title"),
		},
	)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.List(context.TODO(), paperlessnix.DocumentListParams{
		ID:                           paperlessnix.Int(0),
		AddedDateGt:                  paperlessnix.Time(time.Now()),
		AddedDateGte:                 paperlessnix.Time(time.Now()),
		AddedDateLt:                  paperlessnix.Time(time.Now()),
		AddedDateLte:                 paperlessnix.Time(time.Now()),
		AddedDay:                     paperlessnix.Float(0),
		AddedGt:                      paperlessnix.Time(time.Now()),
		AddedGte:                     paperlessnix.Time(time.Now()),
		AddedLt:                      paperlessnix.Time(time.Now()),
		AddedLte:                     paperlessnix.Time(time.Now()),
		AddedMonth:                   paperlessnix.Float(0),
		AddedYear:                    paperlessnix.Float(0),
		ArchiveSerialNumber:          paperlessnix.Int(0),
		ArchiveSerialNumberGt:        paperlessnix.Int(0),
		ArchiveSerialNumberGte:       paperlessnix.Int(0),
		ArchiveSerialNumberIsnull:    paperlessnix.Bool(true),
		ArchiveSerialNumberLt:        paperlessnix.Int(0),
		ArchiveSerialNumberLte:       paperlessnix.Int(0),
		ChecksumIcontains:            paperlessnix.String("checksum__icontains"),
		ChecksumIendswith:            paperlessnix.String("checksum__iendswith"),
		ChecksumIexact:               paperlessnix.String("checksum__iexact"),
		ChecksumIstartswith:          paperlessnix.String("checksum__istartswith"),
		ContentIcontains:             paperlessnix.String("content__icontains"),
		ContentIendswith:             paperlessnix.String("content__iendswith"),
		ContentIexact:                paperlessnix.String("content__iexact"),
		ContentIstartswith:           paperlessnix.String("content__istartswith"),
		CorrespondentID:              paperlessnix.Int(0),
		CorrespondentIDIn:            []int64{0},
		CorrespondentIDNone:          paperlessnix.Int(0),
		CorrespondentIsnull:          paperlessnix.Bool(true),
		CorrespondentNameIcontains:   paperlessnix.String("correspondent__name__icontains"),
		CorrespondentNameIendswith:   paperlessnix.String("correspondent__name__iendswith"),
		CorrespondentNameIexact:      paperlessnix.String("correspondent__name__iexact"),
		CorrespondentNameIstartswith: paperlessnix.String("correspondent__name__istartswith"),
		CreatedDateGt:                paperlessnix.Time(time.Now()),
		CreatedDateGte:               paperlessnix.Time(time.Now()),
		CreatedDateLt:                paperlessnix.Time(time.Now()),
		CreatedDateLte:               paperlessnix.Time(time.Now()),
		CreatedDay:                   paperlessnix.Float(0),
		CreatedGt:                    paperlessnix.Time(time.Now()),
		CreatedGte:                   paperlessnix.Time(time.Now()),
		CreatedLt:                    paperlessnix.Time(time.Now()),
		CreatedLte:                   paperlessnix.Time(time.Now()),
		CreatedMonth:                 paperlessnix.Float(0),
		CreatedYear:                  paperlessnix.Float(0),
		CustomFieldQuery:             paperlessnix.String("x"),
		CustomFieldsIcontains:        paperlessnix.String("x"),
		CustomFieldsIDAll:            paperlessnix.Int(0),
		CustomFieldsIDIn:             paperlessnix.Int(0),
		CustomFieldsIDNone:           paperlessnix.Int(0),
		DocumentTypeID:               paperlessnix.Int(0),
		DocumentTypeIDIn:             []int64{0},
		DocumentTypeIDNone:           paperlessnix.Int(0),
		DocumentTypeIsnull:           paperlessnix.Bool(true),
		DocumentTypeNameIcontains:    paperlessnix.String("document_type__name__icontains"),
		DocumentTypeNameIendswith:    paperlessnix.String("document_type__name__iendswith"),
		DocumentTypeNameIexact:       paperlessnix.String("document_type__name__iexact"),
		DocumentTypeNameIstartswith:  paperlessnix.String("document_type__name__istartswith"),
		Fields:                       []string{"string"},
		FullPerms:                    paperlessnix.Bool(true),
		HasCustomFields:              paperlessnix.Bool(true),
		IDIn:                         []int64{0},
		IsInInbox:                    paperlessnix.Bool(true),
		IsTagged:                     paperlessnix.Bool(true),
		MimeType:                     paperlessnix.String("mime_type"),
		ModifiedDateGt:               paperlessnix.Time(time.Now()),
		ModifiedDateGte:              paperlessnix.Time(time.Now()),
		ModifiedDateLt:               paperlessnix.Time(time.Now()),
		ModifiedDateLte:              paperlessnix.Time(time.Now()),
		ModifiedDay:                  paperlessnix.Float(0),
		ModifiedGt:                   paperlessnix.Time(time.Now()),
		ModifiedGte:                  paperlessnix.Time(time.Now()),
		ModifiedLt:                   paperlessnix.Time(time.Now()),
		ModifiedLte:                  paperlessnix.Time(time.Now()),
		ModifiedMonth:                paperlessnix.Float(0),
		ModifiedYear:                 paperlessnix.Float(0),
		Ordering:                     paperlessnix.String("ordering"),
		OriginalFilenameIcontains:    paperlessnix.String("original_filename__icontains"),
		OriginalFilenameIendswith:    paperlessnix.String("original_filename__iendswith"),
		OriginalFilenameIexact:       paperlessnix.String("original_filename__iexact"),
		OriginalFilenameIstartswith:  paperlessnix.String("original_filename__istartswith"),
		OwnerID:                      paperlessnix.Int(0),
		OwnerIDIn:                    []int64{0},
		OwnerIDNone:                  paperlessnix.Int(0),
		OwnerIsnull:                  paperlessnix.Bool(true),
		Page:                         paperlessnix.Int(0),
		PageSize:                     paperlessnix.Int(0),
		Search:                       paperlessnix.String("search"),
		SharedByID:                   paperlessnix.Bool(true),
		StoragePathID:                paperlessnix.Int(0),
		StoragePathIDIn:              []int64{0},
		StoragePathIDNone:            paperlessnix.Int(0),
		StoragePathIsnull:            paperlessnix.Bool(true),
		StoragePathNameIcontains:     paperlessnix.String("storage_path__name__icontains"),
		StoragePathNameIendswith:     paperlessnix.String("storage_path__name__iendswith"),
		StoragePathNameIexact:        paperlessnix.String("storage_path__name__iexact"),
		StoragePathNameIstartswith:   paperlessnix.String("storage_path__name__istartswith"),
		TagsID:                       paperlessnix.Int(0),
		TagsIDAll:                    paperlessnix.Int(0),
		TagsIDIn:                     paperlessnix.Int(0),
		TagsIDNone:                   paperlessnix.Int(0),
		TagsNameIcontains:            paperlessnix.String("tags__name__icontains"),
		TagsNameIendswith:            paperlessnix.String("tags__name__iendswith"),
		TagsNameIexact:               paperlessnix.String("tags__name__iexact"),
		TagsNameIstartswith:          paperlessnix.String("tags__name__istartswith"),
		TitleIcontains:               paperlessnix.String("title__icontains"),
		TitleIendswith:               paperlessnix.String("title__iendswith"),
		TitleIexact:                  paperlessnix.String("title__iexact"),
		TitleIstartswith:             paperlessnix.String("title__istartswith"),
		TitleContent:                 paperlessnix.String("x"),
	})
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	err := client.Documents.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentBulkDownloadWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.BulkDownload(context.TODO(), paperlessnix.DocumentBulkDownloadParams{
		Documents:        []int64{0},
		Compression:      paperlessnix.CompressionEnumNone,
		Content:          paperlessnix.ContentEnumArchive,
		FollowFormatting: paperlessnix.Bool(true),
	})
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentBulkEditWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.BulkEdit(context.TODO(), paperlessnix.DocumentBulkEditParams{
		Documents: []int64{0},
		Method:    paperlessnix.DocumentBulkEditParamsMethodSetCorrespondent,
		Parameters: map[string]any{
			"foo": "bar",
		},
	})
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentDownloadWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Download(
		context.TODO(),
		0,
		paperlessnix.DocumentDownloadParams{
			Original: paperlessnix.Bool(true),
		},
	)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentEmailWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Email(
		context.TODO(),
		0,
		paperlessnix.DocumentEmailParams{
			Addresses:         "x",
			Message:           "x",
			Subject:           "x",
			UseArchiveVersion: paperlessnix.Bool(true),
		},
	)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentHistoryWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.History(
		context.TODO(),
		0,
		paperlessnix.DocumentHistoryParams{
			Page:     paperlessnix.Int(0),
			PageSize: paperlessnix.Int(0),
		},
	)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentMetadata(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Metadata(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentNextAsn(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.NextAsn(context.TODO())
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentPreview(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Preview(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentSelectionData(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.SelectionData(context.TODO(), paperlessnix.DocumentSelectionDataParams{
		Documents: []int64{0},
	})
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentShareLinks(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.ShareLinks(context.TODO(), "id")
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentSuggestions(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Suggestions(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentThumbnail(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Thumbnail(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentUploadWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Upload(context.TODO(), paperlessnix.DocumentUploadParams{
		Document:            io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		ArchiveSerialNumber: paperlessnix.Int(0),
		Correspondent:       paperlessnix.Int(0),
		Created:             paperlessnix.Time(time.Now()),
		CustomFields:        []int64{0},
		DocumentType:        paperlessnix.Int(0),
		FromWebui:           paperlessnix.Bool(true),
		StoragePath:         paperlessnix.Int(0),
		Tags:                []int64{0},
		Title:               paperlessnix.String("x"),
	})
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
