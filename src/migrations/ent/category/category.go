// Code generated by ent, DO NOT EDIT.

package category

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "category_id"
	// FieldCategoryName holds the string denoting the category_name field in the database.
	FieldCategoryName = "category_name"
	// FieldCategoryIcon holds the string denoting the category_icon field in the database.
	FieldCategoryIcon = "category_icon"
	// FieldCategoryMemo holds the string denoting the category_memo field in the database.
	FieldCategoryMemo = "category_memo"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "user_id"
	// Table holds the table name of the category in the database.
	Table = "category"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "j_user_with_category"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "user"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldCategoryName,
	FieldCategoryIcon,
	FieldCategoryMemo,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "category_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
