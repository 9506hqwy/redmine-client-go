package redmine

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

var (
	activityId           = 1
	attachmentId         = 1
	fileToken            = ""
	groupId              = 1
	issueId              = 1
	issueCategoryId      = 1
	issueRelatedId       = 1
	journalId            = 1
	membershipId         = 1
	nwesId               = 1
	projectId            = 1
	projectIdentifier    = "test-project"
	repositoryIdentifier = "main"
	revision             = ""
	roleId               = 1
	timeEntityId         = 1
	trackerId            = 1
	userId               = 1
	versionId            = 1
	wikiTitle            = "wiki"
	wikiVersion          = 1
)

func basicAuth(ctx context.Context, req *http.Request) error {
	req.SetBasicAuth("admin", "admin")
	return nil
}

func client() (*ClientWithResponses, error) {
	hc := http.Client{}

	return NewClientWithResponses("http://127.0.0.1:3000", WithHTTPClient(&hc))
}

func newClient(t *testing.T) *ClientWithResponses {
	t.Skip()

	c, err := client()
	if err != nil {
		t.Errorf("%v", err)
	}

	return c
}

func assertError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func assertHTTPStatus(t *testing.T, resp *http.Response, msg any) {
	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		t.Errorf("HTTP Status: %v %s", resp.StatusCode, msg)
	}
}

func assertResponseBody(t *testing.T, resp any) {
	if reflect.ValueOf(resp).IsNil() {
		t.Errorf("Error: No Content")
	}
}

func assertJson(t *testing.T, resp any) {
	if reflect.ValueOf(resp).IsNil() {
		t.Errorf("Error: No Content")
	}
}

func assertLength(t *testing.T, data any, source []byte) {
	d, err := json.Marshal(data)
	if err != nil {
		t.Errorf("JSON Encoding: %v", err)
	} else if len(d) != len(source) {
		t.Errorf("JSON Length not match: %d ! = %d", len(d), len(source))
	}
}

