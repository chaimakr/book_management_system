CREATE TABLE `books`
(
    id   bigint auto_increment,
    name varchar(50) NOT NULL,
    author varchar(50) NOT NULL,
    synopsis varchar(255),
    PRIMARY KEY (`id`)
);

INSERT INTO `books` (`name`,`author`,`synopsis`)
VALUES  ('The undomestic goddess','Sophie Kinsella',''),
        ('The Duke and I','Julia Quinn',''),
        ('The love hypothesis','Ali hazelwood',''),
        ('It ends with us','Colleen Hoover','');

    
