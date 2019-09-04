DROP DATABASE IF EXISTS `test`;
CREATE DATABASE `test`;
USE test;

CREATE TABLE `folders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY `folder_pkey` (`id`)
) ENGINE=InnoDB;

CREATE TABLE `articles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `folder_id` int(11) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `body` text DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY `article_pkey` (`id`),
  KEY `article_folder_fkey` (`folder_id`),
  CONSTRAINT `article_folder_fkey` FOREIGN KEY (folder_id) 
    REFERENCES folders(id) 
    ON DELETE CASCADE 
    ON UPDATE CASCADE
) ENGINE=InnoDB;

CREATE TABLE `tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY `tag_pkey` (`id`)
) ENGINE=InnoDB;

CREATE TABLE `article_tags` (
  `article_id` int(11) NOT NULL,
  `tag_id` int(11) NOT NULL,
  PRIMARY KEY `article_tags_pkey` (`article_id`, `tag_id`),
  KEY `article_tags_article_fkey` (`article_id`),
  KEY `article_tags_tag_fkey` (`tag_id`),

  CONSTRAINT `article_tags_article_fkey` FOREIGN KEY (article_id) 
    REFERENCES articles(id) 
    ON DELETE CASCADE 
    ON UPDATE CASCADE,

  CONSTRAINT `article_tags_tag_fkey` FOREIGN KEY (tag_id) 
    REFERENCES tags(id) 
    ON DELETE CASCADE 
    ON UPDATE CASCADE
) ENGINE=InnoDB;