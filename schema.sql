CREATE TABLE "user" (
  id uuid default gen_random_uuid(),
  name text not null,
  email text not null unique,
  username text not null unique,
  verfied boolean default false,

  primary key (id)
);

create table "verification" (
  user_id uuid not null references "user" (id),
  code int not null,
  created_at timestamp default (now() at time zone 'utc'),

  primary key (user_id),
  unique (user_id, created_at)
)