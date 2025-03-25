-- Funkce pro "koupení" slevy uživatelem
CREATE OR REPLACE FUNCTION purchase_discount(
    p_user_id INT,
    p_discount_code TEXT,
    p_price NUMERIC,
    p_valid_until TIMESTAMP 
) 
RETURNS VOID AS $$
BEGIN
    -- Volání: SELECT purchase_discount(22, 'EXAMPLE-CODE', 100, (now() + INTERVAL '1 month')::TIMESTAMP);
    BEGIN
        -- Kontrola, zda uživatel má dostatek bodů
        IF (SELECT points FROM users WHERE id = p_user_id) < p_price THEN
            RAISE EXCEPTION 'Uživatel #% nemá dostatek bodů (potřebuje %, má %).', 
                            p_user_id, 
                            p_price, 
                            (SELECT points FROM users WHERE id = p_user_id);
        END IF;
        -- Vytvoření slevy v tabulce discounts
        INSERT INTO discounts (
            user_id, 
            code, 
            valid_until,
			used,
			price
        ) VALUES (
            p_user_id,
            p_discount_code,
            p_valid_until,
            FALSE,  -- Sleva je zatím nepoužitá
			p_price
        );
        -- Odečtení bodů uživateli
        UPDATE users 
        SET points = points - p_price 
        WHERE id = p_user_id;
    -- Ošetření chyb
    EXCEPTION 
        WHEN unique_violation THEN
            RAISE EXCEPTION 'Slevový kód "%" již existuje!', p_discount_code;
        WHEN others THEN
            RAISE;
    END;
END;
$$ LANGUAGE plpgsql;
