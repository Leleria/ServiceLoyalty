CREATE TABLE IF NOT EXISTS PromoCodes (
    Id	               INTEGER PRIMARY KEY,
    Name	       TEXT	NOT NULL UNIQUE,
    TypeDiscountFK INTEGER NOT NULL,
    ValueDiscount INTEGER NOT NULL,
    DateStartActive TEXT,
    DateFinishActive TEXT,
    MaxCountUses INTEGER
);