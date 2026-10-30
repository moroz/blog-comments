-- +goose Up
create table comments (
    id uuid not null primary key,
    post_path text not null,
    author text not null,
    body text not null,
    website text,
    inserted_at timestamp not null default (unixepoch()),
    check (typeof(id) = 'text' and length(id) = 36), -- UUID in its canonical form is 36 characters long
    check (typeof(post_path) = 'text' and post_path like '/blog/%/'),
    check (typeof(author) = 'text' and length(author) between 1 and 64),
    check (typeof(body) = 'text')
);

create index comments_post_path_idx on comments (post_path);

-- +goose Down
drop table comments;
