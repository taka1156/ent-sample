-- modify "user" table
ALTER TABLE `user` DROP FOREIGN KEY `user_team_user`;
-- modify "user" table
ALTER TABLE `user` MODIFY COLUMN `team_id` int NULL, ADD CONSTRAINT `user_team_user` FOREIGN KEY (`team_id`) REFERENCES `team` (`team_id`) ON DELETE SET NULL;
