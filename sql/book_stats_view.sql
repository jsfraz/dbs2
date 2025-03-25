CREATE OR REPLACE VIEW public.book_stats AS
SELECT 
    b.id AS book_id,
    b.name AS book_name,
    -- Složení celého jména autora (jméno + příjmení) do jednoho sloupce
    (a.first_name || ' '::text) || a.last_name AS author_name,
    count(uob.book_id) AS total_sales,
    sum(b.price) AS total_revenue,
    COALESCE(AVG(r.stars), 0) AS average_rating,
    count(DISTINCT r.id) AS total_reviews
FROM 
    books b
    JOIN authors a ON b.author_id = a.id
    LEFT JOIN user_order_books uob ON b.id = uob.book_id
    LEFT JOIN reviews r ON b.id = r.book_id
GROUP BY 
    -- Skupinování dat podle ID knihy, názvu knihy a jména autora
    b.id, 
    b.name, 
    a.first_name, 
    a.last_name;
