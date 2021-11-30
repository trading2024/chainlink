package resolver

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	gqlerrors "github.com/graph-gophers/graphql-go/errors"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink/core/services/directrequest"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/testdata/testspecs"
)

// This tests the main fields on the job results. Embedded spec testing is done
// in the `spec_test` file
func TestResolver_Jobs(t *testing.T) {
	var (
		externalJobID = uuid.Must(uuid.FromString("00000000-0000-0000-0000-000000000001"))

		query = `
			query GetJobs {
				jobs {
					results {
						id
						createdAt
						externalJobID
						maxTaskDuration
						name
						schemaVersion
						spec {
							__typename
						}
						runs {
							id
							allErrors
							outputs
							createdAt
						}
						observationSource
					}
					metadata {
						total
					}
				}
			}`
	)

	testCases := []GQLTestCase{
		unauthorizedTestCase(GQLTestCase{query: query}, "jobs"),
		{
			name:          "get jobs success",
			authenticated: true,
			before: func(f *gqlTestFramework) {
				plnSpecID := int32(12)

				f.App.On("JobORM").Return(f.Mocks.jobORM)
				f.Mocks.jobORM.On("PipelineRunsByJobsIDs", []int32{plnSpecID}).Return([]pipeline.Run{
					{
						ID:             int64(1),
						PipelineSpecID: plnSpecID,
						State:          pipeline.RunStatusRunning,
						Outputs:        pipeline.JSONSerializable{Valid: false},
						AllErrors:      pipeline.RunErrors{},
						CreatedAt:      f.Timestamp(),
						FinishedAt:     null.Time{},
					},
				}, nil)
				f.Mocks.jobORM.On("FindJobs", 0, 50).Return([]job.Job{
					{
						ID:                          1,
						Name:                        null.StringFrom("job1"),
						SchemaVersion:               1,
						MaxTaskDuration:             models.Interval(1 * time.Second),
						ExternalJobID:               externalJobID,
						CreatedAt:                   f.Timestamp(),
						Type:                        job.OffchainReporting,
						PipelineSpecID:              plnSpecID,
						OffchainreportingOracleSpec: &job.OffchainReportingOracleSpec{},
						PipelineSpec: &pipeline.Spec{
							DotDagSource: "ds1 [type=bridge name=voter_turnout];",
						},
					},
				}, 1, nil)
			},
			query: query,
			result: `
				{
					"jobs": {
						"results": [{
							"id": "1",
							"createdAt": "2021-01-01T00:00:00Z",
							"externalJobID": "00000000-0000-0000-0000-000000000001",
							"maxTaskDuration": "1s",
							"name": "job1",
							"schemaVersion": 1,
							"spec": {
								"__typename": "OCRSpec"
							},
							"runs": [
								{
									"id": "1",
									"allErrors": [],
									"outputs": ["error: unable to retrieve outputs"],
									"createdAt": "2021-01-01T00:00:00Z"
								}
							],
							"observationSource": "ds1 [type=bridge name=voter_turnout];"
						}],
						"metadata": {
							"total": 1
						}
					}
				}`,
		},
	}

	RunGQLTests(t, testCases)
}

