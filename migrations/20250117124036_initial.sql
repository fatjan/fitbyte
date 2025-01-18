-- +goose Up
-- +goose StatementBegin
create type activity_type_enum as enum (
    'Walking',
    'Yoga',
    'Stretching',
    'Cycling',
    'Swimming',
    'Dancing',
    'Hiking',
    'Running',
    'HIIT',
    'JumpRope'
);

create type preference_enum as enum (
    'CARDIO',
    'WEIGHT'
);

create type weight_unit_enum as enum (
    'KG',
    'LBS'
);

create type height_unit_enum as enum (
    'CM',
    'INCH'
);

create table if not exists users (
    id serial primary key,
    email text unique not null,
    password_hash text not null,
    name varchar(60),
    preference preference_enum,
    weight_unit weight_unit_enum,
    height_unit height_unit_enum,
    weight int,
    height int,
    image_uri text,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

CREATE INDEX IF NOT EXISTS users_email_idx ON users USING HASH (email);

create table if not exists activities (
    id serial primary key,
    user_id int references users(id) on delete cascade,
    activity_type activity_type_enum,
    done_at timestamptz default null,
    duration_in_minutes int default 1,
    calories_burned int,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

CREATE INDEX IF NOT EXISTS user_id_idx ON activities USING HASH (user_id);
CREATE INDEX IF NOT EXISTS user_id_doneAt_idx ON activities(user_id, doneAt);
CREATE INDEX IF NOT EXISTS user_id_caloriesBurned_idx ON activities(user_id, caloriesBurned);
CREATE INDEX IF NOT EXISTS user_id_activityType_idx ON activities(user_id, activityType);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists activities;
drop table if exists users;
drop type if exists activity_type_enum;
drop type if exists preference_enum;
drop type if exists weight_unit_enum;
drop type if exists height_unit_enum;
-- +goose StatementEnd
