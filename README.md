# sqlmap

Convert a native go datatype to and from its sql null datatype equivalent.

## Why does this exist?

Recently I started using a tool called [sqlc](https://sqlc.dev/) as a pseudo-ORM in one of my projects. Its a great tool but it has a couple of quirks. This package addresses one of those quirks. 

If your database schema contains a null value, the generated code requires you to transform a pointer value to a `sql.NulXYZ` type.
```golang
    // We have a string pointer but need a sql.NullString type.
    var biz *string

    str := sql.NullString{}

    if biz != nil {
		str.String = *in
		str.Valid = true
	}

    // the str var is only initialized if biz is not null.

    // insertRowInDatabase(context.Context, sql.NullString)
    err := insertRowInDatabase(ctx, str)
```

This package provides helper code to enable you to do the conversion inline.

```golang
    // We have a string pointer but need a sql.NullString type.
    var biz *string

    // insertRowInDatabase(context.Context, sql.NullString) (sql.NullString, error)
    val, _ := insertRowInDatabase(ctx, sqlmap.ToNullString(biz))

	// It also allows you to "unwrap" the resultant `sql.NulXYZ` value to make it actually usable.
	biz = sqlmap.UnwrapString(val)
```


## Example

Say you have a database schema with a null column like so, and would like to insert some data into the table.

```sql
-- schema.sql

CREATE TABLE foo (
	bar VARCHAR(128) NULL,
    biz VARCHAR(400) NOT NULL
);
```
```sql
-- query.sql

-- Insert a Foo into the foo table.
-- name: CreateFoo :one
INSERT INTO foo (
	bar,
    biz
) VALUES (
	$1,
	$2
)
```

This will generate some code:

```golang
// models.go

type Foo struct {
	Bar sql.NullString
	Biz string
}
```

```golang
// query.sql.go

const createFoo = `-- name: CreateFoo :one
INSERT INTO foo (
	bar,
    biz
) VALUES (
	$1,
	$2
)
RETURNING bar, biz
`

type CreateFooParams struct {
	Bar sql.NullString
	Biz string
}

// Insert a Foo into the foo table.
func (q *Queries) CreateFoo(ctx context.Context, arg CreateFooParams) (Foo, error) {
	row := q.db.QueryRowContext(ctx, createFoo, arg.Bar, arg.Biz)
	var i Foo
	err := row.Scan(&i.Bar, &i.Biz)
	return i, err
}
```

Use sqlmap inline with your params to avoid boilerplate code. Super useful for CRUD operations where you need to map a request to your datastores representation.

```golang
func main() {
    var db *sql.db
    // Assume you have connected to db.

    queries := datastore.New(db)

    var biz *string

	foo, err := queries.CreateFoo(ctx, datastore.CreateFooParams{
		Bar: sqlmap.NullString(biz),
		Biz: "This is biz",
	})

	// Note: foo.Bar is of type sql.NullString, which isn't really useful to us. 
	// So we make use of the unwrap function to convert it to a more canonical go datatype.
	biz = sqlmap.UnwrapString(foo.Bar)
}
```

## License

This program is released under the GNU Lesser General Public License v3 or later.
