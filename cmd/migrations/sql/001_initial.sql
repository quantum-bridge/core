-- +migrate Up
create table transactions_history
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

-- +migrate Down
drop table transactions_history;
