
-- test varchar column minus
drop table if exists t1;
create table t1(
a varchar(100),
b varchar(100)
);

insert into t1 values('dddd', 'cccc');
insert into t1 values('aaaa', 'bbbb');
insert into t1 values('eeee', 'aaaa');
insert into t1 values ();
select * from t1;

(select a from t1) minus (select b from t1);
(select a from t1) minus (select a from t1 limit 1);
(select a from t1) minus (select a from t1 limit 2);
(select a from t1) minus (select a from t1 limit 3);
(select a from t1) minus (select a from t1 limit 4);

(select b from t1) minus (select a from t1);
(select b from t1) minus (select b from t1 limit 1);
(select b from t1) minus (select b from t1 limit 2);
(select b from t1) minus (select b from t1 limit 3);
(select b from t1) minus (select b from t1 limit 4);

(select a from t1) minus (select b from t1) minus (select b from t1);
(select a from t1) minus (select b from t1) minus (select b from t1) minus (select a from t1);

((select a from t1) union (select b from t1)) minus (select a from t1);

drop table t1;

-- test date type and time type minus
drop table if exists t2;
create table t2(
col1 date,
col2 datetime,
col3 timestamp
);

insert into t2 values ();
insert into t2 values('2022-01-01', '2022-01-01', '2022-01-01');
insert into t2 values('2022-01-01', '2022-01-01 00:00:00', '2022-01-01 00:00:00.000000');
insert into t2 values('2022-01-01', '2022-01-01 00:00:00.000000', '2022-01-01 23:59:59.999999');
select * from t2;

-- test date type and time type
(select col1 from t2) minus (select col1 from t2 limit 1);
(select col1 from t2) minus (select col2 from t2);
(select col1 from t2) minus (select col3 from t2);

(select col1 from t2) minus (select col1 from t2 limit 0);
(select col1 from t2) minus (select col2 from t2 limit 0);
(select col1 from t2) minus (select col3 from t2 limit 0);

(select col2 from t2) minus (select col1 from t2 limit 0);
(select col2 from t2) minus (select col2 from t2 limit 0);
(select col2 from t2) minus (select col3 from t2 limit 0);

(select col3 from t2) minus (select col1 from t2 limit 0);
(select col3 from t2) minus (select col2 from t2 limit 0);
(select col3 from t2) minus (select col3 from t2 limit 0);

drop table t2;


-- test union and minus
drop table if exists t3;
create table t3(
a int
);

insert into t3 values (20),(10),(30),(-10);

drop table if exists t4;
create table t4(
col1 float,
col2 decimal(5,2)
);

insert into t4 values(100.01,100.01);
insert into t4 values(1.10,1.10);
insert into t4 values(0.0,0.0);
insert into t4 values(127.0,127.0);
insert into t4 values(127.44,127.44);


(select a from t3) union (select col1 from t4) minus (select col1 from t4);
(select a from t3) union (select col2 from t4) minus (select col2 from t4);
(select a from t3) union (select col1 from t4) minus (select a from t3);
(select a from t3) union (select col2 from t4) minus (select a from t3);

(select col1 from t4) minus (select col1 from t4);
(select col1 from t4) minus (select col2 from t4);
(select col2 from t4) minus (select col2 from t4);
(select col2 from t4) minus (select col1 from t4);

drop table t3;
drop table t4;

-- test int type and text type union varchar type and text type
drop table if exists t5;
create table t5(
a int,
b text
);

insert into t5 values (11, 'aa');
insert into t5 values (33, 'bb');
insert into t5 values (44, 'aa');
insert into t5 values (55, 'cc');
insert into t5 values (55, 'dd');

drop table if exists t6;
create table t6 (
col1 varchar(100),
col2 text,
col3 char(100)
);

insert into t6 values ('aa', '11', 'aa');
insert into t6 values ('bb', '22', '11');
insert into t6 values ('cc', '33', 'bb');
insert into t6 values ('dd', '44', '22');

(select a from t5) minus (select col2 from t6);
(select col2 from t6) minus (select a from t5);

(select b from t5) minus (select col1 from t6);
(select b from t5) minus (select col2 from t6);
(select b from t5) minus (select col3 from t6);

(select col1 from t6) minus (select b from t5);
(select col2 from t6) minus (select b from t5);
(select col3 from t6) minus (select b from t5);

drop table t5;
drop table t6;

-- test subquery minus
drop table if exists t7;
CREATE TABLE t7 (
a int not null,
b char (10) not null
);

insert into t7 values(1,'3'),(2,'4'),(3,'5'),(3,'1');

select * from (select a from t7 minus select a from t7) a;

select * from (select a from t7 minus select b from t7) a;
select * from (select b from t7 minus select a from t7) a;

select * from (select a from t7 minus (select b from t7 limit 2)) a;
select * from (select b from t7 minus (select a from t7 limit 2)) a;

select * from (select a from t7 minus select b from t7 limit 1) a;
select * from (select b from t7 minus select a from t7 limit 1) a;

select * from (select a from t7 minus select b from t7 where a > 2) a;
select * from (select b from t7 minus select a from t7 where a > 4) a;

drop table t7;

