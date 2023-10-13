// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlueprintsColumns holds the columns for the "blueprints" table.
	BlueprintsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "blueprint_template", Type: field.TypeBytes},
		{Name: "blueprint_parent_group", Type: field.TypeUUID},
		{Name: "blueprint_virtualization_provider", Type: field.TypeUUID},
	}
	// BlueprintsTable holds the schema information for the "blueprints" table.
	BlueprintsTable = &schema.Table{
		Name:       "blueprints",
		Columns:    BlueprintsColumns,
		PrimaryKey: []*schema.Column{BlueprintsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blueprints_groups_parent_group",
				Columns:    []*schema.Column{BlueprintsColumns[3]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "blueprints_virtualization_providers_virtualization_provider",
				Columns:    []*schema.Column{BlueprintsColumns[4]},
				RefColumns: []*schema.Column{VirtualizationProvidersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DeploymentsColumns holds the columns for the "deployments" table.
	DeploymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "deployment_blueprint", Type: field.TypeUUID},
		{Name: "deployment_requester", Type: field.TypeUUID},
	}
	// DeploymentsTable holds the schema information for the "deployments" table.
	DeploymentsTable = &schema.Table{
		Name:       "deployments",
		Columns:    DeploymentsColumns,
		PrimaryKey: []*schema.Column{DeploymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deployments_blueprints_blueprint",
				Columns:    []*schema.Column{DeploymentsColumns[1]},
				RefColumns: []*schema.Column{BlueprintsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "deployments_users_requester",
				Columns:    []*schema.Column{DeploymentsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "group_children", Type: field.TypeUUID, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_groups_children",
				Columns:    []*schema.Column{GroupsColumns[2]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "component", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:       "permissions",
		Columns:    PermissionsColumns,
		PrimaryKey: []*schema.Column{PermissionsColumns[0]},
	}
	// PermissionPoliciesColumns holds the columns for the "permission_policies" table.
	PermissionPoliciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeEnum, Nullable: true, Enums: []string{"ALLOW", "DENY"}, Default: "ALLOW"},
		{Name: "is_inherited", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "permission_policy_permission", Type: field.TypeUUID},
		{Name: "permission_policy_group", Type: field.TypeUUID},
	}
	// PermissionPoliciesTable holds the schema information for the "permission_policies" table.
	PermissionPoliciesTable = &schema.Table{
		Name:       "permission_policies",
		Columns:    PermissionPoliciesColumns,
		PrimaryKey: []*schema.Column{PermissionPoliciesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "permission_policies_permissions_permission",
				Columns:    []*schema.Column{PermissionPoliciesColumns[3]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "permission_policies_groups_group",
				Columns:    []*schema.Column{PermissionPoliciesColumns[4]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VirtualizationProvidersColumns holds the columns for the "virtualization_providers" table.
	VirtualizationProvidersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "display_name", Type: field.TypeString},
		{Name: "provider_git_url", Type: field.TypeString},
		{Name: "provider_version", Type: field.TypeString},
		{Name: "config_path", Type: field.TypeString},
	}
	// VirtualizationProvidersTable holds the schema information for the "virtualization_providers" table.
	VirtualizationProvidersTable = &schema.Table{
		Name:       "virtualization_providers",
		Columns:    VirtualizationProvidersColumns,
		PrimaryKey: []*schema.Column{VirtualizationProvidersColumns[0]},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "group_id", Type: field.TypeUUID},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0], UserGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_user_id",
				Columns:    []*schema.Column{UserGroupsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_groups_group_id",
				Columns:    []*schema.Column{UserGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlueprintsTable,
		DeploymentsTable,
		GroupsTable,
		PermissionsTable,
		PermissionPoliciesTable,
		UsersTable,
		VirtualizationProvidersTable,
		UserGroupsTable,
	}
)

func init() {
	BlueprintsTable.ForeignKeys[0].RefTable = GroupsTable
	BlueprintsTable.ForeignKeys[1].RefTable = VirtualizationProvidersTable
	DeploymentsTable.ForeignKeys[0].RefTable = BlueprintsTable
	DeploymentsTable.ForeignKeys[1].RefTable = UsersTable
	GroupsTable.ForeignKeys[0].RefTable = GroupsTable
	PermissionPoliciesTable.ForeignKeys[0].RefTable = PermissionsTable
	PermissionPoliciesTable.ForeignKeys[1].RefTable = GroupsTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[1].RefTable = GroupsTable
}
