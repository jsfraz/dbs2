-- Funkce pro automatické dokončení objednávky
CREATE OR REPLACE FUNCTION complete_order_after_fake()
RETURNS trigger
LANGUAGE plpgsql
AS $function$
BEGIN
    -- Aktualizujeme status objednávky na 'done'
    UPDATE orders
    SET status = 'done'
    WHERE id = NEW.id;
    
    RETURN NEW;
END;
$function$;

-- Trigger, který se spustí po vložení nové objednávky se statusem 'pending'
CREATE OR REPLACE TRIGGER complete_order_trigger
    AFTER INSERT
    ON orders
    FOR EACH ROW
    WHEN (NEW.status = 'pending')
    EXECUTE FUNCTION complete_order_after_fake();
