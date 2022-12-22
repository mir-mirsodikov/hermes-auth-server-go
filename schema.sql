CREATE TABLE "user" (
  id uuid default gen_random_uuid(),
  name text not null,
  email text not null unique,
  username text not null unique,

  primary key (id)
);

