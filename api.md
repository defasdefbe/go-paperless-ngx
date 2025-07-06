# BulkEditObjects

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#BulkEditObjectNewResponse">BulkEditObjectNewResponse</a>

Methods:

- <code title="post /api/bulk_edit_objects/">client.BulkEditObjects.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#BulkEditObjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#BulkEditObjectNewParams">BulkEditObjectNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#BulkEditObjectNewResponse">BulkEditObjectNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Config

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ApplicationConfiguration">ApplicationConfiguration</a>

Methods:

- <code title="get /api/config/">client.Config.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ApplicationConfiguration">ApplicationConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/config/{id}/">client.Config.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ConfigService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ConfigUpdateParams">ConfigUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ApplicationConfiguration">ApplicationConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/config/{id}/">client.Config.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ConfigService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="patch /api/config/{id}/">client.Config.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ConfigService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ConfigPatchParams">ConfigPatchParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ApplicationConfiguration">ApplicationConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/config/{id}/">client.Config.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ConfigService.GetByID">GetByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ApplicationConfiguration">ApplicationConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Correspondents

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentRequestParam">CorrespondentRequestParam</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Correspondent">Correspondent</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentListResponse">CorrespondentListResponse</a>

Methods:

- <code title="post /api/correspondents/">client.Correspondents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentNewParams">CorrespondentNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Correspondent">Correspondent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/correspondents/{id}/">client.Correspondents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentGetParams">CorrespondentGetParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Correspondent">Correspondent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/correspondents/{id}/">client.Correspondents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentUpdateParams">CorrespondentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Correspondent">Correspondent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/correspondents/">client.Correspondents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentListParams">CorrespondentListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentListResponse">CorrespondentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/correspondents/{id}/">client.Correspondents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CorrespondentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# CustomFields

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldRequestParam">CustomFieldRequestParam</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DataTypeEnum">DataTypeEnum</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomField">CustomField</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DataTypeEnum">DataTypeEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldListResponse">CustomFieldListResponse</a>

Methods:

- <code title="post /api/custom_fields/">client.CustomFields.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldNewParams">CustomFieldNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomField">CustomField</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/custom_fields/{id}/">client.CustomFields.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomField">CustomField</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/custom_fields/{id}/">client.CustomFields.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldUpdateParams">CustomFieldUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomField">CustomField</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/custom_fields/">client.CustomFields.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldListParams">CustomFieldListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldListResponse">CustomFieldListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/custom_fields/{id}/">client.CustomFields.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# DocumentTypes

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeRequestParam">DocumentTypeRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentType">DocumentType</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeListResponse">DocumentTypeListResponse</a>

Methods:

- <code title="post /api/document_types/">client.DocumentTypes.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeNewParams">DocumentTypeNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentType">DocumentType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/document_types/{id}/">client.DocumentTypes.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeGetParams">DocumentTypeGetParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentType">DocumentType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/document_types/{id}/">client.DocumentTypes.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeUpdateParams">DocumentTypeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentType">DocumentType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/document_types/">client.DocumentTypes.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeListParams">DocumentTypeListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeListResponse">DocumentTypeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/document_types/{id}/">client.DocumentTypes.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentTypeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Documents

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CompressionEnum">CompressionEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ContentEnum">ContentEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CustomFieldInstanceRequestParam">CustomFieldInstanceRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#CompressionEnum">CompressionEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ContentEnum">ContentEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Document">Document</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Notes">Notes</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentListResponse">DocumentListResponse</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentBulkDownloadResponse">DocumentBulkDownloadResponse</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentBulkEditResponse">DocumentBulkEditResponse</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentEmailResponse">DocumentEmailResponse</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentHistoryResponse">DocumentHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentMetadataResponse">DocumentMetadataResponse</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentSelectionDataResponse">DocumentSelectionDataResponse</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentShareLinksResponse">DocumentShareLinksResponse</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentSuggestionsResponse">DocumentSuggestionsResponse</a>

Methods:

