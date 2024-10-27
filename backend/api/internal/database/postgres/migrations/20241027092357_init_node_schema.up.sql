-- create "node" table
CREATE TABLE "node" (
  "id" uuid NOT NULL,
  "type" character varying(255) NOT NULL,
  "name" character varying(255) NOT NULL,
  "parent_id" uuid NULL,
  "owner_id" uuid NOT NULL,
  PRIMARY KEY ("id")
);
-- create "node_closure" table
CREATE TABLE "node_closure" (
  "ancestor_id" uuid NOT NULL,
  "descendant_id" uuid NOT NULL,
  "depth" integer NOT NULL,
  PRIMARY KEY ("ancestor_id", "descendant_id"),
  CONSTRAINT "node_closure_ancestor_id_fkey" FOREIGN KEY ("ancestor_id") REFERENCES "node" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "node_closure_descendant_id_fkey" FOREIGN KEY ("descendant_id") REFERENCES "node" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- create "node_content" table
CREATE TABLE "node_content" (
  "node_id" uuid NOT NULL,
  "content" text NOT NULL,
  PRIMARY KEY ("node_id"),
  CONSTRAINT "node_content_node_id_fkey" FOREIGN KEY ("node_id") REFERENCES "node" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- create "insert_node_closure" function
CREATE FUNCTION insert_node_closure()
RETURNS TRIGGER AS
    $$
    BEGIN
        INSERT INTO node_closure("ancestor_id", "descendant_id", "depth")
        VALUES (NEW.id, NEW.id, 0);

        INSERT INTO node_closure("ancestor_id", "descendant_id", "depth")
            SELECT a.ancestor_id, b.ancestor_id, a.depth + b.depth + 1 FROM node_closure a CROSS JOIN node_closure b
            WHERE b.ancestor_id = NEW.id AND a.descendant_id = NEW.parent_id;

        RETURN null;
    END;
    $$
LANGUAGE plpgsql;

-- create "insert_node_closure_trigger" trigger
CREATE TRIGGER insert_node_closure_trigger
    AFTER INSERT
    ON node
    FOR EACH ROW
    EXECUTE FUNCTION insert_node_closure();