CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  name varchar NOT NULL,
  email varchar UNIQUE NOT NULL,
  phone varchar UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS venues (
  id serial PRIMARY KEY,
  name varchar UNIQUE NOT NULL,
  address varchar UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS sections (
  id serial PRIMARY KEY,
  venue_id int NOT NULL REFERENCES venues(id),
  name varchar NOT NULL,
  row_count int NOT NULL,
  seats_per_row int NOT NULL,
  available bool NOT NULL
);

CREATE TABLE IF NOT EXISTS events (
  id serial PRIMARY KEY,
  name varchar not null,
  venue_id int NOT NULL REFERENCES venues(id),
  start_time timestamp NOT NULL,
  end_time timestamp NOT NULL 
);

CREATE TABLE IF NOT EXISTS event_section_prices (
  id serial PRIMARY KEY,
  section_id int NOT NULL REFERENCES sections(id),
  event_id int NOT NULL REFERENCES events(id),
  price decimal NOT NULL
);

CREATE TABLE IF NOT EXISTS seats (
  id serial PRIMARY KEY,
  section_id int NOT NULL REFERENCES sections(id),
  row_number int NOT NULL,
  seat_number int NOT NULL
);

CREATE TABLE IF NOT EXISTS reservations (
  id serial PRIMARY KEY,
  user_id int NOT NULL REFERENCES users(id),
  event_id int NOT NULL REFERENCES events(id)
);

CREATE TABLE IF NOT EXISTS event_seats (
  id serial PRIMARY KEY,
  event_id int NOT NULL REFERENCES events(id),
  seat_id int NOT NULL REFERENCES seats(id),
  price decimal NOT NULL,
  reserved bool NOT NULL
);

CREATE TABLE IF NOT EXISTS seats_reservations (
  id serial PRIMARY KEY,
  reservation_id int REFERENCES reservations(id),
  event_id int REFERENCES events(id),
  event_seat_id int REFERENCES event_seats(id)
);
