SELECT setval('addresses_id_seq', (SELECT MAX(id) FROM addresses));
SELECT setval('authors_id_seq', (SELECT MAX(id) FROM authors));
SELECT setval('books_id_seq', (SELECT MAX(id) FROM books));
SELECT setval('discounts_id_seq', (SELECT MAX(id) FROM discounts));
SELECT setval('genres_id_seq', (SELECT MAX(id) FROM genres));
SELECT setval('orders_id_seq', (SELECT MAX(id) FROM orders));
SELECT setval('reviews_id_seq', (SELECT MAX(id) FROM reviews));
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));