CREATE OR REPLACE FUNCTION update_book_rating()
RETURNS trigger
LANGUAGE plpgsql
AS $function$
DECLARE
    v_avg_rating DECIMAL(3,2);  -- Proměnná pro uložení průměrného hodnocení
BEGIN
    -- Vypočítá průměrné hodnocení pro knihu na základě recenzí
    SELECT AVG(stars) INTO v_avg_rating
    FROM reviews
    WHERE book_id = COALESCE(NEW.book_id, OLD.book_id) AND approved = true;
    -- Aktualizuje průměrné hodnocení knihy v tabulce books
    UPDATE books
    SET average_rating = COALESCE(v_avg_rating, 0)  -- Pokud není průměr, nastaví 0
    WHERE id = COALESCE(NEW.book_id, OLD.book_id);
    RETURN NEW;
END;
$function$
;

CREATE OR REPLACE TRIGGER update_book_rating_trigger
AFTER INSERT OR UPDATE OR DELETE
ON reviews
FOR EACH ROW
EXECUTE FUNCTION update_book_rating();
