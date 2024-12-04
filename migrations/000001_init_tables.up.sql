START TRANSACTION;

-- Create table otps
CREATE TABLE IF NOT EXISTS otps (
  created_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  updated_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  deleted_at datetime(6) DEFAULT NULL,
  id int(11) NOT NULL AUTO_INCREMENT,
  code varchar(255) NOT NULL,
  identifier varchar(255) NOT NULL,
  trial int(11) NOT NULL,
  is_valid tinyint(4) NOT NULL DEFAULT '0',
  expired_at datetime DEFAULT NULL,
  PRIMARY KEY (id)
);

-- Create table permissions
CREATE TABLE IF NOT EXISTS `permissions` (
  created_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  updated_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  deleted_at datetime(6) DEFAULT NULL,
  id int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `key` varchar(255) NOT NULL,
  PRIMARY KEY (id)
);

-- Create table roles
CREATE TABLE IF NOT EXISTS roles (
  created_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  updated_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  deleted_at datetime(6) DEFAULT NULL,
  id int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `key` varchar(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY IDX_a87cf0659c3ac379b339acf36a (`key`)
);

-- Create table users
CREATE TABLE IF NOT EXISTS users (
  created_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  updated_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  deleted_at datetime(6) DEFAULT NULL,
  id int(11) NOT NULL AUTO_INCREMENT,
  fullname varchar(255) NOT NULL,
  identity_number varchar(255) NOT NULL,
  phone_number varchar(255) NOT NULL,
  email_verified_at datetime DEFAULT NULL,
  one_signal_player_ids json DEFAULT NULL,
  phone_verified_at datetime DEFAULT NULL,
  email varchar(255) DEFAULT NULL,
  birth_date datetime DEFAULT NULL,
  gender varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  image_url varchar(255) DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY IDX_43d2ef62e309fe8f4bae2a67e5 (identity_number),
  UNIQUE KEY IDX_17d1817f241f10a3dbafb169fd (phone_number),
  UNIQUE KEY IDX_97672ac88f789774dd47f7c8be (email)
);

-- Create table in_app_notifications
CREATE TABLE IF NOT EXISTS in_app_notifications (
  created_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  updated_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  deleted_at datetime(6) DEFAULT NULL,
  id int(11) NOT NULL AUTO_INCREMENT,
  target_user_id int(11) NOT NULL,
  `type` varchar(255) NOT NULL,
  title varchar(255) NOT NULL,
  `message` varchar(255) NOT NULL,
  meta json DEFAULT NULL,
  is_read tinyint(4) NOT NULL,
  PRIMARY KEY (id),
  KEY FK_046713440a98830b619c4c649b4 (target_user_id),
  CONSTRAINT FK_046713440a98830b619c4c649b4 FOREIGN KEY (target_user_id) REFERENCES users (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

-- Create table log_activities
CREATE TABLE IF NOT EXISTS log_activities (
  id int(11) NOT NULL AUTO_INCREMENT,
  meta_data json DEFAULT NULL,
  source varchar(255) DEFAULT NULL,
  activity varchar(255) NOT NULL,
  menu varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  created_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  user_id int(11) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY FK_4357a91cbef922677d73d510f70 (user_id),
  KEY log_activity_menu (menu) USING BTREE,
  KEY log_activity_activity (activity) USING BTREE,
  CONSTRAINT FK_4357a91cbef922677d73d510f70 FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

-- Create table role_permissions
CREATE TABLE IF NOT EXISTS role_permissions (
  created_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  updated_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  deleted_at datetime(6) DEFAULT NULL,
  id int(11) NOT NULL AUTO_INCREMENT,
  role_id int(11) DEFAULT NULL,
  permission_id int(11) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY IDX_3d0a7155eafd75ddba5a701336 (role_id),
  KEY IDX_e3a3ba47b7ca00fd23be4ebd6c (permission_id),
  CONSTRAINT FK_3d0a7155eafd75ddba5a7013368 FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT FK_e3a3ba47b7ca00fd23be4ebd6cf FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE CASCADE ON UPDATE NO ACTION
);

-- Create table user_role
CREATE TABLE IF NOT EXISTS user_role (
  created_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  updated_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  deleted_at datetime(6) DEFAULT NULL,
  id int(11) NOT NULL AUTO_INCREMENT,
  role_id int(11) DEFAULT NULL,
  user_id int(11) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY IDX_d0e5815877f7395a198a4cb0a4 (user_id),
  KEY IDX_32a6fc2fcb019d8e3a8ace0f55 (role_id),
  CONSTRAINT FK_32a6fc2fcb019d8e3a8ace0f55f FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT FK_d0e5815877f7395a198a4cb0a46 FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);

COMMIT;