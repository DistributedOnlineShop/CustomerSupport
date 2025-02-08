-- name: CreateCustomerSupportCase :one
INSERT INTO customer_support (
    USER_ID,
    ORDER_ID,
    SUBJECT,
    MESSAGE,
    STATUS
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetCustomerSupportCaseList :many
SELECT 
    * 
FROM customer_support;

-- name: GetCustomerSupportCase :one
SELECT 
    * 
FROM customer_support 
WHERE CS_ID = $1;

-- name: UpdateCustomerSupportCaseStatus :one
UPDATE customer_support 
SET 
    STATUS = COALESCE($2,Status),
    MESSAGE = COALESCE($3,Message),
    UPDATED_AT = NOW()
WHERE 
    CS_ID = $1 RETURNING *;