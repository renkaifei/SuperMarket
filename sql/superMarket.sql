create database if not exists superMarket;


use superMarket;

create table if not exists goods(
	goodsId int not null auto_increment,
	goodsBarCode varchar(50) not null,
	goodsName varchar(256) not null,
	goodsSpecification varchar(256) not null,
	goodsDescription int not null,
	goodsTrademark varchar(128) not null default '',
	company varchar(256) not null,
	primary key(goodsId)
)CHARACTER SET utf8 COLLATE utf8_general_ci;

create table if not exists merchantGoods(
	merchantGoodsId int not null auto_increment,
	goodsId int ,
	merchantId int,
	price decimal(10,2),
	discount decimal(10,2)
	primary key(merchantGoodsId);
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
	mobilePhone varchar(50) not null,
	pwd varchar(50) not null,
	IDNumber varchar(20) not null,
	isAdmin int not null,
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

create table if not exists university(
	universityId int not null auto_increment,
	universityName varchar(128) not null,
	universityCode varchar(20) not null,
	universityAddress varchar(256) not null,
	universityPicture varchar(256) not null,
	primary key(universityId)
);

create table if not exists goodsRecomentdation(
	goodsRecommentdationId int not null auto_increment,
	goodsTitle varchar(216) not null,
	goodsDescription varchar(512) not null,
	goodsPicUrl varchar(512) not null,
	goodsUrl varchar(512) not null
)

create table if not exists Article(
	articleId int not null auto_increment,
	articleTitle varchar(256) not null,
	articleDescription varchar(512) not null,
	articlePicUrl varchar(512) not null,
	aritcleUrl varchar(512) not null,
	primary key(articleId)
)


