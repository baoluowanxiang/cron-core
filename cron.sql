-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2019-12-16 19:35:20
-- 服务器版本： 5.7.27-log
-- PHP 版本： 7.1.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `cron`
--

-- --------------------------------------------------------

--
-- 表的结构 `cron_agent`
--

CREATE TABLE `cron_agent` (
  `id` int(11) NOT NULL,
  `service` varchar(32) NOT NULL COMMENT 'agent 所属服务',
  `ip` varchar(32) NOT NULL COMMENT '机器ip',
  `status` tinyint(4) NOT NULL COMMENT '状态'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `cron_job`
--

CREATE TABLE `cron_job` (
  `id` int(11) NOT NULL,
  `name` varchar(64) CHARACTER SET utf8mb4 NOT NULL,
  `crontab` varchar(16) CHARACTER SET utf8mb4 NOT NULL,
  `service_name` varchar(64) CHARACTER SET utf8mb4 NOT NULL,
  `create_time` timestamp NOT NULL,
  `last_execute_time` timestamp NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci COMMENT='定时任务数据表';

--
-- 转存表中的数据 `cron_job`
--

INSERT INTO `cron_job` (`id`, `name`, `crontab`, `service_name`, `create_time`, `last_execute_time`) VALUES
(1, '测试', '*/3 * * * * *', 'tms', '2019-12-18 16:00:00', '2019-12-18 16:00:00');

--
-- 转储表的索引
--

--
-- 表的索引 `cron_agent`
--
ALTER TABLE `cron_agent`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `cron_job`
--
ALTER TABLE `cron_job`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `cron_agent`
--
ALTER TABLE `cron_agent`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `cron_job`
--
ALTER TABLE `cron_job`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
