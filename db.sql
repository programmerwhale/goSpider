-- 豆瓣电影top250
CREATE TABLE `movie` (
     `Id` bigint(20) NOT NULL AUTO_INCREMENT,
     `Title` varchar(255) DEFAULT NULL,
     `Director` varchar(255) DEFAULT NULL,
     `Picture` varchar(255) DEFAULT NULL,
     `Actor` varchar(255) DEFAULT NULL,
     `Year` varchar(255) DEFAULT NULL,
     `Score` varchar(255) DEFAULT NULL,
     `Quote` varchar(255) DEFAULT NULL,
     `CreatedAt` datetime DEFAULT NULL,
     PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=126 DEFAULT CHARSET=utf8;


-- 书
CREATE TABLE `books` (
     `Id` bigint(20) NOT NULL AUTO_INCREMENT,
     `BookId` varchar(255) DEFAULT NULL,
     `Title` varchar(255) DEFAULT NULL,
     `Author` varchar(255) DEFAULT NULL,
     `Picture` varchar(255) DEFAULT NULL,
     `Year` varchar(255) DEFAULT NULL,
     `Rating` varchar(255) DEFAULT NULL,
     `RateWord` varchar(255) DEFAULT NULL,
     `Page` varchar(255) DEFAULT NULL,
     `IsCollect` tinyint(1) DEFAULT '0',
     `IsAddToNotion` tinyint(1) DEFAULT '0',
     `CreatedAt` datetime DEFAULT NULL,
     PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=302 DEFAULT CHARSET=utf8;