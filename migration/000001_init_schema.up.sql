set timezone to 'Asia/Taipei';  
-- comments table
CREATE TABLE IF NOT EXISTS comments (
  id                  SERIAL PRIMARY KEY,
  uuid                char(36) UNIQUE NOT NULL,
  parentid            char(36) NOT NULL,
  comment             varchar(500),
  author              varchar(50),
  favorite            boolean,
  created_at          TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
  updated_at          TIMESTAMP WITH TIME ZONE    DEFAULT CURRENT_TIMESTAMP
);
create index idx_comments_desc on comments (id desc nulls last); -- desc nulls last : large small null


