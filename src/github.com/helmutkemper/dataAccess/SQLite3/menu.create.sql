CREATE TABLE IF NOT EXISTS
    menu (
             id INTEGER PRIMARY KEY AUTOINCREMENT,
             secondaryId INTEGER,   -- id of parent menu
             text TEXT,             -- menu text
             admin INTEGER,         -- 0: normal user; 1 admin user
             icon TEXT,             -- fontawesome icon [optional]
             url TEXT,			         -- link URL [optional]
             itemOrder INTEGER      -- menu order
);
INSERT INTO main.menu VALUES
(1,  0, 'Kemper.com.br',            0, 'fas fa-code-branch',       '',                                                    0),
(2,  0, 'Códigos',                  0, 'fas fa-code',              '',                                                    1),
(3,  0, 'Admin',                    1, 'fas fa-cogs',              '',                                                    2),
(4,  0, 'Donation',                 0, '',                         '',                                                    3),
(5,  1, 'Sobre mim',                0, 'fas fa-info-circle',       '',                                                    0),
(6,  1, 'Github',                   0, 'fas fa-info-circle',       'https://github.com/helmutkemper',                     1),
(7,  1, 'LinkedIn',                 0, 'fab fa-linkedin',          'https://www.linkedin.com/in/helmut-kemper-93a5441b/', 2),
(8,  2, 'Conversando com o sênior', 0, 'fas fa-fire-extinguisher', '',                                                    0),
(9,  2, 'Migrando para o Go',       0, 'fas fa-fire',              '',                                                    1),
(10, 3, 'Login',                    1, '',                         '',                                                    1),
(11, 3, 'New content',              1, '',                         '',                                                    0);