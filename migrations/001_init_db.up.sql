-- Create users table
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR,
  email VARCHAR,
  role VARCHAR,
  date_of_birth DATE,
  password VARCHAR,
  created_at TIMESTAMP
);

-- Create organization table
CREATE TABLE organization (
  id SERIAL PRIMARY KEY,
  name VARCHAR
);

-- Create organization_role table with partition by organization_id
CREATE TABLE organization_role (
  organization_id INTEGER,
  user_id INTEGER,
  role INTEGER,
  PRIMARY KEY (organization_id, user_id),
  FOREIGN KEY (organization_id) REFERENCES organization(id)
) PARTITION BY LIST (organization_id);

-- Create event table with partition by organization_id
CREATE TABLE event (
  id SERIAL PRIMARY KEY,
  name VARCHAR,
  created_at TIMESTAMP,
  organization_id INTEGER,
  FOREIGN KEY (organization_id) REFERENCES organization(id)
) PARTITION BY LIST (organization_id);


-- Create ticket table with partition by organization_id
CREATE TABLE ticket (
  ticket_id SERIAL PRIMARY KEY, -- Add a primary key for ticket identification
  assistant_id INTEGER,
  event_id INTEGER,
  salesrep_id INTEGER,
  organization_id INTEGER, -- Add organization_id as a column
  passcode VARCHAR(255), -- Column for the passcode
  FOREIGN KEY (assistant_id) REFERENCES users(id),
  FOREIGN KEY (event_id) REFERENCES event(id),
  FOREIGN KEY (salesrep_id) REFERENCES users(id),
  FOREIGN KEY (organization_id) REFERENCES organization(id) -- Reference to the organization table
) PARTITION BY LIST (organization_id);
