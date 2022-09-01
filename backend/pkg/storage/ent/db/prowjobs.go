// Code generated by entc, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/prowjobs"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
)

// ProwJobs is the model entity for the ProwJobs schema.
type ProwJobs struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// JobID holds the value of the "job_id" field.
	JobID string `json:"job_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration float64 `json:"duration,omitempty"`
	// TestsCount holds the value of the "tests_count" field.
	TestsCount int64 `json:"tests_count,omitempty"`
	// FailedCount holds the value of the "failed_count" field.
	FailedCount int64 `json:"failed_count,omitempty"`
	// SkippedCount holds the value of the "skipped_count" field.
	SkippedCount int64 `json:"skipped_count,omitempty"`
	// JobName holds the value of the "job_name" field.
	JobName string `json:"job_name,omitempty"`
	// JobType holds the value of the "job_type" field.
	JobType string `json:"job_type,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// JobURL holds the value of the "job_url" field.
	JobURL string `json:"job_url,omitempty"`
	// CiFailed holds the value of the "ci_failed" field.
	CiFailed int16 `json:"ci_failed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProwJobsQuery when eager-loading is set.
	Edges                ProwJobsEdges `json:"edges"`
	repository_prow_jobs *uuid.UUID
}

// ProwJobsEdges holds the relations/edges for other nodes in the graph.
type ProwJobsEdges struct {
	// ProwJobs holds the value of the prow_jobs edge.
	ProwJobs *Repository `json:"prow_jobs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProwJobsOrErr returns the ProwJobs value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProwJobsEdges) ProwJobsOrErr() (*Repository, error) {
	if e.loadedTypes[0] {
		if e.ProwJobs == nil {
			// The edge prow_jobs was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: repository.Label}
		}
		return e.ProwJobs, nil
	}
	return nil, &NotLoadedError{edge: "prow_jobs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProwJobs) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case prowjobs.FieldDuration:
			values[i] = new(sql.NullFloat64)
		case prowjobs.FieldID, prowjobs.FieldTestsCount, prowjobs.FieldFailedCount, prowjobs.FieldSkippedCount, prowjobs.FieldCiFailed:
			values[i] = new(sql.NullInt64)
		case prowjobs.FieldJobID, prowjobs.FieldJobName, prowjobs.FieldJobType, prowjobs.FieldState, prowjobs.FieldJobURL:
			values[i] = new(sql.NullString)
		case prowjobs.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case prowjobs.ForeignKeys[0]: // repository_prow_jobs
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProwJobs", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProwJobs fields.
func (pj *ProwJobs) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case prowjobs.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pj.ID = int(value.Int64)
		case prowjobs.FieldJobID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field job_id", values[i])
			} else if value.Valid {
				pj.JobID = value.String
			}
		case prowjobs.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pj.CreatedAt = value.Time
			}
		case prowjobs.FieldDuration:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				pj.Duration = value.Float64
			}
		case prowjobs.FieldTestsCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tests_count", values[i])
			} else if value.Valid {
				pj.TestsCount = value.Int64
			}
		case prowjobs.FieldFailedCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field failed_count", values[i])
			} else if value.Valid {
				pj.FailedCount = value.Int64
			}
		case prowjobs.FieldSkippedCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field skipped_count", values[i])
			} else if value.Valid {
				pj.SkippedCount = value.Int64
			}
		case prowjobs.FieldJobName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field job_name", values[i])
			} else if value.Valid {
				pj.JobName = value.String
			}
		case prowjobs.FieldJobType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field job_type", values[i])
			} else if value.Valid {
				pj.JobType = value.String
			}
		case prowjobs.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				pj.State = value.String
			}
		case prowjobs.FieldJobURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field job_url", values[i])
			} else if value.Valid {
				pj.JobURL = value.String
			}
		case prowjobs.FieldCiFailed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ci_failed", values[i])
			} else if value.Valid {
				pj.CiFailed = int16(value.Int64)
			}
		case prowjobs.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field repository_prow_jobs", values[i])
			} else if value.Valid {
				pj.repository_prow_jobs = new(uuid.UUID)
				*pj.repository_prow_jobs = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryProwJobs queries the "prow_jobs" edge of the ProwJobs entity.
func (pj *ProwJobs) QueryProwJobs() *RepositoryQuery {
	return (&ProwJobsClient{config: pj.config}).QueryProwJobs(pj)
}

// Update returns a builder for updating this ProwJobs.
// Note that you need to call ProwJobs.Unwrap() before calling this method if this ProwJobs
// was returned from a transaction, and the transaction was committed or rolled back.
func (pj *ProwJobs) Update() *ProwJobsUpdateOne {
	return (&ProwJobsClient{config: pj.config}).UpdateOne(pj)
}

// Unwrap unwraps the ProwJobs entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pj *ProwJobs) Unwrap() *ProwJobs {
	tx, ok := pj.config.driver.(*txDriver)
	if !ok {
		panic("db: ProwJobs is not a transactional entity")
	}
	pj.config.driver = tx.drv
	return pj
}

// String implements the fmt.Stringer.
func (pj *ProwJobs) String() string {
	var builder strings.Builder
	builder.WriteString("ProwJobs(")
	builder.WriteString(fmt.Sprintf("id=%v", pj.ID))
	builder.WriteString(", job_id=")
	builder.WriteString(pj.JobID)
	builder.WriteString(", created_at=")
	builder.WriteString(pj.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", pj.Duration))
	builder.WriteString(", tests_count=")
	builder.WriteString(fmt.Sprintf("%v", pj.TestsCount))
	builder.WriteString(", failed_count=")
	builder.WriteString(fmt.Sprintf("%v", pj.FailedCount))
	builder.WriteString(", skipped_count=")
	builder.WriteString(fmt.Sprintf("%v", pj.SkippedCount))
	builder.WriteString(", job_name=")
	builder.WriteString(pj.JobName)
	builder.WriteString(", job_type=")
	builder.WriteString(pj.JobType)
	builder.WriteString(", state=")
	builder.WriteString(pj.State)
	builder.WriteString(", job_url=")
	builder.WriteString(pj.JobURL)
	builder.WriteString(", ci_failed=")
	builder.WriteString(fmt.Sprintf("%v", pj.CiFailed))
	builder.WriteByte(')')
	return builder.String()
}

// ProwJobsSlice is a parsable slice of ProwJobs.
type ProwJobsSlice []*ProwJobs

func (pj ProwJobsSlice) config(cfg config) {
	for _i := range pj {
		pj[_i].config = cfg
	}
}