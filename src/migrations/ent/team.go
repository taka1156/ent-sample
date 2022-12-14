// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sample/migrations/ent/team"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Team is the model entity for the Team schema.
type Team struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TeamName holds the value of the "team_name" field.
	TeamName string `json:"team_name,omitempty"`
	// TeamIcon holds the value of the "team_icon" field.
	TeamIcon string `json:"team_icon,omitempty"`
	// TeamMemo holds the value of the "team_memo" field.
	TeamMemo string `json:"team_memo,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeamQuery when eager-loading is set.
	Edges TeamEdges `json:"edges"`
}

// TeamEdges holds the relations/edges for other nodes in the graph.
type TeamEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Team) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case team.FieldID:
			values[i] = new(sql.NullInt64)
		case team.FieldTeamName, team.FieldTeamIcon, team.FieldTeamMemo:
			values[i] = new(sql.NullString)
		case team.FieldCreatedAt, team.FieldUpdatedAt, team.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Team", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Team fields.
func (t *Team) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case team.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case team.FieldTeamName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team_name", values[i])
			} else if value.Valid {
				t.TeamName = value.String
			}
		case team.FieldTeamIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team_icon", values[i])
			} else if value.Valid {
				t.TeamIcon = value.String
			}
		case team.FieldTeamMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team_memo", values[i])
			} else if value.Valid {
				t.TeamMemo = value.String
			}
		case team.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case team.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case team.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Team entity.
func (t *Team) QueryUser() *UserQuery {
	return (&TeamClient{config: t.config}).QueryUser(t)
}

// Update returns a builder for updating this Team.
// Note that you need to call Team.Unwrap() before calling this method if this Team
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Team) Update() *TeamUpdateOne {
	return (&TeamClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Team entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Team) Unwrap() *Team {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Team is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Team) String() string {
	var builder strings.Builder
	builder.WriteString("Team(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("team_name=")
	builder.WriteString(t.TeamName)
	builder.WriteString(", ")
	builder.WriteString("team_icon=")
	builder.WriteString(t.TeamIcon)
	builder.WriteString(", ")
	builder.WriteString("team_memo=")
	builder.WriteString(t.TeamMemo)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(t.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Teams is a parsable slice of Team.
type Teams []*Team

func (t Teams) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
