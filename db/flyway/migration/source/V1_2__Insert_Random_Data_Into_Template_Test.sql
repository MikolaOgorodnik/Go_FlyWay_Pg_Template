insert into Template_Test (id, rnd_str, rnd_float, rnd_int)
select val,
       md5(random()::text)           as rnd_str,
       random()::float8 * pow(10, 3) as rnd_float,
       floor(random() * pow(10, 5))  as rnd_int
from generate_series(1, pow(10, 3)::int) as val;