func assertCustomField(t *testing.T, cf *struct {
	Id       *int         `json:"id,omitempty"`
	Multiple *bool        `json:"multiple,omitempty"`
	Name     *string      `json:"name,omitempty"`
	Value    *interface{} `json:"value,omitempty"`
}) {
	if cf.Multiple != nil && *cf.Multiple {
		arr, ok := (*cf.Value).([]interface{})
		if !ok {
			t.Errorf("%s = %v", *cf.Name, *cf.Value)
		}

		for _, v := range arr {
			_, ok := v.(string)
			if !ok {
				t.Errorf("%s = %v", *cf.Name, *cf.Value)
			}
		}
	} else {
		_, ok := (*cf.Value).(string)
		if !ok {
			t.Errorf("%s = %v", *cf.Name, *cf.Value)
		}
	}
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestAttachmentsDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.AttachmentsDestroyWithResponse(context.TODO(), attachmentId, &AttachmentsDestroyParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestAttachmentsDownloadAllWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.AttachmentsDownloadAllWithResponse(context.TODO(), "issues", issueId, &AttachmentsDownloadAllParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestAttachmentsDownloadWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.AttachmentsDownloadWithResponse(context.TODO(), attachmentId, &AttachmentsDownloadParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestAttachmentsShowWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.AttachmentsShowWithResponse(context.TODO(), attachmentId, &AttachmentsShowParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Attachment)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestAttachmentsThumbnailSizeWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.AttachmentsThumbnailSizeWithResponse(context.TODO(), attachmentId, 64, &AttachmentsThumbnailSizeParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestAttachmentsThumbnailWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.AttachmentsThumbnailWithResponse(context.TODO(), attachmentId, &AttachmentsThumbnailParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestAttachmentsUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	description := t.Name() + "Description"
	body := AttachmentsUpdatePatchJSONRequestBody{
		Attachment: &struct {
			ContentType *string "json:\"content_type,omitempty\""
			Description *string "json:\"description,omitempty\""
			Filename    *string "json:\"filename,omitempty\""
		}{
			Description: &description,
		},
	}
	resp, err := c.AttachmentsUpdatePatchWithResponse(context.TODO(), attachmentId, &AttachmentsUpdatePatchParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestAttachmentsUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	description := t.Name() + "Description"
	filename := t.Name()
	body := AttachmentsUpdatePutJSONRequestBody{
		Attachment: &struct {
			ContentType *string "json:\"content_type,omitempty\""
			Description *string "json:\"description,omitempty\""
			Filename    *string "json:\"filename,omitempty\""
		}{
			Description: &description,
			Filename:    &filename,
		},
	}
	resp, err := c.AttachmentsUpdatePutWithResponse(context.TODO(), attachmentId, &AttachmentsUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestAttachmentsUploadWithBodyWithResponse(t *testing.T) {
	c := newClient(t)

	fileName := "test.txt"
	contentType := "application/octet-stream "
	params := &AttachmentsUploadParams{
		Filename:    &fileName,
		ContentType: &contentType,
	}
	body := bytes.NewReader([]byte("test file content"))

	resp, err := c.AttachmentsUploadWithBodyWithResponse(context.TODO(), params, contentType, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON201)
	assertJson(t, resp.JSON201.Upload)
	assertLength(t, resp.JSON201, resp.Body)
}

func TestCustomFieldsIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.CustomFieldsIndexWithResponse(context.TODO(), &CustomFieldsIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.CustomFields)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestEnumerationsIndexDocumentCategoryWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.EnumerationsIndexDocumentCategoryWithResponse(context.TODO(), &EnumerationsIndexDocumentCategoryParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.DocumentCategories)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestEnumerationsIndexIssuePriorityWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.EnumerationsIndexIssuePriorityWithResponse(context.TODO(), &EnumerationsIndexIssuePriorityParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.IssuePriorities)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestEnumerationsIndexTimeEntryActivityWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.EnumerationsIndexTimeEntryActivityWithResponse(context.TODO(), &EnumerationsIndexTimeEntryActivityParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.TimeEntryActivities)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestFilesCreateWithResponse(t *testing.T) {
	c := newClient(t)

	description := t.Name() + "Description"
	filename := t.Name()
	body := FilesCreateJSONRequestBody{
		File: &struct {
			Description *string "json:\"description,omitempty\""
			Filename    *string "json:\"filename,omitempty\""
			Token       *string "json:\"token,omitempty\""
			VersionId   *int    "json:\"version_id,omitempty\""
		}{
			Description: &description,
			Filename:    &filename,
			Token:       &fileToken,
			VersionId:   &versionId,
		},
	}
	resp, err := c.FilesCreateWithResponse(context.TODO(), projectIdentifier, &FilesCreateParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestFilesIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.FilesIndexWithResponse(context.TODO(), projectIdentifier, &FilesIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Files)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestGanttsShowPdfWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.GanttsShowPdfWithResponse(context.TODO(), &GanttsShowPdfParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestGanttsShowPngWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.GanttsShowPngWithResponse(context.TODO(), &GanttsShowPngParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestGanttsShowProjectPdfWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.GanttsShowProjectPdfWithResponse(context.TODO(), projectIdentifier, &GanttsShowProjectPdfParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestGanttsShowProjectPngWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.GanttsShowProjectPngWithResponse(context.TODO(), projectIdentifier, &GanttsShowProjectPngParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestGroupsAddUsersWithResponse(t *testing.T) {
	c := newClient(t)

	userids := []int{userId}
	body := GroupsAddUsersJSONRequestBody{
		UserIds: &userids,
	}
	resp, err := c.GroupsAddUsersWithResponse(context.TODO(), groupId, &GroupsAddUsersParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestGroupsCreateWithResponse(t *testing.T) {
	c := newClient(t)

	name := t.Name()
	body := GroupsCreateJSONRequestBody{
		Group: &struct {
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			Name          *string "json:\"name,omitempty\""
			TwofaRequired *bool   "json:\"twofa_required,omitempty\""
			UserIds       *[]int  "json:\"user_ids,omitempty\""
		}{
			Name: &name,
		},
	}
	resp, err := c.GroupsCreateWithResponse(context.TODO(), &GroupsCreateParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON201)
	assertJson(t, resp.JSON201.Group)
	assertLength(t, resp.JSON201, resp.Body)
}

func TestGroupsDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.GroupsDestroyWithResponse(context.TODO(), groupId, &GroupsDestroyParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestGroupsIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.GroupsIndexWithResponse(context.TODO(), &GroupsIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Groups)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestGroupsRemoveUserWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.GroupsRemoveUserWithResponse(context.TODO(), groupId, userId, &GroupsRemoveUserParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestGroupsShowWithResponse(t *testing.T) {
	c := newClient(t)

	include := []string{
		"users",
		"memberships",
	}
	params := GroupsShowParams{
		Include: &include,
	}
	resp, err := c.GroupsShowWithResponse(context.TODO(), groupId, &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Group)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestGroupsUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	name := t.Name()
	body := GroupsUpdatePatchJSONRequestBody{
		Group: &struct {
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			Name          *string "json:\"name,omitempty\""
			TwofaRequired *bool   "json:\"twofa_required,omitempty\""
			UserIds       *[]int  "json:\"user_ids,omitempty\""
		}{
			Name: &name,
		},
	}
	resp, err := c.GroupsUpdatePatchWithResponse(context.TODO(), groupId, &GroupsUpdatePatchParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestGroupsUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	custom_values := map[string]interface{}{
		"19": "aaa",
	}
	custom_field_id := 20
	custom_field_value := interface{}("bbb")
	name := t.Name()
	twofaRequired := true
	userIds := []int{userId}
	body := GroupsUpdatePutJSONRequestBody{
		Group: &struct {
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			Name          *string "json:\"name,omitempty\""
			TwofaRequired *bool   "json:\"twofa_required,omitempty\""
			UserIds       *[]int  "json:\"user_ids,omitempty\""
		}{
			CustomFieldValues: &custom_values,
			CustomFields: &[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			}{
				{
					Id:    &custom_field_id,
					Value: &custom_field_value,
				},
			},
			Name:          &name,
			TwofaRequired: &twofaRequired,
			UserIds:       &userIds,
		},
	}
	resp, err := c.GroupsUpdatePutWithResponse(context.TODO(), groupId, &GroupsUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestIssueCategoriesCreateWithResponse(t *testing.T) {
	c := newClient(t)

	name := t.Name()
	body := IssueCategoriesCreateJSONRequestBody{
		IssueCategory: &struct {
			AssignedToId *int    "json:\"assigned_to_id,omitempty\""
			Name         *string "json:\"name,omitempty\""
		}{
			Name: &name,
		},
	}
	resp, err := c.IssueCategoriesCreateWithResponse(context.TODO(), projectIdentifier, &IssueCategoriesCreateParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON201)
	assertJson(t, resp.JSON201.IssueCategory)
	assertLength(t, resp.JSON201, resp.Body)
}

func TestIssueCategoriesDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssueCategoriesDestroyWithResponse(context.TODO(), issueCategoryId, &IssueCategoriesDestroyParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestIssueCategoriesIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssueCategoriesIndexWithResponse(context.TODO(), projectIdentifier, &IssueCategoriesIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.IssueCategories)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestIssueCategoriesShowWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssueCategoriesShowWithResponse(context.TODO(), issueCategoryId, &IssueCategoriesShowParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.IssueCategory)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestIssueCategoriesUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	name := t.Name()
	body := IssueCategoriesUpdatePatchJSONRequestBody{
		IssueCategory: &struct {
			AssignedToId *int    `json:"assigned_to_id,omitempty"`
			Name         *string `json:"name,omitempty"`
		}{
			Name: &name,
		},
	}
	resp, err := c.IssueCategoriesUpdatePatchWithResponse(context.TODO(), issueCategoryId, &IssueCategoriesUpdatePatchParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestIssueCategoriesUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	assignedToId := userId
	name := t.Name()
	body := IssueCategoriesUpdatePutJSONRequestBody{
		IssueCategory: &struct {
			AssignedToId *int    `json:"assigned_to_id,omitempty"`
			Name         *string `json:"name,omitempty"`
		}{
			AssignedToId: &assignedToId,
			Name:         &name,
		},
	}
	resp, err := c.IssueCategoriesUpdatePutWithResponse(context.TODO(), issueCategoryId, &IssueCategoriesUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestIssueRelationsCreateWithResponse(t *testing.T) {
	c := newClient(t)

	issueToId := strconv.Itoa(issueId + 1)
	relationType := "relates"
	body := IssueRelationsCreateJSONRequestBody{
		Relation: &struct {
			IssueToId    *string "json:\"issue_to_id,omitempty\""
			RelationType *string "json:\"relation_type,omitempty\""
		}{
			IssueToId:    &issueToId,
			RelationType: &relationType,
		},
	}
	resp, err := c.IssueRelationsCreateWithResponse(context.TODO(), issueId, &IssueRelationsCreateParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestIssueRelationsDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssueRelationsDestroyWithResponse(context.TODO(), issueRelatedId, &IssueRelationsDestroyParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestIssueRelationsIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssueRelationsIndexWithResponse(context.TODO(), issueId, &IssueRelationsIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Relations)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestIssueRelationsShowWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssueRelationsShowWithResponse(context.TODO(), issueRelatedId, &IssueRelationsShowParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Relation)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestIssuesCreateProjectWithResponse(t *testing.T) {
	c := newClient(t)

	subject := t.Name()
	body := IssuesCreateProjectJSONRequestBody{
		Issue: &struct {
			AssignedToId      *int                    "json:\"assigned_to_id,omitempty\""
			CategoryId        *int                    "json:\"category_id,omitempty\""
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			DeletedAttachmentIds *[]int              "json:\"deleted_attachment_ids,omitempty\""
			Description          *string             "json:\"description,omitempty\""
			DoneRatio            *int                "json:\"done_ratio,omitempty\""
			DueDate              *openapi_types.Date "json:\"due_date,omitempty\""
			EstimatedHours       *float32            "json:\"estimated_hours,omitempty\""
			FixedVersionId       *int                "json:\"fixed_version_id,omitempty\""
			IsPrivate            *bool               "json:\"is_private,omitempty\""
			Notes                *string             "json:\"notes,omitempty\""
			ParentIssueId        *int                "json:\"parent_issue_id,omitempty\""
			PriorityId           *int                "json:\"priority_id,omitempty\""
			PrivateNotes         *bool               "json:\"private_notes,omitempty\""
			ProjectId            *string             "json:\"project_id,omitempty\""
			StartDate            *openapi_types.Date "json:\"start_date,omitempty\""
			StatusId             *int                "json:\"status_id,omitempty\""
			Subject              *string             "json:\"subject,omitempty\""
			TrackerId            *int                "json:\"tracker_id,omitempty\""
			WatcherUserIds       *[]int              "json:\"watcher_user_ids,omitempty\""
		}{
			Subject: &subject,
		},
	}
	resp, err := c.IssuesCreateProjectWithResponse(context.TODO(), projectIdentifier, &IssuesCreateProjectParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON201)
	assertJson(t, resp.JSON201.Issue)
	//assertLength(t, resp.JSON201, resp.Body)
}

func TestIssuesCreateWithResponse(t *testing.T) {
	c := newClient(t)

	subject := t.Name()
	body := IssuesCreateJSONRequestBody{
		Issue: &struct {
			AssignedToId      *int                    `json:"assigned_to_id,omitempty"`
			CategoryId        *int                    `json:"category_id,omitempty"`
			CustomFieldValues *map[string]interface{} `json:"custom_field_values,omitempty"`
			CustomFields      *[]struct {
				Id       *int         `json:"id,omitempty"`
				Multiple *bool        `json:"multiple,omitempty"`
				Name     *string      `json:"name,omitempty"`
				Value    *interface{} `json:"value,omitempty"`
			} `json:"custom_fields,omitempty"`
			DeletedAttachmentIds *[]int              `json:"deleted_attachment_ids,omitempty"`
			Description          *string             `json:"description,omitempty"`
			DoneRatio            *int                `json:"done_ratio,omitempty"`
			DueDate              *openapi_types.Date `json:"due_date,omitempty"`
			EstimatedHours       *float32            `json:"estimated_hours,omitempty"`
			FixedVersionId       *int                `json:"fixed_version_id,omitempty"`
			IsPrivate            *bool               `json:"is_private,omitempty"`
			Notes                *string             `json:"notes,omitempty"`
			ParentIssueId        *int                `json:"parent_issue_id,omitempty"`
			PriorityId           *int                `json:"priority_id,omitempty"`
			PrivateNotes         *bool               `json:"private_notes,omitempty"`
			ProjectId            *string             `json:"project_id,omitempty"`
			StartDate            *openapi_types.Date `json:"start_date,omitempty"`
			StatusId             *int                `json:"status_id,omitempty"`
			Subject              *string             `json:"subject,omitempty"`
			TrackerId            *int                `json:"tracker_id,omitempty"`
			WatcherUserIds       *[]int              `json:"watcher_user_ids,omitempty"`
		}{
			Subject:   &subject,
			ProjectId: &projectIdentifier,
			TrackerId: &trackerId,
		},
	}
	resp, err := c.IssuesCreateWithResponse(context.TODO(), &IssuesCreateParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON201)
	assertJson(t, resp.JSON201.Issue)
	//assertLength(t, resp.JSON201, resp.Body)
}

func TestIssuesDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssuesDestroyWithResponse(context.TODO(), issueId, &IssuesDestroyParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestIssuesIndexCsvWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssuesIndexCsvWithResponse(context.TODO(), &IssuesIndexCsvParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestIssuesIndexPdfWithResponse(t *testing.T) {
	c := newClient(t)

	include := []string{
		"allowed_statuses",
		"attachments",
		"changesets",
		"children",
		"journals",
		"relations",
		"watchers",
	}
	params := IssuesIndexPdfParams{
		Include: &include,
	}
	resp, err := c.IssuesIndexPdfWithResponse(context.TODO(), &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestIssuesIndexProjectCsvWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssuesIndexProjectCsvWithResponse(context.TODO(), projectIdentifier, &IssuesIndexProjectCsvParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestIssuesIndexProjectPdfWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssuesIndexProjectPdfWithResponse(context.TODO(), projectIdentifier, &IssuesIndexProjectPdfParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestIssuesIndexProjectWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssuesIndexProjectWithResponse(context.TODO(), projectIdentifier, &IssuesIndexProjectParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Issues)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestIssuesIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssuesIndexWithResponse(context.TODO(), &IssuesIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Issues)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestIssuesShowPdfWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssuesShowPdfWithResponse(context.TODO(), issueId, &IssuesShowPdfParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestIssuesShowWithResponse(t *testing.T) {
	c := newClient(t)

	include := []string{
		"allowed_statuses",
		"attachments",
		"changesets",
		"children",
		"journals",
		"relations",
		"watchers",
	}
	params := IssuesShowParams{
		Include: &include,
	}

	resp, err := c.IssuesShowWithResponse(context.TODO(), issueId, &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Issue)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestIssueStatusesIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.IssueStatusesIndexWithResponse(context.TODO(), &IssueStatusesIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.IssueStatuses)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestIssuesUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	subject := t.Name()
	body := IssuesUpdatePatchJSONRequestBody{
		Issue: &struct {
			AssignedToId      *int                    "json:\"assigned_to_id,omitempty\""
			CategoryId        *int                    "json:\"category_id,omitempty\""
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			DeletedAttachmentIds *[]int              "json:\"deleted_attachment_ids,omitempty\""
			Description          *string             "json:\"description,omitempty\""
			DoneRatio            *int                "json:\"done_ratio,omitempty\""
			DueDate              *openapi_types.Date "json:\"due_date,omitempty\""
			EstimatedHours       *float32            "json:\"estimated_hours,omitempty\""
			FixedVersionId       *int                "json:\"fixed_version_id,omitempty\""
			IsPrivate            *bool               "json:\"is_private,omitempty\""
			Notes                *string             "json:\"notes,omitempty\""
			ParentIssueId        *int                "json:\"parent_issue_id,omitempty\""
			PriorityId           *int                "json:\"priority_id,omitempty\""
			PrivateNotes         *bool               "json:\"private_notes,omitempty\""
			ProjectId            *string             "json:\"project_id,omitempty\""
			StartDate            *openapi_types.Date "json:\"start_date,omitempty\""
			StatusId             *int                "json:\"status_id,omitempty\""
			Subject              *string             "json:\"subject,omitempty\""
			TrackerId            *int                "json:\"tracker_id,omitempty\""
			WatcherUserIds       *[]int              "json:\"watcher_user_ids,omitempty\""
		}{
			Subject:   &subject,
			ProjectId: &projectIdentifier,
		},
	}
	resp, err := c.IssuesUpdatePatchWithResponse(context.TODO(), issueId, &IssuesUpdatePatchParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestIssuesUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	custom_values := map[string]interface{}{
		"3": "1",
	}
	custom_field_id := 4
	custom_field_value := interface{}(versionId)
	description := t.Name() + "Description"
	doneRatio := 10
	estimatedHours := float32(5)
	fixedVersionId := versionId
	isPrivate := true
	notes := t.Name() + "Notes"
	parentIssueId := 1
	priorityId := 1
	privateNotes := true
	statusId := 1
	subject := t.Name()
	trackerId := trackerId
	watcherUserIds := []int{userId}
	body := IssuesUpdatePutJSONRequestBody{
		Issue: &struct {
			AssignedToId      *int                    "json:\"assigned_to_id,omitempty\""
			CategoryId        *int                    "json:\"category_id,omitempty\""
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			DeletedAttachmentIds *[]int              "json:\"deleted_attachment_ids,omitempty\""
			Description          *string             "json:\"description,omitempty\""
			DoneRatio            *int                "json:\"done_ratio,omitempty\""
			DueDate              *openapi_types.Date "json:\"due_date,omitempty\""
			EstimatedHours       *float32            "json:\"estimated_hours,omitempty\""
			FixedVersionId       *int                "json:\"fixed_version_id,omitempty\""
			IsPrivate            *bool               "json:\"is_private,omitempty\""
			Notes                *string             "json:\"notes,omitempty\""
			ParentIssueId        *int                "json:\"parent_issue_id,omitempty\""
			PriorityId           *int                "json:\"priority_id,omitempty\""
			PrivateNotes         *bool               "json:\"private_notes,omitempty\""
			ProjectId            *string             "json:\"project_id,omitempty\""
			StartDate            *openapi_types.Date "json:\"start_date,omitempty\""
			StatusId             *int                "json:\"status_id,omitempty\""
			Subject              *string             "json:\"subject,omitempty\""
			TrackerId            *int                "json:\"tracker_id,omitempty\""
			WatcherUserIds       *[]int              "json:\"watcher_user_ids,omitempty\""
		}{
			AssignedToId:      &userId,
			CategoryId:        &issueCategoryId,
			CustomFieldValues: &custom_values,
			CustomFields: &[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			}{
				{
					Id:    &custom_field_id,
					Value: &custom_field_value,
				},
			},
			Description: &description,
			DoneRatio:   &doneRatio,
			DueDate: &openapi_types.Date{
				Time: time.Now().AddDate(0, 0, 3),
			},
			EstimatedHours: &estimatedHours,
			FixedVersionId: &fixedVersionId,
			IsPrivate:      &isPrivate,
			Notes:          &notes,
			ParentIssueId:  &parentIssueId,
			PriorityId:     &priorityId,
			PrivateNotes:   &privateNotes,
			ProjectId:      &projectIdentifier,
			StartDate: &openapi_types.Date{
				Time: time.Now(),
			},
			StatusId:       &statusId,
			Subject:        &subject,
			TrackerId:      &trackerId,
			WatcherUserIds: &watcherUserIds,
		},
	}
	resp, err := c.IssuesUpdatePutWithResponse(context.TODO(), issueId, &IssuesUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestJournalsUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	notes := t.Name()
	body := JournalsUpdatePatchJSONRequestBody{
		Journal: &struct {
			Notes        *string "json:\"notes,omitempty\""
			PrivateNotes *bool   "json:\"private_notes,omitempty\""
		}{
			Notes: &notes,
		},
	}
	resp, err := c.JournalsUpdatePatchWithResponse(context.TODO(), journalId, &JournalsUpdatePatchParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestJournalsUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	notes := t.Name()
	privateNotes := true
	body := JournalsUpdatePutJSONRequestBody{
		Journal: &struct {
			Notes        *string "json:\"notes,omitempty\""
			PrivateNotes *bool   "json:\"private_notes,omitempty\""
		}{
			Notes:        &notes,
			PrivateNotes: &privateNotes,
		},
	}
	resp, err := c.JournalsUpdatePutWithResponse(context.TODO(), journalId, &JournalsUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestMembersCreateWithResponse(t *testing.T) {
	c := newClient(t)

	body := MembersCreateJSONRequestBody{
		Membership: &struct {
			RoleIds *[]int "json:\"role_ids,omitempty\""
			UserId  *int   "json:\"user_id,omitempty\""
			UserIds *[]int "json:\"user_ids,omitempty\""
		}{
			RoleIds: &[]int{roleId},
			UserId:  &userId,
		},
	}
	resp, err := c.MembersCreateWithResponse(context.TODO(), projectIdentifier, &MembersCreateParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON201)
	assertJson(t, resp.JSON201.Membership)
	assertLength(t, resp.JSON201, resp.Body)
}

func TestMembersDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.MembersDestroyWithResponse(context.TODO(), membershipId, &MembersDestroyParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestMembersIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.MembersIndexWithResponse(context.TODO(), projectIdentifier, &MembersIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Memberships)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestMembersShowWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.MembersShowWithResponse(context.TODO(), membershipId, &MembersShowParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
}

func TestMembersUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	body := MembersUpdatePatchJSONRequestBody{
		Membership: &struct {
			RoleIds *[]int "json:\"role_ids,omitempty\""
			UserId  *int   "json:\"user_id,omitempty\""
			UserIds *[]int "json:\"user_ids,omitempty\""
		}{
			RoleIds: &[]int{roleId},
		},
	}
	resp, err := c.MembersUpdatePatchWithResponse(context.TODO(), membershipId, &MembersUpdatePatchParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestMembersUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	body := MembersUpdatePutJSONRequestBody{
		Membership: &struct {
			RoleIds *[]int "json:\"role_ids,omitempty\""
			UserId  *int   "json:\"user_id,omitempty\""
			UserIds *[]int "json:\"user_ids,omitempty\""
		}{
			RoleIds: &[]int{roleId},
		},
	}
	resp, err := c.MembersUpdatePutWithResponse(context.TODO(), membershipId, &MembersUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestMyAccountPutWithResponse(t *testing.T) {
	c := newClient(t)

	firstNam := t.Name()
	body := MyAccountPutJSONRequestBody{
		User: &struct {
			Admin             *bool                   "json:\"admin,omitempty\""
			AuthSourceId      *int                    "json:\"auth_source_id,omitempty\""
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			Firstname          *string   "json:\"firstname,omitempty\""
			GeneratePassword   *bool     "json:\"generate_password,omitempty\""
			GroupIds           *[]int    "json:\"group_ids,omitempty\""
			Language           *string   "json:\"language,omitempty\""
			Lastname           *string   "json:\"lastname,omitempty\""
			Login              *string   "json:\"login,omitempty\""
			Mail               *string   "json:\"mail,omitempty\""
			MailNotification   *string   "json:\"mail_notification,omitempty\""
			MustChangePasswd   *bool     "json:\"must_change_passwd,omitempty\""
			NotifiedProjectIds *[]string "json:\"notified_project_ids,omitempty\""
			Password           *string   "json:\"password,omitempty\""
			Status             *int      "json:\"status,omitempty\""
		}{
			Firstname: &firstNam,
		},
	}
	resp, err := c.MyAccountPutWithResponse(context.TODO(), &MyAccountPutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestMyAccountWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.MyAccountWithResponse(context.TODO(), &MyAccountParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.User)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestNewsCreateProjectWithResponse(t *testing.T) {
	c := newClient(t)

	title := t.Name()
	description := t.Name() + "Description"
	body := NewsCreateProjectJSONRequestBody{
		News: &struct {
			Description *string "json:\"description,omitempty\""
			Summary     *string "json:\"summary,omitempty\""
			Title       *string "json:\"title,omitempty\""
		}{
			Title:       &title,
			Description: &description,
		},
	}
	resp, err := c.NewsCreateProjectWithResponse(context.TODO(), projectIdentifier, &NewsCreateProjectParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestNewsCreateWithResponse(t *testing.T) {
	c := newClient(t)

	title := t.Name()
	description := t.Name() + "Description"
	body := NewsCreateJSONRequestBody{
		News: &struct {
			Description *string "json:\"description,omitempty\""
			Summary     *string "json:\"summary,omitempty\""
			Title       *string "json:\"title,omitempty\""
		}{
			Title:       &title,
			Description: &description,
		},
		ProjectId: &projectIdentifier,
	}
	resp, err := c.NewsCreateWithResponse(context.TODO(), &NewsCreateParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestNewsDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.NewsDestroyWithResponse(context.TODO(), nwesId, &NewsDestroyParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestNewsIndexProjectWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.NewsIndexProjectWithResponse(context.TODO(), projectIdentifier, &NewsIndexProjectParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.News)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestNewsIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.NewsIndexWithResponse(context.TODO(), &NewsIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.News)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestNewsShowWithResponse(t *testing.T) {
	c := newClient(t)

	include := []string{
		"attachments",
		"comments",
	}
	paramms := NewsShowParams{
		Include: &include,
	}
	resp, err := c.NewsShowWithResponse(context.TODO(), nwesId, &paramms, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.News)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestNewsUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	title := t.Name()
	body := NewsUpdatePatchJSONRequestBody{
		News: &struct {
			Description *string "json:\"description,omitempty\""
			Summary     *string "json:\"summary,omitempty\""
			Title       *string "json:\"title,omitempty\""
		}{
			Title: &title,
		},
	}
	resp, err := c.NewsUpdatePatchWithResponse(context.TODO(), nwesId, &NewsUpdatePatchParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestNewsUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	description := t.Name() + "Description"
	summary := t.Name() + "Summary"
	title := t.Name() + "Title"
	body := NewsUpdatePutJSONRequestBody{
		News: &struct {
			Description *string "json:\"description,omitempty\""
			Summary     *string "json:\"summary,omitempty\""
			Title       *string "json:\"title,omitempty\""
		}{
			Description: &description,
			Summary:     &summary,
			Title:       &title,
		},
	}
	resp, err := c.NewsUpdatePutWithResponse(context.TODO(), nwesId, &NewsUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func ProjectsArchivePostWithResponse(t *testing.T) {
	c := newClient(t)

	params := ProjectsArchivePostParams{}
	resp, err := c.ProjectsArchivePostWithResponse(context.TODO(), projectIdentifier, &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestProjectsArchivePutWithResponse(t *testing.T) {
	c := newClient(t)

	params := ProjectsArchivePutParams{}
	resp, err := c.ProjectsArchivePutWithResponse(context.TODO(), projectIdentifier, &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestProjectsCreateWithResponse(t *testing.T) {
	c := newClient(t)

	params := ProjectsCreateParams{}

	name := t.Name()
	identifier := projectIdentifier + "-01"
	custom_values := map[string]interface{}{
		"1": "aaa",
		"2": &[]string{"a", "c"},
	}
	body := ProjectsCreateJSONRequestBody{}
	body.Project = &struct {
		CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
		CustomFields      *[]struct {
			Id       *int         "json:\"id,omitempty\""
			Multiple *bool        "json:\"multiple,omitempty\""
			Name     *string      "json:\"name,omitempty\""
			Value    *interface{} "json:\"value,omitempty\""
		} "json:\"custom_fields,omitempty\""
		DefaultAssignedToId *int      "json:\"default_assigned_to_id,omitempty\""
		DefaultIssueQueryId *int      "json:\"default_issue_query_id,omitempty\""
		DefaultVersionId    *int      "json:\"default_version_id,omitempty\""
		Description         *string   "json:\"description,omitempty\""
		EnabledModuleNames  *[]string "json:\"enabled_module_names,omitempty\""
		Homepage            *string   "json:\"homepage,omitempty\""
		Identifier          *string   "json:\"identifier,omitempty\""
		InheritMembers      *bool     "json:\"inherit_members,omitempty\""
		IsPublic            *bool     "json:\"is_public,omitempty\""
		IssueCustomFieldIds *[]int    "json:\"issue_custom_field_ids,omitempty\""
		Name                *string   "json:\"name,omitempty\""
		ParentId            *int      "json:\"parent_id,omitempty\""
		TrackerIds          *[]int    "json:\"tracker_ids,omitempty\""
	}{
		CustomFieldValues: &custom_values,
		Name:              &name,
		Identifier:        &identifier,
	}

	resp, err := c.ProjectsCreateWithResponse(context.TODO(), &params, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON201)
	assertJson(t, resp.JSON201.Project)
	//assertLength(t, resp.JSON201, resp.Body)
}

func TestProjectsDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	params := ProjectsDestroyParams{}
	resp, err := c.ProjectsDestroyWithResponse(context.TODO(), projectIdentifier+"-01", &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestProjectsIndexCsvWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.ProjectsIndexCsvWithResponse(context.TODO(), &ProjectsIndexCsvParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestProjectsIndexWithResponse(t *testing.T) {
	c := newClient(t)

	include := []string{
		"enabled_modules",
		"issue_categories",
		"issue_custom_fields",
		"time_entry_activities",
		"trackers",
	}
	params := ProjectsIndexParams{}
	params.Include = &include
	resp, err := c.ProjectsIndexWithResponse(context.TODO(), &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Projects)
	//assertLength(t, resp.JSON200, resp.Body)

	// Check custom fields.
	for _, p := range *resp.JSON200.Projects {
		for _, cf := range *p.CustomFields {
			assertCustomField(t, &cf)
		}
	}
}

func TestProjectsShowWithResponse(t *testing.T) {
	c := newClient(t)

	include := []string{
		"enabled_modules",
		"issue_categories",
		"issue_custom_fields",
		"time_entry_activities",
		"trackers",
	}
	params := ProjectsShowParams{}
	params.Include = &include
	resp, err := c.ProjectsShowWithResponse(context.TODO(), projectIdentifier, &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Project)
	assertLength(t, resp.JSON200, resp.Body)
}

func ProjectsUnarchivePostWithResponse(t *testing.T) {
	c := newClient(t)

	params := ProjectsUnarchivePostParams{}
	resp, err := c.ProjectsUnarchivePostWithResponse(context.TODO(), projectIdentifier, &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestProjectsUnarchivePutWithResponse(t *testing.T) {
	c := newClient(t)

	params := ProjectsUnarchivePutParams{}
	resp, err := c.ProjectsUnarchivePutWithResponse(context.TODO(), projectIdentifier, &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestProjectsUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	params := ProjectsUpdatePatchParams{}
	name := t.Name()
	body := ProjectsUpdatePatchJSONRequestBody{
		Project: &struct {
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			DefaultAssignedToId *int      "json:\"default_assigned_to_id,omitempty\""
			DefaultIssueQueryId *int      "json:\"default_issue_query_id,omitempty\""
			DefaultVersionId    *int      "json:\"default_version_id,omitempty\""
			Description         *string   "json:\"description,omitempty\""
			EnabledModuleNames  *[]string "json:\"enabled_module_names,omitempty\""
			Homepage            *string   "json:\"homepage,omitempty\""
			Identifier          *string   "json:\"identifier,omitempty\""
			InheritMembers      *bool     "json:\"inherit_members,omitempty\""
			IsPublic            *bool     "json:\"is_public,omitempty\""
			IssueCustomFieldIds *[]int    "json:\"issue_custom_field_ids,omitempty\""
			Name                *string   "json:\"name,omitempty\""
			ParentId            *int      "json:\"parent_id,omitempty\""
			TrackerIds          *[]int    "json:\"tracker_ids,omitempty\""
		}{
			Name: &name,
		},
	}
	resp, err := c.ProjectsUpdatePatchWithResponse(context.TODO(), projectIdentifier+"-01", &params, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestProjectsUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	custom_values := map[string]interface{}{
		"1": "aaa",
	}
	custom_field_id := 2
	custom_field_value := interface{}("a")
	defaultIssueQueryId := 2
	description := t.Name() + "Description"
	enabledModuleNames := []string{"issue_tracking", "wiki"}
	homepage := "ccc"
	inheritMembers := true
	isPublic := true
	issueCustomFieldIds := []int{4}
	name := t.Name()
	parentId := projectId
	trackerIds := []int{trackerId}
	body := ProjectsUpdatePutJSONRequestBody{
		Project: &struct {
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			DefaultAssignedToId *int      "json:\"default_assigned_to_id,omitempty\""
			DefaultIssueQueryId *int      "json:\"default_issue_query_id,omitempty\""
			DefaultVersionId    *int      "json:\"default_version_id,omitempty\""
			Description         *string   "json:\"description,omitempty\""
			EnabledModuleNames  *[]string "json:\"enabled_module_names,omitempty\""
			Homepage            *string   "json:\"homepage,omitempty\""
			Identifier          *string   "json:\"identifier,omitempty\""
			InheritMembers      *bool     "json:\"inherit_members,omitempty\""
			IsPublic            *bool     "json:\"is_public,omitempty\""
			IssueCustomFieldIds *[]int    "json:\"issue_custom_field_ids,omitempty\""
			Name                *string   "json:\"name,omitempty\""
			ParentId            *int      "json:\"parent_id,omitempty\""
			TrackerIds          *[]int    "json:\"tracker_ids,omitempty\""
		}{
			CustomFieldValues: &custom_values,
			CustomFields: &[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			}{
				{
					Id:    &custom_field_id,
					Value: &custom_field_value,
				},
			},
			DefaultAssignedToId: &userId,
			DefaultIssueQueryId: &defaultIssueQueryId,
			DefaultVersionId:    &versionId,
			Description:         &description,
			EnabledModuleNames:  &enabledModuleNames,
			Homepage:            &homepage,
			InheritMembers:      &inheritMembers,
			IsPublic:            &isPublic,
			IssueCustomFieldIds: &issueCustomFieldIds,
			Name:                &name,
			ParentId:            &parentId,
			TrackerIds:          &trackerIds,
		},
	}

	resp, err := c.ProjectsUpdatePutWithResponse(context.TODO(), projectIdentifier+"-01", &ProjectsUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestQueriesIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.QueriesIndexWithResponse(context.TODO(), &QueriesIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Queries)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestRepositoriesAddRelatedIssueWithResponse(t *testing.T) {
	c := newClient(t)

	body := RepositoriesAddRelatedIssueJSONRequestBody{
		IssueId: &issueId,
	}
	resp, err := c.RepositoriesAddRelatedIssueWithResponse(context.TODO(), projectIdentifier, repositoryIdentifier, revision, &RepositoriesAddRelatedIssueParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestRepositoriesRemoveRelatedIssueWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.RepositoriesRemoveRelatedIssueWithResponse(context.TODO(), projectIdentifier, repositoryIdentifier, revision, issueId, &RepositoriesRemoveRelatedIssueParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestRolesIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.RolesIndexWithResponse(context.TODO(), &RolesIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Roles)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestRolesShowWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.RolesShowWithResponse(context.TODO(), roleId, &RolesShowParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Role)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestSearchIndexProjectWithResponse(t *testing.T) {
	c := newClient(t)

	params := SearchIndexProjectParams{
		Q: "test",
	}
	resp, err := c.SearchIndexProjectWithResponse(context.TODO(), projectIdentifier, &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Results)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestSearchIndexWithResponse(t *testing.T) {
	c := newClient(t)

	params := SearchIndexParams{
		Q: "test",
	}
	resp, err := c.SearchIndexWithResponse(context.TODO(), &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Results)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestTimelogCreateIssueWithResponse(t *testing.T) {
	c := newClient(t)

	hours := float32(1.0)
	body := TimelogCreateIssueJSONRequestBody{
		TimeEntry: &struct {
			ActivityId        *int                    "json:\"activity_id,omitempty\""
			Comments          *string                 "json:\"comments,omitempty\""
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			Hours     *float32            "json:\"hours,omitempty\""
			IssueId   *int                "json:\"issue_id,omitempty\""
			ProjectId *int                "json:\"project_id,omitempty\""
			SpentOn   *openapi_types.Date "json:\"spent_on,omitempty\""
			UserId    *int                "json:\"user_id,omitempty\""
		}{
			ActivityId: &activityId,
			Hours:      &hours,
		},
	}
	resp, err := c.TimelogCreateIssueWithResponse(context.TODO(), issueId, &TimelogCreateIssueParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestTimelogCreateProjectWithResponse(t *testing.T) {
	c := newClient(t)

	hours := float32(1.0)
	body := TimelogCreateProjectJSONRequestBody{
		TimeEntry: &struct {
			ActivityId        *int                    "json:\"activity_id,omitempty\""
			Comments          *string                 "json:\"comments,omitempty\""
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			Hours     *float32            "json:\"hours,omitempty\""
			IssueId   *int                "json:\"issue_id,omitempty\""
			ProjectId *int                "json:\"project_id,omitempty\""
			SpentOn   *openapi_types.Date "json:\"spent_on,omitempty\""
			UserId    *int                "json:\"user_id,omitempty\""
		}{
			ActivityId: &activityId,
			Hours:      &hours,
		},
	}
	resp, err := c.TimelogCreateProjectWithResponse(context.TODO(), projectIdentifier, &TimelogCreateProjectParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestTimelogCreateWithResponse(t *testing.T) {
	c := newClient(t)

	hours := float32(0.5)
	body := TimelogCreateJSONRequestBody{
		TimeEntry: &struct {
			ActivityId        *int                    "json:\"activity_id,omitempty\""
			Comments          *string                 "json:\"comments,omitempty\""
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			Hours     *float32            "json:\"hours,omitempty\""
			IssueId   *int                "json:\"issue_id,omitempty\""
			ProjectId *int                "json:\"project_id,omitempty\""
			SpentOn   *openapi_types.Date "json:\"spent_on,omitempty\""
			UserId    *int                "json:\"user_id,omitempty\""
		}{
			ActivityId: &activityId,
			Hours:      &hours,
			ProjectId:  &projectId,
		},
	}
	resp, err := c.TimelogCreateWithResponse(context.TODO(), &TimelogCreateParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestTimelogDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.TimelogDestroyWithResponse(context.TODO(), timeEntityId, &TimelogDestroyParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestTimelogIndexCsvWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.TimelogIndexCsvWithResponse(context.TODO(), &TimelogIndexCsvParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestTimelogIndexProjectCsvWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.TimelogIndexProjectCsvWithResponse(context.TODO(), projectIdentifier, &TimelogIndexProjectCsvParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestTimelogIndexProjectWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.TimelogIndexProjectWithResponse(context.TODO(), projectIdentifier, &TimelogIndexProjectParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.TimeEntries)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestTimelogIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.TimelogIndexWithResponse(context.TODO(), &TimelogIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.TimeEntries)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestTimelogShowWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.TimelogShowWithResponse(context.TODO(), timeEntityId, &TimelogShowParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.TimeEntry)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestTimelogUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	hours := float32(1.5)
	body := TimelogUpdatePatchJSONRequestBody{
		TimeEntry: &struct {
			ActivityId        *int                    "json:\"activity_id,omitempty\""
			Comments          *string                 "json:\"comments,omitempty\""
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			Hours     *float32            "json:\"hours,omitempty\""
			IssueId   *int                "json:\"issue_id,omitempty\""
			ProjectId *int                "json:\"project_id,omitempty\""
			SpentOn   *openapi_types.Date "json:\"spent_on,omitempty\""
			UserId    *int                "json:\"user_id,omitempty\""
		}{
			Hours:     &hours,
			ProjectId: &projectId,
		},
	}
	resp, err := c.TimelogUpdatePatchWithResponse(context.TODO(), timeEntityId, &TimelogUpdatePatchParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestTimelogUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	activityId := 9
	comments := t.Name() + "Comments"
	custom_values := map[string]interface{}{
		"5": "1",
	}
	custom_field_id := 6
	custom_field_value := interface{}(userId)
	hours := float32(2.0)
	body := TimelogUpdatePutJSONRequestBody{
		TimeEntry: &struct {
			ActivityId        *int                    "json:\"activity_id,omitempty\""
			Comments          *string                 "json:\"comments,omitempty\""
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			Hours     *float32            "json:\"hours,omitempty\""
			IssueId   *int                "json:\"issue_id,omitempty\""
			ProjectId *int                "json:\"project_id,omitempty\""
			SpentOn   *openapi_types.Date "json:\"spent_on,omitempty\""
			UserId    *int                "json:\"user_id,omitempty\""
		}{
			ActivityId:        &activityId,
			Comments:          &comments,
			CustomFieldValues: &custom_values,
			CustomFields: &[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			}{
				{
					Id:    &custom_field_id,
					Value: &custom_field_value,
				},
			},
			Hours:   &hours,
			IssueId: &issueId,
			UserId:  &userId,
		},
	}
	resp, err := c.TimelogUpdatePutWithResponse(context.TODO(), timeEntityId, &TimelogUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestTrackersIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.TrackersIndexWithResponse(context.TODO(), &TrackersIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Trackers)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestUsersCreateWithResponse(t *testing.T) {
	c := newClient(t)

	params := UsersCreateParams{}

	login := "test-user"
	password := "password"
	firstname := "Test"
	lastname := "User"
	mail := "testuser@example.com"
	body := UsersCreateJSONRequestBody{}
	body.User = &struct {
		Admin             *bool                   "json:\"admin,omitempty\""
		AuthSourceId      *int                    "json:\"auth_source_id,omitempty\""
		CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
		CustomFields      *[]struct {
			Id       *int         "json:\"id,omitempty\""
			Multiple *bool        "json:\"multiple,omitempty\""
			Name     *string      "json:\"name,omitempty\""
			Value    *interface{} "json:\"value,omitempty\""
		} "json:\"custom_fields,omitempty\""
		Firstname          *string   "json:\"firstname,omitempty\""
		GeneratePassword   *bool     "json:\"generate_password,omitempty\""
		GroupIds           *[]int    "json:\"group_ids,omitempty\""
		Language           *string   "json:\"language,omitempty\""
		Lastname           *string   "json:\"lastname,omitempty\""
		Login              *string   "json:\"login,omitempty\""
		Mail               *string   "json:\"mail,omitempty\""
		MailNotification   *string   "json:\"mail_notification,omitempty\""
		MustChangePasswd   *bool     "json:\"must_change_passwd,omitempty\""
		NotifiedProjectIds *[]string "json:\"notified_project_ids,omitempty\""
		Password           *string   "json:\"password,omitempty\""
		Status             *int      "json:\"status,omitempty\""
	}{
		Login:     &login,
		Password:  &password,
		Firstname: &firstname,
		Lastname:  &lastname,
		Mail:      &mail,
	}

	resp, err := c.UsersCreateWithResponse(context.TODO(), &params, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON201)
	assertJson(t, resp.JSON201.User)
	//assertLength(t, resp.JSON201, resp.Body)
}

func TestUsersDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	params := UsersDestroyParams{}
	resp, err := c.UsersDestroyWithResponse(context.TODO(), strconv.Itoa(userId), &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestUsersIndexCsvWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.UsersIndexCsvWithResponse(context.TODO(), &UsersIndexCsvParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestUsersIndexWithResponse(t *testing.T) {
	c := newClient(t)

	include := []string{
		"auth_source",
		"groups",
		"memberships",
	}
	params := UsersIndexParams{}
	params.Include = &include
	resp, err := c.UsersIndexWithResponse(context.TODO(), &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Users)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestUsersShowWithResponse(t *testing.T) {
	c := newClient(t)

	include := []string{
		"auth_source",
		"groups",
		"memberships",
	}
	params := UsersShowParams{}
	params.Include = &include
	resp, err := c.UsersShowWithResponse(context.TODO(), strconv.Itoa(userId), &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.User)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestUsersUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	firstname := t.Name()[:30]
	body := UsersUpdatePatchJSONRequestBody{
		User: &struct {
			Admin             *bool                   "json:\"admin,omitempty\""
			AuthSourceId      *int                    "json:\"auth_source_id,omitempty\""
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			Firstname          *string   "json:\"firstname,omitempty\""
			GeneratePassword   *bool     "json:\"generate_password,omitempty\""
			GroupIds           *[]int    "json:\"group_ids,omitempty\""
			Language           *string   "json:\"language,omitempty\""
			Lastname           *string   "json:\"lastname,omitempty\""
			Login              *string   "json:\"login,omitempty\""
			Mail               *string   "json:\"mail,omitempty\""
			MailNotification   *string   "json:\"mail_notification,omitempty\""
			MustChangePasswd   *bool     "json:\"must_change_passwd,omitempty\""
			NotifiedProjectIds *[]string "json:\"notified_project_ids,omitempty\""
			Password           *string   "json:\"password,omitempty\""
			Status             *int      "json:\"status,omitempty\""
		}{
			Firstname: &firstname,
		},
	}
	resp, err := c.UsersUpdatePatchWithResponse(context.TODO(), strconv.Itoa(userId), &UsersUpdatePatchParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestUsersUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	autoWatchOn := []string{"issue_contributed_to"}
	commentsSorting := "asc"
	defaultIssueQuery := 3
	defaultProjectQuery := 6
	hideMail := false
	historyDefaultTab := "changesets"
	noSelfNotified := false
	notifyAboutHighPriorityIssues := true
	recentlyUsedProjects := 6
	textareaFont := "monospace"
	timeZone := "Tokyo"
	toolbarLanguageOptions := "c"
	warnOnLeavingUnsaved := "1" // true(1) or false(0)
	sendInformation := false
	admin := true
	custom_values := map[string]interface{}{
		"11": "1", // boot: true(1) or false(0)
	}
	custom_field_id := 12
	custom_field_value := interface{}("text")
	firstname := (t.Name() + "Firstname")[:30]
	generatePassword := false
	groupIds := []int{groupId}
	language := "ja" // 2code
	lastname := (t.Name() + "Lastname")[:30]
	mail := t.Name() + "@example.com"
	mailNotification := "all"
	mustChangePasswd := false
	notifiedProjectIds := []string{projectIdentifier}
	status := 1
	body := UsersUpdatePutJSONRequestBody{
		Pref: &struct {
			AutoWatchOn                   *[]string "json:\"auto_watch_on,omitempty\""
			CommentsSorting               *string   "json:\"comments_sorting,omitempty\""
			DefaultIssueQuery             *int      "json:\"default_issue_query,omitempty\""
			DefaultProjectQuery           *int      "json:\"default_project_query,omitempty\""
			HideMail                      *bool     "json:\"hide_mail,omitempty\""
			HistoryDefaultTab             *string   "json:\"history_default_tab,omitempty\""
			NoSelfNotified                *bool     "json:\"no_self_notified,omitempty\""
			NotifyAboutHighPriorityIssues *bool     "json:\"notify_about_high_priority_issues,omitempty\""
			RecentlyUsedProjects          *int      "json:\"recently_used_projects,omitempty\""
			TextareaFont                  *string   "json:\"textarea_font,omitempty\""
			TimeZone                      *string   "json:\"time_zone,omitempty\""
			ToolbarLanguageOptions        *string   "json:\"toolbar_language_options,omitempty\""
			WarnOnLeavingUnsaved          *string   "json:\"warn_on_leaving_unsaved,omitempty\""
		}{
			AutoWatchOn:                   &autoWatchOn,
			CommentsSorting:               &commentsSorting,
			DefaultIssueQuery:             &defaultIssueQuery,
			DefaultProjectQuery:           &defaultProjectQuery,
			HideMail:                      &hideMail,
			HistoryDefaultTab:             &historyDefaultTab,
			NoSelfNotified:                &noSelfNotified,
			NotifyAboutHighPriorityIssues: &notifyAboutHighPriorityIssues,
			RecentlyUsedProjects:          &recentlyUsedProjects,
			TextareaFont:                  &textareaFont,
			TimeZone:                      &timeZone,
			ToolbarLanguageOptions:        &toolbarLanguageOptions,
			WarnOnLeavingUnsaved:          &warnOnLeavingUnsaved,
		},
		SendInformation: &sendInformation,
		User: &struct {
			Admin             *bool                   "json:\"admin,omitempty\""
			AuthSourceId      *int                    "json:\"auth_source_id,omitempty\""
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			Firstname          *string   "json:\"firstname,omitempty\""
			GeneratePassword   *bool     "json:\"generate_password,omitempty\""
			GroupIds           *[]int    "json:\"group_ids,omitempty\""
			Language           *string   "json:\"language,omitempty\""
			Lastname           *string   "json:\"lastname,omitempty\""
			Login              *string   "json:\"login,omitempty\""
			Mail               *string   "json:\"mail,omitempty\""
			MailNotification   *string   "json:\"mail_notification,omitempty\""
			MustChangePasswd   *bool     "json:\"must_change_passwd,omitempty\""
			NotifiedProjectIds *[]string "json:\"notified_project_ids,omitempty\""
			Password           *string   "json:\"password,omitempty\""
			Status             *int      "json:\"status,omitempty\""
		}{
			Admin:             &admin,
			CustomFieldValues: &custom_values,
			CustomFields: &[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			}{
				{
					Id:    &custom_field_id,
					Value: &custom_field_value,
				},
			},
			Firstname:          &firstname,
			GeneratePassword:   &generatePassword,
			GroupIds:           &groupIds,
			Language:           &language,
			Lastname:           &lastname,
			Mail:               &mail,
			MailNotification:   &mailNotification,
			MustChangePasswd:   &mustChangePasswd,
			NotifiedProjectIds: &notifiedProjectIds,
			Status:             &status,
		},
	}

	resp, err := c.UsersUpdatePutWithResponse(context.TODO(), strconv.Itoa(userId), &UsersUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestVersionsCreateWithResponse(t *testing.T) {
	c := newClient(t)

	name := t.Name()
	body := VersionsCreateJSONRequestBody{
		Version: &struct {
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			DefaultProjectVersion *bool               "json:\"default_project_version,omitempty\""
			Description           *string             "json:\"description,omitempty\""
			DueDate               *openapi_types.Date "json:\"due_date,omitempty\""
			EffectiveDate         *openapi_types.Date "json:\"effective_date,omitempty\""
			Name                  *string             "json:\"name,omitempty\""
			Sharing               *string             "json:\"sharing,omitempty\""
			Status                *string             "json:\"status,omitempty\""
			WikiPageTitle         *string             "json:\"wiki_page_title,omitempty\""
		}{
			Name: &name,
		},
	}
	resp, err := c.VersionsCreateWithResponse(context.TODO(), projectIdentifier, &VersionsCreateParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON201)
}

func TestVersionsDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.VersionsDestroyWithResponse(context.TODO(), versionId, &VersionsDestroyParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestVersionsIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.VersionsIndexWithResponse(context.TODO(), projectIdentifier, &VersionsIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Versions)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestVersionsShowTxtWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.VersionsShowTxtWithResponse(context.TODO(), versionId, &VersionsShowTxtParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestVersionsShowWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.VersionsShowWithResponse(context.TODO(), versionId, &VersionsShowParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.Version)
	//assertLength(t, resp.JSON200, resp.Body)
}

func TestVersionsUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	name := t.Name()
	body := VersionsUpdatePatchJSONRequestBody{
		Version: &struct {
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			DefaultProjectVersion *bool               "json:\"default_project_version,omitempty\""
			Description           *string             "json:\"description,omitempty\""
			DueDate               *openapi_types.Date "json:\"due_date,omitempty\""
			EffectiveDate         *openapi_types.Date "json:\"effective_date,omitempty\""
			Name                  *string             "json:\"name,omitempty\""
			Sharing               *string             "json:\"sharing,omitempty\""
			Status                *string             "json:\"status,omitempty\""
			WikiPageTitle         *string             "json:\"wiki_page_title,omitempty\""
		}{
			Name: &name,
		},
	}
	resp, err := c.VersionsUpdatePatchWithResponse(context.TODO(), versionId, &VersionsUpdatePatchParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestVersionsUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	defaultProjectVersion := true
	description := t.Name() + "Description"
	name := t.Name()
	sharing := "tree"
	status := "open"
	wikiTitle := "wiki"
	custom_values := map[string]interface{}{
		"7": "aaaa",
	}
	custom_field_id := 8
	custom_field_value := interface{}(0.1)
	body := VersionsUpdatePutJSONRequestBody{
		Version: &struct {
			CustomFieldValues *map[string]interface{} "json:\"custom_field_values,omitempty\""
			CustomFields      *[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			} "json:\"custom_fields,omitempty\""
			DefaultProjectVersion *bool               "json:\"default_project_version,omitempty\""
			Description           *string             "json:\"description,omitempty\""
			DueDate               *openapi_types.Date "json:\"due_date,omitempty\""
			EffectiveDate         *openapi_types.Date "json:\"effective_date,omitempty\""
			Name                  *string             "json:\"name,omitempty\""
			Sharing               *string             "json:\"sharing,omitempty\""
			Status                *string             "json:\"status,omitempty\""
			WikiPageTitle         *string             "json:\"wiki_page_title,omitempty\""
		}{
			CustomFieldValues: &custom_values,
			CustomFields: &[]struct {
				Id       *int         "json:\"id,omitempty\""
				Multiple *bool        "json:\"multiple,omitempty\""
				Name     *string      "json:\"name,omitempty\""
				Value    *interface{} "json:\"value,omitempty\""
			}{
				{
					Id:    &custom_field_id,
					Value: &custom_field_value,
				},
			},
			DefaultProjectVersion: &defaultProjectVersion,
			Description:           &description,
			EffectiveDate: &openapi_types.Date{
				Time: time.Now(),
			},
			Name:          &name,
			Sharing:       &sharing,
			Status:        &status,
			WikiPageTitle: &wikiTitle,
		},
	}
	resp, err := c.VersionsUpdatePutWithResponse(context.TODO(), versionId, &VersionsUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestWatchersCreateIssueWithResponse(t *testing.T) {
	c := newClient(t)

	body := WatchersCreateIssueJSONRequestBody{
		Watcher: &struct {
			UserId  *int   "json:\"user_id,omitempty\""
			UserIds *[]int "json:\"user_ids,omitempty\""
		}{
			UserId: &userId,
		},
	}
	resp, err := c.WatchersCreateIssueWithResponse(context.TODO(), issueId, &WatchersCreateIssueParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestWatchersCreateWithResponse(t *testing.T) {
	c := newClient(t)

	objectType := "issue"
	body := WatchersCreateJSONRequestBody{
		ObjectId:   &issueId,
		ObjectType: &objectType,
		Watcher: &struct {
			UserId  *int   "json:\"user_id,omitempty\""
			UserIds *[]int "json:\"user_ids,omitempty\""
		}{
			UserIds: &[]int{userId},
		},
	}
	resp, err := c.WatchersCreateWithResponse(context.TODO(), &WatchersCreateParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestWatchersDestroyIssueWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.WatchersDestroyIssueWithResponse(context.TODO(), issueId, userId, &WatchersDestroyIssueParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestWatchersDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	objectType := "issue"
	params := WatchersDestroyParams{
		ObjectType: objectType,
		ObjectId:   issueId,
		UserId:     userId,
	}
	resp, err := c.WatchersDestroyWithResponse(context.TODO(), &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestWikiDestroyWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.WikiDestroyWithResponse(context.TODO(), projectIdentifier, wikiTitle, &WikiDestroyParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestWikiIndexWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.WikiIndexWithResponse(context.TODO(), projectIdentifier, &WikiIndexParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.WikiPages)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestWikiShowPdfWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.WikiShowPdfWithResponse(context.TODO(), projectIdentifier, wikiTitle, &WikiShowPdfParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestWikiShowRootWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.WikiShowRootWithResponse(context.TODO(), projectIdentifier, &WikiShowRootParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.WikiPage)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestWikiShowTxtWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.WikiShowTxtWithResponse(context.TODO(), projectIdentifier, wikiTitle, &WikiShowTxtParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestWikiShowVersionPdfWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.WikiShowVersionPdfWithResponse(context.TODO(), projectIdentifier, wikiTitle, wikiVersion, &WikiShowVersionPdfParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestWikiShowVersionTxtWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.WikiShowVersionTxtWithResponse(context.TODO(), projectIdentifier, wikiTitle, wikiVersion, &WikiShowVersionTxtParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.Body)
}

func TestWikiShowVersionWithResponse(t *testing.T) {
	c := newClient(t)

	resp, err := c.WikiShowVersionWithResponse(context.TODO(), projectIdentifier, wikiTitle, wikiVersion, &WikiShowVersionParams{}, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.WikiPage)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestWikiShowWithResponse(t *testing.T) {
	c := newClient(t)

	include := []string{
		"attachments",
	}
	params := WikiShowParams{
		Include: &include,
	}
	resp, err := c.WikiShowWithResponse(context.TODO(), projectIdentifier, wikiTitle, &params, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON200)
	assertJson(t, resp.JSON200.WikiPage)
	assertLength(t, resp.JSON200, resp.Body)
}

func TestWikiUpdatePatchWithResponse(t *testing.T) {
	c := newClient(t)

	text := t.Name()
	body := WikiUpdatePatchJSONRequestBody{
		WikiPage: &struct {
			Comments    *string "json:\"comments,omitempty\""
			ParentTitle *string "json:\"parent_title,omitempty\""
			Text        *string "json:\"text,omitempty\""
			Version     *int    "json:\"version,omitempty\""
		}{
			Text: &text,
		},
	}
	resp, err := c.WikiUpdatePatchWithResponse(context.TODO(), projectIdentifier, wikiTitle, &WikiUpdatePatchParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
}

func TestWikiUpdatePutWithResponse(t *testing.T) {
	c := newClient(t)

	comments := t.Name() + "Comments"
	text := t.Name()
	parentTitle := "wiki"
	version := 1
	body := WikiUpdatePutJSONRequestBody{
		WikiPage: &struct {
			Comments    *string `json:"comments,omitempty"`
			ParentTitle *string `json:"parent_title,omitempty"`
			Text        *string `json:"text,omitempty"`
			Version     *int    `json:"version,omitempty"`
		}{
			Comments:    &comments,
			ParentTitle: &parentTitle,
			Text:        &text,
			Version:     &version,
		},
	}
	resp, err := c.WikiUpdatePutWithResponse(context.TODO(), projectIdentifier, wikiTitle+"01", &WikiUpdatePutParams{}, body, basicAuth)

	assertError(t, err)
	assertHTTPStatus(t, resp.HTTPResponse, resp.Body)
	assertResponseBody(t, resp.JSON201)
	assertJson(t, resp.JSON201.WikiPage)
	//assertLength(t, resp.JSON201, resp.Body)
}
