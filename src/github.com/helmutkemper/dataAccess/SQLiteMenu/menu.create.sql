DROP TABLE IF EXISTS menuList;
CREATE TABLE IF NOT EXISTS
    menuList(
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                menu TEXT,          -- menu name
                admin INTEGER       -- 0: normal user; 1 admin user
    );
INSERT INTO main.menuList VALUES
(1, 'Menu Principal', 0);

DROP TABLE IF EXISTS menu;
CREATE TABLE IF NOT EXISTS
    menu (
             id INTEGER PRIMARY KEY AUTOINCREMENT,
             menuId INTEGER,        -- id menu list
             secondaryId INTEGER,   -- id of parent menu
             text TEXT,             -- menu text
             admin INTEGER,         -- 0: normal user; 1 admin user
             icon TEXT,             -- fontawesome icon [optional]
             url TEXT,			    -- link URL [optional]
             itemOrder INTEGER,     -- menu order
             FOREIGN KEY(menuId) REFERENCES menuList(id) ON DELETE restrict ON UPDATE restrict
);
INSERT INTO main.menu VALUES
(1,  1, 0, 'Kemper.com.br',            0, 'fas fa-code-branch',       '',                                                    0),
(2,  1, 0, 'Códigos',                  0, 'fas fa-code',              '',                                                    1),
(3,  1, 0, 'Admin',                    1, 'fas fa-cogs',              '',                                                    2),
(4,  1, 0, 'Donation',                 0, '',                         '',                                                    3),
(5,  1, 1, 'Sobre mim',                0, 'fas fa-info-circle',       '',                                                    0),
(6,  1, 1, 'Github',                   0, 'fas fa-info-circle',       'https://github.com/helmutkemper',                     1),
(7,  1, 1, 'LinkedIn',                 0, 'fab fa-linkedin',          'https://www.linkedin.com/in/helmut-kemper-93a5441b/', 2),
(8,  1, 2, 'Conversando com o sênior', 0, 'fas fa-fire-extinguisher', '',                                                    0),
(9,  1, 2, 'Migrando para o Go',       0, 'fas fa-fire',              '',                                                    1),
(10, 1, 3, 'Login',                    1, '',                         '',                                                    1),
(11, 1, 3, 'New content',              1, '',                         '',                                                    0);