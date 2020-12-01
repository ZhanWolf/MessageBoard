SET NAMES utf8;

CREATE DATABASE IF NOT EXISTS stu_info;
USE stu_info;
DROP TABLE IF EXISTS `student`;
DROP TABLE IF EXISTS `score`;

CREATE TABLE `student` (
    `student_id` BIGINT NOT NULL DEFAULT 0 PRIMARY KEY,
    `real_name` VARCHAR(255) NULL DEFAULT NULL,
    `password` VARCHAR(255)  NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

CREATE TABLE `class` (
    `student_id` BIGINT NOT NULL DEFAULT 0 PRIMARY KEY,
    `university` VARCHAR(255) NULL,
    `college` VARCHAR(255) NULL,
    `class` VARCHAR(255) NULL,
    FOREIGN KEY(student_id) REFERENCES student(student_id)
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

CREATE TABLE `point` (
    `student_id` BIGINT NOT NULL DEFAULT 0,
    `subject` VARCHAR(255) NOT NULL,
    `point` DOUBLE NOT NULL,
    FOREIGN KEY(student_id) REFERENCES student(student_id),
    PRIMARY KEY (`subject`, `student_id`)
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

INSERT INTO `student` VALUES (2020233233, '王小美', 'Wangxiaomei011224');
INSERT INTO `class` VALUES (2020233233, "重庆邮电大学", "移动学院", "2班");
INSERT INTO `point` VALUES (2020233233, "art", 3.6);
INSERT INTO `point` VALUES (2020233233, "math", 3.2);
INSERT INTO `point` VALUES (2020233233, "sports", 1.0);

-- select avg of point and list

SELECT student.student_id, student.real_name,
MAX(CASE WHEN point.subject = 'art' THEN point ELSE 0 END) AS art,
MAX(CASE WHEN point.subject = 'math' THEN point ELSE 0 END) AS math,
MAX(CASE WHEN point.subject = 'sports' THEN point ELSE 0 END) AS sports,
AVG(point) AS `avg_point`
FROM student LEFT JOIN point ON student.student_id = point.student_id GROUP BY student.student_id;

/* output
+------------+-----------+------+------+--------+-----------+
| student_id | real_name | art  | math | sports | avg_point |
+------------+-----------+------+------+--------+-----------+
| 2020233233 | 王小美    |  3.6 |  3.2 |      1 |       2.6 |
+------------+-----------+------+------+--------+-----------+
*/

-- selest good subject

SELECT `subject` AS `good_subjects` FROM `point` WHERE `student_id` = 2020233233 AND `point` > 3.0;

/* output
+---------------+
| good_subjects |
+---------------+
| art           |
| math          |
+---------------+
 */

