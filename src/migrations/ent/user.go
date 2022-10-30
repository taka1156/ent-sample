// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sample/migrations/ent/goal"
	"sample/migrations/ent/role"
	"sample/migrations/ent/team"
	"sample/migrations/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID int `json:"role_id,omitempty"`
	// TeamID holds the value of the "team_id" field.
	TeamID int `json:"team_id,omitempty"`
	// GoalID holds the value of the "goal_id" field.
	GoalID int `json:"goal_id,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// UserIcon holds the value of the "user_icon" field.
	UserIcon string `json:"user_icon,omitempty"`
	// UserMemo holds the value of the "user_memo" field.
	UserMemo string `json:"user_memo,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Category holds the value of the category edge.
	Category []*Category `json:"category,omitempty"`
	// Role holds the value of the role edge.
	Role *Role `json:"role,omitempty"`
	// Team holds the value of the team edge.
	Team *Team `json:"team,omitempty"`
	// Goal holds the value of the goal edge.
	Goal *Goal `json:"goal,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CategoryOrErr() ([]*Category, error) {
	if e.loadedTypes[0] {
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) RoleOrErr() (*Role, error) {
	if e.loadedTypes[1] {
		if e.Role == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) TeamOrErr() (*Team, error) {
	if e.loadedTypes[2] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// GoalOrErr returns the Goal value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) GoalOrErr() (*Goal, error) {
	if e.loadedTypes[3] {
		if e.Goal == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: goal.Label}
		}
		return e.Goal, nil
	}
	return nil, &NotLoadedError{edge: "goal"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldRoleID, user.FieldTeamID, user.FieldGoalID:
			values[i] = new(sql.NullInt64)
		case user.FieldUserName, user.FieldEmail, user.FieldUserIcon, user.FieldUserMemo:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				u.RoleID = int(value.Int64)
			}
		case user.FieldTeamID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field team_id", values[i])
			} else if value.Valid {
				u.TeamID = int(value.Int64)
			}
		case user.FieldGoalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field goal_id", values[i])
			} else if value.Valid {
				u.GoalID = int(value.Int64)
			}
		case user.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				u.UserName = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldUserIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_icon", values[i])
			} else if value.Valid {
				u.UserIcon = value.String
			}
		case user.FieldUserMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_memo", values[i])
			} else if value.Valid {
				u.UserMemo = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// QueryCategory queries the "category" edge of the User entity.
func (u *User) QueryCategory() *CategoryQuery {
	return (&UserClient{config: u.config}).QueryCategory(u)
}

// QueryRole queries the "role" edge of the User entity.
func (u *User) QueryRole() *RoleQuery {
	return (&UserClient{config: u.config}).QueryRole(u)
}

// QueryTeam queries the "team" edge of the User entity.
func (u *User) QueryTeam() *TeamQuery {
	return (&UserClient{config: u.config}).QueryTeam(u)
}

// QueryGoal queries the "goal" edge of the User entity.
func (u *User) QueryGoal() *GoalQuery {
	return (&UserClient{config: u.config}).QueryGoal(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("role_id=")
	builder.WriteString(fmt.Sprintf("%v", u.RoleID))
	builder.WriteString(", ")
	builder.WriteString("team_id=")
	builder.WriteString(fmt.Sprintf("%v", u.TeamID))
	builder.WriteString(", ")
	builder.WriteString("goal_id=")
	builder.WriteString(fmt.Sprintf("%v", u.GoalID))
	builder.WriteString(", ")
	builder.WriteString("user_name=")
	builder.WriteString(u.UserName)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("user_icon=")
	builder.WriteString(u.UserIcon)
	builder.WriteString(", ")
	builder.WriteString("user_memo=")
	builder.WriteString(u.UserMemo)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(u.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}