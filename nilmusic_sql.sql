/*
 Source Server         : suvvm_aliyun_beijing
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 101.201.70.76:3306
 Source Schema         : nilmusic

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 21/11/2020 17:29:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for album_music
-- ----------------------------
DROP TABLE IF EXISTS `album_music`;
CREATE TABLE `album_music` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID，主键',
  `aid` int(11) NOT NULL COMMENT '歌单ID，外键',
  `mid` int(11) NOT NULL COMMENT '歌曲ID，外键',
  PRIMARY KEY (`id`),
  KEY `aid` (`aid`),
  KEY `mid` (`mid`),
  CONSTRAINT `album_music_ibfk_1` FOREIGN KEY (`aid`) REFERENCES `albums` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `album_music_ibfk_2` FOREIGN KEY (`mid`) REFERENCES `musics` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=50000000 DEFAULT CHARSET=utf8mb4 COMMENT='歌单歌曲关系表';

-- ----------------------------
-- Table structure for albums
-- ----------------------------
DROP TABLE IF EXISTS `albums`;
CREATE TABLE `albums` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID，主键',
  `name` varchar(200) CHARACTER SET utf8 NOT NULL DEFAULT 'new album' COMMENT '专辑名称',
  `poster` varchar(500) CHARACTER SET utf8 NOT NULL DEFAULT 'https://www.suvvm.work/images/ortrait.jpg' COMMENT '专辑封面URL',
  `playnum` varchar(20) CHARACTER SET utf8 NOT NULL DEFAULT '0万' COMMENT '专辑播放量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20000000 DEFAULT CHARSET=utf8mb4 COMMENT='专辑信息表';

-- ----------------------------
-- Table structure for musics
-- ----------------------------
DROP TABLE IF EXISTS `musics`;
CREATE TABLE `musics` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID，主键',
  `name` varchar(200) CHARACTER SET utf8 NOT NULL DEFAULT 'new music' COMMENT '音乐名称',
  `poster` varchar(500) CHARACTER SET utf8 NOT NULL DEFAULT 'https://www.suvvm.work/images/ortrait.jpg' COMMENT '音乐封面URL',
  `path` varchar(500) CHARACTER SET utf8 NOT NULL DEFAULT 'http://m8.music.126.net/20201119220648/17233129086daaf596237f43b218beb5/ymusic/1a32/22d0/301e/3964f63dc993257f280cb214cefc403a.mp3' COMMENT '音乐外链URL',
  `author` varchar(200) CHARACTER SET utf8 NOT NULL DEFAULT 'suvvm' COMMENT '音乐作者',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30000000 DEFAULT CHARSET=utf8mb4 COMMENT='歌曲信息表';

-- ----------------------------
-- Table structure for user_album
-- ----------------------------
DROP TABLE IF EXISTS `user_album`;
CREATE TABLE `user_album` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID，主键',
  `uid` int(11) NOT NULL COMMENT '用户ID，外键',
  `aid` int(11) NOT NULL COMMENT '歌单ID，外键',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `aid` (`aid`),
  CONSTRAINT `user_album_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `user_album_ibfk_2` FOREIGN KEY (`aid`) REFERENCES `albums` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=40000000 DEFAULT CHARSET=utf8mb4 COMMENT='用户歌单关系表';

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID，主键',
  `pnum` varchar(20) CHARACTER SET utf8 NOT NULL DEFAULT '11111111111' COMMENT '用户手机号',
  `password` varchar(200) CHARACTER SET utf8 NOT NULL DEFAULT 'poiuytrewq' COMMENT '用户密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000000 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

SET FOREIGN_KEY_CHECKS = 1;
