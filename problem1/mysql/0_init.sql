CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `name` varchar(64) DEFAULT '' NOT NULL,
  PRIMARY KEY (`id`)
);
-- user1 user2
CREATE TABLE `friend_link` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user1_id` int(11) NOT NULL,
  `user2_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
);
-- user1 user2 block
CREATE TABLE `block_list` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user1_id` int(11) NOT NULL,
  `user2_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO users (user_id, name) VALUES (1,"A");
INSERT INTO users (user_id, name) VALUES (2,"B");
INSERT INTO users (user_id, name) VALUES (3,"C");
INSERT INTO users (user_id, name) VALUES (4,"D");
INSERT INTO users (user_id, name) VALUES (5,"E");
INSERT INTO users (user_id, name) VALUES (6,"F");
INSERT INTO users (user_id, name) VALUES (7,"G");
INSERT INTO users (user_id, name) VALUES (8,"H");
INSERT INTO users (user_id, name) VALUES (9,"I");
INSERT INTO users (user_id, name) VALUES (10,"J");

INSERT INTO friend_link (user1_id, user2_id) VALUES (1,2);
INSERT INTO friend_link (user1_id, user2_id) VALUES (1,3);
INSERT INTO friend_link (user1_id, user2_id) VALUES (1,4);
INSERT INTO friend_link (user1_id, user2_id) VALUES (1,5);
INSERT INTO friend_link (user1_id, user2_id) VALUES (1,7);
INSERT INTO friend_link (user1_id, user2_id) VALUES (1,8);
INSERT INTO friend_link (user1_id, user2_id) VALUES (1,9);
INSERT INTO friend_link (user1_id, user2_id) VALUES (1,10);
INSERT INTO friend_link (user1_id, user2_id) VALUES (2,3);
INSERT INTO friend_link (user1_id, user2_id) VALUES (2,6);
INSERT INTO friend_link (user1_id, user2_id) VALUES (2,7);
INSERT INTO friend_link (user1_id, user2_id) VALUES (2,8);
INSERT INTO friend_link (user1_id, user2_id) VALUES (3,6);
INSERT INTO friend_link (user1_id, user2_id) VALUES (3,9);
INSERT INTO friend_link (user1_id, user2_id) VALUES (3,10);
;
INSERT INTO block_list (user1_id, user2_id) VALUES (1,10);
INSERT INTO block_list (user1_id, user2_id) VALUES (3,6);