-- +goose Up
CREATE TABLE IF NOT EXISTS songs_info
(
    id int NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name varchar(50) NOT NULL,                                
    group_name varchar(50),                                   
    link text,                                                
    release_date varchar(50)                                        
);

-- +goose Down
DROP TABLE IF EXISTS songs_info;
