# glng_ks08_kelompok5_final_Project_1


create table todo (
todo_id serial primary key,
title varchar (225) not null,
completed bool,
Created_at timestamptz default now(),
Updated_at timestamptz default now()
);
