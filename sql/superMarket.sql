create database if not exists superMarket;


use superMarket;

create table if not exists goods(
	goodsId int not null auto_increment,
	goodsName varchar(256) not null,
	goodsCategoryId int not null,
	primary key(goodsId)
);

create table if not exists goodsCategory(
	categoryId int not null auto_increment,
	categoryName varchar(256) not null,
	parentId int not null,
	primary key(categoryId)
);


