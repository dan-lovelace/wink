-- new tables
create table user (
    id integer not null primary key autoincrement,
	created_at text,
	updated_at text,
    active boolean default true not null,

    username text unique not null
);

create table project (
    id integer not null primary key autoincrement,
	created_at text,
	updated_at text,
    active boolean default true not null,

    user_id integer references user(id) on delete cascade not null,
    
    name text not null
);

create table time_block (
    id integer not null primary key autoincrement,
	created_at text,
	updated_at text,
    active boolean default true not null,

    project_id integer references project(id) on delete cascade not null,

    uid text not null,
    start_time text,
    stop_time text,
    description text
);

-- user triggers
drop trigger if exists user_row_inserted;
create trigger user_row_inserted
after insert on user
for each row
begin
    update user
    set
        created_at = datetime('NOW'),
	    updated_at = datetime('NOW'),
	    active = true
    where id = new.id;
end;

drop trigger if exists user_row_updated;
create trigger user_row_updated
after update on user
for each row
begin
    update user
    set updated_at = datetime('NOW')
    where id = new.id;
end;

-- project triggers
drop trigger if exists project_row_inserted;
create trigger project_row_inserted
after insert on project
for each row
begin
    update project
    set
        created_at = datetime('NOW'),
	    updated_at = datetime('NOW'),
	    active = true
    where id = new.id;
end;

drop trigger if exists project_row_updated;
create trigger project_row_updated
after update on project
for each row
begin
    update project
    set updated_at = datetime('NOW')
    where id = new.id;
end;

-- time_block triggers
drop trigger if exists time_block_row_inserted;
create trigger time_block_row_inserted
after insert on time_block
for each row
begin
    update time_block
    set
        created_at = datetime('NOW'),
	    updated_at = datetime('NOW'),
	    active = true
    where id = new.id;
end;

drop trigger if exists time_block_row_updated;
create trigger time_block_row_updated
after update on time_block
for each row
begin
    update time_block
    set updated_at = datetime('NOW')
    where id = new.id;
end;