func TestResolver_Job(t *testing.T) {
	var (
		id            = int32(1)
		externalJobID = uuid.Must(uuid.FromString("00000000-0000-0000-0000-000000000001"))

		query = `
			query GetJob {
				job(id: "1") {
					... on Job {
						id
						createdAt
						externalJobID
						maxTaskDuration
						name
						schemaVersion
						spec {
							__typename
						}
						observationSource
					}
					... on NotFoundError {
						code
						message
					}
				}
			}
		`
	)

	testCases := []GQLTestCase{
		unauthorizedTestCase(GQLTestCase{query: query}, "job"),
		{
			name:          "success",
			authenticated: true,
			before: func(f *gqlTestFramework) {
				f.App.On("JobORM").Return(f.Mocks.jobORM)
				f.Mocks.jobORM.On("FindJobTx", id).Return(job.Job{
					ID:                          1,
					Name:                        null.StringFrom("job1"),
					SchemaVersion:               1,
					MaxTaskDuration:             models.Interval(1 * time.Second),
					ExternalJobID:               externalJobID,
					CreatedAt:                   f.Timestamp(),
					Type:                        job.OffchainReporting,
					OffchainreportingOracleSpec: &job.OffchainReportingOracleSpec{},
					PipelineSpec: &pipeline.Spec{
						DotDagSource: "ds1 [type=bridge name=voter_turnout];",
					},
				}, nil)
			},
			query: query,
			result: `
				{
					"job": {
						"id": "1",
						"createdAt": "2021-01-01T00:00:00Z",
						"externalJobID": "00000000-0000-0000-0000-000000000001",
						"maxTaskDuration": "1s",
						"name": "job1",
						"schemaVersion": 1,
						"spec": {
							"__typename": "OCRSpec"
						},
						"observationSource": "ds1 [type=bridge name=voter_turnout];"
					}
				}
			`,
		},
		{
			name:          "not found",
			authenticated: true,
			before: func(f *gqlTestFramework) {
				f.App.On("JobORM").Return(f.Mocks.jobORM)
				f.Mocks.jobORM.On("FindJobTx", id).Return(job.Job{}, sql.ErrNoRows)
			},
			query: query,
			result: `
				{
					"job": {
						"code": "NOT_FOUND",
						"message": "job not found"
					}
				}
			`,
		},
	}

	RunGQLTests(t, testCases)
}

func TestResolver_CreateJob(t *testing.T) {
	t.Parallel()

	mutation := `
		mutation CreateJob($input: CreateJobInput!) {
			createJob(input: $input) {
				... on CreateJobSuccess {
					job {
						id
						createdAt
						externalJobID
						maxTaskDuration
						name
						schemaVersion
					}
				}
				... on InputErrors {
					errors {
						path
						message
						code
					}
				}
			}
		}`
	variables := map[string]interface{}{
		"input": map[string]interface{}{
			"TOML": testspecs.DirectRequestSpec,
		},
	}
	invalid := map[string]interface{}{
		"input": map[string]interface{}{
			"TOML": "some wrong value",
		},
	}
	jb, err := directrequest.ValidatedDirectRequestSpec(testspecs.DirectRequestSpec)
	assert.NoError(t, err)

	d, err := json.Marshal(map[string]interface{}{
		"createJob": map[string]interface{}{
			"job": map[string]interface{}{
				"id":              "0",
				"maxTaskDuration": "0s",
				"name":            jb.Name,
				"schemaVersion":   1,
				"createdAt":       "0001-01-01T00:00:00Z",
				"externalJobID":   jb.ExternalJobID.String(),
			},
		},
	})
	assert.NoError(t, err)
	expected := string(d)

	gError := errors.New("error")

	testCases := []GQLTestCase{
		unauthorizedTestCase(GQLTestCase{query: mutation, variables: variables}, "createJob"),
		{
			name:          "success",
			authenticated: true,
			before: func(f *gqlTestFramework) {
				f.App.On("GetConfig").Return(f.Mocks.cfg)
				f.App.On("AddJobV2", mock.Anything, &jb).Return(nil)
			},
			query:     mutation,
			variables: variables,
			result:    expected,
		},
		{
			name:          "invalid TOML error",
			authenticated: true,
			query:         mutation,
			variables:     invalid,
			result: `
				{
					"createJob": {
						"errors": [{
							"code": "INVALID_INPUT",
							"message": "failed to parse TOML: (1, 6): was expecting token =, but got \"wrong\" instead",
							"path": "TOML spec"
						}]
					}
				}`,
		},
		{
			name:          "generic error when adding the job",
			authenticated: true,
			before: func(f *gqlTestFramework) {
				f.App.On("GetConfig").Return(f.Mocks.cfg)
				f.App.On("AddJobV2", mock.Anything, &jb).Return(gError)
			},
			query:     mutation,
			variables: variables,
			result:    `null`,
			errors: []*gqlerrors.QueryError{
				{
					Extensions:    nil,
					ResolverError: gError,
					Path:          []interface{}{"createJob"},
					Message:       gError.Error(),
				},
			},
		},
	}

	RunGQLTests(t, testCases)
}
