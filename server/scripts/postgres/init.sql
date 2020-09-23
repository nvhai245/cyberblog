create table if not exists posts (
id serial primary key,
author_id integer not null,
parent_id integer,
title text default 'Không có chủ đề',
published bool default false,
up_vote integer default 0,
content text not null,
created_at integer default date_part('epoch', now()),
updated_at integer,
published_at integer
);

create table if not exists comments (
id serial primary key,
post_id integer not null,
author_id integer not null,
content text default 'Hay!',
up_vote integer default 0,
created_at integer default date_part('epoch', now()),
updated_at integer
);

create table if not exists category (
id serial primary key,
title text not null,
slug text not null,
content text default 'Viết nội dung cho thể loại này.',
created_at integer default date_part('epoch', now()),
updated_at integer
);

create table if not exists post_category (
post_id integer not null,
category_id integer not null
);