CREATE OR REPLACE FUNCTION after_order_completion()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
BEGIN
    UPDATE users 
    SET points = points + FLOOR(NEW.total_price / 5)
    WHERE id = NEW.user_id::INT;
    RETURN NEW;
END;
$function$
;

CREATE OR REPLACE TRIGGER after_order_completion 
    AFTER UPDATE
    ON orders
    FOR EACH ROW
    -- Trigger se aktivuje pouze tehdy, když nový stav objednávky ("status") je "done"
    WHEN (NEW.status = 'done'::TEXT)
    EXECUTE FUNCTION after_order_completion();
