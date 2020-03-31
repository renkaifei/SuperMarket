create database if not exists superMarket;


use superMarket;

create table if not exists goods(
	goodsId int not null auto_increment,
	goodsName varchar(256) not null,
	goodsCode varchar(256) not null,
	goodsCategoryId int not null,
	goodsBarCode varchar(128) not null default '',
	merchantId int not null,
	primary key(goodsId)
);

create table if not exists goodsCategory (
	categoryId int not null auto_increment,
	categoryName varchar(256) not null,
	parentId int not null,
	primary key(categoryId)
);

create table if not exists merchanter (
	merchanterId int not null auto_increment,
	merchanterOpenId varchar(256) not null,
	merchanterName varchar(256) not null,
	merchantId int not null,
	primary key(merchanterId)
);

create table if not exists merchant (
	merchantId int not null auto_increment,
	merchantName varchar(256) not null,
	merchantAlias varchar(256) not null,
	merchantAddress varchar(512) not null,
	socialCreditCode varchar(50) not null,
	ceo varchar(50) not null,
	mobilePhone varchar(20) not null,
	primary key(merchantId)
);
