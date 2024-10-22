-- +migrate Up
create table withdrawals_history
(
    id                  bigserial primary key,
    withdrawal_tx_hash  text    not null unique,
    deposit_tx_hash     text    not null unique,
    block_number        bigint  not null,
    token_address       text    not null,
    token_id            text,
    amount              text    not null,
    to_address          text    not null,
    is_mintable         boolean not null,
    destination_network text    not null
);

-- +migrate Down
drop table withdrawals_history;
