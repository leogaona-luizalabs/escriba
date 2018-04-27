create table drafts
(
	url varchar(500) not null primary key,
	approvals int default '0' not null,
	created_at datetime default CURRENT_TIMESTAMP not null,
	published_at datetime null,
	constraint drafts_draft_url_uindex unique (url)
)
engine=InnoDB;