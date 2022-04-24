CREATE TABLE organizations (
    id SERIAL PRIMARY KEY,
    login varchar(50) NOT NULL,
    name varchar(100) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    deleted_at timestamptz
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login varchar(50) NOT NULL,
    name varchar(100) NOT NULL,
    followers int NOT NULL DEFAULT 0,
    follows int NOT NULL DEFAULT 0,
    avatar_url text,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    deleted_at timestamptz
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    user_id int,
    organization_id int NOT NULL,
    text text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    deleted_at timestamptz
);

CREATE TABLE memberships (
    user_id int NOT NULL,
    organization_id int NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    deleted_at timestamptz,
    PRIMARY KEY (user_id, organization_id)
);

ALTER TABLE comments ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE comments ADD FOREIGN KEY (organization_id) REFERENCES organizations (id);

ALTER TABLE memberships ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE memberships ADD FOREIGN KEY (organization_id) REFERENCES organizations (id);
