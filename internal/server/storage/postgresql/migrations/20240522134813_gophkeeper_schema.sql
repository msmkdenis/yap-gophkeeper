-- +goose Up
-- +goose StatementBegin
create schema if not exists gophkeeper;

create table if not exists gophkeeper.user
(
    id                      text,
    login                   text not null,
    password                bytea not null,
    crypt_key               bytea not null,
    created_at              timestamp not null,
    updated_at              timestamp not null,
    constraint pk_user primary key (id),
    constraint ux_user__login unique (login)
);

create type gophkeeper.data_type as enum
    ('credit_card', 'text_data', 'credentials', 'binary_data');

create table if not exists gophkeeper.data
(
    id                      text,
    owner_id                text not null,
    type                    gophkeeper.data_type not null,
    data                    bytea not null,
    metadata                text,
    created_at              timestamp not null,
    updated_at              timestamp not null,
    constraint pk_credit_card primary key (id, type),
    constraint fk_owner_id foreign key (owner_id) references gophkeeper.user (id)
) partition by list (type);

create index if not exists gophkeeper_data_type_idx on gophkeeper.data (type);

create table if not exists gophkeeper.data_credit_card partition of gophkeeper.data
    for values in ('credit_card');

create table if not exists gophkeeper.data_text_data partition of gophkeeper.data
    for values in ('text_data');

create table if not exists gophkeeper.data_credentials partition of gophkeeper.data
    for values in ('credentials');

create table if not exists gophkeeper.data_binary_data partition of gophkeeper.data
    for values in ('binary_data');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
