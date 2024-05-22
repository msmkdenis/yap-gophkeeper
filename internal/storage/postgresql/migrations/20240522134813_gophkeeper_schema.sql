-- +goose Up
-- +goose StatementBegin
create schema if not exists gophkeeper;

create table if not exists gophkeeper.user
(
    id                      text,
    login                   text unique not null,
    password                bytea not null,
    created_at              timestamp not null,
    updated_at              timestamp not null,
    constraint pk_user primary key (id),
    constraint ux_user__login unique (login)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
