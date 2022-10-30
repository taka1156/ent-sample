-- reverse: modify "user" table
ALTER TABLE `user` DROP FOREIGN KEY `user_team_user`, MODIFY COLUMN `team_id` int NOT NULL;
-- reverse: modify "user" table
ALTER TABLE `user` ADD CONSTRAINT `user_team_user` FOREIGN KEY (`team_id`) REFERENCES `sample`.`team` (`team_id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
