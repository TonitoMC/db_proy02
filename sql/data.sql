-- Usuarios

INSERT INTO users (name, email, phone) VALUES
    ('Alice', 'alice@example.com', '1000000001'),
    ('Bob', 'bob@example.com', '1000000002'),
    ('Carol', 'carol@example.com', '1000000003'),
    ('Dave', 'dave@example.com', '1000000004'),
    ('Eve', 'eve@example.com', '1000000005'),
    ('Frank', 'frank@example.com', '1000000006'),
    ('Grace', 'grace@example.com', '1000000007'),
    ('Heidi', 'heidi@example.com', '1000000008'),
    ('Ivan', 'ivan@example.com', '1000000009'),
    ('Judy', 'judy@example.com', '1000000010'),
    ('Mallory', 'mallory@example.com', '1000000011'),
    ('Niaj', 'niaj@example.com', '1000000012'),
    ('Olivia', 'olivia@example.com', '1000000013'),
    ('Peggy', 'peggy@example.com', '1000000014'),
    ('Trent', 'trent@example.com', '1000000015'),
    ('George', 'george@example.com', '1000000016'),
    ('Fred', 'fred@example.com', '1000000017'),
    ('Fisher', 'fisher@example.com', '1000000018'),
    ('Nick', 'nick@example.com', '1000000019'),
    ('Oscar', 'oscar@example.com', '1000000020'),
    ('Paul', 'paul@example.com', '1000000021'),
    ('Quincy', 'quincy@example.com', '1000000022'),
    ('Rita', 'rita@example.com', '1000000023'),
    ('Sam', 'sam@example.com', '1000000024'),
    ('Tina', 'tina@example.com', '1000000025'),
    ('Uma', 'uma@example.com', '1000000026'),
    ('Vince', 'vince@example.com', '1000000027'),
    ('Wendy', 'wendy@example.com', '1000000028'),
    ('Xander', 'xander@example.com', '1000000029'),
    ('Yara', 'yara@example.com', '1000000030');


-- Un venue para este ejemplo
INSERT INTO venues (name, address) VALUES ('Main Hall', '123 Street');

-- 3 secciones dentro del venue, con 2 filas y 5 asientos por fila
INSERT INTO sections (venue_id, name, row_count, seats_per_row, available) VALUES
  (1, 'General', 2, 5, true),
  (1, 'VIP', 2, 5, true),
  (1, 'Balcon', 2, 5, true);

-- La tabla de seats se pobla por medio de un trigger (no estoy muy seguro de la eficiencia)
-- pero ayuda a mantener la integridad y ahorra trabajo en los inserts :D

-- Eventos
INSERT INTO events (venue_id, name, start_time, end_time) VALUES
  (1, 'Ozzy Osbourne', NOW() + INTERVAL '7 days', NOW() + INTERVAL '7 days 3 hours');

-- Precio por cada seccion, los seats del evento se llenan automaticamente con el precio
-- por seccion
INSERT INTO event_section_prices (section_id, event_id, price) VALUES
  (1, 1, 20.0),
  (2, 1, 40.0),
  (3, 1, 30.0);

