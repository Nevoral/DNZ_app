-- Enable foreign key enforcement
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS ProductMenu (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Date DATETIME NOT NULL,
    StartRegister INTEGER NOT NULL,
    Activity TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Product (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    ProductMenuID INTEGER NOT NULL,
    Title TEXT NOT NULL,
    Price INTEGER NOT NULL,
    Served INTEGER NOT NULL,
    Category TEXT NOT NULL,
    UNIQUE (Title),
    FOREIGN KEY(ProductMenuID) REFERENCES ProductMenu(ID) ON DELETE CASCADE
);

CREATE INDEX idx_product_ProductMenuID ON Product(ProductMenuID);

CREATE TABLE IF NOT EXISTS OpenOrder (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    ProductMenuID INTEGER NOT NULL,
    CustomerID INTEGER,
    Date DATETIME NOT NULL,
    Summary INTEGER NOT NULL,
    Status TEXT NOT NULL,
    ItemsOrdered TEXT NOT NULL,
    FOREIGN KEY(ProductMenuID) REFERENCES ProductMenu(ID) ON DELETE CASCADE,
    FOREIGN KEY(CustomerID) REFERENCES Customer(ID) ON DELETE CASCADE
);

CREATE INDEX idx_OpenOrder_ProductMenuID ON OpenOrder(ProductMenuID);
CREATE INDEX idx_OpenOrder_CustomerID ON OpenOrder(CustomerID);

CREATE TABLE IF NOT EXISTS Customer (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Name TEXT NOT NULL,
    PhoneNumber TEXT,
    Role TEXT NOT NULL,
    UNIQUE (Name, PhoneNumber)
);
