// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_schemas/ent/website"
	"_schemas/ent/websiteevent"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// WebsiteEvent is the model entity for the WebsiteEvent schema.
type WebsiteEvent struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// EventName holds the value of the "event_name" field.
	EventName string `json:"event_name,omitempty"`
	// URLPath holds the value of the "url_path" field.
	URLPath string `json:"url_path,omitempty"`
	// URLQuery holds the value of the "url_query" field.
	URLQuery string `json:"url_query,omitempty"`
	// ReferrerPath holds the value of the "referrer_path" field.
	ReferrerPath string `json:"referrer_path,omitempty"`
	// ReferrerQuery holds the value of the "referrer_query" field.
	ReferrerQuery string `json:"referrer_query,omitempty"`
	// ReferrerDomain holds the value of the "referrer_domain" field.
	ReferrerDomain string `json:"referrer_domain,omitempty"`
	// PageTitle holds the value of the "page_title" field.
	PageTitle string `json:"page_title,omitempty"`
	// PageData holds the value of the "page_data" field.
	PageData string `json:"page_data,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WebsiteEventQuery when eager-loading is set.
	Edges        WebsiteEventEdges `json:"edges"`
	website_id   *uuid.UUID
	selectValues sql.SelectValues
}

// WebsiteEventEdges holds the relations/edges for other nodes in the graph.
type WebsiteEventEdges struct {
	// Website holds the value of the website edge.
	Website *Website `json:"website,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WebsiteOrErr returns the Website value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WebsiteEventEdges) WebsiteOrErr() (*Website, error) {
	if e.loadedTypes[0] {
		if e.Website == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: website.Label}
		}
		return e.Website, nil
	}
	return nil, &NotLoadedError{edge: "website"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WebsiteEvent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case websiteevent.FieldEventName, websiteevent.FieldURLPath, websiteevent.FieldURLQuery, websiteevent.FieldReferrerPath, websiteevent.FieldReferrerQuery, websiteevent.FieldReferrerDomain, websiteevent.FieldPageTitle, websiteevent.FieldPageData:
			values[i] = new(sql.NullString)
		case websiteevent.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case websiteevent.FieldID:
			values[i] = new(uuid.UUID)
		case websiteevent.ForeignKeys[0]: // website_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WebsiteEvent fields.
func (we *WebsiteEvent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case websiteevent.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				we.ID = *value
			}
		case websiteevent.FieldEventName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_name", values[i])
			} else if value.Valid {
				we.EventName = value.String
			}
		case websiteevent.FieldURLPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url_path", values[i])
			} else if value.Valid {
				we.URLPath = value.String
			}
		case websiteevent.FieldURLQuery:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url_query", values[i])
			} else if value.Valid {
				we.URLQuery = value.String
			}
		case websiteevent.FieldReferrerPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field referrer_path", values[i])
			} else if value.Valid {
				we.ReferrerPath = value.String
			}
		case websiteevent.FieldReferrerQuery:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field referrer_query", values[i])
			} else if value.Valid {
				we.ReferrerQuery = value.String
			}
		case websiteevent.FieldReferrerDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field referrer_domain", values[i])
			} else if value.Valid {
				we.ReferrerDomain = value.String
			}
		case websiteevent.FieldPageTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field page_title", values[i])
			} else if value.Valid {
				we.PageTitle = value.String
			}
		case websiteevent.FieldPageData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field page_data", values[i])
			} else if value.Valid {
				we.PageData = value.String
			}
		case websiteevent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				we.CreatedAt = value.Time
			}
		case websiteevent.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field website_id", values[i])
			} else if value.Valid {
				we.website_id = new(uuid.UUID)
				*we.website_id = *value.S.(*uuid.UUID)
			}
		default:
			we.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WebsiteEvent.
// This includes values selected through modifiers, order, etc.
func (we *WebsiteEvent) Value(name string) (ent.Value, error) {
	return we.selectValues.Get(name)
}

// QueryWebsite queries the "website" edge of the WebsiteEvent entity.
func (we *WebsiteEvent) QueryWebsite() *WebsiteQuery {
	return NewWebsiteEventClient(we.config).QueryWebsite(we)
}

// Update returns a builder for updating this WebsiteEvent.
// Note that you need to call WebsiteEvent.Unwrap() before calling this method if this WebsiteEvent
// was returned from a transaction, and the transaction was committed or rolled back.
func (we *WebsiteEvent) Update() *WebsiteEventUpdateOne {
	return NewWebsiteEventClient(we.config).UpdateOne(we)
}

// Unwrap unwraps the WebsiteEvent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (we *WebsiteEvent) Unwrap() *WebsiteEvent {
	_tx, ok := we.config.driver.(*txDriver)
	if !ok {
		panic("ent: WebsiteEvent is not a transactional entity")
	}
	we.config.driver = _tx.drv
	return we
}

// String implements the fmt.Stringer.
func (we *WebsiteEvent) String() string {
	var builder strings.Builder
	builder.WriteString("WebsiteEvent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", we.ID))
	builder.WriteString("event_name=")
	builder.WriteString(we.EventName)
	builder.WriteString(", ")
	builder.WriteString("url_path=")
	builder.WriteString(we.URLPath)
	builder.WriteString(", ")
	builder.WriteString("url_query=")
	builder.WriteString(we.URLQuery)
	builder.WriteString(", ")
	builder.WriteString("referrer_path=")
	builder.WriteString(we.ReferrerPath)
	builder.WriteString(", ")
	builder.WriteString("referrer_query=")
	builder.WriteString(we.ReferrerQuery)
	builder.WriteString(", ")
	builder.WriteString("referrer_domain=")
	builder.WriteString(we.ReferrerDomain)
	builder.WriteString(", ")
	builder.WriteString("page_title=")
	builder.WriteString(we.PageTitle)
	builder.WriteString(", ")
	builder.WriteString("page_data=")
	builder.WriteString(we.PageData)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(we.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// WebsiteEvents is a parsable slice of WebsiteEvent.
type WebsiteEvents []*WebsiteEvent
