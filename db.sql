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
     PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;