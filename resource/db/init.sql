drop
database easy-frpc;
create
database easy-frpc if not exists;
use
easy-frpc;
create table user
(
    id       int          not null auto_increment primary key,
    username varchar(255) not null,

)