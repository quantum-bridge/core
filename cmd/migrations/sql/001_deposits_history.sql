-- +migrate Up
create table deposits_history
(
    id                  bigserial primary key,
    tx_hash             text    not null unique,
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
drop table deposits_history;
