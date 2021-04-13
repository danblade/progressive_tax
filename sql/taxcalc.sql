create or replace function tax(income numeric) returns numeric
language sql
as $$
select floor(sum(rate * taxable_income)) as tax
from (
select *,
case
  when income < lower(range) then 0
  when income > upper(range) then upper(range) - lower(range)
  else income - lower(range)
  end as taxable_income
from tax_brackets
) t
;
$$
;
