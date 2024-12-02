INSERT INTO
    app_data.restaurants (name, menu)
VALUES
    (
        'Mango Sticky Rice',
        '[{"id": 1, "price": 30.99, "name": "Mango"}, {"id": 2, "price": 12.50, "name": "California roll"}]'::jsonb
    ),
    (
        'Indian Curry House',
        '[{"id": 1, "price": 14.99, "name": "Chicken tikka masala"}, {"id": 2, "price": 9.50, "name": "Vegetable korma"}]'::jsonb
    ),
    (
        'Pork Noodle',
        '[{"id": 1, "price": 12.00, "name": "Noodle"}, {"id": 2, "price": 12.75, "name": "Chili con carne"}]'::jsonb
    );
    

INSERT INTO
    app_data.riders (name)
VALUES
    ('John'),
    ('Yok'),
    ('Nitcha'),
    ('Pub');