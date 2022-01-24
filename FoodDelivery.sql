CREATE TABLE `user` (
  `user_id` int PRIMARY KEY AUTO_INCREMENT,
  `user_name` char(50),
  `user_pass` char(50),
  `user_fullname` varchar(50),
  `user_address` varchar(200),
  `user_phone` numeric
);

CREATE TABLE `seller` (
  `seller_id` int PRIMARY KEY AUTO_INCREMENT,
  `seller_name` char(50),
  `seller_pass` char(50),
  `seller_fullname` varchar(50),
  `seller_address` varchar(200),
  `seller_phone` numeric
);

CREATE TABLE `category` (
  `category_id` int PRIMARY KEY AUTO_INCREMENT,
  `category_name` varchar(50),
  `category_note` varchar(200),
  `category_create_date` timestamp
);

CREATE TABLE `product` (
  `product_id` int PRIMARY KEY AUTO_INCREMENT,
  `product_name` varchar(50),
  `product_note` varchar(200),
  `product_price` numeric,
  `product_create_date` timestamp,
  `product_image` json
);

CREATE TABLE `cart` (
  `cart_id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `product_id` int,
  `product_num` int,
  `discount` float
);

CREATE TABLE `favorites` (
  `product_id` int[pk],
  `user_id` int[pk],
  `isfavorite` bool
);

CREATE TABLE `rating` (
  `product_id` int,
  `user_id` int,
  `rating` float,
  PRIMARY KEY (`product_id`, `user_id`)
);

ALTER TABLE `product` ADD FOREIGN KEY (`product_id`) REFERENCES `category` (`category_id`);

ALTER TABLE `category` ADD FOREIGN KEY (`category_id`) REFERENCES `seller` (`seller_id`);

ALTER TABLE `cart` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`);

ALTER TABLE `cart` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`);

ALTER TABLE `favorites` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`);

ALTER TABLE `favorites` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`);

ALTER TABLE `rating` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`);

ALTER TABLE `rating` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`);

