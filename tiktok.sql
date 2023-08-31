CREATE DATABASE IF NOT EXISTS `tiktok` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

USE tiktok;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account
-- ----------------------------
INSERT INTO `account` VALUES (1, 1, 'admin', 'admin');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `follow_count` int NULL DEFAULT 0,
  `follower_count` int NULL DEFAULT 0,
  `is_follow` tinyint(1) NULL DEFAULT 0,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `total_favorited` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `work_count` int NULL DEFAULT 0,
  `favorite_count` int NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `userName`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'Liion', 0, 0, 0, 'https://www.helloimg.com/images/2023/08/25/oiG7nT.jpg', 'https://www.helloimg.com/images/2023/08/25/oiGQUq.png', '家人们，谁懂啊', '1', 10, 100);

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '视频唯一标识ID',
  `user_id` bigint NULL DEFAULT NULL COMMENT '视频作者ID',
  `play_url` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频播放地址',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频封面地址',
  `favorite_count` bigint NULL DEFAULT NULL COMMENT '视频的点赞总数',
  `comment_count` bigint NULL DEFAULT NULL COMMENT '视频的评论总数',
  `is_favorite` tinyint(1) NULL DEFAULT NULL COMMENT '是否点赞（true-已点赞，false-未点赞）',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频标题',
  `post_time` datetime NULL DEFAULT NULL COMMENT '视频投稿时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1, 1, 'http://vfx.mtime.cn/Video/2021/07/10/mp4/210710094507540173.mp4', 'https://www.helloimg.com/images/2023/08/29/oiEHNM.jpg', 100, 1000, 1, 'Big Buck', '2023-08-22 15:04:25');
INSERT INTO `video` VALUES (2, 1, 'http://www.w3school.com.cn/example/html5/mov_bbb.mp4', 'https://www.helloimg.com/images/2023/08/29/oiEHNM.jpg', 200, 2000, 0, 'Butterfly', '2023-08-23 15:05:50');
INSERT INTO `video` VALUES (3, 1, 'https://media.w3.org/2010/05/sintel/trailer.mp4', 'https://www.helloimg.com/images/2023/08/29/oiEHNM.jpg', 300, 3000, 0, 'Snow Mountain', '2023-08-08 15:06:51');
INSERT INTO `video` VALUES (12, 1, 'https://4b866bc09c8ece71b8354b53bae25b8f-app.1024paas.com/static/3_1693452019.mp4', 'https://www.helloimg.com/images/2023/08/29/oiEHNM.jpg', 400, 4000, 0, 'test', '2023-08-29 12:57:38');

SET FOREIGN_KEY_CHECKS = 1;
