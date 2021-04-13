drop table if exists tax_brackets;

create table tax_brackets(
  range numrange not null,
  rate numeric not null
)
;

insert into tax_brackets
values
  (numrange(0, 10000), 0),
  (numrange(10000, 30000), 0.1),
  (numrange(30000, 100000), 0.25),
  (numrange(100000, null), 0.4)
;
