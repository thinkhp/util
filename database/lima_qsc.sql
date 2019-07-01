CREATE TABLE `age_reduce` (
  `age_reduce_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'lima减龄记录表,主键',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `date` date DEFAULT NULL COMMENT '日期,减龄记录对应日期',
  `step_average` double DEFAULT NULL COMMENT '平均步数,用户前7日平均步数',
  `age_reduce` double DEFAULT NULL COMMENT '减龄(年)',
  `user_only_sign` varchar(100) DEFAULT NULL COMMENT '用户唯一标识号',
  `info_input_id` int(11) DEFAULT NULL COMMENT '用户信息表,information_input表主键',
  PRIMARY KEY (`age_reduce_id`),
  UNIQUE KEY `only_sign_date` (`user_only_sign`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `discount_forage_reference` (
  `discount_forage_reference_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '折扣查询表id',
  `step` int(11) DEFAULT NULL COMMENT '步数',
  `gender` varchar(10) DEFAULT NULL COMMENT '性别',
  `discount` double DEFAULT NULL,
  PRIMARY KEY (`discount_forage_reference_id`),
  UNIQUE KEY `step_gender_discount` (`step`,`gender`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `discount_reference` (
  `discount_reference_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '折扣查询表id',
  `step` int(11) DEFAULT NULL COMMENT '步数',
  `bmi_level` int(2) DEFAULT NULL COMMENT 'bmi 水平:1:low,2:medium,3:height',
  `discount` double DEFAULT NULL,
  PRIMARY KEY (`discount_reference_id`),
  UNIQUE KEY `step_bmi_num` (`step`,`bmi_level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `factor_reference` (
  `factor_reference_id` int(11) NOT NULL COMMENT '扩大因子表 id,减龄用',
  `age` int(11) DEFAULT NULL COMMENT '年龄',
  `gender` varchar(5) DEFAULT NULL COMMENT '性别',
  `factor` double DEFAULT NULL COMMENT '扩大因子',
  PRIMARY KEY (`factor_reference_id`),
  UNIQUE KEY `age_gender_factor` (`age`,`gender`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `information_input` (
  `information_input_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '请求数据记录表',
  `create_time` datetime DEFAULT NULL,
  `only_sign` varchar(100) DEFAULT NULL COMMENT '用户唯一标识号',
  `birth_date` date DEFAULT NULL COMMENT '出生日期',
  `gender` varchar(10) DEFAULT 'NULL' COMMENT '性别:0,女性;1,男性',
  `height` double DEFAULT NULL COMMENT '身高',
  `weight` double DEFAULT NULL COMMENT '体重',
  `age` int(11) DEFAULT NULL COMMENT '年龄',
  `bmi` double DEFAULT NULL COMMENT 'BMI',
  `plan_kind` varchar(100) DEFAULT NULL COMMENT '保险计划',
  `plan_type` varchar(50) DEFAULT NULL COMMENT '保单类型,(0:首次投保或间断投保;1:不间断投保)',
  `plan_effective_date` date DEFAULT NULL COMMENT '保单生效日',
  `user_type` varchar(5) DEFAULT NULL COMMENT '用户类型,(0:试玩用户;1:正式用户)',
  `plan_premium` double DEFAULT NULL COMMENT '原始保费',
  `same` int(3) DEFAULT NULL COMMENT '是否与京东数据相同,(0:相同;1:不同)',
  `status` int(3) DEFAULT NULL COMMENT '用户状态,(1:age<18;age>60)',
  `premium_from_u` double DEFAULT NULL COMMENT '原始保费(合作方:京东,轻松筹)',
  `premium_from_m` double DEFAULT NULL COMMENT '原始保费(自查)',
  `pay_method` varchar(5) DEFAULT NULL COMMENT '支付方式(年缴:year;月缴:year)',
  PRIMARY KEY (`information_input_id`),
  KEY `only_sign` (`only_sign`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8

CREATE TABLE `information_output` (
  `information_output_id` int(11) NOT NULL COMMENT '京东小程序输出表',
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL,
  `information_input_id` int(11) NOT NULL COMMENT '原始保费',
  `premium` double NOT NULL COMMENT '用户原始保费',
  `step_average` double NOT NULL COMMENT '平均步数(-1:没有平均步数)',
  `customer_bmi` double NOT NULL COMMENT '用户BMI',
  `discount` double NOT NULL COMMENT '保单折扣',
  `premium_discount` double NOT NULL COMMENT '折后保单',
  `age_reduce_initial` double NOT NULL COMMENT '初始减龄幅度',
  `factor` double NOT NULL COMMENT '扩大因子',
  `age` int(11) NOT NULL COMMENT '年龄',
  `factor_adjust` double NOT NULL COMMENT '调整后的扩大因子',
  `age_reduce_final` double NOT NULL COMMENT '最终减龄幅度',
  `age_output` double NOT NULL COMMENT '最终输出的生理年龄',
  PRIMARY KEY (`information_output_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `premium_award` (
  `premium_award_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户保费折扣表主键',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `info_input_id` int(11) DEFAULT NULL COMMENT '用户信息表,information_input表主键',
  `user_only_sign` varchar(100) DEFAULT NULL COMMENT '用户唯一标识',
  `discount` double DEFAULT NULL COMMENT '折扣,每周保费折扣',
  `start_date` date DEFAULT NULL COMMENT '开始日期,保费折扣开始日期',
  `end_date` date DEFAULT NULL COMMENT '结束日期,保费折扣结束日期',
  `effective_num` int(11) DEFAULT NULL COMMENT '有效天数',
  `step_average` double DEFAULT NULL COMMENT '平均步数',
  `plan_effective_date` date DEFAULT NULL COMMENT '保单生效日',
  `premium_discount_num` double DEFAULT NULL COMMENT '折扣*原始保费',
  `premium_discount_type` varchar(45) DEFAULT NULL COMMENT '保费折扣类型(16岁,70%,普通折扣)',
  `technical_discount_num` double DEFAULT NULL COMMENT 'technical折扣',
  `technical_discount_type` varchar(45) DEFAULT NULL COMMENT '保费折扣类型(16岁,70%,普通折扣)',
  `first_week` double DEFAULT NULL COMMENT '首单,第一周保单有效日数/7,对应时间折扣',
  PRIMARY KEY (`premium_award_id`),
  UNIQUE KEY `unique` (`user_only_sign`,`end_date`,`plan_effective_date`),
  KEY `only_sign` (`user_only_sign`),
  KEY `end_date` (`end_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `premium_rate_reference` (
  `premium_rate_reference_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '原始保费率查询表id,用于减龄计算',
  `gender` varchar(6) DEFAULT NULL COMMENT '性别:female(男),male(女)',
  `age` int(3) DEFAULT NULL COMMENT '年龄,存在年龄为0',
  `premium_rate` double DEFAULT NULL COMMENT ' 原始保费率',
  PRIMARY KEY (`premium_rate_reference_id`),
  UNIQUE KEY `unique` (`gender`,`age`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `premium_reference` (
  `premium_reference_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '保单原始费用查询表id',
  `plan_kind` varchar(100) DEFAULT NULL COMMENT '保单保险计划',
  `gender` varchar(6) DEFAULT NULL COMMENT '保单人性别:female(男),male(女)',
  `age` int(3) DEFAULT NULL COMMENT '保单人年龄,存在年龄为0',
  `premium` double DEFAULT NULL COMMENT '保单原始费用',
  `plan_type` varchar(50) DEFAULT NULL COMMENT '间断投保类型(0:间断或首次;1:不间断)',
  PRIMARY KEY (`premium_reference_id`),
  UNIQUE KEY `unique` (`plan_kind`,`gender`,`age`,`plan_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `step` (
  `step_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '微信步数表主键',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `date` date DEFAULT NULL COMMENT '日期',
  `step_num` int(11) DEFAULT NULL COMMENT '步数',
  `user_only_sign` varchar(100) DEFAULT NULL COMMENT '用户唯一标识号',
  `info_input_id` int(11) DEFAULT NULL COMMENT '用户表,information_input表主键',
  PRIMARY KEY (`step_id`),
  UNIQUE KEY `unique_day_onlysign` (`user_only_sign`,`date`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8

CREATE TABLE `step_update` (
  `step_update_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户平均步数表主键',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `step_average_num` double DEFAULT NULL COMMENT '用户7天平均步数',
  `step_average_day` int(11) DEFAULT NULL COMMENT '用户连续更新天数',
  `user_only_sign` varchar(100) DEFAULT NULL COMMENT '用户唯一标志号',
  PRIMARY KEY (`step_update_id`),
  KEY `only_sign` (`user_only_sign`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8