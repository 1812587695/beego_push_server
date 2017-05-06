CREATE TABLE `websocket_order_msg` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `local_id` varchar(255) NOT NULL DEFAULT '',
  `time` datetime DEFAULT NULL,
  `msg` varchar(255) DEFAULT NULL,
  `status` smallint(1) unsigned DEFAULT '0' COMMENT '0É¾³ý',
  PRIMARY KEY (`id`),
  KEY `local_id` (`local_id`),
  KEY `time` (`time`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;

