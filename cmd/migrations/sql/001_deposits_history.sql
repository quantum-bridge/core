-- +migrate Up
create table deposits_history
(
    id            bigserial primary key,
    chain_id      int     not null,
    tx_hash       text    not null unique,
    block_number  bigint  not null,
    token_address text    not null,
    token_id      text,
    amount        text    not null,
    from_address  text    not null,
    to_address    text    not null,
    to_network    text    not null,
    is_mintable   boolean not null
);

create table withdrawals_history
(
    id            bigserial primary key,
    chain_id      int     not null,
    tx_hash       text    not null references deposits_history(tx_hash),
    token_address text    not null,
    token_id      text,
    amount        text    not null,
    to_address    text    not null,
    token_uri     text,
    is_mintable   boolean not null
);

-- +migrate Down
drop table deposits_history;
drop table withdrawals_history;