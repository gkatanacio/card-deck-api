CREATE OR REPLACE FUNCTION set_updated_at_to_now()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TABLE decks (
  id UUID NOT NULL,
  shuffled BOOLEAN NOT NULL,
  cards JSONB NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP NOT NULL DEFAULT now(),
  PRIMARY KEY(id)
);

CREATE TRIGGER set_updated_at_decks
  BEFORE UPDATE
  ON decks
  FOR EACH ROW
EXECUTE FUNCTION set_updated_at_to_now();
