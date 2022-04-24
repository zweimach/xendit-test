INSERT INTO organizations (
    login,
    name
)
VALUES (
    'xendit',
    'Xendit'
);

INSERT INTO comments (
    organization_id,
    text
)
SELECT o.id, 'Xendit is a leading payment gateway for Indonesia, the Philippines and Southeast Asia.'
FROM organizations AS o
WHERE o.login = 'xendit'
ORDER BY o.id
LIMIT 1;

INSERT INTO users (
    login,
    name,
    followers,
    follows,
    avatar_url
)
VALUES (
    'sandy',
    'Sandy Kurnia',
    300,
    15,
    'https://image.com/profile.jpg'
);

INSERT INTO memberships (
    user_id,
    organization_id
)
SELECT
    (SELECT id FROM users WHERE login = 'sandy' LIMIT 1),
    (SELECT id FROM organizations WHERE login = 'xendit' LIMIT 1);
