-- +goose Up
CREATE TABLE IF NOT EXISTS songs_text
(
    id int NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY, 
    track_id int NOT NULL,                                    
    couplet_number int NOT NULL,                              
    couplet_text text NOT NULL,                              
    CONSTRAINT fk_track FOREIGN KEY (track_id) REFERENCES songs_info(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS songs_text;