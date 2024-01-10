CREATE TABLE IF NOT EXISTS "user" (  
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  user_email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS summary (
  id SERIAL PRIMARY KEY,
  artifact_url VARCHAR(255) NOT NULL,
  user_id UUID REFERENCES "user"(id) ON DELETE CASCADE, 
  summary JSONB
);