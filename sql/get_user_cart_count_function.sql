CREATE OR REPLACE FUNCTION get_user_cart_count(p_user_id INT)
RETURNS INT AS $$
DECLARE
    v_cart_count INT;
BEGIN
    SELECT COUNT(*)
    INTO v_cart_count
    FROM carts
    WHERE user_id = p_user_id;
    -- Vrácení počtu knih
    RETURN v_cart_count;
END;
$$ LANGUAGE plpgsql;
