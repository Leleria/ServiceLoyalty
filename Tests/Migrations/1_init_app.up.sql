    insert into CashBackTypes (NameType)
    values ('за покупки'),

        ON CONFLICT DO NOTHING;

    insert into TypesOfDiscounts (NameType)
    values ('процентная'),
           ('фиксированная сумма')
        ON CONFLICT DO NOTHING;

    insert into TypesOfGroups (NameType)
    values ('мужчины'),
           ('женщины')
        ON CONFLICT DO NOTHING;

    insert into UserActions (NameAction)
    values ('оставил отзыв')
        ON CONFLICT DO NOTHING;

    insert into LoyaltyLevels (NameLevel, UserActionFK, CountBonuses)
    values ('начальный', 1, 200)
        ON CONFLICT DO NOTHING;

    insert into Clients (Name, Email, CountBonuses, LoyaltyLevelFK)
    values ('User', 'user@gmail.com', 200, 1)
        ON CONFLICT DO NOTHING;