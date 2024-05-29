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

create table if not exists gophkeeper.credit_card
(
    id                      text,
    owner_id                text not null,
    data                    bytea not null,
    metadata                text,
    created_at              timestamp not null,
    updated_at              timestamp not null,
    constraint pk_credit_card primary key (id),
    constraint fk_owner_id foreign key (owner_id) references gophkeeper.user (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
