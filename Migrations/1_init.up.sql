CREATE TABLE IF NOT EXISTS TypesOfDiscounts (
    Id	               INTEGER PRIMARY KEY,
    NameType	       TEXT	NOT NULL
);
CREATE TABLE IF NOT EXISTS TypesOfGroups (
    Id	               INTEGER PRIMARY KEY,
    NameType	       TEXT	NOT NULL
);
CREATE TABLE IF NOT EXISTS CashBackTypes (
    Id	               INTEGER PRIMARY KEY,
    NameType	       TEXT	NOT NULL
);

CREATE TABLE IF NOT EXISTS CashBack (
    Id	               INTEGER PRIMARY KEY,
    Budget	           INTEGER	NOT NULL,
    TypeCashBackFK	   INTEGER	NOT NULL,
    ValueCondition	       TEXT	NOT NULL,
    FOREIGN KEY(TypeCashBackFK) REFERENCES CashBackTypes(Id)
);

CREATE TABLE IF NOT EXISTS UserActions (
    Id	               INTEGER PRIMARY KEY,
    NameAction	       TEXT	NOT NULL
);

CREATE TABLE IF NOT EXISTS TypesOperations (
    Id	               INTEGER PRIMARY KEY,
    NameTypeOperation  TEXT	NOT NULL
);


CREATE TABLE IF NOT EXISTS LoyaltyLevels (
    Id	               INTEGER PRIMARY KEY,
    NameLevel	       TEXT	NOT NULL,
    UserActionFK INTEGER NOT NULL,
    CountBonuses INTEGER NOT NULL,
    FOREIGN KEY(UserActionFK) REFERENCES UserActions(Id)
);
CREATE TABLE IF NOT EXISTS Clients (
    Id	               INTEGER PRIMARY KEY,
    Name	       TEXT	NOT NULL UNIQUE,
    Email TEXT NOT NULL,
    CountBonuses INTEGER NOT NULL,
    LoyaltyLevelFK INTEGER NOT NULL,
    FOREIGN KEY(LoyaltyLevelFK) REFERENCES LoyaltyLevels(Id)
);

CREATE TABLE IF NOT EXISTS Operations (
    Id	               INTEGER PRIMARY KEY,
    TypeOperationFK    INTEGER	NOT NULL,
    ClientFK             INTEGER	NOT NULL,
    CountBonuses       INTEGER	NOT NULL,
    DateAndTimeOperation      TEXT NOT NULL,
    FOREIGN KEY(TypeOperationFK) REFERENCES TypesOperations(Id),
    FOREIGN KEY(ClientFK) REFERENCES Clients(Id)
);

CREATE TABLE IF NOT EXISTS PromoCodes (
    Id	               INTEGER PRIMARY KEY,
    Name	       TEXT	NOT NULL UNIQUE,
    TypeDiscountFK INTEGER NOT NULL,
    ValueDiscount INTEGER NOT NULL,
    DateStartActive TEXT,
    DateFinishActive TEXT,
    MaxCountUses INTEGER,
    FOREIGN KEY(TypeDiscountFK) REFERENCES TypesOfDiscounts(Id)
);

CREATE TABLE IF NOT EXISTS PersonalPromoCodes (
    Id INTEGER PRIMARY KEY,
    ClientFK INTEGER	NOT NULL,
    GroupFK INTEGER NOT NULL,
    NamePromoCode TEXT,
    TypeDiscountFK INTEGER NOT NULL,
    ValueDiscount INTEGER NOT NULL,
    DateStartActive TEXT,
    DateFinishActive TEXT,
    FOREIGN KEY(TypeDiscountFK) REFERENCES TypesOfDiscounts(Id),
    FOREIGN KEY(ClientFK) REFERENCES Clients(Id),
    FOREIGN KEY(GroupFK) REFERENCES TypesOfGroups(Id)
);


