# MongoDB Oplog to PostgreSQL Processor

This Go program reads MongoDB oplog entries and converts them into equivalent SQL statements for PostgreSQL. It supports insert, update, and delete operations. The program uses the GORM library to interact with PostgreSQL and MongoDB's official Go driver to read the oplogs.

## Prerequisites

- Go 1.18+ installed
- MongoDB running locally or accessible remotely
- PostgreSQL running locally or accessible remotely
- The following Go modules should be installed:
  - `go.mongodb.org/mongo-driver`
  - `gorm.io/gorm`
  - `gorm.io/driver/postgres`

## Setup

1. Clone this repository or download the Go file.

2. Install the necessary Go modules:
   ```bash
   go mod tidy
   ```

3. Modify the connection details for MongoDB and PostgreSQL in the code (in the main function):
   - **MongoDB URI**: Update the MongoDB URI if your MongoDB instance is running remotely or on a different port.
   - **PostgreSQL DSN**: Update the PostgreSQL Data Source Name (DSN) with the correct database credentials.
     - Example DSN format:
       ```text
       host=localhost user=postgres password=secret dbname=test port=5432 sslmode=disable
       ```

4. Ensure the MongoDB oplog collection (`oplog.rs`) is accessible. The program will read the oplogs and process each entry to generate the corresponding SQL statement.

## Program Overview

### Oplog Entry

The program expects the following fields in each MongoDB oplog entry:

- **Timestamp**: The timestamp of the operation.
- **Operation**: The type of operation (`i` for insert, `u` for update, `d` for delete).
- **Namespace**: The namespace (database and collection) of the document being operated on.
- **Document**: The document involved in the operation (for inserts and deletes).
- **UpdateFields**: The fields being updated (for update operations).

### Operations

- **Insert (`i`)**: An `INSERT` SQL statement is generated for the corresponding table.
- **Update (`u`)**: An `UPDATE` SQL statement is generated based on the filter and update fields.
- **Delete (`d`)**: A `DELETE` SQL statement is generated for the corresponding table.

### SQL Generation

The program dynamically generates SQL statements for each operation:

- **INSERT**: Creates an `INSERT` statement with the fields and values from the MongoDB document.
- **UPDATE**: Creates an `UPDATE` statement using the filter and update fields.
- **DELETE**: Creates a `DELETE` statement based on the filter.

## Example

Assume MongoDB oplog entry:

```json
{
  "ts": 1234567890,
  "op": "i",
  "ns": "mydb.mycollection",
  "o": {
    "_id": 1,
    "name": "John",
    "age": 30
  }
}
```

This entry will be converted to an SQL statement like:

```sql
INSERT INTO mycollection (_id, name, age) VALUES ('1', 'John', '30');
```

## Running the Program

To run the program, execute the following command:

```bash
go run main.go
```

The program will:

1. Connect to MongoDB and PostgreSQL.
2. Start reading the oplog from the `oplog.rs` collection.
3. Process each oplog entry and generate the corresponding SQL statements.
4. Execute the generated SQL on the PostgreSQL database.

## Error Handling

The program logs errors when:

- An oplog entry cannot be decoded.
- SQL execution fails.

Errors are logged using `log.Println` and `log.Printf`.