-- test minus prepare
drop table if exists t8;
create table t8 (
a int primary key,
b int
);

insert into t8 values (1,5),(2,4),(3,3);
set @a=1;

prepare s1 from '(select a from t8 where a>?) minus (select b from t8 where b>?)';
prepare s2 from '(select a from t8 where a>?)';

execute s1 using @a;
execute s1 using @a, @a;
execute s2 using @a;
execute s2 using @a, @a;

deallocate prepare s1;
deallocate prepare s2;

drop table t8;

-- test minus join, left join, right join
drop table if exists t9;
create table t9(
a int,
b varchar
);

insert into t9 values (1, 'a'), (2, 'b'), (3,'c'), (4, 'd');

drop table if exists t10;
create table t10(
c int,
d varchar
);

insert into t10 values (1, 'a'), (10, 'b'), (2,'b'), (2, 'e');

(select a from t9) minus (select tab1.a from t9 as tab1 join t10 as tab2 on tab1.a=tab2.c);
(select a from t9) minus (select tab1.a from t9 as tab1 left join t10 as tab2 on tab1.a=tab2.c);
(select a from t9) minus (select tab1.a from t9 as tab1 right join t10 as tab2 on tab1.a=tab2.c);

(select b from t9) minus (select tab1.b from t9 as tab1 join t10 as tab2 on tab1.b=tab2.d);
(select b from t9) minus (select tab1.b from t9 as tab1 left join t10 as tab2 on tab1.b=tab2.d);
(select b from t9) minus (select tab1.b from t9 as tab1 right join t10 as tab2 on tab1.b=tab2.d);

(select c from t10) minus (select tab1.a from t9 as tab1 join t10 as tab2 on tab1.a=tab2.c);
(select c from t10) minus (select tab1.a from t9 as tab1 left join t10 as tab2 on tab1.a=tab2.c);
(select c from t10) minus (select tab1.a from t9 as tab1 right join t10 as tab2 on tab1.a=tab2.c);

(select d from t10) minus (select tab1.b from t9 as tab1 join t10 as tab2 on tab1.b=tab2.d);
(select d from t10) minus (select tab1.b from t9 as tab1 left join t10 as tab2 on tab1.b=tab2.d);
(select d from t10) minus (select tab1.b from t9 as tab1 right join t10 as tab2 on tab1.b=tab2.d);

drop table t9;
drop table t10;


drop table if exists t11;
create table t11 (
RID int(11) not null default '0',
IID int(11) not null default '0',
nada varchar(50)  not null,
NAME varchar(50) not null,
PHONE varchar(50) not null);

insert into t11 ( RID,IID,nada,NAME,PHONE) values
(1, 1, 'main', 'a', '111'),
(2, 1, 'main', 'b', '222'),
(3, 1, 'main', 'c', '333'),
(4, 1, 'main', 'd', '444'),
(5, 1, 'main', 'e', '555'),
(6, 2, 'main', 'c', '333'),
(7, 2, 'main', 'd', '454'),
(8, 2, 'main', 'e', '555'),
(9, 2, 'main', 'f', '666'),
(10, 2, 'main', 'g', '777');

select A.NAME, A.PHONE, B.NAME, B.PHONE from t11 A
left join t11 B on A.NAME = B.NAME and B.IID = 2 where A.IID = 1 and (A.PHONE <> B.PHONE or B.NAME is null)
minus
select A.NAME, A.PHONE, B.NAME, B.PHONE from t11 B left join t11 A on B.NAME = A.NAME and A.IID = 1
where B.IID = 2 and (A.PHONE <> B.PHONE or A.NAME is null);

select A.NAME, A.PHONE, B.NAME, B.PHONE from t11 B left join t11 A on B.NAME = A.NAME and A.IID = 1
where B.IID = 2 and (A.PHONE <> B.PHONE or A.NAME is null)
minus
select A.NAME, A.PHONE, B.NAME, B.PHONE from t11 A
left join t11 B on A.NAME = B.NAME and B.IID = 2 where A.IID = 1 and (A.PHONE <> B.PHONE or B.NAME is null);

set @val1=1;
set @val2=2;

prepare s1 from
'select A.NAME, A.PHONE, B.NAME, B.PHONE from t11 A
left join t11 B on A.NAME = B.NAME and B.IID = ? where A.IID = ? and (A.PHONE <> B.PHONE or B.NAME is null)
minus
select A.NAME, A.PHONE, B.NAME, B.PHONE from t11 B left join t11 A on B.NAME = A.NAME and A.IID = ?
where B.IID = ? and (A.PHONE <> B.PHONE or A.NAME is null)';

execute s1 using @val2, @val1, @val1, @val2;

deallocate prepare s1;

drop table t11;

-- test primary key minus
drop table if exists t12;

create table t12(
a int primary key,
b int auto_increment
);

insert into t12(a) values (1);
insert into t12(a) values (2);
insert into t12(a) values (3);
insert into t12(a) values (10);
insert into t12(a) values (20);

(select a from t12 ) minus (select b from t12);
(select a from t12 ) minus (select a from t12);

(select b from t12 ) minus (select a from t12);
(select b from t12 ) minus (select b from t12);

drop table t12;
