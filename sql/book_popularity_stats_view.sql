CREATE OR REPLACE VIEW public.book_popularity_stats AS
SELECT 
    b.id AS book_id,
    b.name AS book_name,
    -- Kombinace jména a příjmení autora do jednoho sloupce
    (a.first_name || ' ' || a.last_name) AS author_name,
    -- Pouze od schválených recenzí
    (SELECT count(rr.id) AS count
           FROM reviews rr
          WHERE rr.approved = true AND rr.book_id = b.id) AS total_reviews,
    b.average_rating,
    g.name AS genre_name
FROM 
    books b
    JOIN authors a ON b.author_id = a.id
    LEFT JOIN reviews r ON b.id = r.book_id
    LEFT JOIN book_genres bg ON b.id = bg.book_id
    LEFT JOIN genres g ON bg.genre_id = g.id
GROUP BY 
    -- Skupinování dat podle ID knihy, názvu knihy, jména autora a názvu žánru
    b.id, 
    b.name, 
    a.first_name, 
    a.last_name, 
    g.name;
