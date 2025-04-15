-- Trigger para llenar los asientos del evento cuando se crea
CREATE OR REPLACE FUNCTION populate_event_seats()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO event_seats (event_id, seat_id, price, reserved)
  SELECT
    NEW.event_id,
    s.id,
    NEW.price,
    false
  FROM seats s
  WHERE s.section_id = NEW.section_id;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_populate_event_seats ON event_section_prices;

CREATE TRIGGER trg_populate_event_seats
AFTER INSERT ON event_section_prices
FOR EACH ROW
EXECUTE FUNCTION populate_event_seats();

-- Trigger para crear asientos de cada seccion basado en la cantidad
-- de filas y asientos por fila de la seccion
CREATE OR REPLACE FUNCTION generate_seats_for_section()
RETURNS TRIGGER AS $$
DECLARE
  r INT;
  s INT;
BEGIN
  FOR r IN 1..NEW.row_count LOOP
    FOR s IN 1..NEW.seats_per_row LOOP
      INSERT INTO seats (section_id, row_number, seat_number)
      VALUES (NEW.id, r, s);
    END LOOP;
  END LOOP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_generate_seats ON sections;

CREATE TRIGGER trg_generate_seats
AFTER INSERT ON sections
FOR EACH ROW
EXECUTE FUNCTION generate_seats_for_section();

