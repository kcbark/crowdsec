// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AlertsColumns holds the columns for the "alerts" table.
	AlertsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "scenario", Type: field.TypeString},
		{Name: "bucket_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "events_count", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "stopped_at", Type: field.TypeTime, Nullable: true},
		{Name: "source_ip", Type: field.TypeString, Nullable: true},
		{Name: "source_range", Type: field.TypeString, Nullable: true},
		{Name: "source_as_number", Type: field.TypeString, Nullable: true},
		{Name: "source_as_name", Type: field.TypeString, Nullable: true},
		{Name: "source_country", Type: field.TypeString, Nullable: true},
		{Name: "source_latitude", Type: field.TypeFloat32, Nullable: true},
		{Name: "source_longitude", Type: field.TypeFloat32, Nullable: true},
		{Name: "source_scope", Type: field.TypeString, Nullable: true},
		{Name: "source_value", Type: field.TypeString, Nullable: true},
		{Name: "capacity", Type: field.TypeInt32, Nullable: true},
		{Name: "leak_speed", Type: field.TypeString, Nullable: true},
		{Name: "scenario_version", Type: field.TypeString, Nullable: true},
		{Name: "scenario_hash", Type: field.TypeString, Nullable: true},
		{Name: "simulated", Type: field.TypeBool, Default: false},
		{Name: "machine_alerts", Type: field.TypeInt, Nullable: true},
	}
	// AlertsTable holds the schema information for the "alerts" table.
	AlertsTable = &schema.Table{
		Name:       "alerts",
		Columns:    AlertsColumns,
		PrimaryKey: []*schema.Column{AlertsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "alerts_machines_alerts",
				Columns:    []*schema.Column{AlertsColumns[23]},
				RefColumns: []*schema.Column{MachinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "alert_id",
				Unique:  false,
				Columns: []*schema.Column{AlertsColumns[0]},
			},
		},
	}
	// BouncersColumns holds the columns for the "bouncers" table.
	BouncersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "api_key", Type: field.TypeString},
		{Name: "revoked", Type: field.TypeBool},
		{Name: "ip_address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "type", Type: field.TypeString, Nullable: true},
		{Name: "version", Type: field.TypeString, Nullable: true},
		{Name: "until", Type: field.TypeTime, Nullable: true},
		{Name: "last_pull", Type: field.TypeTime},
	}
	// BouncersTable holds the schema information for the "bouncers" table.
	BouncersTable = &schema.Table{
		Name:       "bouncers",
		Columns:    BouncersColumns,
		PrimaryKey: []*schema.Column{BouncersColumns[0]},
	}
	// DecisionsColumns holds the columns for the "decisions" table.
	DecisionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "until", Type: field.TypeTime},
		{Name: "scenario", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "start_ip", Type: field.TypeInt64, Nullable: true},
		{Name: "end_ip", Type: field.TypeInt64, Nullable: true},
		{Name: "start_suffix", Type: field.TypeInt64, Nullable: true},
		{Name: "end_suffix", Type: field.TypeInt64, Nullable: true},
		{Name: "ip_size", Type: field.TypeInt64, Nullable: true},
		{Name: "scope", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "origin", Type: field.TypeString},
		{Name: "simulated", Type: field.TypeBool, Default: false},
		{Name: "alert_decisions", Type: field.TypeInt, Nullable: true},
	}
	// DecisionsTable holds the schema information for the "decisions" table.
	DecisionsTable = &schema.Table{
		Name:       "decisions",
		Columns:    DecisionsColumns,
		PrimaryKey: []*schema.Column{DecisionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "decisions_alerts_decisions",
				Columns:    []*schema.Column{DecisionsColumns[15]},
				RefColumns: []*schema.Column{AlertsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "time", Type: field.TypeTime},
		{Name: "serialized", Type: field.TypeString, Size: 8191},
		{Name: "alert_events", Type: field.TypeInt, Nullable: true},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_alerts_events",
				Columns:    []*schema.Column{EventsColumns[5]},
				RefColumns: []*schema.Column{AlertsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MachinesColumns holds the columns for the "machines" table.
	MachinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "last_push", Type: field.TypeTime, Nullable: true},
		{Name: "machine_id", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "ip_address", Type: field.TypeString},
		{Name: "scenarios", Type: field.TypeString, Nullable: true, Size: 4095},
		{Name: "version", Type: field.TypeString, Nullable: true},
		{Name: "is_validated", Type: field.TypeBool, Default: false},
		{Name: "status", Type: field.TypeString, Nullable: true},
	}
	// MachinesTable holds the schema information for the "machines" table.
	MachinesTable = &schema.Table{
		Name:       "machines",
		Columns:    MachinesColumns,
		PrimaryKey: []*schema.Column{MachinesColumns[0]},
	}
	// MetaColumns holds the columns for the "meta" table.
	MetaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString, Size: 4095},
		{Name: "alert_metas", Type: field.TypeInt, Nullable: true},
	}
	// MetaTable holds the schema information for the "meta" table.
	MetaTable = &schema.Table{
		Name:       "meta",
		Columns:    MetaColumns,
		PrimaryKey: []*schema.Column{MetaColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "meta_alerts_metas",
				Columns:    []*schema.Column{MetaColumns[5]},
				RefColumns: []*schema.Column{AlertsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AlertsTable,
		BouncersTable,
		DecisionsTable,
		EventsTable,
		MachinesTable,
		MetaTable,
	}
)

func init() {
	AlertsTable.ForeignKeys[0].RefTable = MachinesTable
	DecisionsTable.ForeignKeys[0].RefTable = AlertsTable
	EventsTable.ForeignKeys[0].RefTable = AlertsTable
	MetaTable.ForeignKeys[0].RefTable = AlertsTable
}
