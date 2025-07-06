// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx_test

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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Get(
		context.TODO(),
		0,
		paperlessngx.DocumentGetParams{
			Fields:    []string{"string"},
			FullPerms: paperlessngx.Bool(true),
		},
	)
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Update(
		context.TODO(),
		0,
		paperlessngx.DocumentUpdateParams{
			ArchiveSerialNumber: paperlessngx.Int(0),
			Content:             paperlessngx.String("content"),
			Correspondent:       paperlessngx.Int(0),
			Created:             paperlessngx.Time(time.Now()),
			CreatedDate:         paperlessngx.Time(time.Now()),
			CustomFields: []paperlessngx.CustomFieldInstanceRequestParam{{
				Field: 0,
				Value: paperlessngx.CustomFieldInstanceRequestValueUnionParam{
					OfString: paperlessngx.String("string"),
				},
			}},
			DeletedAt:       paperlessngx.Time(time.Now()),
			DocumentType:    paperlessngx.Int(0),
			Owner:           paperlessngx.Int(0),
			RemoveInboxTags: paperlessngx.Bool(true),
			SetPermissions: paperlessngx.DocumentUpdateParamsSetPermissions{
				Change: paperlessngx.DocumentUpdateParamsSetPermissionsChange{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessngx.DocumentUpdateParamsSetPermissionsView{
					Groups: []int64{0},
					Users:  []int64{0},
				},
			},
			StoragePath: paperlessngx.Int(0),
			Tags:        []int64{0},
			Title:       paperlessngx.String("title"),
		},
	)
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.List(context.TODO(), paperlessngx.DocumentListParams{
		ID:                           paperlessngx.Int(0),
		AddedDateGt:                  paperlessngx.Time(time.Now()),
		AddedDateGte:                 paperlessngx.Time(time.Now()),
		AddedDateLt:                  paperlessngx.Time(time.Now()),
		AddedDateLte:                 paperlessngx.Time(time.Now()),
		AddedDay:                     paperlessngx.Float(0),
		AddedGt:                      paperlessngx.Time(time.Now()),
		AddedGte:                     paperlessngx.Time(time.Now()),
		AddedLt:                      paperlessngx.Time(time.Now()),
		AddedLte:                     paperlessngx.Time(time.Now()),
		AddedMonth:                   paperlessngx.Float(0),
		AddedYear:                    paperlessngx.Float(0),
		ArchiveSerialNumber:          paperlessngx.Int(0),
		ArchiveSerialNumberGt:        paperlessngx.Int(0),
		ArchiveSerialNumberGte:       paperlessngx.Int(0),
		ArchiveSerialNumberIsnull:    paperlessngx.Bool(true),
		ArchiveSerialNumberLt:        paperlessngx.Int(0),
		ArchiveSerialNumberLte:       paperlessngx.Int(0),
		ChecksumIcontains:            paperlessngx.String("checksum__icontains"),
		ChecksumIendswith:            paperlessngx.String("checksum__iendswith"),
		ChecksumIexact:               paperlessngx.String("checksum__iexact"),
		ChecksumIstartswith:          paperlessngx.String("checksum__istartswith"),
		ContentIcontains:             paperlessngx.String("content__icontains"),
		ContentIendswith:             paperlessngx.String("content__iendswith"),
		ContentIexact:                paperlessngx.String("content__iexact"),
		ContentIstartswith:           paperlessngx.String("content__istartswith"),
		CorrespondentID:              paperlessngx.Int(0),
		CorrespondentIDIn:            []int64{0},
		CorrespondentIDNone:          paperlessngx.Int(0),
		CorrespondentIsnull:          paperlessngx.Bool(true),
		CorrespondentNameIcontains:   paperlessngx.String("correspondent__name__icontains"),
		CorrespondentNameIendswith:   paperlessngx.String("correspondent__name__iendswith"),
		CorrespondentNameIexact:      paperlessngx.String("correspondent__name__iexact"),
		CorrespondentNameIstartswith: paperlessngx.String("correspondent__name__istartswith"),
		CreatedDateGt:                paperlessngx.Time(time.Now()),
		CreatedDateGte:               paperlessngx.Time(time.Now()),
		CreatedDateLt:                paperlessngx.Time(time.Now()),
		CreatedDateLte:               paperlessngx.Time(time.Now()),
		CreatedDay:                   paperlessngx.Float(0),
		CreatedGt:                    paperlessngx.Time(time.Now()),
		CreatedGte:                   paperlessngx.Time(time.Now()),
		CreatedLt:                    paperlessngx.Time(time.Now()),
		CreatedLte:                   paperlessngx.Time(time.Now()),
		CreatedMonth:                 paperlessngx.Float(0),
		CreatedYear:                  paperlessngx.Float(0),
		CustomFieldQuery:             paperlessngx.String("x"),
		CustomFieldsIcontains:        paperlessngx.String("x"),
		CustomFieldsIDAll:            paperlessngx.Int(0),
		CustomFieldsIDIn:             paperlessngx.Int(0),
		CustomFieldsIDNone:           paperlessngx.Int(0),
		DocumentTypeID:               paperlessngx.Int(0),
		DocumentTypeIDIn:             []int64{0},
		DocumentTypeIDNone:           paperlessngx.Int(0),
		DocumentTypeIsnull:           paperlessngx.Bool(true),
		DocumentTypeNameIcontains:    paperlessngx.String("document_type__name__icontains"),
		DocumentTypeNameIendswith:    paperlessngx.String("document_type__name__iendswith"),
		DocumentTypeNameIexact:       paperlessngx.String("document_type__name__iexact"),
		DocumentTypeNameIstartswith:  paperlessngx.String("document_type__name__istartswith"),
		Fields:                       []string{"string"},
		FullPerms:                    paperlessngx.Bool(true),
		HasCustomFields:              paperlessngx.Bool(true),
		IDIn:                         []int64{0},
		IsInInbox:                    paperlessngx.Bool(true),
		IsTagged:                     paperlessngx.Bool(true),
		MimeType:                     paperlessngx.String("mime_type"),
		ModifiedDateGt:               paperlessngx.Time(time.Now()),
		ModifiedDateGte:              paperlessngx.Time(time.Now()),
		ModifiedDateLt:               paperlessngx.Time(time.Now()),
		ModifiedDateLte:              paperlessngx.Time(time.Now()),
		ModifiedDay:                  paperlessngx.Float(0),
		ModifiedGt:                   paperlessngx.Time(time.Now()),
		ModifiedGte:                  paperlessngx.Time(time.Now()),
		ModifiedLt:                   paperlessngx.Time(time.Now()),
		ModifiedLte:                  paperlessngx.Time(time.Now()),
		ModifiedMonth:                paperlessngx.Float(0),
		ModifiedYear:                 paperlessngx.Float(0),
		Ordering:                     paperlessngx.String("ordering"),
		OriginalFilenameIcontains:    paperlessngx.String("original_filename__icontains"),
		OriginalFilenameIendswith:    paperlessngx.String("original_filename__iendswith"),
		OriginalFilenameIexact:       paperlessngx.String("original_filename__iexact"),
		OriginalFilenameIstartswith:  paperlessngx.String("original_filename__istartswith"),
		OwnerID:                      paperlessngx.Int(0),
		OwnerIDIn:                    []int64{0},
		OwnerIDNone:                  paperlessngx.Int(0),
		OwnerIsnull:                  paperlessngx.Bool(true),
		Page:                         paperlessngx.Int(0),
		PageSize:                     paperlessngx.Int(0),
		Search:                       paperlessngx.String("search"),
		SharedByID:                   paperlessngx.Bool(true),
		StoragePathID:                paperlessngx.Int(0),
		StoragePathIDIn:              []int64{0},
		StoragePathIDNone:            paperlessngx.Int(0),
		StoragePathIsnull:            paperlessngx.Bool(true),
		StoragePathNameIcontains:     paperlessngx.String("storage_path__name__icontains"),
		StoragePathNameIendswith:     paperlessngx.String("storage_path__name__iendswith"),
		StoragePathNameIexact:        paperlessngx.String("storage_path__name__iexact"),
		StoragePathNameIstartswith:   paperlessngx.String("storage_path__name__istartswith"),
		TagsID:                       paperlessngx.Int(0),
		TagsIDAll:                    paperlessngx.Int(0),
		TagsIDIn:                     paperlessngx.Int(0),
		TagsIDNone:                   paperlessngx.Int(0),
		TagsNameIcontains:            paperlessngx.String("tags__name__icontains"),
		TagsNameIendswith:            paperlessngx.String("tags__name__iendswith"),
		TagsNameIexact:               paperlessngx.String("tags__name__iexact"),
		TagsNameIstartswith:          paperlessngx.String("tags__name__istartswith"),
		TitleIcontains:               paperlessngx.String("title__icontains"),
		TitleIendswith:               paperlessngx.String("title__iendswith"),
		TitleIexact:                  paperlessngx.String("title__iexact"),
		TitleIstartswith:             paperlessngx.String("title__istartswith"),
		TitleContent:                 paperlessngx.String("x"),
	})
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	err := client.Documents.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.BulkDownload(context.TODO(), paperlessngx.DocumentBulkDownloadParams{
		Documents:        []int64{0},
		Compression:      paperlessngx.CompressionEnumNone,
		Content:          paperlessngx.ContentEnumArchive,
		FollowFormatting: paperlessngx.Bool(true),
	})
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.BulkEdit(context.TODO(), paperlessngx.DocumentBulkEditParams{
		Documents: []int64{0},
		Method:    paperlessngx.DocumentBulkEditParamsMethodSetCorrespondent,
		Parameters: map[string]any{
			"foo": "bar",
		},
	})
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Download(
		context.TODO(),
		0,
		paperlessngx.DocumentDownloadParams{
			Original: paperlessngx.Bool(true),
		},
	)
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Email(
		context.TODO(),
		0,
		paperlessngx.DocumentEmailParams{
			Addresses:         "x",
			Message:           "x",
			Subject:           "x",
			UseArchiveVersion: paperlessngx.Bool(true),
		},
	)
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.History(
		context.TODO(),
		0,
		paperlessngx.DocumentHistoryParams{
			Page:     paperlessngx.Int(0),
			PageSize: paperlessngx.Int(0),
		},
	)
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Metadata(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.NextAsn(context.TODO())
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Preview(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.SelectionData(context.TODO(), paperlessngx.DocumentSelectionDataParams{
		Documents: []int64{0},
	})
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.ShareLinks(context.TODO(), "id")
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Suggestions(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Thumbnail(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
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
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Documents.Upload(context.TODO(), paperlessngx.DocumentUploadParams{
		Document:            io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		ArchiveSerialNumber: paperlessngx.Int(0),
		Correspondent:       paperlessngx.Int(0),
		Created:             paperlessngx.Time(time.Now()),
		CustomFields:        []int64{0},
		DocumentType:        paperlessngx.Int(0),
		FromWebui:           paperlessngx.Bool(true),
		StoragePath:         paperlessngx.Int(0),
		Tags:                []int64{0},
		Title:               paperlessngx.String("x"),
	})
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
