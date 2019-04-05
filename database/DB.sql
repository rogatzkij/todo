CREATE DATABASE todo;
USE todo;
/*
	Пользователи
*/
CREATE TABLE E1_Users
(
	login	VARCHAR(20)	NOT NULL,
	email	TINYTEXT	NOT NULL,
	hash	VARCHAR(20)	NOT NULL,
	PRIMARY KEY (login)
);

/*
	Доступные сессии
*/
CREATE TABLE E2_Session
(
	idSession            INTEGER AUTO_INCREMENT,
	cookie               VARCHAR(20) NULL,
	login                VARCHAR(20) NOT NULL,
	PRIMARY KEY (idSession),
	FOREIGN KEY (login) REFERENCES E1_Users(login)
);

/*
	Активные Дела
*/
CREATE TABLE E3_Tasks
(
	idTask               INTEGER AUTO_INCREMENT,
	description          VARCHAR(20) NULL,
	title                INTEGER NULL,
	defer                INTEGER NULL,
	dateEnd              DATE NULL,
	login                VARCHAR(20) NOT NULL,
	PRIMARY KEY (idTask),
	FOREIGN KEY (login) REFERENCES E1_Users(login)
);

/*
	Архивные Дела
*/
CREATE TABLE E4_Archive
(
	idArchive            INTEGER AUTO_INCREMENT,
	title                VARCHAR(20) NULL,
	description          VARCHAR(20) NULL,
	dateEnd              DATE NULL,
	dateReady            DATE NULL,
	login                VARCHAR(20) NOT NULL,
	PRIMARY KEY (idArchive),
	FOREIGN KEY (login) REFERENCES E1_Users(login)
);

/*
	Дополнительная информация
*/
CREATE TABLE E5_AdditionalInformation
(
	login                VARCHAR(20) NOT NULL,
	firstName            VARCHAR(20) NULL,
	lastName             VARCHAR(20) NULL,
	dateRegistration     DATE NULL,
	PRIMARY KEY (login),
	FOREIGN KEY (login) REFERENCES E1_Users(login)
);

/*
	Награды
*/
CREATE TABLE E6_Achievement
(
	idAchievement          INTEGER AUTO_INCREMENT,
	achievementName        VARCHAR(20) NULL,
	achievementDescription VARCHAR(20) NULL,
	PRIMARY KEY (idAchievement)
);

/*
	Награды пользователя
*/
CREATE TABLE E7_UsersAchievement
(
	login                VARCHAR(20) NOT NULL,
	idAchievement          INTEGER NOT NULL,
	dateAchivment        DATE NULL,
	PRIMARY KEY (login,idAchievement),
	FOREIGN KEY (login) REFERENCES E1_Users(login),
	FOREIGN KEY (idAchievement) REFERENCES E6_Achievement(idAchievement)
);



