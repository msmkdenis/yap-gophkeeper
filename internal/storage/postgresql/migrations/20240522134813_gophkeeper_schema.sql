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
    number                  bytea not null,
    owner_name              bytea not null,
    expires_at              bytea,
    cvv_code                bytea,
    pin_code                bytea,
    metadata                text,
    created_at              timestamp not null,
    updated_at              timestamp not null,
    constraint pk_credit_card primary key (id),
    constraint fk_owner_id foreign key (owner_id) references gophkeeper.user (id),
    constraint ux_credit_card__number unique (number)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
