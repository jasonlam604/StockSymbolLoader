--
-- Table structure for table `symbols`
--
CREATE TABLE `symbols` (
  `id` int(10) UNSIGNED NOT NULL,
  `exchange` enum('UNKNOWN','NASDAQ','TSX','TSXV') NOT NULL DEFAULT 'UNKNOWN',
  `code` varchar(10) NOT NULL,
  `status` enum('ACTIVE','INACTIVE','PENDING_REVIEW') NOT NULL,
  `company_name` varchar(256) NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

ALTER TABLE `symbols`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `exchange` (`exchange`,`code`);

ALTER TABLE `symbols`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;

