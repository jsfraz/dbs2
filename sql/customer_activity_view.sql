CREATE OR REPLACE VIEW public.customer_activity AS
SELECT 
    u.id AS user_id,
    u.first_name || ' ' || u.last_name AS full_name,
    u.mail AS email,
    -- Počet unikátních objednávek vytvořených uživatelem
    COUNT(DISTINCT o.id) AS total_orders,
    -- Počet unikátních knih v košíku uživatele
    COUNT(DISTINCT c.book_id) AS total_cart_books,
    -- Počet unikátních knih na seznamu přání uživatele
    COUNT(DISTINCT w.book_id) AS total_wishlist_books,
    -- Počet recenzí napsaných uživatelem
    COUNT(DISTINCT r.id) AS total_reviews,
    COALESCE(SUM(o.total_price), 0) AS total_spent,
    u.points AS points
FROM 
    users u
    LEFT JOIN orders o ON u.id = o.user_id
    LEFT JOIN carts c ON u.id = c.user_id
    LEFT JOIN wishlists w ON u.id = w.user_id
    LEFT JOIN reviews r ON u.id = r.user_id
WHERE 
    u.role = 'customer'
GROUP BY 
    -- Skupinování dat podle ID uživatele a dalších atributů, aby bylo možné použít agregační funkce jako COUNT a SUM
    u.id, u.first_name, u.last_name, u.mail, u.role, u.points;