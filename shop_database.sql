-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 06, 2022 at 07:01 AM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `shop_databse`
--

-- --------------------------------------------------------

--
-- Table structure for table `code_coupon`
--

CREATE TABLE `code_coupon` (
  `ID` int(10) UNSIGNED NOT NULL,
  `ID_COUPON` int(20) NOT NULL,
  `UNIQUE_ID` varchar(50) NOT NULL,
  `ACTIVE_USE_STATUS` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `coupon`
--

CREATE TABLE `coupon` (
  `ID` int(20) UNSIGNED NOT NULL,
  `ID_COUPON` int(20) NOT NULL,
  `ID_CUSTOMER` varchar(50) NOT NULL,
  `GENERATE_DATE` varchar(50) NOT NULL,
  `REVENUE` int(50) NOT NULL,
  `ACTIVE_DATE_STATUS` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `customer`
--

CREATE TABLE `customer` (
  `ID` int(20) UNSIGNED NOT NULL,
  `UNIQ_ID` varchar(50) NOT NULL,
  `NAME` varchar(100) NOT NULL,
  `JOIN_DATE` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `customer`
--

INSERT INTO `customer` (`ID`, `UNIQ_ID`, `NAME`, `JOIN_DATE`) VALUES
(1, 'MIN-BGR1-021222163858', 'Mina', '2022-12-02 16:38:58.'),
(2, 'MOM-BGR1-021222163858', 'Momo', '2022-12-02 16:38:58.'),
(3, 'SAN-BGR2-021222163858', 'Sana', '2022-12-02 16:38:58.'),
(4, 'MIN-BGR1-021222164020', 'Mina', '2022-12-02 16:40:20.'),
(5, 'MOM-BGR1-021222164020', 'Momo', '2022-12-02 16:40:20.'),
(6, 'SAN-BGR2-021222164020', 'Sana', '2022-12-02 16:40:20.');

-- --------------------------------------------------------

--
-- Table structure for table `inventory`
--

CREATE TABLE `inventory` (
  `ID` int(20) UNSIGNED NOT NULL,
  `ID_PRODUCT` int(20) NOT NULL,
  `PRODUCT_NAME` varchar(100) NOT NULL,
  `BRAND` varchar(100) NOT NULL,
  `PRICE` int(50) NOT NULL,
  `CATEGORY` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `inventory`
--

INSERT INTO `inventory` (`ID`, `ID_PRODUCT`, `PRODUCT_NAME`, `BRAND`, `PRICE`, `CATEGORY`) VALUES
(1, 100001, 'Playstation 5', 'Sony', 11000000, 'New Console'),
(2, 100002, 'Playstation 4', 'Sony', 5500000, 'New Console'),
(3, 100003, 'Nintendo Switch', 'Nintendo', 4000000, 'New Console'),
(4, 200001, 'GTA Collection PS5', 'Rockstar Games', 1200000, 'New Game'),
(5, 200002, 'FIFA 2022 PS4', 'EA', 700000, 'New Game'),
(6, 200003, 'RESIDENT EVIL 7 PS5', 'CAPCOM', 750000, 'New Game'),
(7, 300001, 'RESIDENT EVIL 7 PS5', 'CAPCOM', 400000, 'Second Game'),
(8, 300002, 'Story of Seasons : Mineral Town Wii', 'Natsume', 250000, 'Second Game'),
(9, 400001, 'Joystick PS5', 'Sony', 880000, 'Accessories Console'),
(10, 400002, 'Cable HDMI High-Speed', 'Vention', 70000, 'Accessories Console'),
(11, 400003, 'Cable HDMI Ultra', 'Vention', 190000, 'Accessories Console'),
(12, 500001, 'Upgrade SSD M.2 Sata', 'Service', 900000, 'Service Console'),
(13, 500002, 'Service & Cleaning', 'Service', 150000, 'Service Console');

-- --------------------------------------------------------

--
-- Table structure for table `settlement`
--

CREATE TABLE `settlement` (
  `ID` int(20) UNSIGNED NOT NULL,
  `USER` int(20) NOT NULL,
  `CUSTOMER` int(20) NOT NULL,
  `COUPON_ID` int(20) NOT NULL,
  `CODE_TRANSACTION` varchar(50) NOT NULL,
  `TOTAL_PRICE` int(50) NOT NULL,
  `STATUS_TRANSACTION` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `settlement`
--

INSERT INTO `settlement` (`ID`, `USER`, `CUSTOMER`, `COUPON_ID`, `CODE_TRANSACTION`, `TOTAL_PRICE`, `STATUS_TRANSACTION`) VALUES
(1, 0, 0, 0, 'BGR01122022014610', 25000000, 0),
(2, 0, 0, 0, 'BGR01122022015110', 25000000, 0),
(3, 0, 0, 0, 'BGR01122022015908', 22000000, 1),
(4, 0, 0, 0, '001122022205838', 27000000, 0),
(5, 0, 0, 0, '001122022205907', 22000000, 1),
(6, 0, 0, 0, '001122022210615', 3200000, 1);

-- --------------------------------------------------------

--
-- Table structure for table `settlement_purchased`
--

CREATE TABLE `settlement_purchased` (
  `ID` int(20) UNSIGNED NOT NULL,
  `ID_SETTLEMENT` int(50) NOT NULL,
  `ID_ITEM` int(50) NOT NULL,
  `QTY` int(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `ID` int(20) UNSIGNED NOT NULL,
  `ID_USER` int(20) NOT NULL,
  `NAME` varchar(100) NOT NULL,
  `OUTLET_CODE` varchar(20) NOT NULL,
  `STATUS` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`ID`, `ID_USER`, `NAME`, `OUTLET_CODE`, `STATUS`) VALUES
(1, 7001, 'Joe', 'BGR1', 'Owner'),
(2, 1001, 'Yuna', 'BGR1', 'Employee'),
(3, 1002, 'Yeji', 'BGR1', 'Employee'),
(4, 1003, 'Ryujin', 'BGR1', 'Employee'),
(5, 7002, 'Juria', 'BGR2', 'Owner'),
(6, 1004, 'Lia', 'BGR2', 'Employee'),
(7, 1005, 'Chae', 'BGR2', 'Employee');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `code_coupon`
--
ALTER TABLE `code_coupon`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `ID_SETTLEMENT` (`ID_COUPON`),
  ADD UNIQUE KEY `UNIQ_ID_COUPON` (`ID_COUPON`);

--
-- Indexes for table `coupon`
--
ALTER TABLE `coupon`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `ID_SETTLEMENT` (`ID_COUPON`),
  ADD UNIQUE KEY `ID_COUPON` (`ID_COUPON`);

--
-- Indexes for table `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `UNIQ_ID` (`UNIQ_ID`);

--
-- Indexes for table `inventory`
--
ALTER TABLE `inventory`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `ID_SETTLEMENT` (`ID_PRODUCT`),
  ADD UNIQUE KEY `ID_PRODUCT` (`ID_PRODUCT`);

--
-- Indexes for table `settlement`
--
ALTER TABLE `settlement`
  ADD PRIMARY KEY (`ID`);

--
-- Indexes for table `settlement_purchased`
--
ALTER TABLE `settlement_purchased`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `ID_SETTLEMENT` (`ID_SETTLEMENT`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`ID`),
  ADD UNIQUE KEY `ID_SETTLEMENT` (`ID_USER`),
  ADD UNIQUE KEY `ID_USER` (`ID_USER`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `code_coupon`
--
ALTER TABLE `code_coupon`
  MODIFY `ID` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `coupon`
--
ALTER TABLE `coupon`
  MODIFY `ID` int(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `customer`
--
ALTER TABLE `customer`
  MODIFY `ID` int(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `inventory`
--
ALTER TABLE `inventory`
  MODIFY `ID` int(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT for table `settlement`
--
ALTER TABLE `settlement`
  MODIFY `ID` int(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `settlement_purchased`
--
ALTER TABLE `settlement_purchased`
  MODIFY `ID` int(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `ID` int(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
