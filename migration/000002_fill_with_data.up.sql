-- scorpion:1qaz2wsx
-- kitana:123456789
INSERT INTO public.users (id, name, username, password)
VALUES  (1, 'Scorpion', 'scorpion', '6d326c6d2362676962266d7472657469626572746261366d73642c76617262216774c6922b6ba9e0939583f973bc1682493351ad4fe8'),
        (2, 'Kitana', 'kitana', '6d326c6d2362676962266d7472657469626572746261366d73642c76617262216774f7c3bc1d808e04732adf679965ccc34ca7ae3441');

INSERT INTO public.lists (id, title, description)
VALUES  (2, 'Health', 'Plans for increasing current results'),
        (3, 'Reading', 'Annual reading plan'),
        (4, 'Shoping', 'List of things to buy'),
        (1, 'Development', 'Improve knowledge in backend field');

INSERT INTO public.users_lists (id, user_id, list_id)
VALUES  (1, 1, 1),
        (2, 1, 2),
        (3, 1, 3),
        (4, 2, 4);

INSERT INTO public.items (id, title, description, done)
VALUES  (1, 'Learn Golang', 'Need extend amount of programming languages', false),
        (2, 'Polish SQL', 'ORM is not a good idea', false),
        (3, 'Learn Rust after Golang', 'Chance to enter to blockchain', false),
        (5, 'Milk', 'For ice cream', false),
        (6, 'Coffee', 'Need only LAVAZZA brand', false),
        (7, 'Rum', 'Only Havana Club', false);

INSERT INTO public.lists_items (id, item_id, list_id)
VALUES  (1, 1, 1),
        (2, 2, 1),
        (3, 3, 1),
        (6, 6, 4),
        (7, 7, 4),
        (5, 5, 4);
