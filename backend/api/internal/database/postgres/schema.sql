CREATE TABLE node (
    id UUID PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    parent_id UUID,
    owner_id UUID NOT NULL
);

CREATE TABLE node_content(
    node_id UUID PRIMARY KEY,
    content TEXT NOT NULL,
    FOREIGN KEY (node_id) REFERENCES node(id)
);

CREATE TABLE node_closure (
    ancestor_id UUID NOT NULL,
    descendant_id UUID NOT NULL,
    depth INT NOT NULL,
    PRIMARY KEY (ancestor_id, descendant_id),
    FOREIGN KEY (ancestor_id) REFERENCES node(id) ON DELETE CASCADE,
    FOREIGN KEY (descendant_id) REFERENCES node(id) ON DELETE CASCADE
);