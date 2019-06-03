CREATE TABLE google_like_search_engine.page (
    `id`    INTEGER PRIMARY KEY AUTO_INCREMENT,
    `page_title` TEXT,
    `page_url`   TEXT,
    `page_html`  LONGTEXT
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

