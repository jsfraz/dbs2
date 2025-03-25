CREATE OR REPLACE PROCEDURE delete_user(p_user_id INT)
LANGUAGE plpgsql AS $$
BEGIN
    -- Zahájení explicitní transakce
    BEGIN
        -- Smazání všech objednávek uživatele
        DELETE FROM orders WHERE user_id = p_user_id;
        -- Smazání všech recenzí uživatele
        DELETE FROM reviews WHERE user_id = p_user_id;
        -- Smazání všech položek v košíku uživatele
        DELETE FROM carts WHERE user_id = p_user_id;
        -- Smazání všech položek na seznamu přání uživatele
        DELETE FROM wishlists WHERE user_id = p_user_id;
        -- Smazání všech slev přiřazených uživateli
        DELETE FROM discounts WHERE user_id = p_user_id;
        -- Nakonec smažeme samotného uživatele
        DELETE FROM users WHERE id = p_user_id;
    EXCEPTION WHEN OTHERS THEN
        -- Pokud dojde k jakékoli chybě, vrátíme všechny změny zpět
        RAISE NOTICE 'Došlo k chybě při mazání uživatele s ID %. Změny byly vráceny zpět.', p_user_id;
        ROLLBACK;
    END;
END;
$$;