- <code title="get /api/documents/{id}/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentGetParams">DocumentGetParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Document">Document</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/documents/{id}/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentUpdateParams">DocumentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Document">Document</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/documents/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentListParams">DocumentListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentListResponse">DocumentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/documents/{id}/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/documents/bulk_download/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.BulkDownload">BulkDownload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentBulkDownloadParams">DocumentBulkDownloadParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentBulkDownloadResponse">DocumentBulkDownloadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/documents/bulk_edit/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.BulkEdit">BulkEdit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentBulkEditParams">DocumentBulkEditParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentBulkEditResponse">DocumentBulkEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/documents/{id}/download/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentDownloadParams">DocumentDownloadParams</a>) (<a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/documents/{id}/email/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.Email">Email</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentEmailParams">DocumentEmailParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentEmailResponse">DocumentEmailResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/documents/{id}/history/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.History">History</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentHistoryParams">DocumentHistoryParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentHistoryResponse">DocumentHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/documents/{id}/metadata/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.Metadata">Metadata</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentMetadataResponse">DocumentMetadataResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/documents/next_asn/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.NextAsn">NextAsn</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#int64">int64</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/documents/{id}/preview/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.Preview">Preview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/documents/selection_data/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.SelectionData">SelectionData</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentSelectionDataParams">DocumentSelectionDataParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentSelectionDataResponse">DocumentSelectionDataResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/documents/{id}/share_links/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.ShareLinks">ShareLinks</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentShareLinksResponse">DocumentShareLinksResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/documents/{id}/suggestions/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.Suggestions">Suggestions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentSuggestionsResponse">DocumentSuggestionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/documents/{id}/thumb/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.Thumbnail">Thumbnail</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/documents/post_document/">client.Documents.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentUploadParams">DocumentUploadParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Notes

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#PaginatedNotesList">PaginatedNotesList</a>

Methods:

- <code title="post /api/documents/{id}/notes/">client.Documents.Notes.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentNoteService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentNoteNewParams">DocumentNoteNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#PaginatedNotesList">PaginatedNotesList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/documents/{id}/notes/">client.Documents.Notes.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentNoteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentNoteListParams">DocumentNoteListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#PaginatedNotesList">PaginatedNotesList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/documents/{id}/notes/">client.Documents.Notes.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentNoteService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#DocumentNoteDeleteParams">DocumentNoteDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#PaginatedNotesList">PaginatedNotesList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Groups

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#GroupRequestParam">GroupRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Group">Group</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#GroupListResponse">GroupListResponse</a>

Methods:

- <code title="post /api/groups/">client.Groups.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#GroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#GroupNewParams">GroupNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Group">Group</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/groups/{id}/">client.Groups.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#GroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Group">Group</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/groups/{id}/">client.Groups.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#GroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#GroupUpdateParams">GroupUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Group">Group</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/groups/">client.Groups.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#GroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#GroupListParams">GroupListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#GroupListResponse">GroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/groups/{id}/">client.Groups.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#GroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Logs

Methods:

- <code title="get /api/logs/{id}/">client.Logs.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#LogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/logs/">client.Logs.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#LogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# MailAccounts

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountRequestParam">MailAccountRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccount">MailAccount</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountListResponse">MailAccountListResponse</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountProcessResponse">MailAccountProcessResponse</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountTestResponse">MailAccountTestResponse</a>

Methods:

- <code title="post /api/mail_accounts/">client.MailAccounts.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountNewParams">MailAccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccount">MailAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/mail_accounts/{id}/">client.MailAccounts.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccount">MailAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/mail_accounts/{id}/">client.MailAccounts.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountUpdateParams">MailAccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccount">MailAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/mail_accounts/">client.MailAccounts.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountListParams">MailAccountListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountListResponse">MailAccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/mail_accounts/{id}/">client.MailAccounts.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/mail_accounts/{id}/process/">client.MailAccounts.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountService.Process">Process</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountProcessParams">MailAccountProcessParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountProcessResponse">MailAccountProcessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/mail_accounts/test/">client.MailAccounts.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountService.Test">Test</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountTestParams">MailAccountTestParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailAccountTestResponse">MailAccountTestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# MailRules

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRuleRequestParam">MailRuleRequestParam</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRule">MailRule</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRuleListResponse">MailRuleListResponse</a>

Methods:

- <code title="post /api/mail_rules/">client.MailRules.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRuleNewParams">MailRuleNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRule">MailRule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/mail_rules/{id}/">client.MailRules.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRuleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRule">MailRule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/mail_rules/{id}/">client.MailRules.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRuleUpdateParams">MailRuleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRule">MailRule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/mail_rules/">client.MailRules.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRuleListParams">MailRuleListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRuleListResponse">MailRuleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/mail_rules/{id}/">client.MailRules.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#MailRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# OAuth

Methods:

- <code title="get /api/oauth/callback/">client.OAuth.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#OAuthService.Callback">Callback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Profile

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Profile">Profile</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileListSocialAccountProvidersResponse">ProfileListSocialAccountProvidersResponse</a>

Methods:

- <code title="get /api/profile/">client.Profile.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Profile">Profile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/profile/">client.Profile.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileUpdateParams">ProfileUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Profile">Profile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/profile/disconnect_social_account/">client.Profile.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileService.DisconnectSocialAccount">DisconnectSocialAccount</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileDisconnectSocialAccountParams">ProfileDisconnectSocialAccountParams</a>) (<a href="https://pkg.go.dev/builtin#int64">int64</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/profile/generate_auth_token/">client.Profile.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileService.GenerateAuthToken">GenerateAuthToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/profile/social_account_providers/">client.Profile.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileService.ListSocialAccountProviders">ListSocialAccountProviders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileListSocialAccountProvidersResponse">ProfileListSocialAccountProvidersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Totp

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileTotpGenerateResponse">ProfileTotpGenerateResponse</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileTotpValidateResponse">ProfileTotpValidateResponse</a>

Methods:

- <code title="delete /api/profile/totp/">client.Profile.Totp.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileTotpService.Deactivate">Deactivate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/profile/totp/">client.Profile.Totp.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileTotpService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileTotpGenerateResponse">ProfileTotpGenerateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/profile/totp/">client.Profile.Totp.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileTotpService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileTotpValidateParams">ProfileTotpValidateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ProfileTotpValidateResponse">ProfileTotpValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RemoteVersion

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#RemoteVersionGetResponse">RemoteVersionGetResponse</a>

Methods:

- <code title="get /api/remote_version/">client.RemoteVersion.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#RemoteVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#RemoteVersionGetResponse">RemoteVersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SavedViews

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewFilterRuleParam">SavedViewFilterRuleParam</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewRequestParam">SavedViewRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedView">SavedView</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewListResponse">SavedViewListResponse</a>

Methods:

- <code title="post /api/saved_views/">client.SavedViews.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewNewParams">SavedViewNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedView">SavedView</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/saved_views/{id}/">client.SavedViews.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedView">SavedView</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/saved_views/{id}/">client.SavedViews.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewUpdateParams">SavedViewUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedView">SavedView</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/saved_views/">client.SavedViews.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewListParams">SavedViewListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewListResponse">SavedViewListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/saved_views/{id}/">client.SavedViews.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SavedViewService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Search

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SearchGetResponse">SearchGetResponse</a>

Methods:

- <code title="get /api/search/">client.Search.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SearchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SearchGetParams">SearchGetParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SearchGetResponse">SearchGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/search/autocomplete/">client.Search.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SearchService.ListTags">ListTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#SearchListTagsParams">SearchListTagsParams</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ShareLinks

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#FileVersionEnum">FileVersionEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLinkRequestParam">ShareLinkRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#FileVersionEnum">FileVersionEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLink">ShareLink</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLinkListResponse">ShareLinkListResponse</a>

Methods:

- <code title="post /api/share_links/">client.ShareLinks.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLinkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLinkNewParams">ShareLinkNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLink">ShareLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/share_links/{id}/">client.ShareLinks.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLinkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLink">ShareLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/share_links/{id}/">client.ShareLinks.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLinkService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLinkUpdateParams">ShareLinkUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLink">ShareLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/share_links/">client.ShareLinks.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLinkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLinkListParams">ShareLinkListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLinkListResponse">ShareLinkListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/share_links/{id}/">client.ShareLinks.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ShareLinkService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Statistics

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StatisticGetResponse">StatisticGetResponse</a>

Methods:

- <code title="get /api/statistics/">client.Statistics.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StatisticService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StatisticGetResponse">StatisticGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Status

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StatusGetResponse">StatusGetResponse</a>

Methods:

- <code title="get /api/status/">client.Status.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StatusGetResponse">StatusGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# StoragePaths

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathRequestParam">StoragePathRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePath">StoragePath</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathListResponse">StoragePathListResponse</a>

Methods:

- <code title="post /api/storage_paths/">client.StoragePaths.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathNewParams">StoragePathNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePath">StoragePath</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/storage_paths/{id}/">client.StoragePaths.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathGetParams">StoragePathGetParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePath">StoragePath</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/storage_paths/{id}/">client.StoragePaths.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathUpdateParams">StoragePathUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePath">StoragePath</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/storage_paths/">client.StoragePaths.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathListParams">StoragePathListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathListResponse">StoragePathListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/storage_paths/{id}/">client.StoragePaths.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/storage_paths/test/">client.StoragePaths.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathService.Test">Test</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePathTestParams">StoragePathTestParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StoragePath">StoragePath</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tags

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagRequestParam">TagRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Tag">Tag</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagListResponse">TagListResponse</a>

Methods:

- <code title="post /api/tags/">client.Tags.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagNewParams">TagNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Tag">Tag</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/tags/{id}/">client.Tags.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagGetParams">TagGetParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Tag">Tag</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/tags/{id}/">client.Tags.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagUpdateParams">TagUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Tag">Tag</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/tags/">client.Tags.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagListParams">TagListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagListResponse">TagListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/tags/{id}/">client.Tags.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TagService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Tasks

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StatusEnum">StatusEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TasksViewTypeEnum">TasksViewTypeEnum</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#StatusEnum">StatusEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TasksView">TasksView</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TasksViewTypeEnum">TasksViewTypeEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TaskAcknowledgeResponse">TaskAcknowledgeResponse</a>

Methods:

- <code title="get /api/tasks/{id}/">client.Tasks.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TaskService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TasksView">TasksView</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/tasks/">client.Tasks.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TaskService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TaskListParams">TaskListParams</a>) ([]<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TasksView">TasksView</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/tasks/acknowledge/">client.Tasks.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TaskService.Acknowledge">Acknowledge</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TaskAcknowledgeParams">TaskAcknowledgeParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TaskAcknowledgeResponse">TaskAcknowledgeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/tasks/run/">client.Tasks.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TaskService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TaskRunParams">TaskRunParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TasksView">TasksView</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Token

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TokenNewResponse">TokenNewResponse</a>

Methods:

- <code title="post /api/token/">client.Token.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TokenNewParams">TokenNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TokenNewResponse">TokenNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Trash

Methods:

- <code title="post /api/trash/">client.Trash.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TrashService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TrashNewParams">TrashNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/trash/">client.Trash.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TrashService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#TrashListParams">TrashListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# UiSettings

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UiSettingsView">UiSettingsView</a>

Methods:

- <code title="post /api/ui_settings/">client.UiSettings.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UiSettingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UiSettingNewParams">UiSettingNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UiSettingsView">UiSettingsView</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/ui_settings/">client.UiSettings.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UiSettingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UiSettingsView">UiSettingsView</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserRequestParam">UserRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#User">User</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserListResponse">UserListResponse</a>

Methods:

- <code title="post /api/users/">client.Users.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserNewParams">UserNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/users/{id}/">client.Users.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/users/{id}/">client.Users.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserUpdateParams">UserUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/users/">client.Users.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserListParams">UserListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserListResponse">UserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/users/{id}/">client.Users.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/users/{id}/deactivate_totp/">client.Users.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#UserService.DeactivateTotp">DeactivateTotp</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WorkflowActions

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionEmailParam">WorkflowActionEmailParam</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionRequestParam">WorkflowActionRequestParam</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionWebhookParam">WorkflowActionWebhookParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowAction">WorkflowAction</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionListResponse">WorkflowActionListResponse</a>

Methods:

- <code title="post /api/workflow_actions/">client.WorkflowActions.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionNewParams">WorkflowActionNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowAction">WorkflowAction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/workflow_actions/{id}/">client.WorkflowActions.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowAction">WorkflowAction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/workflow_actions/{id}/">client.WorkflowActions.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionUpdateParams">WorkflowActionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowAction">WorkflowAction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/workflow_actions/">client.WorkflowActions.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionListParams">WorkflowActionListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionListResponse">WorkflowActionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/workflow_actions/{id}/">client.WorkflowActions.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowActionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# WorkflowTriggers

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ScheduleDateFieldEnum">ScheduleDateFieldEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTriggerRequestParam">WorkflowTriggerRequestParam</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#ScheduleDateFieldEnum">ScheduleDateFieldEnum</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTrigger">WorkflowTrigger</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#int64">int64</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTriggerListResponse">WorkflowTriggerListResponse</a>

Methods:

- <code title="post /api/workflow_triggers/">client.WorkflowTriggers.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTriggerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTriggerNewParams">WorkflowTriggerNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTrigger">WorkflowTrigger</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/workflow_triggers/{id}/">client.WorkflowTriggers.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTriggerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTrigger">WorkflowTrigger</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/workflow_triggers/{id}/">client.WorkflowTriggers.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTriggerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTriggerUpdateParams">WorkflowTriggerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTrigger">WorkflowTrigger</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/workflow_triggers/">client.WorkflowTriggers.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTriggerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTriggerListParams">WorkflowTriggerListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTriggerListResponse">WorkflowTriggerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/workflow_triggers/{id}/">client.WorkflowTriggers.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowTriggerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Workflows

Params Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowRequestParam">WorkflowRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Workflow">Workflow</a>
- <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowListResponse">WorkflowListResponse</a>

Methods:

- <code title="post /api/workflows/">client.Workflows.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowNewParams">WorkflowNewParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Workflow">Workflow</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/workflows/{id}/">client.Workflows.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Workflow">Workflow</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/workflows/{id}/">client.Workflows.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowUpdateParams">WorkflowUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#Workflow">Workflow</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/workflows/">client.Workflows.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowListParams">WorkflowListParams</a>) (<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx">paperlessngx</a>.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowListResponse">WorkflowListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/workflows/{id}/">client.Workflows.<a href="https://pkg.go.dev/github.com/defasdefbe/go-paperless-ngx#WorkflowService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
