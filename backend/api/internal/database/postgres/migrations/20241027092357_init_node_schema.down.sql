-- reverse: create "node_content" table
DROP TABLE "node_content";
-- reverse: create "node_closure" table
DROP TABLE "node_closure";
-- reverse: create "node" table
DROP TABLE "node";
-- reverse: create "insert_node_closure_trigger" trigger
DROP TRIGGER IF EXISTS insert_node_closure_trigger ON node;
-- reverse: create "insert_node_closure" function
DROP FUNCTION IF EXISTS insert_node_closure();
