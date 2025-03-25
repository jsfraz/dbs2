CREATE OR REPLACE PROCEDURE delete_author(p_author_id INT)
LANGUAGE plpgsql AS $$
DECLARE
    v_author_name TEXT;
    book_id INT;
BEGIN
    -- Kontrola existence autora
    SELECT first_name || ' ' || last_name INTO v_author_name
    FROM authors WHERE id = p_author_id;
    IF v_author_name IS NULL THEN
        RAISE EXCEPTION 'Autor s ID % neexistuje.', p_author_id;
    END IF;
    -- Smazání všech knih autora pomocí procedury delete_book
    FOR book_id IN SELECT id FROM books WHERE author_id = p_author_id
    LOOP
        CALL delete_book(book_id);
    END LOOP;
    -- Smazání autora
    DELETE FROM authors WHERE id = p_author_id;
EXCEPTION WHEN OTHERS THEN
    RAISE NOTICE 'Chyba při mazání autora "%" (ID: %): %', v_author_name, p_author_id, SQLERRM;
    ROLLBACK;
    RAISE;
END;
$$;
