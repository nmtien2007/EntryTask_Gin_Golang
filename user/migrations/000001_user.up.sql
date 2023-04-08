BEGIN;
   CREATE TABLE IF NOT EXISTS user(
      id bigserial PRIMARY KEY,
      username VARCHAR (50) UNIQUE NOT NULL,
      password VARCHAR (150) NOT NULL,
      first_name VARCHAR (50) NOT NULL,
      last_name VARCHAR(50) NOT NULL,
      created_at timestamp with time zone
      updated_at timestamp with time zone
   );

COMMIT;
