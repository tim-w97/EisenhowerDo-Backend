SELECT id, userID, title, text, isImportant, isUrgent
FROM todo
WHERE userID = ?