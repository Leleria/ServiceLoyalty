CREATE TABLE IF NOT EXISTS TypesOfDiscounts (
    Id	               INTEGER PRIMARY KEY,
    NameType	       TEXT	NOT NULL
);
CREATE TABLE IF NOT EXISTS TypesOfGroups (
    Id	               INTEGER PRIMARY KEY,
    NameType	       TEXT	NOT NULL
);
CREATE TABLE IF NOT EXISTS LoyaltyLevels (
    Id	               INTEGER PRIMARY KEY,
    NameLevel	       TEXT	NOT NULL,
    MinBalance INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS Clients (
    Id	               INTEGER PRIMARY KEY,
    Name	       TEXT	NOT NULL UNIQUE,
    Email TEXT NOT NULL,
    CountBonuses INTEGER NOT NULL,
    LoyaltyLevelFK INTEGER NOT NULL,
    FOREIGN KEY(LoyaltyLevelFK) REFERENCES LoyaltyLevels(Id)
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
    PromoCodeFK INTEGER NOT NULL,
    FOREIGN KEY(ClientFK) REFERENCES Clients(Id),
    FOREIGN KEY(GroupFK) REFERENCES TypesOfGroups(Id),
    FOREIGN KEY(PromoCodeFK) REFERENCES PromoCodes(Id)
);


