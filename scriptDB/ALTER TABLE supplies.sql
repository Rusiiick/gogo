ALTER TABLE supplies
ADD category_id integer,
    ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES category(category_id);