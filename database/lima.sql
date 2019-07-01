CREATE TABLE `age_discount` (
  `age_discount_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '减龄生理年龄记录表,主键',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `date` date DEFAULT NULL COMMENT '日期,减龄记录对应日期',
  `step_average` double DEFAULT NULL COMMENT '平均步数,用户前7日平均步数',
  `age_reduce` double DEFAULT NULL COMMENT '减龄(年)',
  `user_only_sign` varchar(100) CHARACTER SET utf8 DEFAULT NULL COMMENT '用户唯一标识号',
  `info_input_id` int(11) DEFAULT NULL COMMENT '用户信息表,information_input表主键',
  PRIMARY KEY (`age_discount_id`),
  UNIQUE KEY `only_sign_date` (`user_only_sign`,`date`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

CREATE TABLE `discount_reference` (
  `discount_reference_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '折扣查询表id',
  `discount_steps` int(11) DEFAULT NULL COMMENT '步数',
  `discount_bmi_level` int(2) DEFAULT NULL COMMENT 'bmi 水平:1:low,2:medium,3:height',
  `discount_num` double DEFAULT NULL,
  PRIMARY KEY (`discount_reference_id`),
  UNIQUE KEY `step_bmi_num` (`discount_steps`,`discount_bmi_level`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

CREATE TABLE `information_input` (
  `information_input_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '京东小程序传输数据记录表',
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL,
  `only_sign` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户唯一标识号',
  `birth_date` date NOT NULL COMMENT '出生日期',
  `gender` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'NULL' COMMENT '性别:0,女性;1,男性',
  `height` double NOT NULL COMMENT '身高',
  `weight` double NOT NULL COMMENT '体重',
  `age` int(11) NOT NULL COMMENT '年龄',
  `bmi` double NOT NULL COMMENT 'BMI',
  `plan_kind` varchar(100) CHARACTER SET utf8 NOT NULL COMMENT '保险计划',
  `plan_type` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '保单类型,(0:首次投保或间断投保;1:不间断投保)',
  `plan_effective_date` date NOT NULL COMMENT '保单生效日',
  `user_type` varchar(5) CHARACTER SET utf8mb4 NOT NULL COMMENT '用户类型,(0:试玩用户;1:正式用户)',
  `plan_premium` double NOT NULL COMMENT '原始保费',
  `same` int(3) DEFAULT NULL COMMENT '是否与京东数据相同,(0:相同;1:不同)',
  `status` int(3) DEFAULT NULL COMMENT '用户状态,(1:age<18;age>60)',
  `premium_from_jd` double DEFAULT NULL COMMENT '原始保费(京东)',
  `premium_from_table` double DEFAULT NULL COMMENT '原始保费(自查)',
  PRIMARY KEY (`information_input_id`),
  KEY `only_sign` (`only_sign`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

CREATE TABLE `premium_discount` (
  `premium_discount_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户保费折扣表主键',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `info_input_id` int(11) DEFAULT NULL COMMENT '用户信息表,information_input表主键',
  `user_only_sign` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户唯一标识',
  `discount` double DEFAULT NULL COMMENT '折扣,每周保费折扣',
  `start_date` date DEFAULT NULL COMMENT '开始日期,保费折扣开始日期',
  `end_date` date DEFAULT NULL COMMENT '结束日期,保费折扣结束日期',
  `effective_num` int(11) DEFAULT NULL COMMENT '有效天数',
  `step_average` double DEFAULT NULL COMMENT '平均步数',
  `plan_effective_date` date DEFAULT NULL COMMENT '保单生效日',
  `premium_discount_num` double DEFAULT NULL COMMENT '折扣*原始保费',
  `premium_discount_type` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '保费折扣类型(16岁,70%,普通折扣)',
  `technical_discount_num` double DEFAULT NULL COMMENT 'technical折扣',
  `technical_discount_type` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '保费折扣类型(16岁,70%,普通折扣)',
  `first_week` double DEFAULT NULL COMMENT '首单,第一周保单有效日数/7,对应时间折扣',
  PRIMARY KEY (`premium_discount_id`),
  UNIQUE KEY `unique` (`user_only_sign`,`end_date`,`plan_effective_date`),
  KEY `only_sign` (`user_only_sign`),
  KEY `end_date` (`end_date`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

CREATE TABLE `premium_reference` (
  `premium_reference_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '保单原始费用查询表id',
  `premium_plan_kind` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '保单保险计划',
  `premium_sex` varchar(6) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '保单人性别:female(男),male(女)',
  `premium_age` int(3) DEFAULT NULL COMMENT '保单人年龄,存在年龄为0',
  `premium_price` double DEFAULT NULL COMMENT '保单原始费用',
  `premium_plan_type` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '间断投保类型(0:间断或首次;1:不间断)',
  PRIMARY KEY (`premium_reference_id`),
  UNIQUE KEY `unique` (`premium_plan_kind`,`premium_sex`,`premium_age`,`premium_plan_type`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

CREATE TABLE `step` (
  `step_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '微信步数表主键',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `date` date DEFAULT NULL COMMENT '日期',
  `step_num` int(11) DEFAULT NULL COMMENT '步数',
  `user_only_sign` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户唯一标识号',
  `info_input_id` int(11) DEFAULT NULL COMMENT '用户表,information_input表主键',
  PRIMARY KEY (`step_id`),
  UNIQUE KEY `unique_day_onlysign` (`user_only_sign`,`date`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

CREATE TABLE `step_update` (
  `step_update_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户平均步数表主键',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `step_average_num` double DEFAULT NULL COMMENT '用户7天平均步数',
  `step_average_day` int(11) DEFAULT NULL COMMENT '用户连续更新天数',
  `user_only_sign` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户唯一标志号',
  PRIMARY KEY (`step_update_id`),
  KEY `only_sign` (`user_only_sign`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci