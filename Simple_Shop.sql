-- phpMyAdmin SQL Dump
-- version 4.8.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 23, 2020 at 09:14 AM
-- Server version: 10.1.34-MariaDB
-- PHP Version: 7.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `Simple_Shop`
--

-- --------------------------------------------------------

--
-- Table structure for table `Daftar_Barang`
--

CREATE TABLE `Daftar_Barang` (
  `ID_Barang` int(11) NOT NULL,
  `Nama_Barang` varchar(20) NOT NULL,
  `Harga_Barang` int(20) NOT NULL,
  `ID_User` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `Daftar_Barang`
--

INSERT INTO `Daftar_Barang` (`ID_Barang`, `Nama_Barang`, `Harga_Barang`, `ID_User`) VALUES
(1, 'Sandal', 9000, 1),
(2, 'Sepatu', 200000, 1),
(3, 'Pensil', 5000, 1),
(4, 'Mouse', 20000, 1),
(6, 'Charger HP', 30000, 1),
(7, 'Masker', 3000, 1);

-- --------------------------------------------------------

--
-- Table structure for table `User`
--

CREATE TABLE `User` (
  `ID_User` int(11) NOT NULL,
  `U_Username` varchar(20) NOT NULL,
  `U_Password` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `User`
--

INSERT INTO `User` (`ID_User`, `U_Username`, `U_Password`) VALUES
(1, 'Admin', 'Admin');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Daftar_Barang`
--
ALTER TABLE `Daftar_Barang`
  ADD PRIMARY KEY (`ID_Barang`),
  ADD KEY `Relasi_User_DaftarBarang` (`ID_User`);

--
-- Indexes for table `User`
--
ALTER TABLE `User`
  ADD PRIMARY KEY (`ID_User`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `Daftar_Barang`
--
ALTER TABLE `Daftar_Barang`
  MODIFY `ID_Barang` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `User`
--
ALTER TABLE `User`
  MODIFY `ID_User` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Daftar_Barang`
--
ALTER TABLE `Daftar_Barang`
  ADD CONSTRAINT `Relasi_User_DaftarBarang` FOREIGN KEY (`ID_User`) REFERENCES `User` (`ID_User`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
