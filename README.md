# CRUD Categories

This project is a simple CRUD application for managing Categories and Products.

## Base URLs

- **Local Development**: `http://localhost:8181`
- **Production**: `https://crud-categories-production.up.railway.app`

## Endpoints

### Categories

#### `GET` /categories
Get all categories.

#### `GET` /categories/:id
Get a single category by ID.

#### `POST` /categories
Create a new category.

**Request Body:**

```json
{
  "name": "string",
  "description": "string"
}
```

#### `PUT` /categories/:id
Update a category.

**Request Body:**

```json
{
  "name": "string",
  "description": "string"
}
```

#### `DELETE` /categories/:id
Delete a category.

### Products

#### `GET` /products
Get all products.

#### `GET` /products/:id
Get a single product by ID.

#### `POST` /products
Create a new product.

**Request Body:**

```json
{
  "name": "string",
  "stock": int,
  "price": int,
  "category_id": int
}
```

#### `PUT` /products/:id
Update a product.

**Request Body:**

```json
{
  "name": "string",
  "stock": int,
  "price": int,
  "category_id": int
}
```

#### `DELETE` /products/:id
Delete a product.

### Transactions

#### `POST` /api/checkout
Create a new transaction.

**Request Body:**

```json
{
  "items": [
    {
      "product_id": int,
      "quantity": int
    }
  ]
}
```

#### `GET` /api/report/hari-ini
Get report for today.

#### `GET` /api/report
Get report with date range.

**Query Params:**
- `start_date` (YYYY-MM-DD)
- `end_date` (YYYY-MM-DD)

## Postman Collection

A Postman collection is included in this repository to help you test the endpoints.

- **File**: `codeWithUmam.postman_collection.json`
- **How to use**:
  1. Open Postman.
  2. Click **Import**.
  3. Drag and drop the `codeWithUmam.postman_collection.json` file.
  4. The collection "codeWithUmam" will appear in your workspace.
  5. The collection is pre-configured with a `host` variable for the production URL. You can add a `host_local` variable or edit the `host` variable to point to `http://localhost:8181` for local testing.
