create TABLE supplies (
    id serial PRIMARY KEY,
    title text,
    description text,
    price NUMERIC(10, 2),
    quantity int not null
);