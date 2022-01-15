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

    user_id integer references user(id) on delete cascade not null
);

-- triggers
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
