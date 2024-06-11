    insert into CashBackTypes (NameType)
    values ('за покупки'),
           ('за отзывы'),
           ('за привлечение новых клиентов'),
           ('в честь дня рождения'),
           ('в честь юбилея компании')
        ON CONFLICT DO NOTHING;

    insert into TypesOfDiscounts (NameType)
    values ('процентная'),
           ('фиксированная сумма')
        ON CONFLICT DO NOTHING;

    insert into TypesOperations (NameTypeOperation)
    values ('начисление'),
           ('списание')
        ON CONFLICT DO NOTHING;

    insert into TypesOfGroups (NameType)
    values ('мужчины'),
           ('женщины'),
           ('постоянные клиенты'),
           ('VIP-клиенты'),
           ('новые клиенты')
        ON CONFLICT DO NOTHING;

    insert into UserActions (NameAction)
    values ('подписка'),
           ('участие в конкурсах'),
           ('отзывы и комментарии'),
           ('заполнение анкет и опросов'),
           ('настройка профиля')
        ON CONFLICT DO NOTHING;

    insert into LoyaltyLevels (NameLevel, UserActionFK, CountBonuses)
    values ('бронзовый', 1, 200),
           ('серебряный', 2, 500),
           ('золотой', 3, 1000)
        ON CONFLICT DO NOTHING;

    insert into Clients (Name, Email, CountBonuses, LoyaltyLevelFK)
    values ('User1', 'user1@gmail.com', 200, 1),
           ('User2', 'user2@gmail.com', 200, 1),
           ('User3', 'user3@gmail.com', 200, 1),
           ('User4', 'user4@gmail.com', 500, 2),
           ('User5', 'user5@gmail.com', 1000, 3)
        ON CONFLICT DO NOTHING;

    insert into CashBack (Budget, TypeCashBackFK, ValueCondition)
    values (70, 1, 'в течение праздничной недели'),
           (50, 2, 'в конце месяца'),
           (30, 1, 'в течении майских праздников'),
           (40, 1, 'в период нового года'),
           (90, 1, 'в начале месяца')
        ON CONFLICT DO NOTHING;

    insert into PromoCodes (Name, TypeDiscountFK, ValueDiscount, DateStartActive, DateFinishActive, MaxCountUses)
    values ('PGFTD', 1, 12, '2022-12-11', '2024-01-01', 30),
           ('HGCTY', 2, 200, '2022-11-12', '2024-02-01', 50),
           ('NhgGb', 1, 10, '2022-02-10', '2025-05-05', 10),
           ('PuhtG', 2, 300, '2021-12-09', '2023-01-01', 60),
           ('ojhYG', 1, 15, '2022-12-05', '2023-12-05', 75)
        ON CONFLICT DO NOTHING;

    insert into PersonalPromoCodes (ClientFK, GroupFK, NamePromoCode, TypeDiscountFK, ValueDiscount, DateStartActive, DateFinishActive)
    values (1, 1, 'PGFTD', 1, 12, '2022-12-11', '2024-01-01'),
           (2, 2, 'HGCTY', 2, 200, '2022-11-12', '2024-02-01'),
           (3, 3, 'NhgGb', 1, 10, '2022-02-10', '2025-05-05'),
           (4, 4, 'PuhtG', 2, 300, '2021-12-09', '2023-01-01'),
           (5, 5, 'ojhYG', 1, 15, '2022-12-05', '2023-12-05')
        ON CONFLICT DO NOTHING;

    insert into Operations (TypeOperationFK, ClientFK, CountBonuses, DateAndTimeOperation)
    values (1, 1, 50, '2024-01-01 19:03:05'),
           (2, 2, 200, '2022-11-12 22:05:05'),
           (1, 3, 10, '2022-02-10 21:12:00'),
           (2, 4, 300, '2021-12-09 14:03:59'),
           (1, 5, 10, '2022-12-05 13:06:06')
        ON CONFLICT DO NOTHING;