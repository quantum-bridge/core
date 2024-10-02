-- +migrate Up
create table withdrawals_history
(
    id                  bigserial primary key,
    tx_hash             text    not null unique references deposits_history (tx_hash),
    block_number        bigint  not null,
    token_address       text    not null,
    token_id            text,
    amount              text    not null,
    from_address        text    not null,
    to_address          text    not null,
    is_mintable         boolean not null,
    source_network      text    not null,
    destination_network text    not null
);

-- +migrate Down
drop table withdrawals_history;
