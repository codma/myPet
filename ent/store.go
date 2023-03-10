// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"main/ent/store"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Store is the model entity for the Store schema.
type Store struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// IsUsed holds the value of the "is_used" field.
	IsUsed bool `json:"is_used,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// StoreID holds the value of the "store_id" field.
	StoreID string `json:"store_id,omitempty"`
	// AppAccessKey holds the value of the "app_access_key" field.
	AppAccessKey string `json:"app_access_key,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StoreQuery when eager-loading is set.
	Edges StoreEdges `json:"edges"`
}

// StoreEdges holds the relations/edges for other nodes in the graph.
type StoreEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e StoreEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Store) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case store.FieldIsUsed:
			values[i] = new(sql.NullBool)
		case store.FieldID:
			values[i] = new(sql.NullInt64)
		case store.FieldStoreID, store.FieldAppAccessKey:
			values[i] = new(sql.NullString)
		case store.FieldCreatedAt, store.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Store", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Store fields.
func (s *Store) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case store.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint64(value.Int64)
		case store.FieldIsUsed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_used", values[i])
			} else if value.Valid {
				s.IsUsed = value.Bool
			}
		case store.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case store.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case store.FieldStoreID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field store_id", values[i])
			} else if value.Valid {
				s.StoreID = value.String
			}
		case store.FieldAppAccessKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_access_key", values[i])
			} else if value.Valid {
				s.AppAccessKey = value.String
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Store entity.
func (s *Store) QueryUsers() *UserQuery {
	return NewStoreClient(s.config).QueryUsers(s)
}

// Update returns a builder for updating this Store.
// Note that you need to call Store.Unwrap() before calling this method if this Store
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Store) Update() *StoreUpdateOne {
	return NewStoreClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Store entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Store) Unwrap() *Store {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Store is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Store) String() string {
	var builder strings.Builder
	builder.WriteString("Store(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("is_used=")
	builder.WriteString(fmt.Sprintf("%v", s.IsUsed))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("store_id=")
	builder.WriteString(s.StoreID)
	builder.WriteString(", ")
	builder.WriteString("app_access_key=")
	builder.WriteString(s.AppAccessKey)
	builder.WriteByte(')')
	return builder.String()
}

// Stores is a parsable slice of Store.
type Stores []*Store

func (s Stores) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
