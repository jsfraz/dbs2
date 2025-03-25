CREATE OR REPLACE PROCEDURE delete_book(p_book_id INT)
LANGUAGE plpgsql AS $$
DECLARE
    v_book_name TEXT;
BEGIN
    BEGIN
        SELECT name INTO v_book_name FROM books WHERE id = p_book_id;
        IF v_book_name IS NULL THEN
            RAISE EXCEPTION 'Kniha s ID % neexistuje.', p_book_id;
        END IF;
        -- Smazání záznamů z tabulky user_order_books
        DELETE FROM user_order_books WHERE book_id = p_book_id;
        -- Smazání záznamů z tabulky carts
        DELETE FROM carts WHERE book_id = p_book_id;
        -- Smazání záznamů z tabulky wishlists
        DELETE FROM wishlists WHERE book_id = p_book_id;
        -- Smazání záznamů z tabulky reviews
        DELETE FROM reviews WHERE book_id = p_book_id;
        -- Smazání záznamů z tabulky book_genres
        DELETE FROM book_genres WHERE book_id = p_book_id;
        -- Smazání knihy
        DELETE FROM books WHERE id = p_book_id;
    EXCEPTION WHEN OTHERS THEN
        -- Pokud dojde k jakékoli chybě, vrátíme všechny změny zpět
        RAISE NOTICE 'Došlo k chybě při mazání knihy "%" s ID %. Změny byly vráceny zpět. Chyba: %', v_book_name, p_book_id, SQLERRM;
        ROLLBACK;
    END;
END;
$$;
