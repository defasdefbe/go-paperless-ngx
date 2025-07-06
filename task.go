// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/paperless-nix-go/internal/apijson"
	"github.com/stainless-sdks/paperless-nix-go/internal/apiquery"
	"github.com/stainless-sdks/paperless-nix-go/internal/requestconfig"
	"github.com/stainless-sdks/paperless-nix-go/option"
	"github.com/stainless-sdks/paperless-nix-go/packages/param"
	"github.com/stainless-sdks/paperless-nix-go/packages/respjson"
)

// TaskService contains methods and other services that help with interacting with
// the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskService] method instead.
type TaskService struct {
	Options []option.RequestOption
}

// NewTaskService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaskService(opts ...option.RequestOption) (r TaskService) {
	r = TaskService{}
	r.Options = opts
	return
}

func (r *TaskService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *TasksView, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/tasks/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *TaskService) List(ctx context.Context, query TaskListParams, opts ...option.RequestOption) (res *[]TasksView, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/tasks/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Acknowledge a list of tasks
func (r *TaskService) Acknowledge(ctx context.Context, body TaskAcknowledgeParams, opts ...option.RequestOption) (res *TaskAcknowledgeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/tasks/acknowledge/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *TaskService) Run(ctx context.Context, body TaskRunParams, opts ...option.RequestOption) (res *TasksView, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/tasks/run/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// - `FAILURE` - FAILURE
// - `PENDING` - PENDING
// - `RECEIVED` - RECEIVED
// - `RETRY` - RETRY
// - `REVOKED` - REVOKED
// - `STARTED` - STARTED
// - `SUCCESS` - SUCCESS
type StatusEnum string

const (
	StatusEnumFailure  StatusEnum = "FAILURE"
	StatusEnumPending  StatusEnum = "PENDING"
	StatusEnumReceived StatusEnum = "RECEIVED"
	StatusEnumRetry    StatusEnum = "RETRY"
	StatusEnumRevoked  StatusEnum = "REVOKED"
	StatusEnumStarted  StatusEnum = "STARTED"
	StatusEnumSuccess  StatusEnum = "SUCCESS"
)

type TasksView struct {
	ID              int64  `json:"id,required"`
	RelatedDocument string `json:"related_document,required"`
	// Celery ID for the Task that was run
	TaskID string `json:"task_id,required"`
	// If the task is acknowledged via the frontend or API
	Acknowledged bool `json:"acknowledged"`
	// Datetime field when the task result was created in UTC
	DateCreated time.Time `json:"date_created,nullable" format:"date-time"`
	// Datetime field when the task was completed in UTC
	DateDone time.Time `json:"date_done,nullable" format:"date-time"`
	Owner    int64     `json:"owner,nullable"`
	// The data returned by the task
	Result string `json:"result,nullable"`
	// - `FAILURE` - FAILURE
	// - `PENDING` - PENDING
	// - `RECEIVED` - RECEIVED
	// - `RETRY` - RETRY
	// - `REVOKED` - REVOKED
	// - `STARTED` - STARTED
	// - `SUCCESS` - SUCCESS
	//
	// Any of "FAILURE", "PENDING", "RECEIVED", "RETRY", "REVOKED", "STARTED",
	// "SUCCESS".
	Status StatusEnum `json:"status"`
	// Name of the file which the Task was run for
	TaskFileName string `json:"task_file_name,nullable"`
	// Name of the task that was run
	//
	// - `consume_file` - Consume File
	// - `train_classifier` - Train Classifier
	// - `check_sanity` - Check Sanity
	// - `index_optimize` - Index Optimize
	//
	// Any of "consume_file", "train_classifier", "check_sanity", "index_optimize".
	TaskName TasksViewTaskName `json:"task_name,nullable"`
	// - `auto_task` - Auto Task
	// - `scheduled_task` - Scheduled Task
	// - `manual_task` - Manual Task
	//
	// Any of "auto_task", "scheduled_task", "manual_task".
	Type TasksViewTypeEnum `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		RelatedDocument respjson.Field
		TaskID          respjson.Field
		Acknowledged    respjson.Field
		DateCreated     respjson.Field
		DateDone        respjson.Field
		Owner           respjson.Field
		Result          respjson.Field
		Status          respjson.Field
		TaskFileName    respjson.Field
		TaskName        respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TasksView) RawJSON() string { return r.JSON.raw }
func (r *TasksView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Name of the task that was run
//
// - `consume_file` - Consume File
// - `train_classifier` - Train Classifier
// - `check_sanity` - Check Sanity
// - `index_optimize` - Index Optimize
type TasksViewTaskName string

const (
	TasksViewTaskNameConsumeFile     TasksViewTaskName = "consume_file"
	TasksViewTaskNameTrainClassifier TasksViewTaskName = "train_classifier"
	TasksViewTaskNameCheckSanity     TasksViewTaskName = "check_sanity"
	TasksViewTaskNameIndexOptimize   TasksViewTaskName = "index_optimize"
)

// - `auto_task` - Auto Task
// - `scheduled_task` - Scheduled Task
// - `manual_task` - Manual Task
type TasksViewTypeEnum string

const (
	TasksViewTypeEnumAutoTask      TasksViewTypeEnum = "auto_task"
	TasksViewTypeEnumScheduledTask TasksViewTypeEnum = "scheduled_task"
	TasksViewTypeEnumManualTask    TasksViewTypeEnum = "manual_task"
)

type TaskAcknowledgeResponse struct {
	Result int64 `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAcknowledgeResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAcknowledgeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListParams struct {
	// Acknowledged
	Acknowledged param.Opt[bool] `query:"acknowledged,omitzero" json:"-"`
	// Which field to use when ordering the results.
	Ordering param.Opt[string] `query:"ordering,omitzero" json:"-"`
	// Name of the task that was run
	//
	// - `consume_file` - Consume File
	// - `train_classifier` - Train Classifier
	// - `check_sanity` - Check Sanity
	// - `index_optimize` - Index Optimize
	//
	// Any of "check_sanity", "consume_file", "index_optimize", "train_classifier".
	TaskName TaskListParamsTaskName `query:"task_name,omitzero" json:"-"`
	// Current state of the task being run
	//
	// - `FAILURE` - FAILURE
	// - `PENDING` - PENDING
	// - `RECEIVED` - RECEIVED
	// - `RETRY` - RETRY
	// - `REVOKED` - REVOKED
	// - `STARTED` - STARTED
	// - `SUCCESS` - SUCCESS
	//
	// Any of "FAILURE", "PENDING", "RECEIVED", "RETRY", "REVOKED", "STARTED",
	// "SUCCESS".
	Status TaskListParamsStatus `query:"status,omitzero" json:"-"`
	// The type of task that was run
	//
	// - `auto_task` - Auto Task
	// - `scheduled_task` - Scheduled Task
	// - `manual_task` - Manual Task
	//
	// Any of "auto_task", "manual_task", "scheduled_task".
	Type TaskListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskListParams]'s query parameters as `url.Values`.
func (r TaskListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Current state of the task being run
//
// - `FAILURE` - FAILURE
// - `PENDING` - PENDING
// - `RECEIVED` - RECEIVED
// - `RETRY` - RETRY
// - `REVOKED` - REVOKED
// - `STARTED` - STARTED
// - `SUCCESS` - SUCCESS
type TaskListParamsStatus string

const (
	TaskListParamsStatusFailure  TaskListParamsStatus = "FAILURE"
	TaskListParamsStatusPending  TaskListParamsStatus = "PENDING"
	TaskListParamsStatusReceived TaskListParamsStatus = "RECEIVED"
	TaskListParamsStatusRetry    TaskListParamsStatus = "RETRY"
	TaskListParamsStatusRevoked  TaskListParamsStatus = "REVOKED"
	TaskListParamsStatusStarted  TaskListParamsStatus = "STARTED"
	TaskListParamsStatusSuccess  TaskListParamsStatus = "SUCCESS"
)

// Name of the task that was run
//
// - `consume_file` - Consume File
// - `train_classifier` - Train Classifier
// - `check_sanity` - Check Sanity
// - `index_optimize` - Index Optimize
type TaskListParamsTaskName string

const (
	TaskListParamsTaskNameCheckSanity     TaskListParamsTaskName = "check_sanity"
	TaskListParamsTaskNameConsumeFile     TaskListParamsTaskName = "consume_file"
	TaskListParamsTaskNameIndexOptimize   TaskListParamsTaskName = "index_optimize"
	TaskListParamsTaskNameTrainClassifier TaskListParamsTaskName = "train_classifier"
)

// The type of task that was run
//
// - `auto_task` - Auto Task
// - `scheduled_task` - Scheduled Task
// - `manual_task` - Manual Task
type TaskListParamsType string

const (
	TaskListParamsTypeAutoTask      TaskListParamsType = "auto_task"
	TaskListParamsTypeManualTask    TaskListParamsType = "manual_task"
	TaskListParamsTypeScheduledTask TaskListParamsType = "scheduled_task"
)

type TaskAcknowledgeParams struct {
	Tasks []int64 `json:"tasks,omitzero,required"`
	paramObj
}

func (r TaskAcknowledgeParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskAcknowledgeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAcknowledgeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskRunParams struct {
	// Celery ID for the Task that was run
	TaskID string `json:"task_id,required"`
	// Datetime field when the task result was created in UTC
	DateCreated param.Opt[time.Time] `json:"date_created,omitzero" format:"date-time"`
	// Datetime field when the task was completed in UTC
	DateDone param.Opt[time.Time] `json:"date_done,omitzero" format:"date-time"`
	Owner    param.Opt[int64]     `json:"owner,omitzero"`
	// The data returned by the task
	Result param.Opt[string] `json:"result,omitzero"`
	// Name of the file which the Task was run for
	TaskFileName param.Opt[string] `json:"task_file_name,omitzero"`
	// If the task is acknowledged via the frontend or API
	Acknowledged param.Opt[bool] `json:"acknowledged,omitzero"`
	// Name of the task that was run
	//
	// - `consume_file` - Consume File
	// - `train_classifier` - Train Classifier
	// - `check_sanity` - Check Sanity
	// - `index_optimize` - Index Optimize
	//
	// Any of "consume_file", "train_classifier", "check_sanity", "index_optimize".
	TaskName TaskRunParamsTaskName `json:"task_name,omitzero"`
	// - `FAILURE` - FAILURE
	// - `PENDING` - PENDING
	// - `RECEIVED` - RECEIVED
	// - `RETRY` - RETRY
	// - `REVOKED` - REVOKED
	// - `STARTED` - STARTED
	// - `SUCCESS` - SUCCESS
	//
	// Any of "FAILURE", "PENDING", "RECEIVED", "RETRY", "REVOKED", "STARTED",
	// "SUCCESS".
	Status StatusEnum `json:"status,omitzero"`
	// - `auto_task` - Auto Task
	// - `scheduled_task` - Scheduled Task
	// - `manual_task` - Manual Task
	//
	// Any of "auto_task", "scheduled_task", "manual_task".
	Type TasksViewTypeEnum `json:"type,omitzero"`
	paramObj
}

func (r TaskRunParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Name of the task that was run
//
// - `consume_file` - Consume File
// - `train_classifier` - Train Classifier
// - `check_sanity` - Check Sanity
// - `index_optimize` - Index Optimize
type TaskRunParamsTaskName string

const (
	TaskRunParamsTaskNameConsumeFile     TaskRunParamsTaskName = "consume_file"
	TaskRunParamsTaskNameTrainClassifier TaskRunParamsTaskName = "train_classifier"
	TaskRunParamsTaskNameCheckSanity     TaskRunParamsTaskName = "check_sanity"
	TaskRunParamsTaskNameIndexOptimize   TaskRunParamsTaskName = "index_optimize"
)
