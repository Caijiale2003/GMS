DROP TABLE IF EXISTS `enroll_models`;
DROP TABLE IF EXISTS `score_models`;
DROP TABLE IF EXISTS `team_models`;
DROP TABLE IF EXISTS `user_models`;


-- 创建用户表，主键为学号
CREATE TABLE `user_models` (
                               `id` varchar(20) NOT NULL  COMMENT '学号',
                               `name` varchar(32) DEFAULT NULL COMMENT '姓名',
                               `password` varchar(256) COMMENT '密码',
                               `gender` varchar(16) COMMENT '性别',
                               `academy` longtext COMMENT '学院',
                               `major` varchar(128) COMMENT '专业',
                               `role` bigint(4) COMMENT '权限',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `user_models` WRITE;
UNLOCK TABLES;


-- 创建登陆数据表
DROP TABLE IF EXISTS `login_data_models`;
CREATE TABLE `login_data_models` (
                                     `id` varchar(20) NOT NULL  COMMENT '学号',
                                     `name` varchar(32) DEFAULT NULL COMMENT '姓名',
                                     `academy` longtext COMMENT '学院',
                                     `role` varchar(16) COMMENT '权限',
                                     `token` varchar(256) COMMENT 'token',
                                     `time` varchar(256) COMMENT 'time'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `login_data_models` WRITE;
UNLOCK TABLES;


-- 创建比赛表，主键为名称
DROP TABLE IF EXISTS `game_models`;
CREATE TABLE `game_models` (
                              `name` varchar(32) NOT NULL COMMENT '比赛名称',
                              `organizer` longtext COMMENT '主办方',
                              `start_time` varchar(256) COMMENT '比赛开始时间',
                              `end_time` varchar(256) COMMENT '比赛结束时间',
                              `address` varchar(128) COMMENT '比赛地点',
                              `prize` varchar(128) COMMENT '比赛奖品',
                              `creator` varchar(128) COMMENT '创建者',
                            primary key (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `game_models` WRITE;
UNLOCK TABLES;


-- 触发器：删除比赛表中的数据时，相关联的表相应数据也一并删除
CREATE trigger before_delete_game
    before delete ON game_models
    for each row
begin
    delete from enroll_models WHERE name = enroll_models.game_name;
END;

drop trigger before_delete_game;

-- 创建报名表，比赛名称和学号为外键
CREATE TABLE `enroll_models` (
                               `game_name` varchar(32) NOT NULL COMMENT '比赛名称',
                               `id` varchar(20) NOT NULL  COMMENT '学号',
                               `name` varchar(32) DEFAULT NULL COMMENT '姓名',
                               `teacher` varchar(32) DEFAULT NULL COMMENT '指导老师',
                               `team_name` varchar(32) DEFAULT NULL COMMENT '队伍名称'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

alter table enroll_models add constraint fk_game_name foreign key (game_name) references game_models(name) on delete cascade ;
alter table enroll_models add constraint fk_id foreign key (id) references user_models(id);

LOCK TABLES `enroll_models` WRITE;
UNLOCK TABLES;


-- 创建队伍表，学号为外键
CREATE TABLE `team_models` (
                                 `team_name` varchar(32) NOT NULL COMMENT '队伍名称',
                                 `id` varchar(20) NOT NULL  COMMENT '学号',
                                 `name` varchar(32) DEFAULT NULL COMMENT '姓名'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
alter table team_models add constraint fk_team_id foreign key (id) references user_models(id);
LOCK TABLES `team_models` WRITE;
UNLOCK TABLES;


CREATE TABLE `score_models` (
                               `game_name` varchar(32) NOT NULL COMMENT '比赛名称',
                               `id` varchar(20) NOT NULL  COMMENT '学号',
                               `score` int DEFAULT NULL COMMENT '分数'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
alter table score_models add constraint fk_score_id foreign key (id) references user_models(id);
alter table score_models add constraint fk_score_name foreign key (game_name) references game_models(name);
LOCK TABLES `score_models` WRITE;
UNLOCK TABLES;