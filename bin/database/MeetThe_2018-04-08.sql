# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.6.33)
# Database: MeetThe
# Generation Time: 2018-04-08 11:31:09 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table Answers
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Answers`;

CREATE TABLE `Answers` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `questionPeopleId` int(11) unsigned NOT NULL,
  `answer` varchar(200) NOT NULL DEFAULT '',
  `isCorrect` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `questionPeopleId` (`questionPeopleId`),
  CONSTRAINT `answers_ibfk_1` FOREIGN KEY (`questionPeopleId`) REFERENCES `PeopleQuestions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table Games
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Games`;

CREATE TABLE `Games` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `personId` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `personId` (`personId`),
  CONSTRAINT `games_ibfk_1` FOREIGN KEY (`personId`) REFERENCES `People` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table GameTrack
# ------------------------------------------------------------

DROP TABLE IF EXISTS `GameTrack`;

CREATE TABLE `GameTrack` (
  `gameId` int(11) unsigned NOT NULL,
  `peopleId` int(11) unsigned NOT NULL DEFAULT '0',
  `peopleQuestionId` int(11) unsigned NOT NULL DEFAULT '0',
  `result` tinyint(1) NOT NULL DEFAULT '0',
  `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`gameId`,`peopleId`,`peopleQuestionId`),
  KEY `peopleId` (`peopleId`),
  KEY `peopleQuestionId` (`peopleQuestionId`),
  CONSTRAINT `gametrack_ibfk_1` FOREIGN KEY (`gameId`) REFERENCES `Games` (`id`),
  CONSTRAINT `gametrack_ibfk_2` FOREIGN KEY (`peopleId`) REFERENCES `People` (`id`),
  CONSTRAINT `gametrack_ibfk_3` FOREIGN KEY (`peopleQuestionId`) REFERENCES `PeopleQuestions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table Images
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Images`;

CREATE TABLE `Images` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(500) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table Levels
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Levels`;

CREATE TABLE `Levels` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `description` varchar(100) NOT NULL DEFAULT '',
  `sort` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table People
# ------------------------------------------------------------

DROP TABLE IF EXISTS `People`;

CREATE TABLE `People` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(80) NOT NULL DEFAULT '',
  `active` tinyint(1) NOT NULL DEFAULT '1',
  `photoId` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `photoId` (`photoId`),
  CONSTRAINT `people_ibfk_1` FOREIGN KEY (`photoId`) REFERENCES `Images` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table PeopleQuestions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `PeopleQuestions`;

CREATE TABLE `PeopleQuestions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `questionId` int(11) unsigned NOT NULL,
  `peopleId` int(11) unsigned NOT NULL,
  PRIMARY KEY (`peopleId`,`questionId`),
  UNIQUE KEY `id` (`id`),
  KEY `questionId` (`questionId`),
  CONSTRAINT `peoplequestions_ibfk_1` FOREIGN KEY (`questionId`) REFERENCES `Questions` (`id`),
  CONSTRAINT `peoplequestions_ibfk_2` FOREIGN KEY (`peopleId`) REFERENCES `People` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table Questions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Questions`;

CREATE TABLE `Questions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `question` varchar(400) NOT NULL DEFAULT '',
  `insertedBy` int(11) unsigned NOT NULL,
  `photoId` int(11) unsigned NOT NULL,
  `showImageAtBeginning` tinyint(1) NOT NULL DEFAULT '1',
  `allPeople` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `insertedBy` (`insertedBy`),
  KEY `photoId` (`photoId`),
  CONSTRAINT `questions_ibfk_2` FOREIGN KEY (`insertedBy`) REFERENCES `People` (`id`),
  CONSTRAINT `questions_ibfk_3` FOREIGN KEY (`photoId`) REFERENCES `Images` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table QuestionsLevel
# ------------------------------------------------------------

DROP TABLE IF EXISTS `QuestionsLevel`;

CREATE TABLE `QuestionsLevel` (
  `questionId` int(11) unsigned NOT NULL,
  `levelId` int(11) unsigned NOT NULL,
  PRIMARY KEY (`questionId`,`levelId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
