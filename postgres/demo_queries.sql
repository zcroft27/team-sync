SELECT ctid, id, name, value FROM demo_table;
 
-- Show the physical order (ctid shows the physical location)
SELECT ctid, id, name, value FROM demo_table ORDER BY ctid;
 
-- Create a secondary index on 'value' column
CREATE INDEX idx_value ON demo_table(value);
 
-- Show that secondary index doesn't change physical order
SELECT ctid, id, name, value FROM demo_table ORDER BY ctid;
 
-- Now add a primary key constraint
ALTER TABLE demo_table ADD CONSTRAINT pk_demo PRIMARY KEY (id);
 
-- Show physical order is still the same (PostgreSQL doesn't auto-cluster)
SELECT ctid, id, name, value FROM demo_table ORDER BY ctid;
 
-- CLUSTER the table by primary key to demonstrate clustering
CLUSTER demo_table USING pk_demo;
 
-- Now show the physical order after clustering
SELECT ctid, id, name, value FROM demo_table ORDER BY ctid;
 
-- Compare with ordering by index
SELECT ctid, id, name, value FROM demo_table ORDER BY id;
