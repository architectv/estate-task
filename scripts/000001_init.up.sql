CREATE TABLE rooms (
    id serial PRIMARY KEY,
    description text NOT NULL,
    price int NOT NULL CHECK (price > 0)
);

CREATE TABLE bookings (
    id serial PRIMARY KEY,
    room_id int REFERENCES rooms (id) ON DELETE CASCADE NOT NULL,
    date_start date NOT NULL,
    date_end date NOT NULL CHECK (date_start < date_end)
);
