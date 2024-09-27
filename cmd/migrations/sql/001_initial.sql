-- +migrate Up
create table transactions_history
(
    id            bigserial primary key,
    tx_hash       text      not null,
    token_address text      not null,
    token_id      text,
    amount        text      not null,
    from_address  text      not null,
    to_address    text      not null,
    network_from  text      not null,
    network_to    text      not null
);

-- +migrate Down
drop table transactions_history;
