-- name: CreateProductMenuAndReturnIt :one
INSERT INTO ProductMenu (Date, StartRegister, Activity)
VALUES (?, ?, ?)
RETURNING *;

-- name: SetProductMenuActivityByID :exec
UPDATE ProductMenu
SET Activity = ?
WHERE ID = ?;

-- name: GetListProductMenu :many
SELECT * FROM ProductMenu;

-- name: GetListProductMenuActive :many
SELECT * FROM ProductMenu
WHERE Activity=?;


-- name: CreateProductAndReturnIt :one
INSERT INTO Product (ProductMenuID, Title, Price, Served, Category)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: SetProductCounterByID :exec
UPDATE Product
SET Served = ?
WHERE ID = ?;

-- name: GetProductByID :one
SELECT * FROM Product
WHERE ID=?;

-- name: GetProductByTitle :one
SELECT * FROM Product
WHERE Title=?;

-- name: GetProductListByProductMenuIDAndCategory :many
SELECT * FROM Product
WHERE ProductMenuID=? AND Category=?;

-- name: GetProductListByProductMenuID :many
SELECT * FROM Product
WHERE ProductMenuID=?;

-- name: CreateOpenOrderAndReturnIt :one
INSERT INTO OpenOrder (ProductMenuID, CustomerID, Date, Summary, Status, ItemsOrdered)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: SetOpenOrderSummaryByID :exec
UPDATE OpenOrder
SET Summary = ?
WHERE ID = ?;

-- name: SetOpenOrderStatusByID :exec
UPDATE OpenOrder
SET Status = ?
WHERE ID = ?;

-- name: SetOpenOrderItemsOrdered :exec
UPDATE OpenOrder
SET ItemsOrdered = ?
WHERE ID = ?;

-- name: GetOpenOrderByID :one
SELECT * FROM OpenOrder
WHERE ID=?;

-- name: CreateCustomerAndReturnIt :one
INSERT INTO Customer (Name, PhoneNumber, Role)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetCustomerByID :one
SELECT * FROM Customer
WHERE ID=?;